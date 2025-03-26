package menu

import (
	"bytes"
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed PressStart2P-Regular.ttf
var pressStart2P []byte
var pressStart2PFaceSource *text.GoTextFaceSource

func ScreenDraw(size int, x, y float64, r, g, b, a float32, spacing float64, screen *ebiten.Image, line string) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(pressStart2P))
	if err != nil {
		log.Fatal(err)
	}
	pressStart2PFaceSource = s

	textFace := &text.GoTextFace{
		Source: pressStart2PFaceSource,
		Size:   float64(size),
	}

	textOptions := &text.DrawOptions{}
	textOptions.GeoM.Translate(x, y)
	textOptions.ColorScale.Scale(r, g, b, a)
	textOptions.LineSpacing = spacing

	text.Draw(screen, line, textFace, textOptions)
}

func MeasureText(option MenuOption, spacing float64) (float64, float64) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(pressStart2P))
	if err != nil {
		log.Fatal(err)
	}
	pressStart2PFaceSource = s

	textFace := &text.GoTextFace{
		Source: pressStart2PFaceSource,
		Size:   13,
	}

	boundsX, boundsY := text.Measure(option.Label, textFace, spacing)
	return boundsX, boundsY
}
