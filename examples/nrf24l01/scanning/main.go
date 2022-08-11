package main

import (
	"machine"

	nrf24l01 "tinygo.org/x/drivers/nrf24l01"
)

const (
	BUFF_LENGTH = 4
)

var (
	err error
)

func main() {
	println("----------------RX MODE scanning----------------\n")
	spi := machine.SPI0
	err = spi.Configure(machine.SPIConfig{})
	if err != nil {
		println("failed with spi.Configure(machine.SPIConfig{}):", err)
	}

	ce := machine.D9   // Digital Input	Chip Enable Activates RX or TX mode
	csn := machine.D10 // Digital Input	SPI Chip Select

	nrf := nrf24l01.New(&spi, &ce, &csn)
	err = nrf.Configure()
	if err != nil {
		println("failed with nrf.Configure():", err)
	}

	err = nrf.SetRXMode()
	if err != nil {
		println("failed with nrf.SetRXMode():", err)
	}

	err = nrf.SetRF1MBPS()
	// err = nrf.SetRF2MBPS()
	// err = nrf.SetRF250KBPS()
	if err != nil {
		println("failed with nrf.SetRF1MBPS():", err)
	}

	res, err := nrf.GetRegisterState(nrf24l01.CONFIG)
	println("CONFIG:     ", res)
	res, err = nrf.GetRegisterState(nrf24l01.SETUP_AW)
	println("SETUP_AW:   ", res)
	res, err = nrf.GetRegisterState(nrf24l01.RF_SETUP)
	println("RF_SETUP:   ", res)
	status, err := nrf.GetRegisterState(nrf24l01.STATUS)
	println("STATUS:     ", status)
	fifoStatus, err := nrf.GetRegisterState(nrf24l01.FIFO_STATUS)
	println("FIFO_STATUS:", fifoStatus)
	res, err = nrf.GetRegisterState(nrf24l01.FEATURE)
	println("FEATURE:    ", res)
	println("--------------------------------------------------")

	err = nrf.HearChannels()
	if err != nil { // Don't do that. Handle all errors separately )
		println(err)
	}

}
