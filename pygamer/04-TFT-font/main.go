package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/st7735"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/tinyfont/freesans"
)

var black = color.RGBA{0, 0, 0, 0}
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

	for {
		display.FillScreen(black)
		tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 10, 10, "No rotation", red)
		time.Sleep(1 * time.Second)

		display.FillScreen(black)
		tinyfont.WriteLineRotated(&display, &freemono.Bold9pt7b, 10, 10, "COUCOU", red, tinyfont.ROTATION_90)
		time.Sleep(1 * time.Second)

		display.FillScreen(black)
		tinyfont.WriteLineRotated(&display, &freemono.Bold9pt7b, 100, 50, "Coucou", green, tinyfont.ROTATION_90)
		time.Sleep(1 * time.Second)
	}
}
