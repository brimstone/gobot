package gopigo3

// Lots stolen from https://github.com/DexterInd/GoPiGo3/blob/master/Software/Python/gopigo3.py

import (
	"fmt"
	"log"

	"gobot.io/x/gobot/drivers/spi"
	expspi "golang.org/x/exp/io/spi"
)

const (
	// NAME
	NAME spi.SPI_MESSAGE_TYPE = iota
	GET_MANUFACTURER
	GET_NAME
	GET_HARDWARE_VERSION
	// GET_FIRMWARE_VERSION is a 32bit int
	GET_FIRMWARE_VERSION
	GET_ID
	SET_LED

	GET_VOLTAGE_5V
	GET_VOLTAGE_VCC

	SET_SERVO

	SET_MOTOR_PWM
)

const (
	BUS_ADDRESS byte = 8
	MOTOR_LEFT  byte = 1
	MOTOR_RIGHT byte = 2
)

type GoPiGo3 struct {
	name       string
	Connection string
	Device     string
	spi        *spi.SpiDriver
}

func NewGoPiGo3Driver() *GoPiGo3 {
	log.Println("Got this far")
	return &GoPiGo3{
		name: "GoPiGo3",
		spi: spi.NewSpiDriver(&expspi.Devfs{
			Dev:      "/dev/spidev0.1",
			Mode:     expspi.Mode0,
			MaxSpeed: 500000,
		}),
	}
}

func (me *GoPiGo3) GetFirmwareVersion() string {
	response := me.spi.Read32(BUS_ADDRESS, GET_FIRMWARE_VERSION)
	major := response / 1000000
	minor := response / 1000 % 1000
	patch := response % 1000
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

func (me *GoPiGo3) GoForward(speed int) {
	if speed > 127 {
		speed = 127
	}
	if speed < -128 {
		speed = -128
	}
	me.spi.TransmitArray([]byte{
		BUS_ADDRESS,
		SET_MOTOR_PWM,
		MOTOR_LEFT | MOTOR_RIGHT,
		speed,
	})
}
