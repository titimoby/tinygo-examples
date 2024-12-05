package main

import (
	"image/color"
	"math/rand"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinygba"
)

// background 155,186,90 like the real nokia 3310 - nokiaBG
// https://playsnake.org/
// Display the splash start screen with the same fonts and please press start button
// Change apple to spider
// TODO: Faire bordure comme Telecran
// TODO: add score at the top?
// TODO: possibility to change the speed? press L button reduce speed, press R button increase speed?
// BUG - Fix si le snake move a droite, et qu'on appuie sur la fleche gauche alors erreur ça plante le jeu

const (
	GameSplash = iota
	GameStart
	GamePlay
	GameOver
)

const (
	// Directions
	UP    = "up"
	DOWN  = "down"
	LEFT  = "left"
	RIGHT = "right"

	WIDTHBLOCKS  = 24
	HEIGHTBLOCKS = 16

	snakeDefaultLength = 3
)

var (
	// Colors
	white     = color.RGBA{255, 255, 255, 255}
	nokiaBG   = color.RGBA{155, 186, 90, 255}
	snakeFont = color.RGBA{0, 67, 0, 255}

	// The array of color that represent the spider pixelized
	spiderBuf = []color.RGBA{
		nokiaBG, nokiaBG, snakeFont, snakeFont, snakeFont, snakeFont, snakeFont, snakeFont, nokiaBG, nokiaBG,
		snakeFont, snakeFont, snakeFont, snakeFont, snakeFont, snakeFont, snakeFont, snakeFont, snakeFont, snakeFont,
		snakeFont, nokiaBG, snakeFont, snakeFont, snakeFont, snakeFont, snakeFont, snakeFont, nokiaBG, snakeFont,
		snakeFont, nokiaBG, snakeFont, snakeFont, snakeFont, snakeFont, snakeFont, snakeFont, nokiaBG, snakeFont,
		snakeFont, nokiaBG, snakeFont, snakeFont, snakeFont, snakeFont, snakeFont, snakeFont, nokiaBG, snakeFont,
		snakeFont, nokiaBG, nokiaBG, snakeFont, nokiaBG, nokiaBG, snakeFont, nokiaBG, nokiaBG, snakeFont,
	}
)

type Snake struct {
	body      [104][2]int16
	length    int16
	direction string
}

type Game struct {
	snake            Snake
	spiderX, spiderY int16
	Status           uint8
	score            int
	frame, delay     int
}

var splashed = false
var scoreStr = []byte("SCORE: 123")

func NewGame() *Game {
	return &Game{
		snake: Snake{
			body: [104][2]int16{
				{0, 3},
				{0, 2},
				{0, 1},
			},
			length:    snakeDefaultLength,
			direction: RIGHT,
		},
		spiderX: 5,
		spiderY: 5,
		Status:  GameSplash,
		delay:   120,
	}
}

func (g *Game) Splash() {
	if !splashed {
		g.drawStartScreen()
		splashed = true
	}
}

func (g *Game) Start() {
	clearScreen()

	g.initSnake()
	g.drawSnake()
	g.createSpider()

	g.Status = GamePlay
}

func (g *Game) Play(direction string) {
	g.frame++
	if g.frame%g.delay > 0 {
		return
	}

	switch direction {
	case LEFT:
		g.snake.direction = direction
	case RIGHT:
		g.snake.direction = direction
	case UP:
		g.snake.direction = direction
	case DOWN:
		g.snake.direction = direction
	}

	g.moveSnake()
}

func (g *Game) Over() {
	clearScreen()
	splashed = false

	g.Status = GameOver
}

func (g *Game) drawStartScreen() {
	clearScreen()

	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, 37, 80, "SNAKE", snakeFont)

	tinyfont.WriteLine(&display, &tinyfont.TomThumb, 84, 100, "Press START button", white)

	//tinygba.FillRectangleWithBuffer(10, 10, 10, 6, spiderBuf)

	if g.score > 0 {
		scoreStr[7] = 48 + uint8((g.score)/100)
		scoreStr[8] = 48 + uint8(((g.score)/10)%10)
		scoreStr[9] = 48 + uint8((g.score)%10)

		tinyfont.WriteLine(&display, &tinyfont.TomThumb, 98, 120, string(scoreStr), snakeFont)
	}
}

// When the game begin, the snake moves to right and have a length of 3 blocks (snakeDefaultLength)
func (g *Game) initSnake() {
	g.snake.body[0][0] = 0
	g.snake.body[0][1] = 3
	g.snake.body[1][0] = 0
	g.snake.body[1][1] = 2
	g.snake.body[2][0] = 0
	g.snake.body[2][1] = 1

	g.snake.length = snakeDefaultLength
	g.snake.direction = RIGHT
}

func (g *Game) collisionWithSnake(x, y int16) bool {
	for i := int16(0); i < g.snake.length; i++ {
		if x == g.snake.body[i][0] && y == g.snake.body[i][1] {
			return true
		}
	}
	return false
}

func (g *Game) createSpider() {
	g.spiderX = int16(rand.Int31n(16))
	g.spiderY = int16(rand.Int31n(13))
	for g.collisionWithSnake(g.spiderX, g.spiderY) {
		g.spiderX = int16(rand.Int31n(16))
		g.spiderY = int16(rand.Int31n(13))
	}
	g.drawSpider(g.spiderX, g.spiderY)
}

func (g *Game) moveSnake() {
	x := g.snake.body[0][0]
	y := g.snake.body[0][1]

	switch g.snake.direction {
	case LEFT:
		x--
		break
	case UP:
		y--
		break
	case DOWN:
		y++
		break
	case RIGHT:
		x++
		break
	}
	if x >= WIDTHBLOCKS {
		x = 0
	}
	if x < 0 {
		x = WIDTHBLOCKS - 1
	}
	if y >= HEIGHTBLOCKS {
		y = 0
	}
	if y < 0 {
		y = HEIGHTBLOCKS - 1
	}

	// Game Over
	if g.collisionWithSnake(x, y) {
		// Score = nb of spider eaten = actual length of the snake - default length
		g.score = int(g.snake.length - snakeDefaultLength)
		g.Over()

		return
	}

	// draw head
	g.drawSnakePartial(x, y, snakeFont)
	if x == g.spiderX && y == g.spiderY {
		g.snake.length++
		g.createSpider()
	} else {
		// remove tail
		g.drawSnakePartial(g.snake.body[g.snake.length-1][0], g.snake.body[g.snake.length-1][1], nokiaBG)
	}
	for i := g.snake.length - 1; i > 0; i-- {
		g.snake.body[i][0] = g.snake.body[i-1][0]
		g.snake.body[i][1] = g.snake.body[i-1][1]
	}
	g.snake.body[0][0] = x
	g.snake.body[0][1] = y
}

func (g *Game) drawSpider(x, y int16) {
	tinygba.FillRectangleWithBuffer(10*x, 10*y, 10, 6, spiderBuf)
}

func (g *Game) drawSnake() {
	for i := int16(0); i < g.snake.length; i++ {
		g.drawSnakePartial(g.snake.body[i][0], g.snake.body[i][1], snakeFont)
	}
}

func (g *Game) drawSnakePartial(x, y int16, c color.RGBA) {
	modY := int16(9)
	if y == 12 {
		modY = 8
	}
	tinygba.FillRectangle(10*x, 10*y, 10, modY, c)
}
