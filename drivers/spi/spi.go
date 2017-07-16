package spi

import (
	"gobot.io/x/gobot/drivers/spi"
	expspi "golang.org/x/exp/io/spi"
)

type SPI_MESSAGE_TYPE byte

type SpiDriver struct {
	dev *expspi.Device
}

func NewSpiDriver(devfs *expspi.Devfs) *SpiDriver {
	dev, err := expspi.Open(devfs)
	if err != nil {
		panic(err)
	}
	return &SpiDriver{
		dev: dev,
	}
}

func (me *SpiDriver) Read32(address byte, msg spi.SPI_MESSAGE_TYPE) int32 {
	// read build two arrays 8 bytes wide, 4 for addressing, and 4 for the response
	w := make([]byte, 8)
	w[0] = address
	w[1] = msg
	r := make([]byte, len(w))
	err := me.dev.Tx(w, r)
	if err != nil {
		panic(err)
	}
	return int(4)<<24 | int(foo[5])<<16 | int(foo[6])<<8 | int(foo[7])
}

func (me *SpiDriver) TransmitArray(w []byte) {
	err := me.dev.Tx(w, nil)
	if err != nil {
		panic(err)
	}
}
