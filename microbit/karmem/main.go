package main

import (
	"fmt"
	"log"
	"machine"
	"sync"
	"time"

	karmem "karmem.org/golang"
)

var writerPool = sync.Pool{New: func() any { return karmem.NewWriter(1024) }}

func main() {
	ledCol1 := machine.LED_COL_1
	ledCol2 := machine.LED_COL_2
	ledRow1 := machine.LED_ROW_1
	ledCol1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledCol2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledRow1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledRow1.High()

	var writer = writerPool.Get().(*karmem.Writer)

	p := Point{"hi there", 23.5}

	if _, err := p.WriteAsRoot(writer); err != nil {
		panic(err)
	}

	encoded := writer.Bytes()

	fmt.Println("bytes: ", encoded)

	go func() {
		for {
			ledCol1.Low()
			time.Sleep(time.Millisecond * time.Duration(p.Value))

			ledCol1.High()
			time.Sleep(time.Millisecond * 500)
		}
		log.Println("blink")
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
