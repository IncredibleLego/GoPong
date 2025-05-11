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
	lastMoveTime time.Time
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

	l := float64(len(n.playerNames[n.activePlayer]))

	height := float64(config.GlobalConfig.ScreenHeight)

	d := 20 - config.GlobalConfig.TextDimension

	playerMessage := "Player " + strconv.Itoa(n.activePlayer+1) + ", insert your name:"
	x1 := utils.XCentered(playerMessage, config.GlobalConfig.TextDimension)

	utils.ScreenDraw(d, x1, height/3, "yellow", screen, playerMessage)

	if time.Since(n.timer) < time.Second && !n.maxLenght {
		utils.ScreenDraw(d, float64(config.GlobalConfig.ScreenWidth)/2-(l*20/2), height/2, "white", screen, n.playerNames[n.activePlayer]+"_")
	} else {
		utils.ScreenDraw(d, float64(config.GlobalConfig.ScreenWidth)/2-(l*20/2), height/2, "white", screen, n.playerNames[n.activePlayer])
	}

	if time.Since(n.timer) > time.Second*2 {
		n.timer = time.Now()
	}

	//utils.ScreenDraw(d, float64(config.GlobalConfig.ScreenWidth)/2-(l*20/2), height/2, "white", screen, n.playerNames[n.activePlayer])

	//utils.ScreenDraw(d, 130, height/3*2, "white", screen, "Press Enter to confirm")

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

	moveInterval := time.Duration(time.Second / 8)

	// Check if max letters has been reached
	if len(n.playerNames[n.activePlayer]) > n.maxLetters {
		n.maxLenght = true
	} else {
		n.maxLenght = false
	}

	// Calculate the time since the last letter or number was pressed
	var timeSinceKey []int
	for key := ebiten.KeyA; key <= ebiten.KeyZ; key++ {
		timeSinceKey = append(timeSinceKey, inpututil.KeyPressDuration(key))
	}
	var timeSinceNumber []int
	for key := ebiten.Key0; key <= ebiten.Key9; key++ {
		timeSinceNumber = append(timeSinceNumber, inpututil.KeyPressDuration(key))
	}

	// Backspace
	if len(n.playerNames[n.activePlayer]) > 0 {
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			n.playerNames[n.activePlayer] = n.playerNames[n.activePlayer][:len(n.playerNames[n.activePlayer])-1]
			n.lastMoveTime = time.Now()

		} else if inpututil.KeyPressDuration(ebiten.KeyBackspace) > 30 && time.Since(n.lastMoveTime) >= moveInterval {
			n.playerNames[n.activePlayer] = n.playerNames[n.activePlayer][:len(n.playerNames[n.activePlayer])-1]
			n.lastMoveTime = time.Now()
		}
	}

	if !n.maxLenght {
		// Alphabetical characters
		for key := ebiten.KeyA; key <= ebiten.KeyZ; key++ {
			if inpututil.IsKeyJustPressed(key) {
				n.playerNames[n.activePlayer] += string('A' + (key - ebiten.KeyA))
				n.lastMoveTime = time.Now()
			} else if timeSinceKey[int(key)] > 30 && time.Since(n.lastMoveTime) >= moveInterval {
				n.playerNames[n.activePlayer] += string('A' + (key - ebiten.KeyA))
				n.lastMoveTime = time.Now()
			}
		}
		// Numerical characters
		for key := ebiten.Key0; key <= ebiten.Key9; key++ {
			if inpututil.IsKeyJustPressed(key) {
				n.playerNames[n.activePlayer] += string('0' + (key - ebiten.Key0))
				n.lastMoveTime = time.Now()
			} else if timeSinceNumber[int(key)-43] > 30 && time.Since(n.lastMoveTime) >= moveInterval {
				n.playerNames[n.activePlayer] += string('0' + (key - ebiten.Key0))
				n.lastMoveTime = time.Now()
			}
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
