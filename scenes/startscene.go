package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type StartScene struct { // is the scene loaded now
	loaded bool
}

func NewStartScene() *StartScene {
	return &StartScene{
		loaded: false,
	}
}

func (s *StartScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 0, 0, 255})
	ebitenutil.DebugPrint(screen, "Pong in Go by IncredibleLego. Press enter to start.")

}

func (s *StartScene) FirstLoad() {
	s.loaded = true
}

func (s *StartScene) IsLoaded() bool {
	return s.loaded
}

func (s *StartScene) OnEnter() {

}

func (s *StartScene) OnExit() {

}

func (s *StartScene) Update() SceneId {
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		return GameSceneId
	}
	return StartSceneId
}

var _ Scene = (*StartScene)(nil)
