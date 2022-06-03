package laptop

type Laptop interface {
	InsertIntoLightningPort(cable LightningCable)
}

type LightningCable struct {
	Voltage int
}
