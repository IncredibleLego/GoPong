package scenes

import (
	"goPong/config"
	"goPong/menu"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

	menu.ScreenDraw(-7, 100, 100, "white", screen, "Player "+strconv.Itoa(n.activePlayer+1)+", insert your name:")
	menu.ScreenDraw(-7, 100, 150, "white", screen, n.playerNames[n.activePlayer])
	menu.ScreenDraw(-7, 100, 200, "white", screen, "Press Enter to confirm")

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
