package scenes

import (
	"bytes"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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

	screen.Fill(color.RGBA{0, 255, 0, 255})
	// Text Options
	s, err := text.NewGoTextFaceSource(bytes.NewReader(pressStart2P))
	if err != nil {
		log.Fatal(err)
	}
	pressStart2PFaceSource = s

	textFace := &text.GoTextFace{
		Source: pressStart2PFaceSource,
		Size:   13,
	}

	textOptions := &text.DrawOptions{}
	textOptions.GeoM.Translate(250, 180)
	textOptions.ColorScale.Scale(1, 1, 1, 1)
	textOptions.LineSpacing = 1.5

	row := "Pause menu"
	text.Draw(screen, row, textFace, textOptions)

	textOptions.GeoM.Translate(-70, 20)

	row = "Press Enter to unpause"
	text.Draw(screen, row, textFace, textOptions)

	textOptions.GeoM.Translate(20, 20)

	row = "Press 'q' to quit"
	text.Draw(screen, row, textFace, textOptions)

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
