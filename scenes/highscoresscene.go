package scenes

import (
	"goPong/config"
	"goPong/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type HighScoresScene struct {
	highscoreSelected int // 0 for solo, 1 for computer, 2 for mulitplayer
	computerSelected  int // 0 for easy, 1 for default, 2 for hard
}

func (h *HighScoresScene) ShouldPreserveState(reason SceneChangeReason) bool {
	return false
}

func NewHighScoresScene() *HighScoresScene {
	return &HighScoresScene{
		highscoreSelected: 0,
		computerSelected:  0,
	}
}

func (h *HighScoresScene) Draw(screen *ebiten.Image) {

	utils.HighscoresTableDraw(screen)

	X := float64(config.GlobalConfig.ScreenHeight) / 7.2
	Y := float64(config.GlobalConfig.ScreenHeight) * 0.090277778
	I := float64(config.GlobalConfig.ScreenHeight) * 0.075

	var scores []string
	var dimension float64
	if h.highscoreSelected == 0 {
		utils.ScreenDraw(config.GlobalConfig.TextDimension/6, X, Y, "sky blue", screen, "Solo Mode High Scores")
		scores = GetSoloHighscoresStrings()
		dimension = config.GlobalConfig.TextDimension / 2
	} else if h.highscoreSelected == 1 {
		utils.ScreenDraw(config.GlobalConfig.TextDimension/30, X, Y, "sky blue", screen, "Computer Mode High Scores")
		scores = GetComputerHighscoresStrings(h.computerSelected)
		dimension = config.GlobalConfig.TextDimension / 1.56
	} else if h.highscoreSelected == 2 {
		utils.ScreenDraw(-(config.GlobalConfig.TextDimension / 15), X, Y, "sky blue", screen, "Multiplayer Mode High Scores")
		scores = GetMultiplayerHighscoresStrings()
		dimension = config.GlobalConfig.TextDimension / 1.56
	}

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
		utils.ScreenDraw(-dimension, X, X+X/2+float64(i)*I, color, screen, scores[i])
	}
}

func (h *HighScoresScene) FirstLoad() {}

func (h *HighScoresScene) OnEnter() {}

func (h *HighScoresScene) Update() SceneId {

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		h.highscoreSelected++
	} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		h.highscoreSelected--
	}

	h.highscoreSelected = utils.Adjust(h.highscoreSelected)

	if h.highscoreSelected != 1 {
		h.computerSelected = 1
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		h.computerSelected--
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		h.computerSelected++
	}
	h.computerSelected = utils.Adjust(h.computerSelected)

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return StartSceneId
	}

	return HighScoresSceneId
}

var _ Scene = (*HighScoresScene)(nil)
