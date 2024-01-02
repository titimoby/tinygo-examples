package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinygba"
)

var (
	// Display from machine
	display = machine.Display

	// Colors
	white = color.RGBA{255, 255, 255, 255}
	black = color.RGBA{0, 0, 0, 0}
)

func initEyes() {

	//eyes centered/"standard"
	tinydraw.FilledCircle(&display, 80, 60, 16, white)
	tinydraw.FilledCircle(&display, 80, 60, 4, black)

	tinydraw.FilledCircle(&display, 150, 60, 16, white)
	tinydraw.FilledCircle(&display, 150, 60, 4, black)
}

func main() {
	// Set up the display
	display.Configure()

	// Draw eyes
	initEyes()

	// Infinite loop to avoid exiting the application
	for {
		tinygba.WaitForVBlank()

		update()
	}
}

func update() {
	key := tinygba.ReadButtons()

	switch {
	case tinygba.ButtonStart.IsPushed(key):

		initEyes()

	case tinygba.ButtonRight.IsPushed(key):

		//eyes to the right
		tinydraw.FilledCircle(&display, 80, 60, 16, white)
		tinydraw.FilledCircle(&display, 90, 60, 4, black)

		tinydraw.FilledCircle(&display, 150, 60, 16, white)
		tinydraw.FilledCircle(&display, 160, 60, 4, black)
	case tinygba.ButtonLeft.IsPushed(key):

		//eyes to the left
		tinydraw.FilledCircle(&display, 80, 60, 16, white)
		tinydraw.FilledCircle(&display, 70, 60, 4, black)

		tinydraw.FilledCircle(&display, 150, 60, 16, white)
		tinydraw.FilledCircle(&display, 140, 60, 4, black)
	case tinygba.ButtonUp.IsPushed(key):

		//eyes to the top
		tinydraw.FilledCircle(&display, 80, 60, 16, white)
		tinydraw.FilledCircle(&display, 80, 50, 4, black)

		tinydraw.FilledCircle(&display, 150, 60, 16, white)
		tinydraw.FilledCircle(&display, 150, 50, 4, black)

	case tinygba.ButtonDown.IsPushed(key):

		//eyes to the bottom
		tinydraw.FilledCircle(&display, 80, 60, 16, white)
		tinydraw.FilledCircle(&display, 80, 70, 4, black)

		tinydraw.FilledCircle(&display, 150, 60, 16, white)
		tinydraw.FilledCircle(&display, 150, 70, 4, black)
	}
}
