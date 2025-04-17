package scenes

import (
	"goPong/config"
	"goPong/utils"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type NameInputScene struct {
	mode         int
	numPlayers   int
	playerNames  [2]string
	activePlayer int
	finished     bool
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
	}
}

func (n *NameInputScene) Draw(screen *ebiten.Image) {
	if n.finished {
		return
	}

	l := float64(len(n.playerNames[n.activePlayer]))
	//width := float64(config.GlobalConfig.ScreenWidth)
	height := float64(config.GlobalConfig.ScreenHeight)

	vector.DrawFilledRect(screen,
		float32(config.GlobalConfig.ScreenWidth/2), float32(config.GlobalConfig.ScreenHeight/2),
		float32(3), float32(12),
		color.White, false,
	)

	d := 20 - config.GlobalConfig.TextDimension

	//utils.ScreenDraw(d, 65, height/3, "white", screen, "Player "+strconv.Itoa(n.activePlayer+1)+", insert your name:")

	playerMessage := "Player " + strconv.Itoa(n.activePlayer+1) + ", insert your name:"
	x1 := utils.XCentered(playerMessage, config.GlobalConfig.TextDimension)
	utils.ScreenDraw(d, x1, height/3, "white", screen, playerMessage)

	utils.ScreenDraw(d, float64(config.GlobalConfig.ScreenWidth)/2-(l*20/2), height/2, "white", screen, n.playerNames[n.activePlayer])

	//utils.ScreenDraw(d, 130, height/3*2, "white", screen, "Press Enter to confirm")

	confirmMessage := "Press Enter to confirm"
	x2 := utils.XCentered(confirmMessage, config.GlobalConfig.TextDimension)
	utils.ScreenDraw(d, x2, (height/3)*2, "white", screen, confirmMessage)

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

	// Backspace
	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) && len(n.playerNames[n.activePlayer]) > 0 {
		n.playerNames[n.activePlayer] = n.playerNames[n.activePlayer][:len(n.playerNames[n.activePlayer])-1]
	}

	// Alphabetic characters
	for key := ebiten.KeyA; key <= ebiten.KeyZ; key++ {
		if inpututil.IsKeyJustPressed(key) {
			n.playerNames[n.activePlayer] += string('A' + (key - ebiten.KeyA))
		}
	}

	// Numerical characters
	for key := ebiten.Key0; key <= ebiten.Key9; key++ {
		if inpututil.IsKeyJustPressed(key) {
			n.playerNames[n.activePlayer] += string('0' + (key - ebiten.Key0))
		}
	}

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
