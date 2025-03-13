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

	scoreTextOptions := &text.DrawOptions{}
	scoreTextOptions.GeoM.Translate(250, 180)
	scoreTextOptions.ColorScale.Scale(1, 1, 1, 1)
	scoreTextOptions.LineSpacing = 1.5

	highScoreTextOptions := &text.DrawOptions{}
	highScoreTextOptions.GeoM.Translate(180, 200)
	highScoreTextOptions.ColorScale.Scale(1, 1, 1, 1)
	highScoreTextOptions.LineSpacing = 1.5

	textFace := &text.GoTextFace{
		Source: pressStart2PFaceSource,
		Size:   13,
	}

	scoreStr := "Pause menu"
	text.Draw(screen, scoreStr, textFace, scoreTextOptions)

	HighScoreStr := "Press Enter to unpause"
	text.Draw(screen, HighScoreStr, textFace, highScoreTextOptions)

}

/*
func (s *PauseScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 255, 0, 255})
	ebitenutil.DebugPrint(screen, "Press enter to unpause.")
} */

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
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return GameSceneId
	}

	return PauseSceneId
}

var _ Scene = (*PauseScene)(nil)
