package main

import (
	"machine"

	"image/color"
	"time"

	"github.com/hybridgroup/gopherbot"
)

var (
	red = color.RGBA{R: 0xff, G: 0x00, B: 0x00}
	green = color.RGBA{R: 0x00, G: 0xff, B: 0x00}
	blue = color.RGBA{R: 0x00, G: 0x00, B: 0xff}
	black = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
)

var (
	uart  = machine.Serial
	tx    = machine.UART_TX_PIN
	rx    = machine.UART_RX_PIN
	input = make([]byte, 0, 64)
	mode = "off"

	backpack = gopherbot.Backpack()
)

func main() {
	uart.Configure(machine.UARTConfig{TX: tx, RX: rx})

	go lights()

	for {
		if uart.Buffered() > 0 {
			data, _ := uart.ReadByte()

			switch data {
			case 13:
				// return key
				mode = string(input)
				input = input[:0]
			default:
				// just capture the character
				input = append(input, data)
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func lights() {
	for {
		switch mode {
		case "green":
			backpack.Green()
		case "red":
			backpack.Red()
		case "talk":
			backpack.Alternate(green, black)
		case "stop":
			backpack.Off()
		default:
			backpack.Off()
		}

		time.Sleep(200 * time.Millisecond)
	}
}
