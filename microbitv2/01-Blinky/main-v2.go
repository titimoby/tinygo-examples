package main

import (
	"time"

	"tinygo.org/x/drivers/microbitmatrix"
)

func main() {
	// set LED display
	display := microbitmatrix.New()
	display.Configure(microbitmatrix.Config{})

	for {
		time.Sleep(time.Millisecond * 500)
		display.SetPixel(0, 0, microbitmatrix.BrightnessFull)
		time.Sleep(time.Millisecond * 500)

		display.Display()
	}
}
