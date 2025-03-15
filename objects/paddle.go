package objects

import (
	"goPong/constants"

	"github.com/hajimehoshi/ebiten/v2"
)

type Paddle struct {
	*Object
}

func (p *Paddle) MoveOnKeyPress() { // Move the paddle based on keypress
	if (ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS)) && p.Y+p.H < constants.ScreenHeight { // can't go below the screen
		p.Y += constants.PaddleSpeed
	}
	if (ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW)) && p.Y > 0 { // can't go above the screen
		p.Y -= constants.PaddleSpeed
	}
}

func (p *Paddle) CollideWithPaddle(b *Ball) bool { // Check if the ball collides with the paddle
	return p.X < b.X+b.W &&
		p.X+p.W > b.X &&
		p.Y < b.Y+b.H &&
		p.Y+p.H > b.Y
}
