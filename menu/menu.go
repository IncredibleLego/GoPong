package menu

import (
	_ "embed"
	"goPong/constants"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MenuOption struct {
	Label   string
	SubMenu *Menu
}

type Menu struct {
	Options      []MenuOption
	Selected     int
	LastMoveTime time.Time
}

func (m *Menu) Update() *Menu {

	// moveInterval could be a constant
	moveInterval := time.Duration(time.Second / constants.MenuOptionsPerSecond)

	arrowUp := inpututil.KeyPressDuration(ebiten.KeyArrowUp)
	keyW := inpututil.KeyPressDuration(ebiten.KeyW)

	arrowDown := inpututil.KeyPressDuration(ebiten.KeyArrowDown)
	keyS := inpututil.KeyPressDuration(ebiten.KeyS)

	if (arrowUp > 0 || keyW > 0) && time.Since(m.LastMoveTime) >= moveInterval {
		m.Selected--
		if m.Selected < 0 {
			m.Selected = len(m.Options) - 1
		}
		m.LastMoveTime = time.Now()
	}
	if (arrowDown > 0 || keyS > 0) && time.Since(m.LastMoveTime) >= moveInterval {
		m.Selected++
		if m.Selected >= len(m.Options) {
			m.Selected = 0
		}
		m.LastMoveTime = time.Now()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if m.Options[m.Selected].SubMenu != nil {
			return m.Options[m.Selected].SubMenu
		}
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
				if m.Options[m.Selected].SubMenu != nil {
					return m.Options[m.Selected].SubMenu
				}
			}
		}
	}
	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	for i, option := range m.Options {

		textWidth, textHeight := MeasureText(m.Options[i], 1.5)

		x := (constants.ScreenWidth - textWidth) / 2   // Centra orizzontalmente
		y := (constants.ScreenHeight - textHeight) / 3 // Centra verticalmente

		if i == m.Selected {
			ScreenDraw(constants.TextDimension, x, y+float64(i*30), 1, 1, 0, 1, 1.5, screen, option.Label)
		} else {
			ScreenDraw(constants.TextDimension, x, y+float64(i*30), 1, 1, 1, 1, 1.5, screen, option.Label)
		}
	}
}
