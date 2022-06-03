package monitor

import (
	"fmt"

	"github.com/taraslis453/design-patterns/patterns/structural/adapter/laptop"
)

type monitorAdapter struct {
	monitor Monitor
}

func NewMonitorAdapter(monitor Monitor) *monitorAdapter {
	return &monitorAdapter{monitor: monitor}
}

func (m *monitorAdapter) InsertIntoLightningPort(lightning laptop.LightningCable) VGACable {
	return convertVGAToLightning(lightning)
}

func convertVGAToLightning(lightning laptop.LightningCable) VGACable {
	fmt.Println("Converting Lightning VGA to lightning")
	if lightning.Voltage == 5 {
		return VGACable{Voltage: 2}
	}

	return VGACable{Voltage: 3}
}
