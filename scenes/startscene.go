package scenes

import (
	"bytes"
	"goPong/menu"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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

	// Text Options
	f, err := text.NewGoTextFaceSource(bytes.NewReader(pressStart2P))
	if err != nil {
		log.Fatal(err)
	}
	pressStart2PFaceSource = f

	textFace := &text.GoTextFace{
		Source: pressStart2PFaceSource,
		Size:   13,
	}

	textOptions := &text.DrawOptions{}
	textOptions.GeoM.Translate(250, 180)
	textOptions.ColorScale.Scale(1, 1, 1, 1)
	textOptions.LineSpacing = 1.5

	row := "Pong in Go"
	text.Draw(screen, row, textFace, textOptions)

	textOptions.GeoM.Translate(-40, 20)

	row = "by IncredibleLego"
	text.Draw(screen, row, textFace, textOptions)

	textOptions.GeoM.Translate(-20, 20)

	row = "Press Enter to start"
	text.Draw(screen, row, textFace, textOptions)

	s.menu.Draw(screen)

}

func (s *StartScene) FirstLoad() {
	s.menu = &menu.Menu{
		Options: []string{
			"Solo",
			"Contro IA",
			"Multiplayer",
		},
		Selected: 0,
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
