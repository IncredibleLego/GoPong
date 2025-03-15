package scenes

import (
	"goPong/menu"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type StartScene struct { // is the scene loaded now
	menu   *menu.Menu
	loaded bool
}

func NewStartScene() *StartScene {
	return &StartScene{
		menu:   nil,
		loaded: false,
	}
}

func (s *StartScene) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{255, 0, 0, 255})
	/*
		menu.ScreenDraw(13, 250, 180, 1, 1, 1, 1, 1.5, screen, "Pong in Go")
		menu.ScreenDraw(13, 210, 200, 1, 1, 1, 1, 1.5, screen, "by IncredibleLego")
		menu.ScreenDraw(13, 250, 220, 1, 1, 1, 1, 1.5, screen, "Press Enter to start") */

	s.menu.Draw(screen)

}

func (s *StartScene) FirstLoad() {
	s.menu = &menu.Menu{
		Options: []string{
			"Solo",
			"Contro IA",
			"Multiplayer",
		},
		Selected:     0,
		LastMoveTime: time.Now(),
	}
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
	s.menu.Update()

	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		return GameSceneId
	}
	return StartSceneId
}

var _ Scene = (*StartScene)(nil)
