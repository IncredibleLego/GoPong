package scenes

import (
	"goPong/config"
	"goPong/utils"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type NameInputScene struct {
	mode         int
	numPlayers   int
	playerNames  [2]string
	activePlayer int
	finished     bool
	maxLetters   int
	maxLenght    bool
	timer        time.Time
}

func NewNameInputScene(mode int) *NameInputScene {
	var numPlayers int
	if mode == 2 {
		numPlayers = 2
	} else {
		numPlayers = 1
	}
	return &NameInputScene{
		mode:         mode,
		numPlayers:   numPlayers,
		activePlayer: 0,
		finished:     false,
		maxLetters:   14,
		maxLenght:    false,
	}
}

func (n *NameInputScene) Draw(screen *ebiten.Image) {
	if n.finished {
		return
	}

	if time.Since(n.timer) > time.Second*2 {
		n.timer = time.Now()
	}
	l := float64(len(n.playerNames[n.activePlayer]))
	height := float64(config.GlobalConfig.ScreenHeight)
	d := 20 - config.GlobalConfig.TextDimension

	playerMessage := "Player " + strconv.Itoa(n.activePlayer+1) + ", insert your name:"
	x1 := utils.XCentered(playerMessage, config.GlobalConfig.TextDimension)
	utils.ScreenDraw(d, x1, height/3, "yellow", screen, playerMessage)

	message := n.playerNames[n.activePlayer]
	if time.Since(n.timer) < time.Second && !n.maxLenght {
		message += "_"
	}
	utils.ScreenDraw(d, float64(config.GlobalConfig.ScreenWidth)/2-(l*20/2), height/2, "white", screen, message)

	confirmMessage := "Press Enter to confirm"
	x2 := utils.XCentered(confirmMessage, config.GlobalConfig.TextDimension)
	utils.ScreenDraw(d, x2, (height/3)*2, "yellow", screen, confirmMessage)

	if n.maxLenght {
		errorMessage := "The name can be max " + strconv.Itoa(n.maxLetters) + " letters"
		x2 := utils.XCentered(confirmMessage, config.GlobalConfig.TextDimension)
		utils.ScreenDraw(d-5, x2, (height/3)*2+50, "red", screen, errorMessage)
	}

}

func (n *NameInputScene) FirstLoad() {

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

	utils.Input(&n.playerNames[n.activePlayer], n.maxLetters)

	// Check if max letters has been reached to print the error message
	n.maxLenght = len(n.playerNames[n.activePlayer]) > n.maxLetters

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

	return NameInputSceneId
}

var _ Scene = (*NameInputScene)(nil)
