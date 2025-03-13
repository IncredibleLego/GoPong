package main

import (
	"bytes"
	_ "embed"
	"image/color"
	"log"
	"main/constants"
	"main/objects"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	paddle      objects.Paddle
	enemyPaddle objects.Paddle
	ball        objects.Ball
	score       int
	highScore   int
}

//go:embed PressStart2P-Regular.ttf
var pressStart2P []byte
var pressStart2PFaceSource *text.GoTextFaceSource

func main() {
	ebiten.SetWindowTitle("Pong in Go")
	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	paddle := objects.Paddle{
		Ojbect: objects.Ojbect{
			X: 600,
			Y: 200,
			W: 15,
			H: 100,
		},
	}
	enemypaddle := objects.Paddle{
		Ojbect: objects.Ojbect{
			X: 40,
			Y: 200,
			W: 15,
			H: 100,
		},
	}
	ball := objects.Ball{
		Ojbect: objects.Ojbect{
			X: 0,
			Y: 0,
			W: 15,
			H: 15,
		},
		Dxdt: constants.BallSpeed,
		Dydt: constants.BallSpeed,
	}
	g := &Game{
		paddle:      paddle,
		enemyPaddle: enemypaddle,
		ball:        ball,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return constants.ScreenWidth, constants.ScreenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen,
		float32(g.paddle.X), float32(g.paddle.Y),
		float32(g.paddle.W), float32(g.paddle.H),
		color.White, false,
	) // Draw the paddle
	vector.DrawFilledRect(screen,
		float32(g.enemyPaddle.X), float32(g.enemyPaddle.Y),
		float32(g.enemyPaddle.W), float32(g.enemyPaddle.H),
		color.White, false,
	) // Draw enemy paddle
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

func (g *Game) Reset() { // Reset the game
	g.ball.X = 0
	g.ball.Y = 0
	g.score = 0
}

func (g *Game) CollideWithWall() { // Check if the ball collides with the wall
	if g.ball.X >= constants.ScreenWidth {
		g.Reset()
	} else if g.ball.X <= 0 {
		g.ball.Dxdt = constants.BallSpeed
	} else if g.ball.Y <= 0 {
		g.ball.Dydt = constants.BallSpeed
	} else if g.ball.Y >= constants.ScreenHeight {
		g.ball.Dydt = -constants.BallSpeed
	}
}

func (g *Game) CollideWithPaddle() { // Check if the ball collides with the paddle
	if g.ball.X >= g.paddle.X && g.ball.Y >= g.paddle.Y && g.ball.Y <= g.paddle.Y+g.paddle.H {
		g.ball.Dxdt = -g.ball.Dxdt
		g.IncreaseScore()

	}
}

func (g *Game) IncreaseScore() {
	g.score++
	if g.score > g.highScore {
		g.highScore = g.score
	}
}
