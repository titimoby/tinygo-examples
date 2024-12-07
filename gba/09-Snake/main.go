package main

import (
	"machine"

	"tinygo.org/x/tinygba"
)

var (
	// Display from machine
	display = machine.Display

	game = NewGame()
)

func main() {
	// Set up the display
	display.Configure()

	clearScreen()

	for {
		tinygba.WaitForVBlank()

		update()
	}
}

func update() {
	key := tinygba.ReadButtons()

	switch game.Status {
	case GameSplash:
		game.Splash()

		splash(key)

	case GamePlay:
		switch {
		case tinygba.ButtonStart.IsPushed(key):
			game.Over()

		case tinygba.ButtonRight.IsPushed(key):
			// Don't move if snake is moving to left
			if game.snake.direction != LEFT {
				//change direction and move snake to this direction
				game.Play(RIGHT)
			}
		case tinygba.ButtonLeft.IsPushed(key):
			// Don't move if snake is moving to right
			if game.snake.direction != RIGHT {
				//change direction and move snake to this direction
				game.Play(LEFT)
			}
		case tinygba.ButtonDown.IsPushed(key):
			// Don't move if snake is moving to up
			if game.snake.direction != UP {
				//change direction and move snake to this direction
				game.Play(DOWN)
			}
		case tinygba.ButtonUp.IsPushed(key):
			// Don't move if snake is moving to up
			if game.snake.direction != DOWN {
				//change direction and move snake to this direction
				game.Play(UP)
			}

		default:
			game.Play("")
		}

	case GameOver:
		game.Splash()

		splash(key)
	}
}

func clearScreen() {
	tinygba.FillScreen(nokiaBG)
}

func splash(key uint16) {
	if tinygba.ButtonStart.IsPushed(key) {
		// START = Start the game
		game.Start()
	} else if tinygba.ButtonSelect.IsPushed(key) {
		// SELECT = Change the level/difficulty of the game
		game.SetLevel()
	}
}
