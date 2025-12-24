package main

import (
	"machine"
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
		println(value)
		time.Sleep(1 * time.Second)
	}
}
