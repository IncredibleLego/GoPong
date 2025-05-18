package objects

import (
	"goPong/audio"
	"goPong/config"
	"image/color"
	"math"
	"math/rand/v2"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Ball struct {
	*Object
	Dxdt int // x velocity per tick
	Dydt int // y velocity per tick
}

func (b *Ball) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen,
		float32(b.X), float32(b.Y),
		float32(b.W), float32(b.H),
		color.White, false,
	)
}

func (b *Ball) Move() { // Move the ball
	b.X += b.Dxdt
	b.Y += b.Dydt
}

func (b *Ball) IncreaseSpeed(increase int) { // Increase the speed of the ball
	if b.Dxdt > 0 {
		b.Dxdt += increase
	} else {
		b.Dxdt -= increase
	}
}

// w1 and w2 are the horizontal walls options that the ball can collide with
func (b *Ball) CollideWithWall(w1, w2 bool) int { // Check if the ball collides with the wall
	if b.X <= 0 {
		if w1 {
			audio.PlayScore()
			b.Reset(true) //true = left player got scored
			return 1
		} else {
			audio.PlayPaddle()
			b.Dxdt = -b.Dxdt
		}
	} else if b.X+b.W >= config.GlobalConfig.ScreenWidth {
		if w2 {
			audio.PlayScore()
			b.Reset(false) //false = right player got scored
			return 2
		} else {
			b.Dxdt = -b.Dxdt
		}
	} else if b.Y <= 0 {
		audio.PlayPaddle()
		b.Dydt = -b.Dydt
	} else if b.Y+b.H >= config.GlobalConfig.ScreenHeight {
		audio.PlayPaddle()
		b.Dydt = -b.Dydt
	}
	return 0
}

func (b *Ball) CollideWithPaddle(p *Paddle, direction bool) bool { // Check if the ball collides with the paddle
	check := false

	// direction is true if the ball is moving to the left, false otherwise
	if direction {
		if p.X < b.X+b.W && p.X+p.W > b.X+b.W && p.Y < b.Y+b.H && p.Y+p.H > b.Y {
			audio.PlayPaddle()
			check = true
		}
	} else {
		if p.X < b.X && p.X+p.W > b.X && p.Y < b.Y+b.H && p.Y+p.H > b.Y {
			audio.PlayPaddle()
			check = true
		}
	}

	if check {
		// Calculate the impact point based on the center of the paddle
		impactPoint := (p.Y + p.H/2) - (b.Y + b.H/2)

		// Normalize the result
		normalizedImpactPoint := float64(impactPoint) / float64(p.H/2)

		// Calculate the new vertical speed based on the normalized impact point
		newDydt := float64(config.GlobalConfig.BallSpeed) * normalizedImpactPoint

		// Ensure the newDydt does not exceed the total speed
		if math.Abs(newDydt) > float64(config.GlobalConfig.BallSpeed) {
			newDydt = float64(config.GlobalConfig.BallSpeed) * math.Copysign(1, newDydt)
		}

		// Calculate the new horizontal speed to maintain the total speed
		newDxdt := math.Sqrt(float64(config.GlobalConfig.BallSpeed*config.GlobalConfig.BallSpeed) - newDydt*newDydt)

		// Ensure the newDxdt does not fall below the minimum speed
		if newDxdt < float64(config.GlobalConfig.BallSpeed) {
			newDxdt = float64(config.GlobalConfig.BallSpeed)
		}

		// Update the ball's velocity
		b.Dydt = -int(newDydt)

		if !direction {
			b.Dxdt = int(newDxdt)
		} else {
			b.Dxdt = -int(newDxdt)
		}

		return true
	}

	return false
}

func (b *Ball) Reset(p bool) { // Reset the ball to the center of the screen
	go func() {
		b.X = config.GlobalConfig.ScreenWidth/2 - b.W/2
		b.Y = config.GlobalConfig.ScreenHeight/2 - b.H/2
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
	b.Dxdt = int(float64(config.GlobalConfig.BallSpeed) * math.Cos(radians))
	b.Dydt = int(float64(config.GlobalConfig.BallSpeed) * math.Sin(radians))
}
