package utils

import (
	"goPong/config"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func HighscoresTableDraw(screen *ebiten.Image) {

	var X float32 = float32(config.GlobalConfig.ScreenWidth / 18)
	var Y float32 = float32(config.GlobalConfig.ScreenHeight / 13)
	var offset float32 = float32(config.GlobalConfig.ScreenHeight / 12)
	//var border float32 = float32(config.GlobalConfig.ScreenHeight / 72) ORIGINAL

	var border float32 = float32(config.GlobalConfig.ScreenHeight / 110)
	var color = color.RGBA{255, 128, 0, 255}

	//var color = color.RGBA{102, 255, 102, 1.0}  LIME

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
