package main

import (
	"fmt"

	"time"

	rpio "github.com/stianeikeland/go-rpio"
	//	"github.com/kidoman/embd"
	//	_ "github.com/kidoman/embd/host/all"
)

func main() {
	str := "Hello"
	fmt.Printf("%s\n", str)
	_ = rpio.Open()
	pin11 := rpio.Pin(11)

	pin11.Output() // Output mode
	pin11.High()   // Set pin High
	pin11.Low()    // Set pin Low
	pin11.Toggle() // Toggle pin (Low -> High -> Low)

	time.Sleep(2 * time.Second)

	pin18 := rpio.Pin(18)

	pin18.Output() // Output mode
	pin18.High()   // Set pin High
	pin18.Low()    // Set pin Low
	pin18.Toggle() // Toggle pin (Low -> High -> Low)

	/*
		embd.InitGPIO()
		defer embd.CloseGPIO()
		embd.SetDirection(11, embd.Out)
		embd.DigitalWrite(11, embd.High)
		embd.SetDirection(13, embd.Out)
		embd.DigitalWrite(13, embd.High)
		embd.SetDirection(15, embd.Out)
		embd.DigitalWrite(15, embd.High)
		time.Sleep(3 * time.Second)

		embd.SetDirection(11, embd.Out)
		embd.DigitalWrite(11, embd.Low)
		embd.SetDirection(13, embd.Out)
		embd.DigitalWrite(13, embd.Low)
		embd.SetDirection(15, embd.Out)
		embd.DigitalWrite(15, embd.Low)
	*/
}
