package scenes

import (
	"goPong/utils"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type HighScoresScene struct {
}

func (h *HighScoresScene) ShouldPreserveState(reason SceneChangeReason) bool {
	return false
}

func NewHighScoresScene() *HighScoresScene {
	return &HighScoresScene{}
}

func (h *HighScoresScene) Draw(screen *ebiten.Image) {

	utils.ScreenDraw(5, 100, 100, "white", screen, "High Scores")

}

func (h *HighScoresScene) FirstLoad() {

}

func (h *HighScoresScene) OnEnter() {}

func (h *HighScoresScene) OnExit() {}

func (h *HighScoresScene) Update() SceneId {

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return StartSceneId
	}

	return HighScoresSceneId
}

var _ Scene = (*HighScoresScene)(nil)
