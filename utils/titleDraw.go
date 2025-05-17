package utils

import (
	"goPong/config"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func TitleDraw(screen *ebiten.Image) {
	//Lettere 82 spazio 21 bordere 14

	// Draw Options

	//var X float32 = 21     // Starting X position
	var Y float32 = float32(config.GlobalConfig.ScreenHeight / 8)       // Y of the letters
	var border float32 = float32(config.GlobalConfig.ScreenHeight / 48) // Border size
	var space float32 = float32(config.GlobalConfig.ScreenHeight / 12)  // Space between letters
	var letterHeight float32 = float32(config.GlobalConfig.ScreenHeight / 6)
	var letterWidth float32 = float32(config.GlobalConfig.ScreenHeight / 8)

	//var titleColor1 = color.RGBA{240, 45, 60, 255}

	var titleColor1 = color.RGBA{77, 153, 255, 255}
	var titleColor2 = color.RGBA{0, 255, 255, 255}

	/*
		var Y float32 = 30     // Y of the letters
		var border float32 = 9 // Border size
		var space float32 = 21 // Space between letters
		var letterHeight float32 = 100
		var letterWidth float32 = 82
	*/

	// 82*6 (492) + 21*7 (147) = 639

	// Change colors to the letters to make more cool
	// Do the same thing but I need a new commit

	// Draw "G"
	vector.DrawFilledRect(screen,
		space, Y, border, letterHeight,
		titleColor1, false,
	)
	vector.DrawFilledRect(screen,
		space, Y, letterWidth, border,
		titleColor1, false,
	)
	vector.DrawFilledRect(screen,
		space, Y+letterHeight-border, letterWidth, border,
		titleColor1, false,
	)
	vector.DrawFilledRect(screen,
		space+letterWidth/2, Y+letterHeight/2-border/2, letterWidth/2, border,
		titleColor1, false,
	)
	vector.DrawFilledRect(screen,
		space+letterWidth-border, Y+letterHeight/2-border/2, border, letterHeight/2,
		titleColor1, false,
	)
	// Draw "O"
	vector.DrawFilledRect(screen,
		space*2+letterWidth, Y, border, letterHeight,
		titleColor2, false,
	)
	vector.DrawFilledRect(screen,
		space*2+letterWidth, Y, letterWidth, border,
		titleColor2, false,
	)
	vector.DrawFilledRect(screen,
		space*2+letterWidth*2-border, Y, border, letterHeight,
		titleColor2, false,
	)
	vector.DrawFilledRect(screen,
		space*2+letterWidth, Y+letterHeight-border, letterWidth, border,
		titleColor2, false,
	)
	// Draw "P"
	vector.DrawFilledRect(screen,
		space*3+letterWidth*2, Y, border, letterHeight,
		titleColor1, false,
	)
	vector.DrawFilledRect(screen,
		space*3+letterWidth*2, Y, letterWidth, border,
		titleColor1, false,
	)
	vector.DrawFilledRect(screen,
		space*3+letterWidth*2, Y+letterHeight/20*9, letterWidth, border,
		titleColor1, false,
	)
	vector.DrawFilledRect(screen,
		space*3+letterWidth*3-border, Y, border, letterHeight/20*9,
		titleColor1, false,
	)
	// Draw "O"
	vector.DrawFilledRect(screen,
		space*4+letterWidth*3, Y, border, letterHeight,
		titleColor2, false,
	)
	vector.DrawFilledRect(screen,
		space*4+letterWidth*3, Y, letterWidth, border,
		titleColor2, false,
	)
	vector.DrawFilledRect(screen,
		space*4+letterWidth*4-border, Y, border, letterHeight,
		titleColor2, false,
	)
	vector.DrawFilledRect(screen,
		space*4+letterWidth*3, Y+letterHeight-border, letterWidth, border,
		titleColor2, false,
	)
	// Draw "N"
	vector.DrawFilledRect(screen,
		space*5+letterWidth*4, Y, border, letterHeight,
		titleColor1, false,
	)
	vector.DrawFilledRect(screen,
		space*5+letterWidth*5-border, Y, border, letterHeight,
		titleColor1, false,
	)
	vector.StrokeLine(screen,
		space*5+letterWidth*4+border/2, Y,
		space*5+letterWidth*5-border, Y+letterHeight,
		border, titleColor1, false,
	)
	// Draw "G"
	vector.DrawFilledRect(screen,
		space*6+letterWidth*5, Y, border, letterHeight,
		titleColor2, false,
	)
	vector.DrawFilledRect(screen,
		space*6+letterWidth*5, Y, letterWidth, border,
		titleColor2, false,
	)
	vector.DrawFilledRect(screen,
		space*6+letterWidth*5, Y+letterHeight-border, letterWidth, border,
		titleColor2, false,
	)
	vector.DrawFilledRect(screen,
		space*6+letterWidth*5+letterWidth/2, Y+letterHeight/2-border/2, letterWidth/2, border,
		titleColor2, false,
	)
	vector.DrawFilledRect(screen,
		space*6+letterWidth*6-border, Y+letterHeight/2-border/2, border, letterHeight/2,
		titleColor2, false,
	)
	// Draw title
	ScreenDraw(-3, float64(config.GlobalConfig.ScreenWidth)/3.6, float64(Y+letterHeight+letterHeight/2), "sky blue", screen, "by IncredibleLego")
}
