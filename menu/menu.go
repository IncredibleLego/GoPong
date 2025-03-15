package menu

import (
	_ "embed"
	"fmt"
	"goPong/constants"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Menu struct {
	Options      []string
	Selected     int
	LastMoveTime time.Time
}

func (m *Menu) Update() {

	moveInterval := time.Duration(time.Second / constants.MenuOptionsPerSecond)

	arrowUp := inpututil.KeyPressDuration(ebiten.KeyArrowUp)
	arrowDown := inpututil.KeyPressDuration(ebiten.KeyArrowDown)

	if arrowUp > 0 && time.Since(m.LastMoveTime) >= moveInterval {
		m.Selected--
		if m.Selected < 0 {
			m.Selected = len(m.Options) - 1
		}
		m.LastMoveTime = time.Now()
	}
	if arrowDown > 0 && time.Since(m.LastMoveTime) >= moveInterval {
		m.Selected++
		if m.Selected >= len(m.Options) {
			m.Selected = 0
		}
		m.LastMoveTime = time.Now()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// Avvia la modalità selezionata
		startGame(m.Selected)
	}

	// Controllo posizione mouse
	mouseX, mouseY := ebiten.CursorPosition()
	baseY := 100  // Posizione Y di partenza delle opzioni
	spacing := 30 // Spazio tra le opzioni

	for i := range m.Options {
		optionY := baseY + i*spacing
		textWidth, textHeight := MeasureText(m.Options[i], 1.5)
		textX := 100.0

		if float64(mouseX) >= textX && float64(mouseX) <= textX+textWidth && float64(mouseY) >= float64(optionY)-textHeight && float64(mouseY) <= float64(optionY) {
			m.Selected = i // Evidenzia l'opzione sotto il mouse
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				startGame(m.Selected)
			}
		}
	}
}

func (m *Menu) Draw(screen *ebiten.Image) {
	for i, option := range m.Options {
		if i == m.Selected {
			ScreenDraw(constants.TextDimension, 100, float64(100+(i*30)), 1, 1, 1, 1, 1.5, screen, option)
		} else {
			ScreenDraw(constants.TextDimension, 100, float64(100+(i*30)), 0, 0, 0, 1, 1.5, screen, option)
		}
	}
}

func startGame(mode int) {
	switch mode {
	case 0:
		fmt.Println("Avvio modalità Solo")
	case 1:
		fmt.Println("Avvio modalità Contro IA")
	case 2:
		fmt.Println("Avvio modalità Multiplayer")
	}
}
