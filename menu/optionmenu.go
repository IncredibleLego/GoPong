package menu

import (
	"goPong/config"
	"goPong/utils"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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

func (o *OptionMenu) Update() {
	// Implement
}
