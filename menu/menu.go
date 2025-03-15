package menu

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
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
		col := color.White
		if i == m.Selected {
			//color = color.NRGBA{255, 255, 0, 255} // Evidenzia l'opzione selezionata in giallo
			col = color.Black
			text.Draw(screen, option, basicfont.Face7x13, 100, 100+(i*30), col)
		} else {
			text.Draw(screen, option, basicfont.Face7x13, 100, 100+(i*30), col)
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
