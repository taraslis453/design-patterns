package monitor

import "fmt"

type dell struct {
	Model string
	Cable VGACable
}

func NewDell(model string, cable VGACable) *dell {
	return &dell{Cable: cable, Model: model}
}

var _ Monitor = (*dell)(nil)

func (d dell) InsertIntoVGAPort(cable VGACable) {
	fmt.Println("Inserting into vga port")
}
