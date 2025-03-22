package objects

import (
	"goPong/constants"
	"math"

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

/*
func (p *Paddle) CollideWithPaddle(b *Ball) bool { // Check if the ball collides with the paddle
	if p.X < b.X+b.W && p.X+p.W > b.X && p.Y < b.Y+b.H && p.Y+p.H > b.Y {
		// Calulate the impact point based on the center of the paddle

		impactPoint := (p.Y + p.H/2) - (b.Y + b.H/2)

		// Normalize the result

		normalizedImpactPoint := float64(impactPoint) / float64(p.H/2)

		// Calculate inflexion angle based on the normalized impact point
		// The ball reflects with a lower vertical speed if the impact point is closer to the center of the paddle

		b.Dydt = int(float64(constants.BallSpeed) * normalizedImpactPoint)
		b.Dxdt = -b.Dxdt

		return true
	}
	return false
} */

func (p *Paddle) CollideWithPaddle(b *Ball, direction bool) bool { // Check if the ball collides with the paddle
	check := false

	// direction is true if the ball is moving to the left, false otherwise
	if direction {
		//fmt.Println("checking case 1")
		if p.X < b.X+b.W && p.X+p.W > b.X+b.W && p.Y < b.Y+b.H && p.Y+p.H > b.Y {
			check = true
		}
	} else {
		//fmt.Println("checking case 2")
		if p.X < b.X && p.X+p.W > b.X && p.Y < b.Y+b.H && p.Y+p.H > b.Y {
			check = true
		}
	}

	if check {
		// Calulate the impact point based on the center of the paddle
		impactPoint := (p.Y + p.H/2) - (b.Y + b.H/2)

		// Normalize the result
		normalizedImpactPoint := float64(impactPoint) / float64(p.H/2)

		// Calculate the new vertical speed based on the normalized impact point
		newDydt := float64(constants.BallSpeed) * normalizedImpactPoint

		// Ensure the newDydt does not exceed the total speed
		if math.Abs(newDydt) > float64(constants.BallSpeed) {
			newDydt = float64(constants.BallSpeed) * math.Copysign(1, newDydt)
		}

		// Calculate the new horizontal speed to maintain the total speed
		newDxdt := math.Sqrt(float64(constants.BallSpeed*constants.BallSpeed) - newDydt*newDydt)

		// Ensure the newDxdt does not fall below the minimum speed
		if newDxdt < float64(constants.BallSpeed) {
			newDxdt = float64(constants.BallSpeed)
		}

		// Update the ball's velocity
		b.Dydt = -int(newDydt)
		//b.Dxdt = -int(newDxdt)

		if !direction {
			b.Dxdt = int(newDxdt)
		} else {
			b.Dxdt = -int(newDxdt)
		}

		return true
	}

	return false
}
