package objects

import (
	"goPong/config"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Paddle struct {
	*Object
	aiTargetY     float64
	aiY           float64
	aiVelocityY   float64
	aiNextThink   time.Time
	aiOffset      int
	aiApproaching bool
}

func (p *Paddle) Draw(screen *ebiten.Image) {
	var colore color.RGBA
	if config.Secret {
		colore = color.RGBA{253, 137, 0, 1}
	} else {
		colore = color.RGBA{255, 255, 255, 255}
	}

	vector.DrawFilledRect(screen,
		float32(p.X), float32(p.Y),
		float32(p.W), float32(p.H),
		colore, false,
	)
}

func (p *Paddle) MoveOnKeyPress(keyUp, keyDown ebiten.Key) bool {
	if ebiten.IsKeyPressed(keyDown) && p.Y+p.H < config.GlobalConfig.ScreenHeight {
		p.Y += config.GlobalConfig.PaddleSpeed
		return true
	}
	if ebiten.IsKeyPressed(keyUp) && p.Y > 0 {
		p.Y -= config.GlobalConfig.PaddleSpeed
		return true
	}
	return false
}

// AiMovement moves the paddle based on the ball's position and the current difficulty level
func (p *Paddle) AiMovement(b *Ball) {
	now := time.Now()
	difficulty := config.GlobalConfig.Difficulty

	var reactionDelay time.Duration
	var maxErrorPct float64
	var speedMultiplier float64
	var accelMultiplier float64

	switch {
	case difficulty < 0.33:
		reactionDelay = 200 * time.Millisecond
		maxErrorPct = 0.26
		speedMultiplier = 0.78
		accelMultiplier = 0.13
	case difficulty < 0.66:
		reactionDelay = 125 * time.Millisecond
		maxErrorPct = 0.16
		speedMultiplier = 1.00
		accelMultiplier = 0.16
	default:
		reactionDelay = 70 * time.Millisecond
		maxErrorPct = 0.08
		speedMultiplier = 1.22
		accelMultiplier = 0.20
	}

	// Generate min and max values for error, speed and acceleration based on the difficulty level
	maxError := int(float64(p.H) * maxErrorPct)
	if maxError < 4 {
		maxError = 4
	}

	maxSpeed := float64(config.GlobalConfig.PaddleSpeed) * speedMultiplier
	if maxSpeed < 1 {
		maxSpeed = 1
	}

	acceleration := maxSpeed * accelMultiplier
	if acceleration < 0.25 {
		acceleration = 0.25
	}

	centerY := float64(config.GlobalConfig.ScreenHeight/2 - p.H/2)

	// If the ball is moving towards the paddle, it tries to intercept it
	if b.Dxdt >= 0 {
		p.aiApproaching = false
		p.aiOffset = 0
		p.aiTargetY = centerY

		delta := p.aiTargetY - p.aiY
		desiredVelocity := clampFloat(delta*0.08, -maxSpeed*0.55, maxSpeed*0.55)
		p.aiVelocityY = approachFloat(p.aiVelocityY, desiredVelocity, acceleration)
		p.aiY += p.aiVelocityY
		p.syncPosition()
		p.clampToScreen()
		return
	}

	// If the ball is moving away from the paddle and the AI is not already approaching it calculates a new target position with some randomness
	if !p.aiApproaching {
		p.aiApproaching = true
		p.aiOffset = randomInRange(-maxError, maxError)
		p.aiTargetY = float64(b.Y + p.aiOffset - p.H/2)
		p.aiNextThink = now
	}

	// Update the target position at certain intervals so simulate human-like reaction time
	if now.After(p.aiNextThink) {
		p.aiTargetY = float64(b.Y + p.aiOffset - p.H/2)
		p.aiNextThink = now.Add(reactionDelay)
	}

	delta := p.aiTargetY - p.aiY

	// If the paddle is close enough to the target it will slow down to avoid overshooting and make the movement more human-like
	if math.Abs(delta) <= 1 {
		p.aiVelocityY = approachFloat(p.aiVelocityY, 0, acceleration)
		p.aiY += p.aiVelocityY
		p.syncPosition()
		p.clampToScreen()
		return
	}

	// Calculate the desired velocity based on the distance to the target with a maximum speed limit
	desiredVelocity := clampFloat(delta*0.12, -maxSpeed, maxSpeed)
	p.aiVelocityY = approachFloat(p.aiVelocityY, desiredVelocity, acceleration)

	p.aiY += p.aiVelocityY
	p.syncPosition()
	p.clampToScreen()
}

// Updates the paddle's Y to match the AI's calculated position
func (p *Paddle) syncPosition() {
	p.Y = int(math.Round(p.aiY))
}

// Mantein the paddle between the boundaries of the screen
func (p *Paddle) clampToScreen() {
	if p.Y < 0 {
		p.Y = 0
		p.aiY = 0
		p.aiVelocityY = 0
	}
	if p.Y+p.H > config.GlobalConfig.ScreenHeight {
		p.Y = config.GlobalConfig.ScreenHeight - p.H
		p.aiY = float64(p.Y)
		p.aiVelocityY = 0
	}
}

// Apply min and max limits to a int value
func clamp(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Apply min and max limits to a float value
func clampFloat(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// If the paddle is abot to reach the target it will slow down to avoid overshooting and make the movement more human-like
func approachFloat(current, target, step float64) float64 {
	if current < target {
		current += step
		if current > target {
			return target
		}
		return current
	}
	if current > target {
		current -= step
		if current < target {
			return target
		}
		return current
	}
	return current
}

// Initializes the AI state to avoid random behavior at the start of the game or after a pause
func (p *Paddle) InitAIStateFromCurrentY() {
	p.aiY = float64(p.Y)
	p.aiTargetY = float64(p.Y)
	p.aiVelocityY = 0
	p.aiApproaching = false
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
