package objects

import (
	"goPong/constants"

	"github.com/hajimehoshi/ebiten/v2"
)

type Paddle struct {
	*Object
}

func (p *Paddle) MoveOnKeyPress(keyUp, keyDown ebiten.Key) bool { // Move the paddle based on keypress
	if ebiten.IsKeyPressed(keyDown) && p.Y+p.H < constants.ScreenHeight { // can't go below the screen
		p.Y += constants.PaddleSpeed
		return true
	}
	if ebiten.IsKeyPressed(keyUp) && p.Y > 0 { // can't go above the screen
		p.Y -= constants.PaddleSpeed
		return true
	}
	return false
}
