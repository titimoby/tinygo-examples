package main

import (
	"machine"

	"image/color"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/tinyfont/freesans"
)

var (
	// Display from machine
	display = machine.Display

	// Colors
	// black = color.RGBA{}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
)

func main() {
	// Set up the display
	display.Configure()

	// Display text message
	drawText()

	// Infinite loop to avoid exiting the application
	for {
	}
}

func drawText() {
	tinyfont.DrawChar(&display, &freemono.Bold24pt7b, 46, 60, 'H', blue)
	tinyfont.DrawChar(&display, &freemono.Bold24pt7b, 75, 60, 'e', white)
	tinyfont.DrawChar(&display, &freemono.Bold24pt7b, 103, 60, 'l', red)
	tinyfont.DrawChar(&display, &freemono.Bold24pt7b, 130, 60, 'l', white)
	tinyfont.DrawChar(&display, &freemono.Bold24pt7b, 157, 60, 'o', blue)

	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 55, 90, "TinyGo lovers!", white)
}
