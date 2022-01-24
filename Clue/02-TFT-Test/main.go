package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/st7789"
)

var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}

func main() {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 16000000,
		SCK:       machine.SPI0_SCK_PIN,
		SDO:       machine.SPI0_SDO_PIN,
		SDI:       machine.SPI0_SDI_PIN,
	})
	display := st7789.New(machine.SPI0, machine.TFT_RESET, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	display.Configure(st7789.Config{})

	i := 0
	for {
		display.FillScreen(green)
		time.Sleep(1 * time.Second)
		display.FillScreen(red)
		time.Sleep(1 * time.Second)
		i++
	}
}
