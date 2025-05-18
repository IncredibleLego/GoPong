package objects

import (
	"goPong/config"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Paddle struct {
	*Object
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen,
		float32(p.X), float32(p.Y),
		float32(p.W), float32(p.H),
		color.White, false,
	)
}

func (p *Paddle) MoveOnKeyPress(keyUp, keyDown ebiten.Key) bool { // Move the paddle based on keypress
	if ebiten.IsKeyPressed(keyDown) && p.Y+p.H < config.GlobalConfig.ScreenHeight { // can't go below the screen
		p.Y += config.GlobalConfig.PaddleSpeed
		return true
	}
	if ebiten.IsKeyPressed(keyUp) && p.Y > 0 { // can't go above the screen
		p.Y -= config.GlobalConfig.PaddleSpeed
		return true
	}
	return false
}

var aiTargetY int

func (p *Paddle) AiMovement(b *Ball) {
	// Difficulty: 0.1 (easy) to 1.0 (hard)
	// The main difference between difficulties is the % of hits, not speed.

	// Calculate max error based on difficulty (easy = more error)
	maxError := int(float64(config.GlobalConfig.PaddleHeight) * (1.1 - config.GlobalConfig.Difficulty) * 0.7)
	if maxError < 4 {
		maxError = 4
	}

	// AI reaction: only update target if ball is moving towards AI
	if b.Dxdt < 0 {
		randomOffset := randomInRange(-maxError, maxError)
		aiTargetY = b.Y + randomOffset - p.H/2
	}

	// Paddle speed: only slightly affected by difficulty
	baseSpeed := float64(config.GlobalConfig.PaddleSpeed)
	aiSpeed := int(baseSpeed * (0.85 + 0.3*config.GlobalConfig.Difficulty)) // 0.85x to 1.15x
	if aiSpeed < 1 {
		aiSpeed = 1
	}

	// Main difficulty: chance to "miss" the ball
	hitChance := config.GlobalConfig.Difficulty*0.7 + 0.25 // 0.25 (easy) to 0.95 (hard)
	if rand.Float64() > hitChance {
		// AI "misses" this frame: move less or not at all
		return
	}

	// Stop if the impact point has been reached (within 1 pixel)
	delta := aiTargetY - p.Y
	if abs(delta) <= 1 {
		return
	}

	// Move towards target, but not too fast
	move := clamp(delta, -aiSpeed, aiSpeed)
	p.Y += move

	// Clamp to screen
	if p.Y < 0 {
		p.Y = 0
	}
	if p.Y+p.H > config.GlobalConfig.ScreenHeight {
		p.Y = config.GlobalConfig.ScreenHeight - p.H
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func randomInRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func clamp(value, min, max int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}
