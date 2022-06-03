package adapter

import (
	"github.com/taraslis453/design-patterns/patterns/structural/adapter/laptop"
	"github.com/taraslis453/design-patterns/patterns/structural/adapter/monitor"
)

// Adapter is structural design pattern that lets you connect two objects that have different interfaces.
func Adapter() {
	macbook := laptop.NewMacbook("m1", laptop.LightningCable{
		Voltage: 20,
	})
	dell := monitor.NewDell("Dell", monitor.VGACable{Voltage: 5})
	// without adapter
	// macbook.InsertIntoLightningPort(dell.Cable)

	monitorAdapter := monitor.NewMonitorAdapter(dell)
	monitorAdapter.InsertIntoLightningPort(macbook.Cable)
}
