package objects

import (
	"goPong/constants"
	"math"
	"math/rand/v2"
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
func (b *Ball) CollideWithWall(w1, w2 bool) int { // Check if the ball collides with the wall
	if b.X <= 0 {
		if w1 {
			b.Reset(true) //true = left player got scored
			return 1
		} else {
			b.Dxdt = -b.Dxdt
		}
	} else if b.X+b.W >= constants.ScreenWidth {
		if w2 {
			b.Reset(false) //false = right player got scored
			return 2
		} else {
			b.Dxdt = -b.Dxdt
		}
	} else if b.Y <= 0 {
		b.Dydt = -b.Dydt
	} else if b.Y+b.H >= constants.ScreenHeight {
		b.Dydt = -b.Dydt
	}
	return 0
}

func (b *Ball) Reset(p bool) { // Reset the ball to the center of the screen
	go func() {
		b.X = constants.ScreenWidth/2 - b.W/2
		b.Y = constants.ScreenHeight/2 - b.H/2
		b.Dxdt = 0
		b.Dydt = 0
		time.Sleep(time.Second)

		b.GenerateRandomDirection()

		// Ensure the ball moves in the correct direction based on the player who scored
		if p {
			b.Dxdt = -b.Dxdt
		}
	}()
}

func (b *Ball) GenerateRandomDirection() {
	// Generate a random angle between -45 and 45 degrees
	angle := rand.Float64()*90 - 45
	radians := angle * (math.Pi / 180)

	// Calculate the new velocities based on the angle
	b.Dxdt = int(float64(constants.BallSpeed) * math.Cos(radians))
	b.Dydt = int(float64(constants.BallSpeed) * math.Sin(radians))
}
