package templatemethod

import "fmt"

// TemplateMethod is a behavioral design pattern that defines the skeleton of an algorithm in the superclass but lets subclasses override specific steps of the algorithm without changing its structure.
func TemplateMethod() {
	t := MessageSender{
		iMessageSender: &telegram{},
	}
	err := t.buildAndSendMessage("Hello Telegram message")
	if err != nil {
		fmt.Println(err)
	}

	v := MessageSender{
		iMessageSender: &viber{},
	}
	err = v.buildAndSendMessage("Hello Viber message")
	if err != nil {
		fmt.Println(err)
	}
}

// iMessageSender is an interface that defines the message sending algorithm.
type iMessageSender interface {
	BuildMessage(message string) string
	SendMessage(message string) error
}

// MessageSender is a concrete implementation of the iMessageSender interface.
type MessageSender struct {
	iMessageSender iMessageSender
}

func (m *MessageSender) buildAndSendMessage(message string) error {
	// message sending algorithm
	// 1. build a message
	// 2. send the message
	// 3. handle error
	m.iMessageSender.BuildMessage(message)
	err := m.iMessageSender.SendMessage(message)
	if err != nil {
		return err
	}

	return nil
}
