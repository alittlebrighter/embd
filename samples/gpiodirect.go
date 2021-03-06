// +build ignore

package main

import (
	"flag"
	"time"

	"github.com/alittlebrighter/embd"

	_ "github.com/alittlebrighter/embd/host/all"
)

func main() {
	flag.Parse()

	if err := embd.InitGPIO(); err != nil {
		panic(err)
	}
	defer embd.CloseGPIO()

	if err := embd.SetDirection(10, embd.Out); err != nil {
		panic(err)
	}
	if err := embd.DigitalWrite(10, embd.High); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	if err := embd.SetDirection(10, embd.In); err != nil {
		panic(err)
	}
}
