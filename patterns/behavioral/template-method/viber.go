package templatemethod

import "fmt"

type viber struct {
	iMessageSender
}

// checking if value implements interface
var _ iMessageSender = (*viber)(nil)

func (v *viber) BuildMessage(message string) string {
	a := "Viber"
	b := "message"
	viberMessage := fmt.Sprintf("%s %s: %s", a, b, message)
	return viberMessage
}

func (t *viber) SendMessage(message string) error {
	fmt.Println("Sending viber message:", message)
	return nil
}
