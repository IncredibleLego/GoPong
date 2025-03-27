package scenes

import (
	"goPong/config"
	"goPong/menu"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type OptionScene struct {
	selectedOption int // Indice dell'opzione attualmente selezionata
}

func NewOptionScene() *OptionScene {
	return &OptionScene{}
}

func (o *OptionScene) Draw(screen *ebiten.Image) {
	options := []string{
		"Screen Width: " + strconv.Itoa(config.GlobalConfig.TextDimension),
		"Screen Height: " + strconv.Itoa(config.GlobalConfig.ScreenHeight),
	}

	menu.ScreenDraw(config.GlobalConfig.TextDimension, 200, 200, 255, 255, 255, 255, 1.5, screen, "OPTIONS")
	menu.ScreenDraw(config.GlobalConfig.TextDimension, 120, 220, 255, 255, 255, 255, 1.5, screen, "Press enter to go back")

	for i, option := range options {
		if i == o.selectedOption {
			menu.ScreenDraw(config.GlobalConfig.TextDimension, 120, float64(240+20*i), 255, 255, 0, 255, 1.5, screen, option)
		} else {
			menu.ScreenDraw(config.GlobalConfig.TextDimension, 120, float64(240+20*i), 255, 255, 255, 255, 1.5, screen, option)
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
	optionsCount := 2 // Numero totale di opzioni

	// Navigazione verticale (freccia su/giù o tasti W/S)
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		o.selectedOption = (o.selectedOption - 1 + optionsCount) % optionsCount
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		o.selectedOption = (o.selectedOption + 1) % optionsCount
	}

	// Modifica dell'opzione selezionata (freccia sinistra/destra o tasti A/D)
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		switch o.selectedOption {
		case 0: // Modifica Screen Width
			config.GlobalConfig.TextDimension += 10
		case 1: // Modifica Screen Height
			config.GlobalConfig.ScreenHeight += 10
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		switch o.selectedOption {
		case 0: // Modifica Screen Width
			config.GlobalConfig.TextDimension -= 10
		case 1: // Modifica Screen Height
			config.GlobalConfig.ScreenHeight -= 10
		}
	}

	// Ritorno alla scena iniziale con il tasto Enter
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return StartSceneId
	}

	return OptionsSceneId
}

var _ Scene = (*OptionScene)(nil)
