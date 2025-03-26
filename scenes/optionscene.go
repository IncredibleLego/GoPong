package scenes

import (
	"goPong/config"
	"goPong/menu"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type OptionScene struct {
	playerName string
}

func NewOptionScene() *OptionScene {
	return &OptionScene{
		playerName: "",
	}
}

func (o *OptionScene) Draw(screen *ebiten.Image) {
	menu.ScreenDraw(config.GlobalConfig.TextDimension, 200, 200, 255, 255, 255, 255, 1.5, screen, "OPTIONS")
	menu.ScreenDraw(config.GlobalConfig.TextDimension, 120, 220, 255, 255, 255, 255, 1.5, screen, "Press enter to go back")
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

	return OptionsSceneId
}

var _ Scene = (*OptionScene)(nil)
