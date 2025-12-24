package main

import (
	"machine"
	"time"
)

func main() {
	button := machine.D2
	button.Configure(machine.PinConfig{Mode: machine.PinInput})

	machine.Serial.Configure(machine.UARTConfig{
		BaudRate: 9600,
	})

	for {
		value := button.Get()
		println(value)
		time.Sleep(1 * time.Millisecond)
	}
}
