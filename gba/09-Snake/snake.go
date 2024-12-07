package main

import (
	"image/color"
	"math/rand"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinygba"
)

// Snake game like the old 3310 Nokia game
// Each time the snake eats a spider, it grows and score++

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

	// Snake information
	WIDTHBLOCKS  = 24
	HEIGHTBLOCKS = 16

	snakeDefaultLength = 3

	// Levels
	levelEasy   = 300
	levelNormal = 120
	levelHard   = 50
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
	snake               Snake
	spiderX, spiderY    int16
	Status              uint8
	score, frame, delay int
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
		delay:   levelNormal,
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
	case LEFT, RIGHT, UP, DOWN:
		g.snake.direction = direction
	}

	g.moveSnake()
}

func (g *Game) Over() {
	clearScreen()
	splashed = false

	g.Status = GameOver
}

func (g *Game) SetLevel() {

	//TODO: Fix Bug passage des levels
	switch {
	case g.delay == levelHard:
		// Define level to easy
		activateEasyLevel()
		deactivateHardLevel()
		g.delay = levelEasy

	case g.delay == levelNormal:
		// Define level to hard
		activateHardLevel()
		deactivateNormalLevel()
		g.delay = levelHard

	case g.delay == levelEasy:
		// Define level to normal
		activateNormalLevel()
		deactivateEasyLevel()
		g.delay = levelNormal
	}
}

func (g *Game) drawStartScreen() {
	clearScreen()

	tinyfont.WriteLine(&display, &freesans.Bold24pt7b, 37, 78, "SNAKE", snakeFont)

	//Depending on the level, select the good difficulty
	displayLevelChoices()

	//v2
	tinyfont.WriteLine(&display, &tinyfont.TomThumb, 84, 150, "Press START button", white)

	displayScore()
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

func activateNormalLevel() {
	tinydraw.FilledCircle(&display, 101, 117, 3, snakeFont)
}

func deactivateNormalLevel() {
	tinydraw.FilledCircle(&display, 101, 117, 3, nokiaBG)
	tinydraw.Circle(&display, 101, 117, 3, snakeFont)
}

func activateEasyLevel() {
	tinydraw.FilledCircle(&display, 66, 117, 3, snakeFont)
}

func deactivateEasyLevel() {
	tinydraw.FilledCircle(&display, 66, 117, 3, nokiaBG)
	tinydraw.Circle(&display, 66, 117, 3, snakeFont)
}

func activateHardLevel() {
	tinydraw.FilledCircle(&display, 143, 117, 3, snakeFont)
}

func deactivateHardLevel() {
	tinydraw.FilledCircle(&display, 143, 117, 3, nokiaBG)
	tinydraw.Circle(&display, 143, 117, 3, snakeFont)
}

func displayLevelChoices() {

	displayEasyLevelText()
	displayNormalLevelText()
	displayHardLevelText()

	// Simulate radio buttons/choices
	switch {
	case game.delay == levelEasy:
		deactivateHardLevel()
		deactivateNormalLevel()
		activateEasyLevel()

	case game.delay == levelNormal:
		deactivateEasyLevel()
		activateNormalLevel()
		deactivateHardLevel()

	case game.delay == levelHard:
		deactivateEasyLevel()
		deactivateNormalLevel()
		activateHardLevel()
	}
}

func displayScore() {
	if game.score > 0 {
		scoreStr[7] = 48 + uint8((game.score)/100)
		scoreStr[8] = 48 + uint8(((game.score)/10)%10)
		scoreStr[9] = 48 + uint8((game.score)%10)

		tinyfont.WriteLine(&display, &tinyfont.TomThumb, 98, 100, string(scoreStr), snakeFont)
	}
}

func displayEasyLevelText() {
	tinyfont.WriteLine(&display, &tinyfont.TomThumb, 72, 120, "EASY", snakeFont)
}

func displayNormalLevelText() {
	tinyfont.WriteLine(&display, &tinyfont.TomThumb, 107, 120, "NORMAL", snakeFont)
}

func displayHardLevelText() {
	tinyfont.WriteLine(&display, &tinyfont.TomThumb, 148, 120, "HARD", snakeFont)
}
