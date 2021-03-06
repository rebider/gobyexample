package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"time"
)

type Message struct {
	name    string
	Content map[string]string `json:"content"`
}

func (m *Message) GetChannel() string {
	return m.name
}

func (m *Message) Resolve() error {
	time.Sleep(time.Second)
	fmt.Printf("consumed %+v\n", m.Content)
	return nil
}

func (m *Message) Marshal() ([]byte, error) {
	return jsoniter.Marshal(m)
}

func (m *Message) Unmarshal(reply []byte) (IMessage, error) {
	var msg Message
	err := jsoniter.Unmarshal(reply, &msg)

	return &msg, err
}
