package main

// This file is automatically generated by gopkg.in/qml.v1/cmd/genqrc

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/qml.v1"
)

func init() {
	var r *qml.Resources
	var err error
	if os.Getenv("QRC_REPACK") == "1" {
		err = qrcRepackResources()
		if err != nil {
			panic("cannot repack qrc resources: " + err.Error())
		}
		r, err = qml.ParseResources(qrcResourcesRepacked)
	} else {
		r, err = qml.ParseResourcesString(qrcResourcesData)
	}
	if err != nil {
		panic("cannot parse bundled resources data: " + err.Error())
	}
	qml.LoadResources(r)
}

func qrcRepackResources() error {
	subdirs := []string{"qml"}
	var rp qml.ResourcesPacker
	for _, subdir := range subdirs {
		err := filepath.Walk(subdir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			rp.Add(filepath.ToSlash(path), data)
			return nil
		})
		if err != nil {
			return err
		}
	}
	qrcResourcesRepacked = rp.Pack().Bytes()
	return nil
}

var qrcResourcesRepacked []byte
var qrcResourcesData = "qres\x00\x00\x00\x01\x00\x00\x18\xca\x00\x00\x00\x14\x00\x00\x18F\x00\x00\t<import QtQuick 2.2\nimport QtQuick.Controls 1.1\nimport QtQuick.Layouts 1.1\n\n\nApplicationWindow {\n\tid: conversationWindow\n\tsignal sendMessage(string message)\n\n    visible: true\n    title: \"Conversation\"\n    property int margin: 10\n    width: mainLayout.implicitWidth + 2 * margin\n    height: mainLayout.implicitHeight + 2 * margin\n    minimumWidth: mainLayout.Layout.minimumWidth + 40 * margin\n    minimumHeight: mainLayout.Layout.minimumHeight + 12 * margin\n\n\tAction {\n\t\tid: sendMessage\n\t\ttext: \"Send &Message\"\n\t\tshortcut: \"Ctrl+Return\"\n\t\tonTriggered: {\n\t\t\tconversationWindow.sendMessage(messageArea.text);\n\t\t\tmessageArea.remove(0, messageArea.length);\n\t\t}\n\t}\n\n\tListModel {\n\t\tid: messageModel\n\t\tobjectName: 'messageModel'\n    \tListElement{\n    \t\tContent: \"test message\"\n    \t\tSender:\"Jane\"\n    \t}\n    \tListElement{\n    \t\tContent: \"test message\"\n    \t\tSender:\"Jane\"\n    \t}\n\n\n\t\tfunction addItem(json) {\n\t\t\tconsole.log(json)\n\t\t\tvar parsed = JSON.parse(json);\n\t\t\tfor (var key in parsed) {\n\t\t\t\tif (parsed.hasOwnProperty(key) && (typeof parsed[key] == 'object')) {\n\t\t\t\t\t\t//console.log(key);\n\t\t\t\t\t\tparsed[key] = parsed[key].toString();\n\t\t\t\t}\n\t\t\t}\n\t\t\tparsed['objectName'] = parsed['Subject'];\n\t\t\tappend(parsed);\n\t\t}\n    }\n\n    ColumnLayout {\n        id: mainLayout\n        anchors.fill: parent\n        anchors.margins: margin\n        spacing:20\n\t\tRowLayout {\n\t\t\tText {\n\t\t\t\ttext: \"To:\"\n\t\t\t\tobjectName: \"to\"\n\t\t\t}\n\t\t}\n\n\t\tRowLayout {\n\t\t\tText {\n\t\t\t\ttext: \"Subject:\"\n\t\t\t\tobjectName: \"subject\"\n\t\t\t}\n\t\t}\n\n\t\tRowLayout {\n\t\t\tListView {\n\t\t\t\tid: messageView\n\t\t        objectName: \"messageView\"\n\n\t\t        anchors.fill: parent\n\t\t        model: messageModel\n\t\t        delegate: RowLayout {\n\t\t        \tText{ \n\t\t        \t\tanchors.top: parent.top\n\t\t        \t\ttext: Sender + \": \"\n\t\t        \t\ttextFormat: Text.PlainText\n\t\t        \t\tfont.bold:true\n\t\t        \t}\n\t\t        \tText{ \n\t\t        \t\tanchors.top: parent.top\n\t\t        \t\ttext: Content\n\t\t        \t\ttextFormat: Text.PlainText\n\t\t        \t}\n\n\t\t        }\n\n\t\t        Layout.minimumWidth: 100\n\t\t        Layout.minimumHeight: 100\n\t\t        Layout.preferredWidth: 300\n\t\t        Layout.preferredHeight: 180\n\t\t\t}\n\t\t}\n\n\n\t\tTextArea {\n\t\t\tid: messageArea \n\t\t\ttext: \"Ctrl + Enter to send a message.\"\n\t\t\tLayout.minimumHeight: 10\n\t\t\tLayout.fillWidth: true\n\t\t\tLayout.fillHeight: true\n\t\t\ttextFormat: TextEdit.PlainText\n\t\t\twrapMode: TextEdit.Wrap\n\t\t}\n    }\n}\n\x00\x00\a\xd7import QtQuick 2.2\nimport QtQuick.Controls 1.1\nimport QtQuick.Layouts 1.1\n\nApplicationWindow {\n\tid: historyWindow\n\n    visible: true\n    title: \"History\"\n    property int margin: 5\n    width: mainLayout.implicitWidth + 2 * margin\n    height: mainLayout.implicitHeight + 2 * margin\n    minimumWidth: mainLayout.Layout.minimumWidth + 40 * margin\n    minimumHeight: mainLayout.Layout.minimumHeight + 12 * margin\n\n\tListModel {\n\t    id: sourceModel\n\t\tobjectName: \"listModel\"\n\n\t\tfunction addItem(json) {\n\t\t\tvar parsed = JSON.parse(json);\n\t\t\tfor (var key in parsed) {\n\t\t\t\tif (parsed.hasOwnProperty(key) && (typeof parsed[key] == 'object')) {\n\t\t\t\t\t\t//console.log(key);\n\t\t\t\t\t\tparsed[key] = parsed[key].toString();\n\t\t\t\t}\n\t\t\t}\n\t\t\tparsed['objectName'] = parsed['Subject'];\n\t\t\tappend(parsed);\n\t\t}\n\t}\n\n\n    ColumnLayout {\n        id: mainLayout\n        anchors.fill: parent\n        anchors.margins: margin\n\n\t    TableView {\n\t        id: tableView\n\t        objectName: \"table\"\n\n\t        frameVisible: false\n\t        sortIndicatorVisible: true\n\n\t        anchors.fill: parent\n\t        model: sourceModel\n\n\t        Layout.minimumWidth: 200\n\t        Layout.minimumHeight: 240\n\t        Layout.preferredWidth: 400\n\t        Layout.preferredHeight: 300\n\n\t        TableViewColumn {\n\t            id: subjectColumn\n\t            title: \"Subject\"\n\t            role: \"Subject\"\n\t            movable: false\n\t            resizable: false\n\t            width: tableView.viewport.width / 4\n\t        }\n\n\t        TableViewColumn {\n\t            id: usersColumn\n\t            title: \"Participants\"\n\t            role: \"Users\"\n\t            movable: false\n\t            resizable: false\n\t            width: tableView.viewport.width / 4\n\t        }\n\n\t        TableViewColumn {\n\t            id: lastMessageColumn\n\t            title: \"Last Message\"\n\t            role: \"LastMessage\"\n\t            movable: false\n\t            resizable: false\n\t            width: tableView.viewport.width - usersColumn.width - subjectColumn.width\n\t        }\n \n\t    }\n    }\n}\n\x00\x00\a\x13import QtQuick 2.2\nimport QtQuick.Controls 1.1\nimport QtQuick.Layouts 1.1\n\n\nApplicationWindow {\n\tid: newConversationWindow\n\tsignal sendMessage(string to, string subject, string message)\n\n    visible: true\n    title: \"New Conversation\"\n    property int margin: 5\n    width: mainLayout.implicitWidth + 2 * margin\n    height: mainLayout.implicitHeight + 2 * margin\n    minimumWidth: mainLayout.Layout.minimumWidth + 40 * margin\n    minimumHeight: mainLayout.Layout.minimumHeight + 12 * margin\n\n\tAction {\n\t\tid: sendMessage\n\t\ttext: \"Send &Message\"\n\t\tshortcut: \"Ctrl+Return\"\n\t\tonTriggered: {\n\t\t\tnewConversationWindow.sendMessage(toField.text, subjectField.text, messageArea.text)\n\t\t\tmessageArea.text = \"\"\n\t\t}\n\t}\n\n    ColumnLayout {\n        id: mainLayout\n        anchors.fill: parent\n        anchors.margins: margin\n\t\tRowLayout {\n\t\t\tText {text: \"To:\"}\n\t\t\t\tTextField {\n\t\t\t\t\tid: toField\n\t\t\t\t\t\tfocus: true\n\t\t\t\t\t\tplaceholderText: \"dename names, comma-separated\"\n\t\t\t\t\t\tLayout.fillWidth: true\n\t\t\t\t\t\tonAccepted: {subjectField.focus = true}\n\t\t\t\t}\n\t\t}\n\n\t\tRowLayout {\n\t\t\tText {text: \"Subject:\"}\n\t\t\t\tTextField {\n\t\t\t\t\tid: subjectField\n\t\t\t\t\tLayout.fillWidth: true\n\t\t\t\t\tonAccepted: {messageArea.selectAll(); messageArea.focus = true}\n\t\t\t\t}\n\t\t}\n\n\n\t\tTextArea {\n\t\t\tid: messageArea \n\t\t\ttext: \"Ctrl + Enter to send a message.\"\n\t\t\tLayout.minimumHeight: 10\n\t\t\tLayout.fillWidth: true\n\t\t\tLayout.fillHeight: true\n\t\t\ttextFormat: TextEdit.PlainText\n\t\t\twrapMode: TextEdit.Wrap\n\t\t\t/* andreser: the following works for me\n\t\t\tkeys.onreturnpressed: {\n\t\t\t\tconsole.log(\"return pressed in main textarea\");\n\t\t\t}\n\t\t\t*/\n\t\t}\n\n\t\t/*\n        GroupBox {\n            id: sendMessageButton\n            Layout.alignment: Qt.AlignRight\n            flat: true\n\t\t\tButton {\n\t\t\t\ttext: \"Send Message\"\n\t\t\t\tonClicked: {sendMessage.trigger()}\n\t\t\t}\n        }\n\t\t*/\n    }\n}\n\x00\x03\x00\x00x<\x00q\x00m\x00l\x00\x14\x00<\xd7|\x00o\x00l\x00d\x00-\x00c\x00o\x00n\x00v\x00e\x00r\x00s\x00a\x00t\x00i\x00o\x00n\x00.\x00q\x00m\x00l\x00\v\x06FE\\\x00h\x00i\x00s\x00t\x00o\x00r\x00y\x00.\x00q\x00m\x00l\x00\x14\a|\xd6|\x00n\x00e\x00w\x00-\x00c\x00o\x00n\x00v\x00e\x00r\x00s\x00a\x00t\x00i\x00o\x00n\x00.\x00q\x00m\x00l\x00\x00\x00\x00\x00\x02\x00\x00\x00\x01\x00\x00\x00\x01\x00\x00\x00\x00\x00\x02\x00\x00\x00\x03\x00\x00\x00\x02\x00\x00\x00\f\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\x00\x00\x00:\x00\x00\x00\x00\x00\x01\x00\x00\t@\x00\x00\x00V\x00\x00\x00\x00\x00\x01\x00\x00\x11\x1b"
