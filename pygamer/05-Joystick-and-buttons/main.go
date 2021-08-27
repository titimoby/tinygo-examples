package main

import (
	"image/color"
	"machine"
	"strconv"

	"tinygo.org/x/drivers/st7735"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

var black = color.RGBA{0, 0, 0, 0}
var red = color.RGBA{255, 0, 0, 255}
var gray = color.RGBA{200, 200, 200, 255}

func main() {
	buttons := NewButtons()
	buttons.Configure()

	joystick := NewJoystick()
	joystick.configure()

	display := st7735.New(machine.SPI1, machine.TFT_RST, machine.TFT_DC, machine.TFT_CS, machine.TFT_LITE)
	machine.SPI1.Configure(machine.SPIConfig{
		Frequency: 16000000,
		SCK:       machine.SPI1_SCK_PIN,
		SDO:       machine.SPI1_SDO_PIN,
		SDI:       machine.SPI1_SDI_PIN,
	})
	display.Configure(st7735.Config{Rotation: st7735.ROTATION_90})
	display.FillScreen(black)

	var buttons_pressed = ""
	var x_axis uint16 = 0
	var y_axis uint16 = 0
	for {
		x_axis = joystick.getX()
		y_axis = joystick.getY()
		buttons_pressed = ""
		pressed, _ := buttons.Read8Input()
		if pressed != 0 {
			println(pressed)
		}
		if pressed&machine.BUTTON_A_MASK > 0 {
			buttons_pressed += "BTN A"
			println("BTN A")
		}
		if pressed&machine.BUTTON_B_MASK > 0 {
			buttons_pressed += "BTN B"
			println("BTN B")
		}
		if pressed&machine.BUTTON_SELECT_MASK > 0 {
			buttons_pressed += "SELECT"
			println("SELECT")
		}
		if pressed&machine.BUTTON_START_MASK > 0 {
			buttons_pressed += "START"
			println("START")
		}

		display.FillScreen(black)
		tinyfont.WriteLine(&display, &freemono.Bold9pt7b, 10, 20, strconv.Itoa(int(x_axis)), gray)
		tinyfont.WriteLine(&display, &freemono.Bold9pt7b, 10, 50, strconv.Itoa(int(y_axis)), gray)
		tinyfont.WriteLine(&display, &freemono.Bold9pt7b, 10, 80, buttons_pressed, gray)
	}
}
