package main

import (
	"bytes"
	_ "embed"
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
	ballSpeed    = 3
	paddleSpeed  = 6
)

type Ojbect struct {
	X, Y, W, H int // X, Y, Width, Height. (0,0) is top-left.
}

type Paddle struct {
	Ojbect
}

type Ball struct {
	Ojbect
	dxdt int // x velocity per tick
	dydt int // y velocity per tick
}

type Game struct {
	paddle    Paddle
	ball      Ball
	score     int
	highScore int
}

//go:embed PressStart2P-Regular.ttf
var pressStart2P []byte
var pressStart2PFaceSource *text.GoTextFaceSource

func main() {
	ebiten.SetWindowTitle("Pong in Go")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	paddle := Paddle{
		Ojbect: Ojbect{
			X: 600,
			Y: 200,
			W: 15,
			H: 100,
		},
	}
	ball := Ball{
		Ojbect: Ojbect{
			X: 0,
			Y: 0,
			W: 15,
			H: 15,
		},
		dxdt: ballSpeed,
		dydt: ballSpeed,
	}
	g := &Game{
		paddle: paddle,
		ball:   ball,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen,
		float32(g.paddle.X), float32(g.paddle.Y),
		float32(g.paddle.W), float32(g.paddle.H),
		color.White, false,
	) // Draw the paddle
	vector.DrawFilledRect(screen,
		float32(g.ball.X), float32(g.ball.Y),
		float32(g.ball.W), float32(g.ball.H),
		color.White, false,
	) // Draw the ball

	// Text Options
	s, err := text.NewGoTextFaceSource(bytes.NewReader(pressStart2P))
	if err != nil {
		log.Fatal(err)
	}
	pressStart2PFaceSource = s

	scoreTextOptions := &text.DrawOptions{}
	scoreTextOptions.GeoM.Translate(10, 10)
	scoreTextOptions.ColorScale.Scale(1, 1, 1, 1)
	scoreTextOptions.LineSpacing = 1.5

	highScoreTextOptions := &text.DrawOptions{}
	highScoreTextOptions.GeoM.Translate(10, 30)
	highScoreTextOptions.ColorScale.Scale(1, 1, 1, 1)
	highScoreTextOptions.LineSpacing = 1.5

	textFace := &text.GoTextFace{
		Source: pressStart2PFaceSource,
		Size:   13,
	}

	scoreStr := "Score: " + strconv.Itoa(g.score)
	text.Draw(screen, scoreStr, textFace, scoreTextOptions)

	HighScoreStr := "High Score: " + strconv.Itoa(g.highScore)
	text.Draw(screen, HighScoreStr, textFace, highScoreTextOptions)

}

func (g *Game) Update() error {
	g.paddle.MoveOnKeyPress()
	g.ball.Move()
	g.CollideWithWall()
	g.CollideWithPaddle()
	return nil
}

func (p *Paddle) MoveOnKeyPress() { // Move the paddle based on keypress
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && p.Y+p.H < screenHeight { // can't go below the screen
		p.Y += paddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.Y > 0 { // can't go above the screen
		p.Y -= paddleSpeed
	}
}

func (b *Ball) Move() { // Move the ball
	b.X += b.dxdt
	b.Y += b.dydt
}

func (g *Game) Reset() { // Reset the game
	g.ball.X = 0
	g.ball.Y = 0
	g.score = 0
}

func (g *Game) CollideWithWall() { // Check if the ball collides with the wall
	if g.ball.X >= screenWidth {
		g.Reset()
	} else if g.ball.X <= 0 {
		g.ball.dxdt = ballSpeed
	} else if g.ball.Y <= 0 {
		g.ball.dydt = ballSpeed
	} else if g.ball.Y >= screenHeight {
		g.ball.dydt = -ballSpeed
	}
}

func (g *Game) CollideWithPaddle() { // Check if the ball collides with the paddle
	if g.ball.X >= g.paddle.X && g.ball.Y >= g.paddle.Y && g.ball.Y <= g.paddle.Y+g.paddle.H {
		g.ball.dxdt = -g.ball.dxdt
		g.IncreaseScore()
		if g.score > g.highScore {
			g.highScore = g.score
		}
	}
}

func (g *Game) IncreaseScore() {
	g.score++
}
