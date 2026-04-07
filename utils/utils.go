package utils

import (
	"bytes"
	_ "embed"
	"goPong/config"
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// This file contains utility functions for the game.

//go:embed PressStart2P-Regular.ttf
var pressStart2P []byte

var (
	pressStart2PFaceOnce   sync.Once
	pressStart2PFaceSource *text.GoTextFaceSource
)

func getPressStart2PFaceSource() *text.GoTextFaceSource {
	pressStart2PFaceOnce.Do(func() {
		s, err := text.NewGoTextFaceSource(bytes.NewReader(pressStart2P))
		if err != nil {
			log.Fatal(err)
		}
		pressStart2PFaceSource = s
	})
	return pressStart2PFaceSource
}

func ScreenDraw(size float64, x, y float64, colorName string, screen *ebiten.Image, line string) {
	textFace := &text.GoTextFace{
		Source: getPressStart2PFaceSource(),
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
	textFace := &text.GoTextFace{
		Source: getPressStart2PFaceSource(),
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

// NEED TO ADAPT
func Net(screen *ebiten.Image) {

	dim := config.GlobalConfig.ScreenHeight / 30
	width := 3 * config.GlobalConfig.Scale

	for i := 0; i < config.GlobalConfig.ScreenHeight; i += dim {
		vector.DrawFilledRect(screen,
			float32(config.GlobalConfig.ScreenWidth/2), float32(i),
			float32(width), float32(dim/2),
			color.White, false,
		)
	}
}

// Prints the table used to show the highscores
func HighscoresTableDraw(screen *ebiten.Image) {

	var X float32 = float32(config.GlobalConfig.ScreenWidth / 18)
	var Y float32 = float32(config.GlobalConfig.ScreenHeight / 13)
	var offset float32 = float32(config.GlobalConfig.ScreenHeight / 12)
	//var border float32 = float32(config.GlobalConfig.ScreenHeight / 72) ORIGINAL

	var border float32 = float32(config.GlobalConfig.ScreenHeight / 110)
	var color = color.RGBA{88, 235, 243, 8.0}

	vector.DrawFilledRect(screen,
		X+border, Y-Y/2, X*16-border*2, border,
		color, false,
	)

	vector.DrawFilledRect(screen,
		X, Y-Y/2, border, Y*11+offset-Y/2+border,
		color, false,
	)
	vector.DrawFilledRect(screen,
		X*17-border, Y-Y/2, border, Y*11+offset-Y/2+border,
		color, false,
	)

	for i := 0; i < 11; i++ {
		vector.DrawFilledRect(screen,
			X+border, Y+float32(i)*Y+offset, X*16-border*2, border,
			color, false,
		)
	}
}

func Kubrick(name string) string {
	if name == "COMPUTER" {
		return "HAL 9000"
	} else {
		return "COMPUTER"
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
	case "gold":
		// RGBA: 255, 215, 0, 255
		return 1, 0.84, 0, 1
	case "silver":
		// RGBA: 192, 192, 192, 255
		return 0.75, 0.75, 0.75, 1
	case "bronze":
		// RGBA: 205, 127, 50, 255
		return 0.8, 0.5, 0.2, 1
	case "soft yellow":
		// RGBA: 255, 255, 204, 1.0
		return 1, 1, 0.8, 1
	default:
		log.Printf("Unknown color: %s", colorName)
		return 0, 0, 0, 255 // Default to black
	}
}
