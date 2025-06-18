package scenes

import (
	"goPong/config"
	"goPong/menu"
	"goPong/utils"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type HighScoresScene struct {
	chooseMenu        *menu.RegularMenu
	actionExecuted    bool
	highscoreSelected int // 0 for menu, 1 for solo, 2 for computer, 3 for multiplayer
}

func (h *HighScoresScene) ShouldPreserveState(reason SceneChangeReason) bool {
	return false
}

func NewHighScoresScene() *HighScoresScene {
	return &HighScoresScene{
		chooseMenu:        nil,
		highscoreSelected: 0,
	}
}

func (h *HighScoresScene) Draw(screen *ebiten.Image) {

	switch h.highscoreSelected {
	case 0:
		h.chooseMenu.Draw(screen)
	case 1:
		utils.HighscoresTableDraw(screen)

		X := float64(config.GlobalConfig.ScreenHeight) / 7.2
		I := float64(config.GlobalConfig.ScreenHeight) * 0.075

		utils.ScreenDraw(5, X, 65, "sky blue", screen, "Solo Mode HighScores")

		scores := GetSoloHighscoresStrings()
		for i := 0; i < len(scores); i++ {
			var color string
			if i == 0 {
				color = "gold"
			} else if i == 1 {
				color = "silver"
			} else if i == 2 {
				color = "bronze"
			} else {
				color = "soft yellow"
			}
			utils.ScreenDraw(-(config.GlobalConfig.TextDimension / 2), X, X+X/2+float64(i)*I, color, screen, scores[i])
		}
	case 2:
		utils.ScreenDraw(5, 100, 100, "white", screen, "Computer Mode HighScores")

		scores := GetComputerHighscoresStrings()
		for i := 0; i < len(scores); i++ {
			utils.ScreenDraw(-15, 100, float64(150+i*40), "white", screen, scores[i])
		}
	case 3:
		utils.ScreenDraw(5, 100, 100, "white", screen, "Multiplayer Mode High Scores")

		scores := GetMultiplayerHighscoresStrings()
		for i := 0; i < len(scores); i++ {
			utils.ScreenDraw(-15, 100, float64(150+i*40), "white", screen, scores[i])
		}
	}
}

func (h *HighScoresScene) FirstLoad() {
	h.chooseMenu = &menu.RegularMenu{
		Options: []menu.MenuOption{
			{Label: "SOLO MODE HIGHSCORES"},
			{Label: "COMPUTER MODE HIGHSCORES"},
			{Label: "MULTIPLAYER MODE HIGHSCORES"},
			{Label: "BACK"},
		},
		Selected:      0,
		LastMoveTime:  time.Now(),
		MainColor:     "blue",
		SelectedColor: "orange",
	}
}

func (h *HighScoresScene) OnEnter() {}

func (h *HighScoresScene) OnExit() {}

func (h *HighScoresScene) Update() SceneId {

	switch h.highscoreSelected {
	case 0:
		h.chooseMenu.Update()

		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && !h.actionExecuted {
			id := h.handleMenuSelection()
			h.actionExecuted = true
			if id != PauseSceneId {
				return id
			}
		}

		if inpututil.KeyPressDuration(ebiten.KeyEnter) == 0 && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			h.actionExecuted = false
		}

	case 1:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			h.highscoreSelected = 0
		}
	case 2:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			h.highscoreSelected = 0
		}
	case 3:
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			h.highscoreSelected = 0
		}
	}

	return HighScoresSceneId
}

var _ Scene = (*HighScoresScene)(nil)

func (h *HighScoresScene) handleMenuSelection() SceneId {
	switch h.chooseMenu.Selected {
	case 0:
		h.highscoreSelected = 1
	case 1:
		h.highscoreSelected = 2
	case 2:
		h.highscoreSelected = 3
	case 3:
		return StartSceneId
	}

	return HighScoresSceneId
}
