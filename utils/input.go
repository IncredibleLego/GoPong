package utils

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var lastMoveTime time.Time

// Input handles keyboard input for a given string.
func Input(inputString *string, maxChar int) {
	inputText := *inputString
	moveInterval := time.Duration(time.Second / 8)
	maxLength := len(inputText) >= maxChar

	// Backspace
	if len(inputText) > 0 {
		if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			inputText = inputText[:len(inputText)-1]
			lastMoveTime = time.Now()

		} else if inpututil.KeyPressDuration(ebiten.KeyBackspace) > 30 && time.Since(lastMoveTime) >= moveInterval {
			inputText = inputText[:len(inputText)-1]
			lastMoveTime = time.Now()
		}
	}

	// Handle key input
	handleKeyInput := func(startKey, endKey ebiten.Key, offset rune) {
		for key := startKey; key <= endKey; key++ {
			if inpututil.IsKeyJustPressed(key) {
				if !maxLength {
					inputText += string(offset + rune(key-startKey))
					lastMoveTime = time.Now()
				}
			} else if inpututil.KeyPressDuration(key) > 30 && time.Since(lastMoveTime) >= moveInterval {
				if !maxLength {
					inputText += string(offset + rune(key-startKey))
					lastMoveTime = time.Now()
				}
			}
		}
	}

	// Alphabetical characters (A-Z)
	handleKeyInput(ebiten.KeyA, ebiten.KeyZ, 'A')

	// Numerical characters (0-9)
	handleKeyInput(ebiten.Key0, ebiten.Key9, '0')

	*inputString = inputText
}
