package scenes

import (
	"goPong/constants"
	"goPong/menu"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PauseScene struct {
	loaded bool
}

func NewPauseScene() *PauseScene {
	return &PauseScene{
		loaded: false,
	}
}

func (g *PauseScene) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{200, 200, 0, 200})

	menu.ScreenDraw(constants.TextDimension, 250, 180, 1, 1, 1, 1, 1.5, screen, "Pause menu")
	menu.ScreenDraw(constants.TextDimension, 180, 200, 1, 1, 1, 1, 1.5, screen, "Press Enter to unpause")
	menu.ScreenDraw(constants.TextDimension, 200, 220, 1, 1, 1, 1, 1.5, screen, "Press 'q' to quit")

}

func (s *PauseScene) FirstLoad() {
	s.loaded = true
}

func (s *PauseScene) IsLoaded() bool {
	return s.loaded
}

func (s *PauseScene) OnEnter() {

}

func (s *PauseScene) OnExit() {

}

func (s *PauseScene) Update() SceneId {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return ExitSceneId
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return GameSceneId
	}

	return PauseSceneId
}

var _ Scene = (*PauseScene)(nil)
