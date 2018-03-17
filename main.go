package main

import (
	"time"

	rpi "github.com/nathan-osman/go-rpigpio"
)

func main() {
	p24, err := rpi.OpenPin(24, rpi.OUT)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	for i := 0; i < 10; i++ {
		p24.Write(rpi.HIGH)
		time.Sleep(300 * time.Millisecond)
		p24.Write(rpi.LOW)
		time.Sleep(100 * time.Millisecond)
	}
}
