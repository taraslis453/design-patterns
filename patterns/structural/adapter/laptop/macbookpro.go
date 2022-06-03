package laptop

import "fmt"

type macbook struct {
	Model string
	Cable LightningCable
}

func NewMacbook(Model string, Cable LightningCable) *macbook {
	return &macbook{Model, Cable}
}

var _ Laptop = (*macbook)(nil)

func (m macbook) InsertIntoLightningPort(cable LightningCable) {
	fmt.Println("Inserting into lightning port")
}
