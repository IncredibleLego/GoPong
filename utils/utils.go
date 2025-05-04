package utils

import (
	"bytes"
	_ "embed"
	"goPong/config"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// This file contains utility functions for the game.

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
	r, g, b, a := Color(colorName)
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

// Gives x coord to place a message in the middle of the screen given the message and the font size
func XCentered(message string, fontSize float64) float64 {
	width := float64(len(message)) * fontSize
	x := (float64(config.GlobalConfig.ScreenWidth) / 2) - (width / 2)
	return x
}

func Net(screen *ebiten.Image) {
	for i := 0; i < config.GlobalConfig.ScreenHeight; i += 24 {
		vector.DrawFilledRect(screen,
			float32(config.GlobalConfig.ScreenWidth/2), float32(i),
			float32(3), float32(12),
			color.White, false,
		)
	}
}

func Color(colorName string) (float32, float32, float32, float32) {
	switch colorName {
	case "white":
		// RGBA: 255, 255, 255, 255
		return 1, 1, 1, 1
	case "black":
		// RGBA: 0, 0, 0, 255
		return 0, 0, 0, 1
	case "red":
		// RGBA: 255, 0, 0, 255
		return 1, 0, 0, 1
	case "green":
		// RGBA: 0, 255, 0, 255
		return 0, 1, 0, 1
	case "blue":
		// RGBA: 0, 0, 255, 255
		return 0, 0, 1, 1
	case "yellow":
		// RGBA: 255, 255, 0, 255
		return 1, 1, 0, 1
	case "cyan":
		// RGBA: 0, 255, 255, 255
		return 0, 1, 1, 1
	case "magenta":
		// RGBA: 255, 0, 255, 255
		return 1, 0, 1, 1
	case "light gray":
		// RGBA: 204, 204, 204, 255
		return 0.8, 0.8, 0.8, 1
	case "dark gray":
		// RGBA: 51, 51, 51, 255
		return 0.2, 0.2, 0.2, 1
	case "orange":
		// RGBA: 255, 128, 0, 255
		return 1, 0.5, 0, 1
	case "pink":
		// RGBA: 255, 128, 179, 255
		return 1, 0.5, 0.7, 1
	case "lime":
		// RGBA: 128, 255, 0, 255
		return 0.5, 1, 0, 1
	case "sky blue":
		// RGBA: 77, 153, 255, 255
		return 0.3, 0.6, 1, 1
	case "purple":
		// RGBA: 153, 0, 255, 255
		return 0.6, 0, 1, 1
	case "brown":
		// RGBA: 153, 77, 0, 255
		return 0.6, 0.3, 0, 1
	case "dark red":
		// RGBA: 128, 0, 0, 255
		return 0.5, 0, 0, 1
	case "dark green":
		// RGBA: 0, 128, 0, 255
		return 0, 0.5, 0, 1
	case "dark blue":
		// RGBA: 0, 0, 128, 255
		return 0, 0, 0.5, 1
	case "dark purple":
		// RGBA: 102, 0, 153, 255
		return 0.4, 0, 0.6, 1
	default:
		log.Printf("Unknown color: %s", colorName)
		return 0, 0, 0, 255 // Default to black
	}
}
