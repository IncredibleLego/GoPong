package menu

import (
	_ "embed"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Menu struct {
	Options  []string
	Selected int
}

func (m *Menu) Update() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		m.Selected--
		if m.Selected < 0 {
			m.Selected = len(m.Options) - 1
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		m.Selected++
		if m.Selected >= len(m.Options) {
			m.Selected = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		// Avvia la modalità selezionata
		startGame(m.Selected)
	}
}

func (m *Menu) Draw(screen *ebiten.Image) {
	for i, option := range m.Options {
		if i == m.Selected {
			ScreenDraw(13, 100, float64(100+(i*30)), 1, 1, 1, 1, 1.5, screen, option)
		} else {
			ScreenDraw(13, 100, float64(100+(i*30)), 0, 0, 0, 1, 1.5, screen, option)
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
