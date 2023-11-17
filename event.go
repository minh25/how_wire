package main

import "fmt"

type Message string

type Greater struct {
	Message Message
}

type Event struct {
	Greater Greater
}

func NewMessage(s string) Message {
	return Message(s)
}

func NewGreater(m Message) Greater {
	return Greater{Message: m}
}

func (g Greater) Greet() Message {
	return g.Message
}

func NewEvent(g Greater) Event {
	return Event{Greater: g}
}

func (e Event) Start() {
	msg := e.Greater.Greet()
	fmt.Println(msg)
}
