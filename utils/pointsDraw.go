package utils

import (
	"goPong/config"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func PointsDraw(screen *ebiten.Image, X, Y float32, score int) {
	// Draw the points
	printPoint(screen, X+70*float32(config.GlobalConfig.Scale), Y, score%10)
	if score/10 != 0 {
		printPoint(screen, X, Y, score/10)
	}
}

func printPoint(screen *ebiten.Image, X, Y float32, number int) {

	var border float32 = 9 * float32(config.GlobalConfig.Scale)
	var numHeight float32 = 80 * float32(config.GlobalConfig.Scale)
	var numWidth float32 = 40 * float32(config.GlobalConfig.Scale)

	switch number {
	case 0:
		vector.DrawFilledRect(screen,
			X, Y, numWidth, border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y, border, numHeight,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight-border, numWidth, border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X+numWidth-border, Y, border, numHeight,
			color.White, false,
		)
	case 1:
		vector.DrawFilledRect(screen,
			X+numWidth/2, Y, border, numHeight,
			color.White, false,
		)
	case 2:
		vector.DrawFilledRect(screen,
			X, Y, numWidth, border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X+numWidth-border, Y, border, numHeight/2-border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight/2-border, numWidth, border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight/2, border, numHeight/2,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight-border, numWidth, border,
			color.White, false,
		)
	case 3:
		vector.DrawFilledRect(screen,
			X, Y, numWidth, border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X+numWidth-border, Y, border, numHeight,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight/2-border, numWidth, border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight-border, numWidth, border,
			color.White, false,
		)
	case 4:
		vector.DrawFilledRect(screen,
			X, Y, border, numHeight/2-border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X+numWidth-border, Y, border, numHeight,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight/2-border, numWidth, border,
			color.White, false,
		)
	case 5:
		vector.DrawFilledRect(screen,
			X, Y, numWidth, border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y, border, numHeight/2-border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight/2-border, numWidth, border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X+numWidth-border, Y+numHeight/2, border, numHeight/2,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight-border, numWidth, border,
			color.White, false,
		)
	case 6:
		vector.DrawFilledRect(screen,
			X, Y, border, numHeight,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight/2-border, numWidth, border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X+numWidth-border, Y+numHeight/2, border, numHeight/2,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight-border, numWidth, border,
			color.White, false,
		)
	case 7:
		vector.DrawFilledRect(screen,
			X, Y, numWidth, border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X+numWidth-border, Y, border, numHeight,
			color.White, false,
		)
	case 8:
		vector.DrawFilledRect(screen,
			X, Y, numWidth, border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y, border, numHeight,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X+numWidth-border, Y, border, numHeight,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight/2-border, numWidth, border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight-border, numWidth, border,
			color.White, false,
		)
	case 9:
		vector.DrawFilledRect(screen,
			X, Y, numWidth, border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y, border, numHeight/2,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X+numWidth-border, Y, border, numHeight,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight/2-border, numWidth, border,
			color.White, false,
		)
	}
}
