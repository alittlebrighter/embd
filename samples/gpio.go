package main

import (
	"time"

	"github.com/kidoman/embd"
)

func main() {
	gpio, err := embd.NewGPIO()
	if err != nil {
		panic(err)
	}
	defer gpio.Close()

	led, err := gpio.DigitalPin(10)
	if err != nil {
		panic(err)
	}

	if err := led.SetDir(embd.Out); err != nil {
		panic(err)
	}
	if err := led.Write(embd.High); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	if err := led.SetDir(embd.In); err != nil {
		panic(err)
	}
}