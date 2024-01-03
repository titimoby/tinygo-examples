package main

import (
	"image/color"
	"machine"
	"runtime/interrupt"
	"runtime/volatile"
	"unsafe"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/gophers"
	"tinygo.org/x/tinygba"
)

var (
	// Register display
	regDISPSTAT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000004)))

	// Display from machine
	display = machine.Display

	// Screen resolution
	// screenWidth, screenHeight = display.Size()
	screenWidth  int16 = 240
	screenHeight int16 = 160

	// Colors
	black = color.RGBA{}
	white = color.RGBA{255, 255, 255, 255}
	red   = color.RGBA{255, 0, 0, 255}

	// Coordinates
	x int16 = screenWidth / 2
	y int16 = screenHeight / 2

	// Border
	border int16 = 16
)

func clearScreen() {
	tinygba.FillScreen(black)
}

func initScreen() {
	tinydraw.FilledRectangle(&display, 0, 0, screenWidth, screenHeight, red)

	tinydraw.FilledRectangle(&display, border, border, screenWidth-(2*border), screenHeight-(2*border), white)

	tinyfont.DrawChar(&display, &gophers.Regular18pt, 2, 80, 'B', white)

	tinyfont.WriteLine(&display, &tinyfont.TomThumb, (screenWidth/2)-15, 11, "TELECRAN ", white)

	tinyfont.DrawChar(&display, &gophers.Regular18pt, screenWidth-13, 80, 'B', white)

	// tinyfont.DrawChar(&display, &gophers.Regular18pt, (screenWidth / 2), 11, 'B', white)

	// Draw left and right circles at the bottom of the red screen
	tinydraw.FilledCircle(&display, 10, 150, 7, white)
	tinydraw.FilledCircle(&display, 230, 150, 7, white)
}

// Clear the draw
func clearGame() {

	x = screenWidth / 2
	y = screenHeight / 2
	tinydraw.FilledRectangle(&display, border, border, screenWidth-(2*border), screenHeight-(2*border), white)
}

func update(interrupt.Interrupt) {

	key := tinygba.ReadButtons()

	switch {
	case tinygba.ButtonStart.IsPushed(key):
		clearGame()

	case tinygba.ButtonLeft.IsPushed(key):

		if (border < x-1) && (x-1 < screenWidth-border) &&
			(border < y) && (y < screenHeight-border) {

			x = x - 1
			tinyfont.DrawChar(&display, &tinyfont.TomThumb, x, y, '.', black)
		}

	case tinygba.ButtonRight.IsPushed(key):

		if (border < x+1) && (x+1 < screenWidth-border) &&
			(border < y) && (y < screenHeight-border) {

			x = x + 1
			tinyfont.DrawChar(&display, &tinyfont.TomThumb, x, y, '.', black)
		}

	case tinygba.ButtonUp.IsPushed(key):

		if (border < x) && (x < screenWidth-border) &&
			(border < y-1) && (y-1 < screenHeight-border) {

			y = y - 1
			tinyfont.DrawChar(&display, &tinyfont.TomThumb, x, y, '.', black)
		}

	case tinygba.ButtonDown.IsPushed(key):

		if (border < x) && (x < screenWidth-border) &&
			(border < y+1) && (y+1 < screenHeight-border) {

			y = y + 1
			tinyfont.DrawChar(&display, &tinyfont.TomThumb, x, y, '.', black)
		}
	}
}

func main() {
	// Set up the display
	display.Configure()

	// Register display status
	regDISPSTAT.SetBits(1<<3 | 1<<4)

	// Display text message and draw our Gophers
	initScreen()

	// Creates an interrupt that will call the "update" fonction below, hardware way to display things on the screen
	interrupt.New(machine.IRQ_VBLANK, update).Enable()

	// Infinite loop to avoid exiting the application
	for {
	}
}
