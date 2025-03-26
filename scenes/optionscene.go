package scenes

import (
	"goPong/config"
	"goPong/menu"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type OptionScene struct {
}

func NewOptionScene() *OptionScene {
	return &OptionScene{}
}

func (o *OptionScene) Draw(screen *ebiten.Image) {
	menu.ScreenDraw(config.GlobalConfig.TextDimension, 200, 200, 255, 255, 255, 255, 1.5, screen, "OPTIONS")
	menu.ScreenDraw(config.GlobalConfig.TextDimension, 120, 220, 255, 255, 255, 255, 1.5, screen, "Press enter to go back")
	menu.ScreenDraw(config.GlobalConfig.TextDimension, 120, 240, 255, 255, 255, 255, 1.5, screen, "Screen Width: "+strconv.Itoa(config.GlobalConfig.TextDimension))
	menu.ScreenDraw(config.GlobalConfig.TextDimension, 120, 260, 255, 255, 255, 255, 1.5, screen, "Screen Height: "+strconv.Itoa(config.GlobalConfig.ScreenHeight))
}

func (o *OptionScene) FirstLoad() {

}

func (o *OptionScene) OnEnter() {

}

func (o *OptionScene) OnExit() {

}

func (o *OptionScene) ShouldPreserveState(reason SceneChangeReason) bool {
	return false
}

func (o *OptionScene) Update() SceneId {

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return StartSceneId
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		config.GlobalConfig.ScreenHeight += 10
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		config.GlobalConfig.ScreenHeight -= 10
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		config.GlobalConfig.TextDimension += 10
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		config.GlobalConfig.TextDimension -= 10
	}

	return OptionsSceneId
}

var _ Scene = (*OptionScene)(nil)
