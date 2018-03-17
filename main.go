package main

import (
	"fmt"

	"time"

	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/all"
)

func main() {
	str := "Hello"
	fmt.Printf("%s\n", str)

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

}
