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
		for {
			green.On()
			time.Sleep(6 * time.Second)
			green.Off()
			yellow.On()
			time.Sleep(3 * time.Second)
			yellow.Off()
			red.On()
			time.Sleep(6 * time.Second)
			red.Off()
			yellow.On()
			time.Sleep(500 * time.Millisecond)
			yellow.Off()
		}
	}

	robot := gobot.NewRobot("trafficBot",
		[]gobot.Connection{r},
		[]gobot.Device{green},
		[]gobot.Device{yellow},
		[]gobot.Device{red},
		work,
	)

	robot.Start()
}