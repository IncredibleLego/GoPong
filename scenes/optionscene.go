package scenes

import (
	"fmt"
	"goPong/config"
	"goPong/menu"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type OptionScene struct {
	selectedOption  int
	previousSceneId SceneId
}

func NewOptionScene(previous SceneId) *OptionScene {
	return &OptionScene{
		selectedOption:  0,
		previousSceneId: previous,
	}
}

func (o *OptionScene) Draw(screen *ebiten.Image) {
	options := []string{
		"Text dimension: " + strconv.Itoa(int(config.GlobalConfig.TextDimension)),
		"Screen Height: " + strconv.Itoa(config.GlobalConfig.ScreenHeight),
	}

	menu.ScreenDraw(0, 200, 200, "white", screen, "OPTIONS")
	menu.ScreenDraw(0, 120, 220, "white", screen, "Press enter to go back")

	for i, option := range options {
		if i == o.selectedOption {
			menu.ScreenDraw(0, 120, float64(240+20*i), "yellow", screen, option)
		} else {
			menu.ScreenDraw(0, 120, float64(240+20*i), "white", screen, option)
		}
	}
}

func (o *OptionScene) FirstLoad() {}

func (o *OptionScene) OnEnter() {}

func (o *OptionScene) OnExit() {}

func (o *OptionScene) ShouldPreserveState(reason SceneChangeReason) bool {
	return false
}

func (o *OptionScene) Update() SceneId {
	optionsCount := 2 // Total number of options

	// Selecting the option
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		o.selectedOption = (o.selectedOption - 1 + optionsCount) % optionsCount
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		o.selectedOption = (o.selectedOption + 1) % optionsCount
	}

	// Modify the selected option
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		handleOptionSelection(o, true)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		handleOptionSelection(o, false)
	}

	// Return to the main menu
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return o.previousSceneId
	}

	return OptionsSceneId
}

func handleOptionSelection(o *OptionScene, mode bool) {
	// If mode is true = +, if false = -
	switch o.selectedOption {
	case 0:
		err := config.UpdateConfig(func(cfg *config.Config) {
			if mode {
				cfg.TextDimension += 10
			} else {
				cfg.TextDimension -= 10
			}
		})
		if err != nil {
			fmt.Println("Error during option saving", err)
		}
	case 1:
		err := config.UpdateConfig(func(cfg *config.Config) {
			if mode {
				cfg.ScreenHeight += 10
			} else {
				cfg.ScreenHeight -= 10
			}
		})
		if err != nil {
			fmt.Println("Error during option saving", err)
		}
	}
}

var _ Scene = (*OptionScene)(nil)
