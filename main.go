package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
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

func main() {
	ebiten.SetWindowTitle("Pong in Go")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	g := &Game{}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Update() error {
	return nil
}
