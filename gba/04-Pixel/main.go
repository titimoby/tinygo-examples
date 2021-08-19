package main

import (
	"image/color"
	"machine"
	"runtime/volatile"
	"unsafe"
)

var (
	// Register display
	regDISPSTAT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000004)))

	// Display from machine
	display = machine.Display

	// Colors
	white = color.RGBA{255, 255, 255, 255}
	// green = color.RGBA{0, 255, 0, 255}
	// red   = color.RGBA{255, 0, 0, 255}

	x int16 = 100
	y int16 = 50
)

func main() {
	// Set up the display
	display.Configure()

	// Register display status
	regDISPSTAT.SetBits(1<<3 | 1<<4)

	//1
	for i := 0; i < 2; i++ {
		display.SetPixel(x, y+int16(i), white)
		display.SetPixel(x+1, y+int16(i), white)
		display.SetPixel(x+2, y+int16(i), white)
		display.SetPixel(x+3, y+int16(i), white)
		display.SetPixel(x+4, y+int16(i), white)
		display.SetPixel(x+5, y+int16(i), white)
		display.SetPixel(x+6, y+int16(i), white)
		display.SetPixel(x+7, y+int16(i), white)

		display.SetPixel(x+13, y+int16(i), white)
		display.SetPixel(x+14, y+int16(i), white)
		display.SetPixel(x+15, y+int16(i), white)
		display.SetPixel(x+16, y+int16(i), white)
		display.SetPixel(x+17, y+int16(i), white)
		display.SetPixel(x+18, y+int16(i), white)
		display.SetPixel(x+19, y+int16(i), white)
		display.SetPixel(x+20, y+int16(i), white)

		i = i + 1
	}

	//2
	for i := 2; i < 4; i++ {
		display.SetPixel(x-2, y+int16(i), white)
		display.SetPixel(x-1, y+int16(i), white)

		display.SetPixel(x+8, y+int16(i), white)
		display.SetPixel(x+9, y+int16(i), white)
		display.SetPixel(x+10, y+int16(i), white)
		display.SetPixel(x+11, y+int16(i), white)
		display.SetPixel(x+12, y+int16(i), white)

		display.SetPixel(x+21, y+int16(i), white)
		display.SetPixel(x+22, y+int16(i), white)

		i = i + 1
	}

	//3
	for i := 4; i < 10; i++ {
		display.SetPixel(x-4, y+int16(i), white)
		display.SetPixel(x-3, y+int16(i), white)

		display.SetPixel(x+23, y+int16(i), white)
		display.SetPixel(x+24, y+int16(i), white)

		i = i + 1
	}

	//3
	for i := 10; i < 18; i++ {
		display.SetPixel(x-2, y+int16(i), white)
		display.SetPixel(x-1, y+int16(i), white)

		display.SetPixel(x+21, y+int16(i), white)
		display.SetPixel(x+22, y+int16(i), white)

		i = i + 1
	}

	//4
	for i := 18; i < 21; i++ {
		display.SetPixel(x, y+int16(i), white)
		display.SetPixel(x+1, y+int16(i), white)

		display.SetPixel(x+19, y+int16(i), white)
		display.SetPixel(x+20, y+int16(i), white)

		i = i + 1
	}

	//5
	for i := 21; i < 24; i++ {
		display.SetPixel(x+2, y+int16(i), white)
		display.SetPixel(x+3, y+int16(i), white)

		display.SetPixel(x+17, y+int16(i), white)
		display.SetPixel(x+18, y+int16(i), white)

		i = i + 1
	}

	//6
	for i := 24; i < 27; i++ {
		display.SetPixel(x+4, y+int16(i), white)
		display.SetPixel(x+5, y+int16(i), white)

		display.SetPixel(x+15, y+int16(i), white)
		display.SetPixel(x+16, y+int16(i), white)

		i = i + 1
	}

	//6
	for i := 27; i < 30; i++ {
		display.SetPixel(x+6, y+int16(i), white)
		display.SetPixel(x+7, y+int16(i), white)

		display.SetPixel(x+13, y+int16(i), white)
		display.SetPixel(x+14, y+int16(i), white)

		i = i + 1
	}

	//7
	for i := 30; i < 32; i++ {
		display.SetPixel(x+8, y+int16(i), white)
		display.SetPixel(x+9, y+int16(i), white)
		display.SetPixel(x+10, y+int16(i), white)
		display.SetPixel(x+11, y+int16(i), white)
		display.SetPixel(x+12, y+int16(i), white)

		i = i + 1
	}

	// Infinite loop to avoid exiting the application
	for {
	}
}
