package scenes

import (
	"fmt"
	"goPong/config"
	"goPong/utils"
	"strconv"
	"strings"

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
		"Reset to default",
	}

	x1 := utils.XCentered("Options", config.GlobalConfig.TextDimension)
	x2 := utils.XCentered("Press enter to go back", config.GlobalConfig.TextDimension-5)
	utils.ScreenDraw(0, x1, 50, "white", screen, "OPTIONS")
	utils.ScreenDraw(-5, x2, 80, "white", screen, "Press enter to go back")

	for i, option := range options {
		x := utils.XCentered(option, config.GlobalConfig.TextDimension)
		if i == o.selectedOption {
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

func (o *OptionScene) FirstLoad() {}

func (o *OptionScene) OnEnter() {}

func (o *OptionScene) OnExit() {}

func (o *OptionScene) ShouldPreserveState(reason SceneChangeReason) bool {
	return false
}

func (o *OptionScene) Update() SceneId {
	optionsCount := 3 // Total number of options

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
	case 2:
		config.SaveConfig(config.DefaultConfig)
		config.InitConfig()
	}
}

var _ Scene = (*OptionScene)(nil)
