package main

import (
	rpi "github.com/nathan-osman/go-rpigpio"
)

const (
	Relay1 = 22
	Relay2 = 23
	Relay3 = 27
	Relay4 = 18
	Relay5 = 24
	Relay6 = 17
)

func main() {

	relay1, _ := rpi.OpenPin(Relay1, rpi.OUT)
	relay2, _ := rpi.OpenPin(Relay2, rpi.OUT)
	relay3, _ := rpi.OpenPin(Relay3, rpi.OUT)
	relay4, _ := rpi.OpenPin(Relay4, rpi.OUT)
	relay5, _ := rpi.OpenPin(Relay5, rpi.OUT)
	relay6, _ := rpi.OpenPin(Relay6, rpi.OUT)
	/*if err != nil {
		panic(err)
	}*/
	defer relay1.Close()
	defer relay2.Close()
	defer relay3.Close()
	defer relay4.Close()
	defer relay5.Close()
	defer relay6.Close()

	for i := 0; i < 5; i++ {
		relay1.Write(rpi.HIGH)
		relay2.Write(rpi.HIGH)
		relay3.Write(rpi.HIGH)
		relay4.Write(rpi.HIGH)
		relay5.Write(rpi.HIGH)
		relay6.Write(rpi.HIGH)

		/*
			time.Sleep(1200 * time.Millisecond)
			p17.Write(rpi.LOW)
			p27.Write(rpi.LOW)
			p22.Write(rpi.LOW)
			p18.Write(rpi.LOW)
			p23.Write(rpi.LOW)
			p24.Write(rpi.LOW)
			time.Sleep(500 * time.Millisecond)
		*/
	}
}
