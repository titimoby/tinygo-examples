package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/st7789"
)

var (
	display st7789.Device
	red     = color.RGBA{255, 0, 0, 255}
	green   = color.RGBA{0, 255, 0, 255}
)

func main() {
	machine.SPI1.Configure(machine.SPIConfig{
		Frequency: 8000000,
		SCK:       machine.TFT_SCK,
		SDO:       machine.TFT_SDO,
		SDI:       machine.TFT_SDO,
		Mode:      0,
	})
	display = st7789.New(machine.SPI1,
		machine.TFT_RESET,
		machine.TFT_DC,
		machine.TFT_CS,
		machine.TFT_LITE)
	display.Configure(st7789.Config{
		Rotation:   st7789.ROTATION_180,
		Height:     320,
		FrameRate:  st7789.FRAMERATE_111,
		VSyncLines: st7789.MAX_VSYNC_SCANLINES,
	})

	i := 0
	for {
		display.FillScreen(green)
		time.Sleep(1 * time.Second)
		display.FillScreen(red)
		time.Sleep(1 * time.Second)
		i++
	}
}
