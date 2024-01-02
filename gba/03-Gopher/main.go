package main

import (
	"machine"
	"time"

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
	screenWidth, screenHeight = display.Size()

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
	// x int16 = 100 //TODO: horizontally center
	// y int16 = 100 //TODO: vertically center
	x int16 = screenWidth / 2
	y int16 = screenHeight / 2

	// Borders TODO: optimize left and right borders
	border int16 = 16
	// leftBorder  int16 = 30
	// rightBorder int16 = screenWidth - 45
	// upBorder    int16 = 60
	upBorder int16 = 25
	// downBorder int16 = screenHeight - 20

	// Game status
	gameStarted = false
)

func main() {
	// Set up the display
	display.Configure()

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
	//TODO: les faire bouger de gauche a droite
	// tinyfont.DrawChar(&display, &gophers.Regular58pt, 5, 140, 'B', green)
	// tinyfont.DrawChar(&display, &gophers.Regular58pt, 195, 140, 'X', red)

	for {
		tinyfont.DrawChar(&display, &gophers.Regular58pt, 5, 140, 'B', green)
		tinyfont.DrawChar(&display, &gophers.Regular58pt, 195, 140, 'X', red)

		time.Sleep(10 * time.Second)

		tinyfont.DrawChar(&display, &gophers.Regular58pt, 10, 140, 'B', green)
		tinyfont.DrawChar(&display, &gophers.Regular58pt, 200, 140, 'X', red)

		time.Sleep(10 * time.Second)

		tinyfont.DrawChar(&display, &gophers.Regular58pt, 5, 140, 'B', green)
		tinyfont.DrawChar(&display, &gophers.Regular58pt, 195, 140, 'X', red)
	}
}

//TODO: use gameStarted

func update() {
	// Display Gopher text message and draw our Gophers
	drawGophers()

	key := tinygba.ReadButtons()

	switch {
	case tinygba.ButtonStart.IsPushed(key):
		gameStarted = true

		clearScreen()

		// Display gopher
		tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', green)
	case tinygba.ButtonSelect.IsPushed(key):
		gameStarted = false

		clearScreen()
		drawGophers()

	case tinygba.ButtonRight.IsPushed(key):

		if gameStarted {
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', black)

			// if x <= rightBorder {
			if x <= screenWidth-((border*2)+10) {
				x = x + 10
			}

			// display gopher at right
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', green)
		}

	case tinygba.ButtonLeft.IsPushed(key):
		if gameStarted {
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', black)

			// if x >= leftBorder {
			if x >= border {
				x = x - 10
			}

			// display gopher at right
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', green)
		}

	case tinygba.ButtonDown.IsPushed(key):
		if gameStarted {
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', black)

			// if y <= downBorder {
			if y <= screenHeight-(border) {
				y = y + 10
			}

			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', green)
		}

	case tinygba.ButtonUp.IsPushed(key):
		if gameStarted {
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', black)

			// if y >= upBorder {
			if y >= ((border * 2) + upBorder) {
				y = y - 10
			}

			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', green)
		}

	case tinygba.ButtonA.IsPushed(key):
		if gameStarted {
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', black)

			// Display the gopher up
			y = y - 20
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', green)

			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', black)
			// Display the gopher down
			y = y + 20
			tinyfont.DrawChar(&display, &gophers.Regular58pt, x, y, 'B', green)
		}
	}
}

func clearScreen() {
	tinygba.FillScreen(black)
}
