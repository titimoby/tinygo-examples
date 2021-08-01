package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var (
	black = color.RGBA{R: 0, G: 0, B: 0}
	red   = color.RGBA{R: 255, G: 0, B: 0}
	green = color.RGBA{R: 0, G: 255, B: 0}
	blue  = color.RGBA{R: 0, G: 0, B: 255}
)

func main() {
	// color intensity does not seems to work but you got the idea ;)
	reds := make([]color.RGBA, 5)
	reds[0] = color.RGBA{R: 50, G: 0, B: 0}
	reds[1] = color.RGBA{R: 100, G: 0, B: 0}
	reds[2] = color.RGBA{R: 150, G: 0, B: 0}
	reds[3] = color.RGBA{R: 200, G: 0, B: 0}
	reds[4] = color.RGBA{R: 255, G: 0, B: 0}

	colors := make([]color.RGBA, 5)

	neopixels := ws2812.New(machine.NEOPIXELS)
	neopixels.Pin.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		for i := range [5]int{0, 1, 2, 3, 4} {
			clear(colors)
			colors[i] = reds[i]
			neopixels.WriteColors(colors)
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func clear(colors []color.RGBA) {
	colors[0] = black
	colors[1] = black
	colors[2] = black
	colors[3] = black
	colors[4] = black
}
