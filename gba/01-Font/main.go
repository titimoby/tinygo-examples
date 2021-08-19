package main

import (
	"image/color"
	"machine"
	"runtime/volatile"
	"unsafe"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
	"tinygo.org/x/tinyfont/freesans"
)

var (
	// Register display
	regDISPSTAT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000004)))

	// Display from machine
	display = machine.Display

	// Colors
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
)

func main() {
	// Set up the display
	display.Configure()

	// Register display status
	regDISPSTAT.SetBits(1<<3 | 1<<4)

	tinyfont.DrawChar(display, &freemono.Bold24pt7b, 46, 60, 'H', blue)
	tinyfont.DrawChar(display, &freemono.Bold24pt7b, 75, 60, 'e', white)
	tinyfont.DrawChar(display, &freemono.Bold24pt7b, 103, 60, 'l', red)
	tinyfont.DrawChar(display, &freemono.Bold24pt7b, 130, 60, 'l', white)
	tinyfont.DrawChar(display, &freemono.Bold24pt7b, 157, 60, 'o', blue)

	tinyfont.WriteLine(display, &freesans.Regular9pt7b, 55, 90, "TinyGo lovers!", white)

	// Infinite loop to avoid exiting the application
	for {
	}
}
