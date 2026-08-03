package scenes

import (
	"crypto/sha256"
	"fmt"
	"goPong/config"
	"goPong/utils"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type CreditsScene struct {
	secret       bool
	secretString string
	placeholder  string
	timer        time.Time
}

func (c *CreditsScene) ShouldPreserveState(reason SceneChangeReason) bool {
	return false
}

func NewCreditsScene() *CreditsScene {
	return &CreditsScene{}
}

func (c *CreditsScene) Draw(screen *ebiten.Image) {

	var X, Y float64

	if !c.secret {

		messages := []string{
			"Made by",
			"IncredibleLego",
			"in about 5 months",
			"",
			"Based On",
			`"PONG" by Alan Alcorn (1972)`,
			"",
			"Thanks to",
			"Friends and Family who",
			"beta-tested the game",
			"",
			"Version 1.1.0 2026",
		}
		for i, message := range messages {
			X = utils.XCentered(message, config.GlobalConfig.TextDimension-3)
			Y = float64(config.GlobalConfig.ScreenHeight) / 14 * float64(i+1)

			var color string
			switch i {
			case 0, 2, 4, 7, 11:
				color = "yellow"
			case 1:
				color = "orange"
			case 5:
				color = "sky blue"
			case 8, 9:
				color = "cyan"
			default:
				color = "white"
			}
			utils.ScreenDraw(-3, X, Y, color, screen, message)
		}
	} else {
		if time.Since(c.timer) > time.Second*2 {
			c.timer = time.Now()
		}

		c.placeholder = ""
		for i := 0; i < len(c.secretString); i++ {
			c.placeholder += "*"
		}

		X = utils.XCentered(c.placeholder, config.GlobalConfig.TextDimension)
		Y = float64(config.GlobalConfig.ScreenHeight) / 2

		if time.Since(c.timer) < time.Second && len(c.placeholder) < 14 {
			c.placeholder += "_"
		}
		utils.ScreenDraw(0, X, Y, "white", screen, c.placeholder)
	}
}

func (c *CreditsScene) FirstLoad() {}

func (c *CreditsScene) OnEnter() {}

func (c *CreditsScene) Update() SceneId {
	if !c.secret {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			return StartSceneId
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyI) {
			c.secret = true
			c.secretString = ""
			c.placeholder = ""
			return CreditsSceneId
		}
	} else {
		utils.Input(&c.secretString, 14)
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {

			sum := fmt.Sprintf("%x", sha256.Sum256([]byte(c.secretString)))

			if sum == "3ed48f34c6ee18a7a2677d5b2856c3ee67628ea8b3db114e9eac7e0ff2da66e1" {
				config.Secret = true
			}
			c.secret = false
			return StartSceneId
		}
	}
	return CreditsSceneId
}

var _ Scene = (*CreditsScene)(nil)
