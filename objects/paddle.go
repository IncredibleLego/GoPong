package objects

import (
	"goPong/config"
	"math/rand"

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

var aiTargetY int

func (p *Paddle) AiMovement(b *Ball) { // AI movement
	/*
		pY := float64(p.Y)
		offset := (rand.Float64() - 0.5) * 100 * (1 - config.GlobalConfig.Difficulty) // Offset casuale maggiore a difficoltà bassa
		targetY := offset

		// Movimento verso il target con una velocità massima
		if pY < targetY && p.Y+p.H < config.GlobalConfig.ScreenHeight {
			p.Y += int(min(6, targetY-pY))
		} else if p.Y > 0 {
			p.Y -= int(min(6, pY-targetY))
		} */

	// Configurazione dei parametri

	maxError := config.GlobalConfig.PaddleHeight / 4

	aiSpeed := int(float64(config.GlobalConfig.PaddleSpeed) * 0.8) // Velocità dipende dalla difficoltà

	// Se la pallina si sta avvicinando, aggiorna il target
	if b.Dxdt < 0 {
		hitOffset := randomInRange(-maxError, maxError)
		aiTargetY = b.Y + hitOffset
	}

	// Movimento fluido della racchetta verso il target
	//p.Y += clamp(aiTargetY-p.Y, -aiSpeed, aiSpeed)

	if p.Y+p.H < config.GlobalConfig.ScreenHeight && p.Y >= 0 { // can't go below the screen
		p.Y += clamp(aiTargetY-p.Y, -aiSpeed, aiSpeed)
	}

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
