package objects

import (
	"goPong/config"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

type Paddle struct {
	*Object
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

func (p *Paddle) AiMovement(bY float64) { // AI movement
	pY := float64(p.Y)
	offset := (rand.Float64() - 0.5) * 100 * (1 - config.GlobalConfig.Difficulty) // Offset casuale maggiore a difficoltà bassa
	targetY := bY + offset

	// Movimento verso il target con una velocità massima
	if pY < targetY && p.Y+p.H < config.GlobalConfig.ScreenHeight {
		p.Y += int(min(6, targetY-pY))
	} else if p.Y > 0 {
		p.Y -= int(min(6, pY-targetY))
	}
}
