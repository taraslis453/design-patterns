package templatemethod

import "fmt"

type telegram struct {
	iMessageSender
}

// checking if value implements interface
var _ iMessageSender = (*telegram)(nil)

func (t *telegram) BuildMessage(message string) string {
	telegramMessage := fmt.Sprintf("Telegram message: %s", message)
	return telegramMessage
}

func (t *telegram) SendMessage(message string) error {
	fmt.Println("Sending telegram message:", message)
	return nil
}
