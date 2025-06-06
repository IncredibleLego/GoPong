package menu

import (
	"goPong/config"
	"goPong/utils"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MenuOption struct {
	Label   string
	SubMenu *RegularMenu
}

type RegularMenu struct {
	Options       []MenuOption
	Selected      int
	LastMoveTime  time.Time
	Offset        float64
	MainColor     string
	SelectedColor string
}

func (m *RegularMenu) Draw(screen *ebiten.Image) {
	for i, option := range m.Options {

		textWidth, textHeight := utils.MeasureText(m.Options[i].Label)

		x := float64(config.GlobalConfig.ScreenWidth)/2 - textWidth/2     // Center orizontally
		y := (float64(config.GlobalConfig.ScreenHeight) - textHeight) / 3 // Center vertically

		textDim := config.GlobalConfig.TextDimension
		spacing := textDim * 1.5

		if m.MainColor == "" {
			m.MainColor = "white"
		}
		if m.SelectedColor == "" {
			m.SelectedColor = "yellow"
		}

		if i == m.Selected {
			utils.ScreenDraw(0, x-(textDim), y+m.Offset+float64(i)*spacing-textDim/4, m.SelectedColor, screen, "◀"+option.Label+"▶")
		} else {
			utils.ScreenDraw(0, x, y+m.Offset+float64(i)*spacing, m.MainColor, screen, option.Label)
		}
	}
}

func (m *RegularMenu) Update() Menu {

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

	// Mouse management
	mouseX, mouseY := ebiten.CursorPosition()

	baseY := config.GlobalConfig.ScreenHeight / 3      // Starting Y position for the first option
	spacing := config.GlobalConfig.TextDimension * 1.5 // Spacing between options

	for i, option := range m.Options {
		textWidth, textHeight := utils.MeasureText(option.Label)
		x := (float64(config.GlobalConfig.ScreenWidth) - textWidth) / 2
		y := baseY + i*int(spacing) + int(m.Offset)

		// Check if the mouse is over the option
		if float64(mouseX) >= x && float64(mouseX) <= x+textWidth &&
			float64(mouseY) >= float64(y)-textHeight && float64(mouseY) <= float64(y) {

			m.Selected = i // Hilights the option under the mouse

			// If the left mouse button is pressed, select the option and return the submenu if it exists
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				if option.SubMenu != nil {
					return option.SubMenu
				}
			}
		}
	}
	return nil
}

var _ Menu = (*RegularMenu)(nil)
