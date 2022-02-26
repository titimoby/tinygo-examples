package main

import (
	"image/color"
	"time"

	"tinygo.org/x/drivers/microbitmatrix"
)

var display microbitmatrix.Device

func main() {
	display = microbitmatrix.New()
	display.Configure(microbitmatrix.Config{})
	c := color.RGBA{255, 255, 255, 255}

	for {
		display.Display()
		time.Sleep(time.Millisecond * 500)
		display.SetPixel(0, 0, c)
		time.Sleep(time.Millisecond * 500)
	}
}
