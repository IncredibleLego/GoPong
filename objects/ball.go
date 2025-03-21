package objects

import (
	"goPong/constants"
	"time"
)

type Ball struct {
	*Object
	Dxdt int // x velocity per tick
	Dydt int // y velocity per tick
}

func (b *Ball) Move() { // Move the ball
	b.X += b.Dxdt
	b.Y += b.Dydt
}

func (b *Ball) IncreaseSpeed(increase int) { // Increase the speed of the ball
	b.Dxdt += increase
	b.Dydt += increase
}

// w1 and w2 are the horizontal walls options that the ball can collide with
func (b *Ball) CollideWithWall(w1, w2 bool) { // Check if the ball collides with the wall
	if b.X <= 0 {
		if w1 {
			b.Reset()
		} else {
			b.Dxdt = -b.Dxdt
		}
	} else if b.X+b.W >= constants.ScreenWidth {
		if w2 {
			b.Reset()
		} else {
			b.Dxdt = -b.Dxdt
		}
	} else if b.Y <= 0 {
		b.Dydt = -b.Dydt
	} else if b.Y >= constants.ScreenHeight {
		b.Dydt = -b.Dydt
	}
}

func (b *Ball) Reset() { // Reset the ball to the center of the screen
	go func() {
		b.X = constants.ScreenWidth / 2
		b.Y = constants.ScreenHeight / 2
		b.Dxdt = 0
		b.Dydt = 0
		time.Sleep(time.Second)
		b.Dxdt = constants.BallSpeed
		b.Dydt = constants.BallSpeed
	}()
}
