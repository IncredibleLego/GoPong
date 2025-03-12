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

func main() {
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
