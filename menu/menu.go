package menu

import (
	_ "embed"
	"goPong/config"
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
	moveInterval := time.Duration(time.Second / config.GlobalConfig.MenuOptionsPerSecond)

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

	// **Gestione del mouse**
	mouseX, mouseY := ebiten.CursorPosition()

	baseY := config.GlobalConfig.ScreenHeight / 3 // Punto di partenza delle opzioni (centrato)
	spacing := 30                                 // Spazio tra le opzioni

	for i, option := range m.Options {
		textWidth, textHeight := MeasureText(option)
		x := (float64(config.GlobalConfig.ScreenWidth) - textWidth) / 2
		y := baseY + i*spacing

		// Controllo se il mouse è sopra il testo
		if float64(mouseX) >= x && float64(mouseX) <= x+textWidth &&
			float64(mouseY) >= float64(y)-textHeight && float64(mouseY) <= float64(y) {

			m.Selected = i // Evidenzia l'opzione sotto il mouse

			// Se il tasto sinistro viene premuto, cambia menu o esegui azione
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				if option.SubMenu != nil {
					return option.SubMenu
				}
			}
		}
	}
	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	for i, option := range m.Options {

		textWidth, textHeight := MeasureText(m.Options[i])

		x := float64(config.GlobalConfig.ScreenWidth)/2 - textWidth/2     // Centra orizzontalmente
		y := (float64(config.GlobalConfig.ScreenHeight) - textHeight) / 3 // Centra verticalmente

		spacing := float64(config.GlobalConfig.TextDimension) * 1.5

		if i == m.Selected {
			//ScreenDraw(config.GlobalConfig.TextDimension, x, y+float64(i*30), 1, 1, 0, 1, 1.5, screen, "◀"+option.Label+"▶")
			//ScreenDraw(config.GlobalConfig.TextDimension, x, y+float64(i*30), 1, 1, 0, 1, 1.5, screen, option.Label)
			ScreenDraw(config.GlobalConfig.TextDimension, x, y+float64(i)*spacing-5, 1, 1, 0, 1, 1.5, screen, option.Label)
		} else {
			//ScreenDraw(config.GlobalConfig.TextDimension, x, y+float64(i*30), 1, 1, 1, 1, 1.5, screen, option.Label)
			ScreenDraw(config.GlobalConfig.TextDimension, x, y+float64(i)*spacing, 1, 1, 1, 1, 1.5, screen, option.Label)
		}
	}
}
