package utils

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func TitleDraw(screen *ebiten.Image) {
	//Lettere 82 spazio 21 bordere 14

	// Draw Options

	var X float32 = 21     // Starting X position
	var Y float32 = 30     // Y of the letters
	var border float32 = 9 // Border size
	var space float32 = 21 // Space between letters

	// Draw "G"
	vector.DrawFilledRect(screen,
		X, Y, border, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21, Y, 82, border,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21, Y+100-border, 82, border,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+41, Y+50-border/2, 41, border,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82-border, Y+50-border/2, border, 50,
		color.White, false,
	)
	// Draw "O"
	vector.DrawFilledRect(screen,
		21+82+space, Y, border, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space, Y, 82, border,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82-border, Y, border, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space, Y+100-border, 82, border,
		color.White, false,
	)
	// Draw "P"
	vector.DrawFilledRect(screen,
		21+82+space+82+space, Y, border, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space, Y, 82, border,
		color.White, false,
	)

	vector.DrawFilledRect(screen,
		21+82+space+82+space, Y+45, 82, border,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82-border, Y, border, 45,
		color.White, false,
	)
	// Draw "O"
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space, Y, border, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space, Y, 82, border,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82-border, Y, border, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space, Y+100-border, 82, border,
		color.White, false,
	)
	// Draw "N"
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82+space, Y, border, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82+space+82-border, Y, border, 100,
		color.White, false,
	)
	vector.StrokeLine(screen,
		21+82+space+82+space+82+space+82+space+border/2, Y,
		21+82+space+82+space+82+space+82+space+82-border, Y+100,
		border, color.White, false,
	)

	// Draw "G"
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82+space+82+space, Y, border, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82+space+82+space, Y, 82, border,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82+space+82+space, Y+100-border, 82, border,
		color.White, false,
	)

	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82+space+82+space+41, Y+50-border/2, 41, border,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82+space+82+space+82-border, Y+50-border/2, border, 50,
		color.White, false,
	)

}
