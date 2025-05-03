package utils

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

//Funzione che dato il punteggio lo disegna
//Seconda funzione con i punteggi disegnati da stampare tra cui selezionare
// In file principale: utils.PointsDraw(screen, X, Y, score)

func PointsDraw(screen *ebiten.Image, X, Y float32, number int) {

	var border float32 = 9 // Border size
	//var space float32 = 40 // Space between letters
	var numHeight float32 = 80
	var numWidth float32 = 60

	switch number {
	case 0:
		// Draw "0"
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
		// Draw "1"
		vector.DrawFilledRect(screen,
			X+numWidth/2, Y, border, numHeight,
			color.White, false,
		)
	case 2:
		// Draw "2"
		vector.DrawFilledRect(screen,
			X, Y, numWidth, border,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X+numWidth-border, Y, border, numHeight/2-border/2,
			color.White, false,
		)
		vector.DrawFilledRect(screen,
			X, Y+numHeight/2-border/2, numWidth, border,
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

	}
}
