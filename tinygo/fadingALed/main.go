package main

import (
	"machine"
	"time"
)

func main() {
	pwm := machine.Timer1
	err := pwm.Configure(machine.PWMConfig{})
	if err != nil {
		println("Failed to configure PWM:", err.Error())
		return
	}

	pin := machine.D9
	channel, err := pwm.Channel(pin)
	if err != nil {
		println("Failed to get PWM channel:", err.Error())
		return
	}

	top := pwm.Top()
	inc := top / 50
	for {
		// Fade in
		for i := uint32(0); i < top; i += inc {
			pwm.Set(channel, i)
			time.Sleep(time.Millisecond * 5)
		}
		// Fade out
		for i := top; i > 0; i -= inc {
			pwm.Set(channel, i)
			time.Sleep(time.Millisecond * 5)
		}
	}
}
