
package main

import (
"time"

"gobot.io/x/gobot"
"gobot.io/x/gobot/drivers/gpio"
"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()
	red    := gpio.NewLedDriver(r, "12")
	yellow := gpio.NewLedDriver(r, "16")
	green  := gpio.NewLedDriver(r, "18")

	work := func() {
		gobot.Every(1*time.Second, func() {
			green.Toggle()
		})
		gobot.Every(2*time.Second, func() {
			yellow.Toggle()
		})
		gobot.Every(3*time.Second, func() {
			red.Toggle()
		})
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{r},
		[]gobot.Device{green},
		work,
	)

	robot.Start()
}