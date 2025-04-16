package menu

import (
	"bytes"
	_ "embed"
	"goPong/config"
	"goPong/utils"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed PressStart2P-Regular.ttf
var pressStart2P []byte
var pressStart2PFaceSource *text.GoTextFaceSource

func ScreenDraw(size float64, x, y float64, colorName string, screen *ebiten.Image, line string) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(pressStart2P))
	if err != nil {
		log.Fatal(err)
	}
	pressStart2PFaceSource = s

	textFace := &text.GoTextFace{
		Source: pressStart2PFaceSource,
		Size:   config.GlobalConfig.TextDimension + size,
	}

	textOptions := &text.DrawOptions{}
	textOptions.GeoM.Translate(x, y)
	r, g, b, a := utils.Color(colorName)
	textOptions.ColorScale.Scale(r, g, b, a)
	textOptions.LineSpacing = float64(size) / 10

	text.Draw(screen, line, textFace, textOptions)
}

func MeasureText(label string) (float64, float64) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(pressStart2P))
	if err != nil {
		log.Fatal(err)
	}
	pressStart2PFaceSource = s

	textFace := &text.GoTextFace{
		Source: pressStart2PFaceSource,
		Size:   config.GlobalConfig.TextDimension,
	}

	boundsX, boundsY := text.Measure(label, textFace, float64(config.GlobalConfig.TextDimension)/10)
	return boundsX, boundsY
}
