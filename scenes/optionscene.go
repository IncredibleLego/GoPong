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
	selectedOption int // Indice dell'opzione attualmente selezionata
}

func NewOptionScene() *OptionScene {
	return &OptionScene{}
}

func (o *OptionScene) Draw(screen *ebiten.Image) {
	options := []string{
		"Screen Width: " + strconv.Itoa(config.GlobalConfig.ScreenWidth),
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
			err := config.UpdateConfig(func(cfg *config.Config) {
				cfg.TextDimension += 10
				fmt.Println("Screen Width:", cfg.TextDimension)
			})
			if err != nil {
				// Gestisci l'errore
				fmt.Println("Errore durante il salvataggio della configurazione:", err)
			}
		case 1: // Modifica Screen Height
			config.GlobalConfig.ScreenHeight += 10
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		switch o.selectedOption {
		case 0: // Modifica Screen Width

			err := config.UpdateConfig(func(cfg *config.Config) {
				cfg.TextDimension -= 10
				fmt.Println("Screen Width:", cfg.TextDimension)
			})
			if err != nil {
				// Gestisci l'errore
				fmt.Println("Errore durante il salvataggio della configurazione:", err)
			}
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
