package main

import (
	"fmt"
	"reflect"
)

func main() {
	var msg = TextMsg{}
	msg.setTextMsg("this is a text message")
	msg.send()

	send(&msg)

	setAndSendMessage(&TextMsg{})
	setAndSendMessage(&ImgMsg{})

	nilInterface()
}

func nilInterface() {
	var i int = 7
	var object = interface{}(i)
	fmt.Println(object, "is", reflect.TypeOf(object).Name())
	s, ok := object.(string)
	if ok {
		fmt.Println(s)
	} else {
		fmt.Println(object, "is not string type")
	}
}

// Message interface
type Message interface {
	getType() string
	send()
}

func send(msg Message) {
	msg.send()
}

func setAndSendMessage(msg Message) {
	switch mp := msg.(type) {
	case *TextMsg:
		mp.setTextMsg("set a text")
	case *ImgMsg:
		mp.setImgMsg("set a img")
	}
	msg.send()
}

// TextMsg obj
type TextMsg struct {
	text string
}

func (tm *TextMsg) getType() string {
	return "text"
}

func (tm *TextMsg) setTextMsg(msg string) {
	tm.text = msg
}

func (tm *TextMsg) send() {
	fmt.Println("sendText -> ", tm.text, "; type is ", tm.getType())
}

type ImgMsg struct {
	img string
}

func (im *ImgMsg) getType() string {
	return "img"
}

func (im *ImgMsg) setImgMsg(img string) {
	im.img = img
}

func (im *ImgMsg) send() {
	fmt.Println("sendImage -> ", im.img, "; type is ", im.getType())
}
