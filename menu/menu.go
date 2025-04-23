package menu

import "github.com/hajimehoshi/ebiten/v2"

type Menu interface {
	Update() Menu
	Draw(screen *ebiten.Image)
}
