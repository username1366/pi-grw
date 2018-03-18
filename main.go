package main

import (
	"time"

	rpi "github.com/nathan-osman/go-rpigpio"
)

func main() {
	p17, _ := rpi.OpenPin(17, rpi.OUT)
	p27, _ := rpi.OpenPin(27, rpi.OUT)
	p22, _ := rpi.OpenPin(22, rpi.OUT)
	p18, _ := rpi.OpenPin(18, rpi.OUT)
	p23, _ := rpi.OpenPin(23, rpi.OUT)
	p24, _ := rpi.OpenPin(24, rpi.OUT)
	/*if err != nil {
		panic(err)
	}*/
	defer p17.Close()
	defer p27.Close()
	defer p22.Close()
	defer p18.Close()
	defer p23.Close()
	defer p24.Close()

	for i := 0; i < 5; i++ {
		p17.Write(rpi.HIGH)
		p27.Write(rpi.HIGH)
		p22.Write(rpi.HIGH)
		p18.Write(rpi.HIGH)
		p23.Write(rpi.HIGH)
		p24.Write(rpi.HIGH)
		time.Sleep(1200 * time.Millisecond)
		p17.Write(rpi.LOW)
		p27.Write(rpi.LOW)
		p22.Write(rpi.LOW)
		p18.Write(rpi.LOW)
		p23.Write(rpi.LOW)
		p24.Write(rpi.LOW)
		time.Sleep(500 * time.Millisecond)
	}
}
