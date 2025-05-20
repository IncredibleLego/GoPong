package utils

import (
	"goPong/config"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Popup should be a structure? How to add the options that will activate during the popup?
// Popup graphics choose. Different popups for different situations?

func Popup(screen *ebiten.Image, text string) {

	// Create a new image for the popup
	popupWidth := 300
	popupHeight := 100
	popupImage := ebiten.NewImage(popupWidth, popupHeight) // Should not call this every draw -> inefficient

	X := float64(config.GlobalConfig.ScreenWidth/2 - popupWidth/2)
	Y := float64(config.GlobalConfig.ScreenHeight/2 - popupHeight/2)

	// Fill the popup with a color (e.g., white)
	popupImage.Fill(color.White)

	// Draw the popup on the screen at a specific position
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(X, Y)
	screen.DrawImage(popupImage, op)

	// Draw the text on the popup
	ScreenDraw(-5, X+float64(popupWidth)/10, Y+float64(popupHeight)/3, "black", screen, text)
}
