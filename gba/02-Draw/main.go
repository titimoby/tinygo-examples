package main

import (
	"image/color"
	"machine"
	"runtime/volatile"
	"unsafe"

	"tinygo.org/x/tinydraw"
)

var (
	// Register display
	regDISPSTAT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000004)))

	// Display from machine
	display = machine.Display

	// Colors
	white = color.RGBA{255, 255, 255, 255}
	green = color.RGBA{0, 255, 0, 255}
	red   = color.RGBA{255, 0, 0, 255}
)

func main() {
	// Set up the display
	display.Configure()

	// Register display status
	regDISPSTAT.SetBits(1<<3 | 1<<4)

	white = color.RGBA{255, 255, 255, 255}
	green = color.RGBA{0, 255, 0, 255}
	red = color.RGBA{255, 0, 0, 255}

	// Draw a red line
	tinydraw.Line(&display, 100, 100, 40, 100, red)

	// Draw a white rectangle and inside a green filled rectangle
	tinydraw.Rectangle(&display, 30, 106, 120, 20, white)
	tinydraw.FilledRectangle(&display, 34, 110, 112, 12, green)

	// Draw a white circle and inside a red filled circle
	tinydraw.Circle(&display, 120, 30, 20, white)
	tinydraw.FilledCircle(&display, 120, 30, 16, red)

	// Draw a white triangle and inside a green filled triangle
	tinydraw.Triangle(&display, 120, 102, 100, 80, 152, 46, white)
	tinydraw.FilledTriangle(&display, 120, 98, 104, 80, 144, 54, green)

	// Infinite loop to avoid exiting the application
	for {
	}
}
