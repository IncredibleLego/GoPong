package scenes

import (
	"fmt"
	"goPong/config"
	"goPong/menu"
	"goPong/utils"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type NameInputScene struct {
	mode               int
	numPlayers         int
	playerNames        [2]string
	activePlayer       int
	finished           bool
	maxLetters         int
	maxLenght          bool
	difficulty         bool
	difficultyMenu     *menu.OptionMenu
	lastEnterPressTime time.Time
	timer              time.Time
}

func NewNameInputScene(mode int) *NameInputScene {
	var numPlayers int
	var difficulty bool
	if mode == 2 {
		numPlayers = 2
	} else {
		numPlayers = 1
		if mode == 3 {
			difficulty = true
		}
	}
	return &NameInputScene{
		mode:         mode,
		numPlayers:   numPlayers,
		activePlayer: 0,
		finished:     false,
		maxLetters:   14,
		maxLenght:    false,
		difficulty:   difficulty,
	}
}

func (n *NameInputScene) Draw(screen *ebiten.Image) {
	if n.finished {
		return
	}

	if n.difficulty {
		n.difficultyMenu.Draw(screen)
	} else {

		if time.Since(n.timer) > time.Second*2 {
			n.timer = time.Now()
		}
		l := float64(len(n.playerNames[n.activePlayer]))
		height := float64(config.GlobalConfig.ScreenHeight)
		d := config.GlobalConfig.TextDimension

		playerMessage := "Player " + strconv.Itoa(n.activePlayer+1) + ", insert your name:"
		x1 := utils.XCentered(playerMessage, config.GlobalConfig.TextDimension)
		utils.ScreenDraw(0, x1, height/3, "yellow", screen, playerMessage)

		message := n.playerNames[n.activePlayer]
		if time.Since(n.timer) < time.Second && !n.maxLenght {
			message += "_"
		}
		utils.ScreenDraw(0, float64(config.GlobalConfig.ScreenWidth)/2-(l*d/2), height/2, "white", screen, message)

		confirmMessage := "Press Enter to confirm"
		x2 := utils.XCentered(confirmMessage, config.GlobalConfig.TextDimension)
		utils.ScreenDraw(0, x2+d, (height/3)*2, "yellow", screen, confirmMessage)

		if n.maxLenght {
			errorMessage := "The name can be max " + strconv.Itoa(n.maxLetters) + " letters"
			x2 := utils.XCentered(confirmMessage, config.GlobalConfig.TextDimension)
			utils.ScreenDraw(-(d / 4), x2, (height/10)*8, "red", screen, errorMessage)
		}
	}
}

func (n *NameInputScene) FirstLoad() {
	n.difficultyMenu = &menu.OptionMenu{
		Options:      []string{"EASY", "DEFAULT", "HARD"},
		Selected:     1,
		LastMoveTime: time.Now(),
		MenuName:     "DIFFICULTY",
		Position:     float64(config.GlobalConfig.ScreenHeight) / 2.5,
	}
	n.lastEnterPressTime = time.Now()
}

func (n *NameInputScene) OnEnter() {

}

func (n *NameInputScene) OnExit() {

}

func (n *NameInputScene) ShouldPreserveState(reason SceneChangeReason) bool {
	return false
}

func (n *NameInputScene) Update() SceneId {
	if n.finished {
		return GameSceneId
	}

	if n.difficulty {
		n.difficultyMenu.Update()

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			var value float64
			if n.difficultyMenu.Selected == 1 {
				value = 0.5
			} else if n.difficultyMenu.Selected == 0 {
				value = 0.2
			} else {
				value = 0.8
			}

			err := config.UpdateConfig(func(cfg *config.Config) {
				cfg.Difficulty = value
			})
			if err != nil {
				fmt.Println("Error during option saving", err)
			}
			n.difficulty = false
		}
	} else {

		utils.Input(&n.playerNames[n.activePlayer], n.maxLetters)

		// Check if max letters has been reached to print the error message
		n.maxLenght = len(n.playerNames[n.activePlayer]) >= n.maxLetters

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			n.activePlayer++
			if n.activePlayer >= n.numPlayers { // Se entrambi hanno inserito il nome
				n.finished = true // Indica che l'input è terminato
				config.GlobalConfig.Player1Name = n.playerNames[0]
				if n.numPlayers == 2 {
					config.GlobalConfig.Player2Name = n.playerNames[1]
				}
				if n.mode == 1 {
					return GameSceneId // Passa direttamente alla scena di gioco
				} else if n.mode == 2 {
					return MultiplayerSceneId // Passa alla scena multiplayer
				} else if n.mode == 3 {
					return ComputerSceneId // Passa alla scena computer
				}
			}
		}
	}

	return NameInputSceneId
}

var _ Scene = (*NameInputScene)(nil)
