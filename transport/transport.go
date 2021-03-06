// Package transport implements an encrypted and authenticated connection for
// transferring byte slices over a network. The key-exchange is like SIGMA-I,
// with the exception that instead of unvirersally verifiable signatures,
// nacl/box between one party's ephemeral key and the other's long-term key is
// used for authentication (this provides deniability). Subsequent messages are
// encrypted using nacl/box with the message counter as an implicit nonce.
package transport

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"time"

	"code.google.com/p/go.crypto/nacl/box"
)

// Conn is an encrypted and authenticated connection that is NOT concurrency-safe
type Conn struct {
	unencrypted           net.Conn
	readNonce, writeNonce uint64
	key                   [32]byte
	readBuf, writeBuf     []byte
	maxFrameSize          int
}

var nullNonce = [24]byte{}

// Handshake establishes an encrypted and authenticated connection. unencrypted
// is the underlying connection that will be used for the handshake and the
// following calls to ReadFrame and WriteFrame. The connection should not be
// used after calling Handshake. pk and sk are the Curve25519 public and
// private keys of the caller. If expectedPK is not nil, pk will not be
// revealed to the other party unless they prove that they hold the secret key
// that corresponds to expectedPK.  Note that both sides of a connection using
// this option will result in a deadlock. The public key of the other party is
// returned along with the wrapped connection.
func Handshake(unencrypted net.Conn, pk, sk, expectedPK *[32]byte, maxFrameSize int) (*Conn, *[32]byte, error) {
	if sk == nil && pk == nil {
		var err error
		pk, sk, err = box.GenerateKey(rand.Reader)
		if err != nil {
			return nil, nil, err
		}
	}
	// All single-letter symbols in this comment represent public DH keys.
	// Capital letters represent long-term and others are per-connection
	// ephemeral. [msg](PK1<>PK2) denotes nacl/box authenticated encryption.
	//	--> a
	//	<-- b
	//	<-- [B,[b,a](B<>a)](a<>b)
	//	--> [A,[a,b](A<>b)](a<>b)
	//	--> [data](a<>b)
	//	<-- [data](a<>b)

	ourEphemeralPublic, ourEphemeralSecret, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	var theirEphemeralPublic, theirPK [32]byte
	var readErr, writeErr error
	writeDone, readDone := make(chan struct{}), make(chan struct{})
	go func() { _, writeErr = unencrypted.Write(ourEphemeralPublic[:]); close(writeDone) }()
	go func() { _, readErr = io.ReadFull(unencrypted, theirEphemeralPublic[:]); close(readDone) }()
	if <-writeDone; writeErr != nil {
		return nil, nil, writeErr
	}
	if <-readDone; readErr != nil {
		return nil, nil, readErr
	}

	ret := &Conn{unencrypted: unencrypted, maxFrameSize: maxFrameSize,
		readBuf:  make([]byte, binary.MaxVarintLen64+box.Overhead+maxFrameSize),
		writeBuf: make([]byte, binary.MaxVarintLen64+box.Overhead+maxFrameSize)}
	if bytes.Compare(ourEphemeralPublic[:], theirEphemeralPublic[:]) < 0 {
		ret.writeNonce = 1
	} else {
		ret.readNonce = 1
	}
	box.Precompute(&ret.key, &theirEphemeralPublic, ourEphemeralSecret)

	writeDone, readDone = make(chan struct{}), make(chan struct{})
	go func() {
		defer close(readDone)
		var theirHandshake [32 + (box.Overhead + 32 + 32)]byte // theirPK, box(theirEphPK, ourEphPK)
		if _, readErr = ret.ReadFrame(theirHandshake[:]); readErr != nil {
			return
		}
		copy(theirPK[:], theirHandshake[:32])
		hs, ok := box.Open(nil, theirHandshake[32:], &nullNonce, &theirPK, ourEphemeralSecret)
		if !ok || !bytes.Equal(hs, append(theirEphemeralPublic[:], ourEphemeralPublic[:]...)) {
			readErr = errors.New("authentication failed (ephemeral pk mismatch)")
			return
		}
		if bytes.Equal(theirPK[:], pk[:]) {
			readErr = errors.New("we are talking to a mirror")
			return
		}
		if expectedPK != nil && !bytes.Equal(theirPK[:], expectedPK[:]) {
			readErr = errors.New("authentication failed (observed pk != expected pk)")
			return
		}
	}()
	go func() {
		defer close(writeDone)
		ourHandshake := box.Seal(pk[:], append(ourEphemeralPublic[:], theirEphemeralPublic[:]...),
			&nullNonce, &theirEphemeralPublic, sk)
		if expectedPK != nil {
			if <-readDone; readErr != nil { // only talk to the right server
				return
			}
		}
		_, writeErr = ret.WriteFrame(ourHandshake)
	}()
	if <-readDone; readErr != nil {
		return nil, nil, readErr
	}
	if <-writeDone; writeErr != nil {
		return nil, nil, writeErr
	}
	for i := range ourEphemeralSecret {
		ourEphemeralSecret[i] = 0
	}
	return ret, &theirPK, nil
}

// WriteFrame(b) writes the frame to the connection in a length-value-encoded
// for so it can be read using ReadFrame on the other side. Returns len(b).
func (c *Conn) WriteFrame(b []byte) (int, error) {
	if len(b) > c.maxFrameSize {
		return 0, errors.New("write frame too large")
	}
	var nonce [24]byte
	binary.LittleEndian.PutUint64(nonce[:], c.writeNonce)
	c.writeNonce += 2
	i := binary.PutUvarint(c.writeBuf, uint64(box.Overhead+len(b)))
	buf := box.SealAfterPrecomputation(c.writeBuf[:i], b, &nonce, &c.key)
	if _, err := c.unencrypted.Write(buf); err != nil {
		return 0, err
	}
	return len(b), nil
}

type byteReader struct{ io.Reader }

func (r byteReader) ReadByte() (byte, error) {
	var ret [1]byte
	_, err := io.ReadFull(r, ret[:])
	return ret[0], err
}

// ReadFrame(b) reads a single frame into b and returns an integer n such that
// b[:n] is the frame after possibly modifying b. If b does not have enough
// space (maxFrameSize bytes), this function may panic.
func (c *Conn) ReadFrame(b []byte) (int, error) {
	var nonce [24]byte
	binary.LittleEndian.PutUint64(nonce[:], c.readNonce)
	c.readNonce += 2
	size, err := binary.ReadUvarint(byteReader{c.unencrypted})
	if err != nil {
		return 0, err
	}
	if _, err := io.ReadFull(c.unencrypted, c.readBuf[:size]); err != nil {
		return 0, err
	}
	b2, ok := box.OpenAfterPrecomputation(b[:0], c.readBuf[:size], &nonce, &c.key)
	if !ok {
		return 0, errors.New("authentication failed")
	}
	if &b[0] != &b2[0] {
		panic("ReadFrame buffer space accounting failed")
	}
	return int(size - box.Overhead), nil
}

func (c *Conn) Close() error {
	for i := range c.key {
		c.key[i] = 0
	}
	return c.unencrypted.Close()
}

func (c *Conn) SetDeadline(t time.Time) error      { return c.unencrypted.SetDeadline(t) }
func (c *Conn) SetWriteDeadline(t time.Time) error { return c.unencrypted.SetWriteDeadline(t) }
func (c *Conn) SetReadDeadline(t time.Time) error  { return c.unencrypted.SetReadDeadline(t) }
