package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/st7735"
)

var red = color.RGBA{255, 0, 0, 255}
var green = color.RGBA{0, 255, 0, 255}

func main() {
	display := st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	machine.SPI1.Configure(machine.SPIConfig{
		Frequency: 16000000,
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
	})
	display.Configure(st7735.Config{})

	i := 0
	for {
		display.FillScreen(green)
		time.Sleep(1 * time.Second)
		display.FillScreen(red)
		time.Sleep(1 * time.Second)
		i++
	}
}
