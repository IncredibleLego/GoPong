package objects

import (
	"main/constants"

	"github.com/hajimehoshi/ebiten/v2"
)

type Paddle struct {
	Ojbect
}

func (p *Paddle) MoveOnKeyPress() { // Move the paddle based on keypress
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && p.Y+p.H < constants.ScreenHeight { // can't go below the screen
		p.Y += constants.PaddleSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && p.Y > 0 { // can't go above the screen
		p.Y -= constants.PaddleSpeed
	}
}
