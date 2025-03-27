package menu

import (
	"bytes"
	_ "embed"
	"goPong/config"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed PressStart2P-Regular.ttf
var pressStart2P []byte
var pressStart2PFaceSource *text.GoTextFaceSource

func ScreenDraw(size int, x, y float64, colorName string, screen *ebiten.Image, line string) {
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
	r, g, b, a := Color(colorName)
	textOptions.ColorScale.Scale(r, g, b, a)
	textOptions.LineSpacing = float64(size) / 10

	text.Draw(screen, line, textFace, textOptions)
}

func MeasureText(option MenuOption) (float64, float64) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(pressStart2P))
	if err != nil {
		log.Fatal(err)
	}
	pressStart2PFaceSource = s

	textFace := &text.GoTextFace{
		Source: pressStart2PFaceSource,
		Size:   float64(config.GlobalConfig.TextDimension),
	}

	boundsX, boundsY := text.Measure(option.Label, textFace, float64(config.GlobalConfig.TextDimension)/10)
	return boundsX, boundsY
}

func Color(colorName string) (float32, float32, float32, float32) {
	switch colorName {
	case "white":
		return 255, 255, 255, 255
	case "black":
		return 0, 0, 0, 255
	case "red":
		return 255, 0, 0, 255
	case "green":
		return 0, 255, 0, 255
	case "blue":
		return 0, 0, 255, 255
	case "yellow":
		return 255, 255, 0, 255
	case "cyan":
		return 0, 255, 255, 255
	case "magenta":
		return 255, 0, 255, 255
	case "light gray":
		return 204, 204, 204, 255
	case "dark gray":
		return 51, 51, 51, 255
	case "orange":
		return 255, 128, 0, 255
	case "pink":
		return 255, 128, 179, 255
	case "lime":
		return 128, 255, 0, 255
	case "sky blue":
		return 77, 153, 255, 255
	case "purple":
		return 153, 0, 255, 255
	case "brown":
		return 153, 77, 0, 255
	case "dark red":
		return 128, 0, 0, 255
	case "dark green":
		return 0, 128, 0, 255
	case "dark blue":
		return 0, 0, 128, 255
	case "dark purple":
		return 102, 0, 153, 255
	default:
		log.Printf("Unknown color: %s", colorName)
		return 0, 0, 0, 255 // Default to black
	}
}
