package utils

import (
	"goPong/config"
	"log"
)

// This file contains utility functions for the game.

// Gives x coord to place a message in the middle of the screen given the message and the font size
func XCentered(message string, fontSize float64) float64 {
	width := float64(len(message)) * fontSize
	x := (float64(config.GlobalConfig.ScreenWidth) / 2) - (width / 2)
	return x
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
