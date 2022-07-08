package main

import (
	"machine"
	"time"
)

func main() {
	ledCol1 := machine.LED_COL_1
	ledCol2 := machine.LED_COL_2
	ledRow1 := machine.LED_ROW_1
	ledCol1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledCol2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledRow1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledRow1.High()
	go func() {
		for {
			ledCol1.Low()
			time.Sleep(time.Millisecond * 500)

			ledCol1.High()
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			ledCol2.Low()
			time.Sleep(time.Millisecond * 100)

			ledCol2.High()
			time.Sleep(time.Millisecond * 100)
		}
	}()

	select {}
}
