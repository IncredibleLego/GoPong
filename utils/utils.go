package utils

import "goPong/config"

// This file contains utility functions for the game.

// Gives x coord to place a message in the middle of the screen given the message and the font size
func XCentered(message string, fontSize float64) float64 {
	width := float64(len(message)) * fontSize
	x := (float64(config.GlobalConfig.ScreenWidth) / 2) - (width / 2)
	return x
}
