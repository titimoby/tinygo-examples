package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ili9341"
)

var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}

func main() {
	display := ili9341.NewParallel(
		machine.LCD_DATA0,
		machine.TFT_WR,
		machine.TFT_DC,
		machine.TFT_CS,
		machine.TFT_RESET,
		machine.TFT_RD,
	)
	display.Configure(ili9341.Config{Width: 320, Height: 480})
	display.SetRotation(ili9341.Rotation270)

	backlight := machine.TFT_BACKLIGHT
	backlight.Configure(machine.PinConfig{machine.PinOutput})
	backlight.High()

	i := 0
	for {
		display.FillScreen(green)
		time.Sleep(1 * time.Second)
		display.FillScreen(red)
		time.Sleep(1 * time.Second)
		i++
	}
}
