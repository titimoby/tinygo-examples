package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var (
	red   = color.RGBA{R: 255, G: 0, B: 0}
	green = color.RGBA{R: 0, G: 255, B: 0}
	blue  = color.RGBA{R: 0, G: 0, B: 255}
)

func main() {
	color := make([]color.RGBA, 1)

	neopixels := ws2812.New(machine.NEOPIXEL)
	neopixels.Pin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		color[0] = red
		neopixels.WriteColors(color)
		time.Sleep(time.Millisecond * 500)

		color[0] = green
		neopixels.WriteColors(color)
		time.Sleep(time.Millisecond * 500)

		color[0] = blue
		neopixels.WriteColors(color)
		time.Sleep(time.Millisecond * 500)
	}
}
