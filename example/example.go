package example

import (
	"fmt"
	"log"

	"github.com/a-h/debounce"
	"github.com/stianeikeland/go-rpio/v4"
)

func main() {
	err := rpio.Open()
	if err != nil {
		log.Fatalln("Could not open GPIOs")
	}
	onClick := func() {
		fmt.Print("Clicked...")
	}
	normallyClosed := true
	sw := debounce.Button(onClick, normallyClosed)
	// Set your GPIO here
	gpio := rpio.Pin(0)
	rpio.PinMode(gpio, rpio.Input)
	for {
		sw.SetState(rpio.ReadPin(gpio) == rpio.High)
	}
}
