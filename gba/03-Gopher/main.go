package main

import (
	"machine"

	"image/color"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/gophers"
	"tinygo.org/x/tinygba"
)

var (
	// Display from machine
	display = machine.Display

	// Screen resolution
	screenWidth  int16 = 240
	screenHeight int16 = 160

	// Colors
	black = color.RGBA{}
	white = color.RGBA{255, 255, 255, 255}
	green = color.RGBA{0, 255, 0, 255}
	red   = color.RGBA{255, 0, 0, 255}

	// Google colors
	gBlue   = color.RGBA{66, 163, 244, 255}
	gRed    = color.RGBA{219, 68, 55, 255}
	gYellow = color.RGBA{244, 160, 0, 255}
	gGreen  = color.RGBA{15, 157, 88, 255}

	// Coordinates
	x int16 = 100 //horizontally center
	y int16 = 100 //vertically center

	// Borders
	border int16 = 16

	gopherColor       = green
	gopherSize  int16 = 30

	// Game status
	gameStarted = false
)

func main() {
	// Set up the display
	display.Configure()

	// Display Gopher text message and draw our Gophers
	drawGophers()

	// Infinite loop to avoid exiting the application
	for {
		tinygba.WaitForVBlank()

		update()
	}
}

func drawGophers() {
	// Display a textual message "Gopher" with Google colors
	tinyfont.DrawChar(&display, &freesans.Bold24pt7b, 36, 60, 'G', gBlue)
	tinyfont.DrawChar(&display, &freesans.Bold24pt7b, 71, 60, 'o', gRed)
	tinyfont.DrawChar(&display, &freesans.Bold24pt7b, 98, 60, 'p', gYellow)
	tinyfont.DrawChar(&display, &freesans.Bold24pt7b, 126, 60, 'h', gGreen)
	tinyfont.DrawChar(&display, &freesans.Bold24pt7b, 154, 60, 'e', gBlue)
	tinyfont.DrawChar(&display, &freesans.Bold24pt7b, 180, 60, 'r', gRed)

	// Display a "press START button" message - center
	tinyfont.WriteLine(&display, &tinyfont.TomThumb, 85, 90, "Press START button", white)

	// Display two gophers
	tinyfont.DrawChar(&display, &gophers.Regular58pt, 5, 140, 'B', green)
	tinyfont.DrawChar(&display, &gophers.Regular58pt, 195, 140, 'X', red)
}

func update() {
	key := tinygba.ReadButtons()

	switch {
	case tinygba.ButtonStart.IsPushed(key):
		gameStarted = true
		gopherColor = green

		clearScreen()

		// Display Gopher at the "center" of the screen
		x = 100
		y = 100
		tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', gopherColor)
	case tinygba.ButtonSelect.IsPushed(key):
		gameStarted = false

		clearScreen()
		drawGophers()

	case tinygba.ButtonRight.IsPushed(key):

		if gameStarted {
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', black)

			// if border right is reached, don't move to the right
			if x+10+gopherSize < screenWidth-border {
				x = x + 10
			}

			// display gopher at right
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', gopherColor)
		}

	case tinygba.ButtonLeft.IsPushed(key):
		if gameStarted {
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', black)

			// if border left is reached, don't move to the left
			if border < x-10 {
				x = x - 10
			}

			// display gopher at right
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', gopherColor)
		}

	case tinygba.ButtonDown.IsPushed(key):
		if gameStarted {
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', black)

			// if border down is reached, don't move to the bottom
			if y+10 < screenHeight-border {
				y = y + 10
			}

			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', gopherColor)
		}

	case tinygba.ButtonUp.IsPushed(key):
		if gameStarted {
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', black)

			// if border up is reached, don't move to the top
			if border < y-10-gopherSize {
				y = y - 10
			}

			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', gopherColor)
		}

	case tinygba.ButtonA.IsPushed(key):
		if gameStarted {
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', black)

			if border < y-20-gopherSize {
				// Display the gopher up
				y = y - 20
				tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', gopherColor)

				tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', black)
				// Display the gopher down
				y = y + 20
				tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', gopherColor)
			}
		}

	case tinygba.ButtonB.IsPushed(key):
		if gameStarted {
			gopherColor = green
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', gopherColor)
		}

	case tinygba.ButtonL.IsPushed(key):
		if gameStarted {
			gopherColor = gYellow
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', gopherColor)
		}

	case tinygba.ButtonR.IsPushed(key):
		if gameStarted {
			gopherColor = gBlue
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', gopherColor)
		}
	}
}

func clearScreen() {
	tinygba.FillScreen(black)
}
