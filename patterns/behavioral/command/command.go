package command

import "fmt"

type Invoker struct {
	command Command
}

type Command interface {
	Execute()
}

type Receiver interface {
	On()
}

type OnCommand struct {
	receiver Receiver
}

func (o OnCommand) Execute() {
	o.receiver.On()
}

type Light struct{}

func (l Light) On() {
	fmt.Println("Light is on")
}

func (l Light) Off() {
	fmt.Println("Light is off")
}

// client code
func RunCommand() {
	lightInvoker := Invoker{OnCommand{Light{}}}
	lightInvoker.command.Execute()
}
