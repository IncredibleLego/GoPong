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
}

func (o *OptionMenu) Draw(screen *ebiten.Image) {
	x := utils.XCentered(o.MenuName, config.GlobalConfig.TextDimension)
	utils.ScreenDraw(0, x, 50, "white", screen, o.MenuName)

	for i, option := range o.Options {
		x = utils.XCentered(option, config.GlobalConfig.TextDimension)
		if i == o.Selected {
			j := strings.Index(option, ": ")
			if j > 0 {
				option = option[:j+2] + "◀" + option[j+2:] + "▶"
				x = x - 20
			}
			utils.ScreenDraw(0, x, float64(120+30*i-5), "cyan", screen, option)
		} else {
			utils.ScreenDraw(0, x, float64(120+30*i), "white", screen, option)
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
