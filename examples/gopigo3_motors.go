// +build example
//
// Do not build by default.

package main

import (
	"fmt"
	"log"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi/gopigo3"
)

func main() {
	gpg := gopigo3.NewGoPiGo3Driver()

	fmt.Println(gpg.GetFirmwareVersion())

	work := func() {
		log.Println("Doing work")
		gpg.GoForward(100)
	}

	robot := gobot.NewRobot("blinkBot",
		[]gobot.Connection{gpg.Connection},
		[]gobot.Device{gpg.Device},
		work,
	)

	robot.Start()
}
