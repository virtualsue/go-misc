package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"time"
)

var (
	red_pin    = rpio.Pin(18)   // physical pin 12
	yellow_pin = rpio.Pin(23)   // physical pin 16
	green_pin  = rpio.Pin(24)   // physical pin 18
)

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	red_pin.Output()
	yellow_pin.Output()
	green_pin.Output()

	// Toggle pin 20 times
	for x := 0; x < 20; x++ {
		fmt.Println(x)
		red_pin.Toggle()
		yellow_pin.Toggle()
		green_pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}
