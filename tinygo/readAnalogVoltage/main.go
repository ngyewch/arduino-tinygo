package main

import (
	"machine"
	"math"
	"time"
)

func main() {
	machine.InitADC()
	adc := machine.ADC{Pin: machine.ADC0}
	adc.Configure(machine.ADCConfig{})

	machine.Serial.Configure(machine.UARTConfig{
		BaudRate: 9600,
	})

	for {
		value := adc.Get()
		voltage := float32(value) * 5 / float32(math.MaxUint16)
		println(voltage)
		time.Sleep(1 * time.Second)
	}
}
