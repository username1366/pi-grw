package main

import (
	"fmt"
	"os"
	"time"

	rpi "github.com/nathan-osman/go-rpigpio"
)

const (
	Relay1 = 23 // blue
	Relay2 = 18
	Relay3 = 27
	Relay4 = 24
	Relay5 = 22
	Relay6 = 17
	Delay  = 65
)

func demo() {
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

	for i := 0; i < 3; i++ {
		relay1.Write(rpi.HIGH)
		time.Sleep(Delay * time.Millisecond)
		relay2.Write(rpi.HIGH)
		time.Sleep(Delay * time.Millisecond)
		relay3.Write(rpi.HIGH)
		time.Sleep(Delay * time.Millisecond)
		relay4.Write(rpi.HIGH)
		time.Sleep(Delay * time.Millisecond)
		relay5.Write(rpi.HIGH)
		time.Sleep(Delay * time.Millisecond)
		relay6.Write(rpi.HIGH)

		time.Sleep(1200 * time.Millisecond)

		relay1.Write(rpi.LOW)
		time.Sleep(Delay * time.Millisecond)
		relay2.Write(rpi.LOW)
		time.Sleep(Delay * time.Millisecond)
		relay3.Write(rpi.LOW)
		time.Sleep(Delay * time.Millisecond)
		relay4.Write(rpi.LOW)
		time.Sleep(Delay * time.Millisecond)
		relay5.Write(rpi.LOW)
		time.Sleep(Delay * time.Millisecond)
		relay6.Write(rpi.LOW)
		time.Sleep(500 * time.Millisecond)

		fmt.Printf("%d\n", i)
	}

}

func main() {
	for _, v := range os.Args {
		if v == "demo" {
			demo()
		}
	}

	relay1, _ := rpi.OpenPin(Relay1, rpi.OUT)
	relay2, _ := rpi.OpenPin(Relay2, rpi.OUT)
	relay3, _ := rpi.OpenPin(Relay3, rpi.OUT)
	relay4, _ := rpi.OpenPin(Relay4, rpi.OUT)
	relay5, _ := rpi.OpenPin(Relay5, rpi.OUT)
	relay6, _ := rpi.OpenPin(Relay6, rpi.OUT)

	defer relay1.Close()
	defer relay2.Close()
	defer relay3.Close()
	defer relay4.Close()
	defer relay5.Close()
	defer relay6.Close()

	fmt.Println("Light off\n")
	relay1.Write(rpi.HIGH)
	relay2.Write(rpi.HIGH)
	relay3.Write(rpi.HIGH)
	relay4.Write(rpi.HIGH)
	relay5.Write(rpi.HIGH)
	relay6.Write(rpi.HIGH)

	time.Sleep(60 * time.Second)

	fmt.Println("Light on\n")
	relay1.Write(rpi.LOW)
	relay2.Write(rpi.LOW)
	relay3.Write(rpi.LOW)
	relay4.Write(rpi.LOW)
	relay5.Write(rpi.LOW)
	relay6.Write(rpi.LOW)

	time.Sleep(18 * time.Hour)

	fmt.Println("Light off\n")
	relay1.Write(rpi.HIGH)
	relay2.Write(rpi.HIGH)
	relay3.Write(rpi.HIGH)
	relay4.Write(rpi.HIGH)
	relay5.Write(rpi.HIGH)
	relay6.Write(rpi.HIGH)
}
