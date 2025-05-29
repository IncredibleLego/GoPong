package menu

import (
	"goPong/config"
	"goPong/utils"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type OptionMenu struct {
	MenuName     string
	Options      []string
	Selected     int
	LastMoveTime time.Time
	Position     float64
}

func (o *OptionMenu) Draw(screen *ebiten.Image) {
	textDim := config.GlobalConfig.TextDimension
	spacing := textDim * 1.5

	x := utils.XCentered(o.MenuName, textDim)
	utils.ScreenDraw(0, x, float64(o.Position)-textDim/0.4, "white", screen, o.MenuName)

	for i, option := range o.Options {
		x = utils.XCentered(option, textDim)
		if i == o.Selected {
			j := strings.Index(option, ": ")
			if j > 0 {
				option = option[:j+2] + "◀" + option[j+2:] + "▶"
				x = x - textDim
			}
			utils.ScreenDraw(0, x, o.Position+float64(i)*spacing-textDim/4, "cyan", screen, option)
		} else {
			utils.ScreenDraw(0, x, o.Position+float64(i)*spacing, "white", screen, option)
		}
	}
}

func (o *OptionMenu) Update() Menu {

	// moveInterval could be a constant
	moveInterval := time.Duration(time.Second / config.GlobalConfig.MenuOptionsPerSecond)

	arrowUp := inpututil.KeyPressDuration(ebiten.KeyArrowUp)
	keyW := inpututil.KeyPressDuration(ebiten.KeyW)

	arrowDown := inpututil.KeyPressDuration(ebiten.KeyArrowDown)
	keyS := inpututil.KeyPressDuration(ebiten.KeyS)

	if (arrowUp > 0 || keyW > 0) && time.Since(o.LastMoveTime) >= moveInterval {
		o.Selected--
		if o.Selected < 0 {
			o.Selected = len(o.Options) - 1
		}
		o.LastMoveTime = time.Now()
	}
	if (arrowDown > 0 || keyS > 0) && time.Since(o.LastMoveTime) >= moveInterval {
		o.Selected++
		if o.Selected >= len(o.Options) {
			o.Selected = 0
		}
		o.LastMoveTime = time.Now()
	}

	/*
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			if o.Options[o.Selected].SubMenu != nil {
				return o.Options[o.Selected].SubMenu
			}
		} */

	// **Gestione del mouse**
	mouseX, mouseY := ebiten.CursorPosition()

	baseY := config.GlobalConfig.ScreenHeight / 3 // Punto di partenza delle opzioni (centrato)
	spacing := 30                                 // Spazio tra le opzioni

	for i, option := range o.Options {
		textWidth, textHeight := utils.MeasureText(option)
		x := (float64(config.GlobalConfig.ScreenWidth) - textWidth) / 2
		y := baseY + i*spacing

		// Controllo se il mouse è sopra il testo
		if float64(mouseX) >= x && float64(mouseX) <= x+textWidth &&
			float64(mouseY) >= float64(y)-textHeight && float64(mouseY) <= float64(y) {

			o.Selected = i // Evidenzia l'opzione sotto il mouse

			// Se il tasto sinistro viene premuto, cambia menu o esegui azione
			/*if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				if option.SubMenu != nil {
					return option.SubMenu
				}
			} */
		}
	}
	return nil
}

var _ Menu = (*OptionMenu)(nil)
