package monitor

type Monitor interface {
	InsertIntoVGAPort(cable VGACable)
}

type VGACable struct {
	Voltage int
}
