package scenes

import (
	"fmt"
	"goPong/config"
	"goPong/menu"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type OptionScene struct {
	currentMenu        menu.Menu
	mainMenu           *menu.RegularMenu
	gameMenu           *menu.OptionMenu
	lastEnterPressTime time.Time
	actionExecuted     bool
	previousSceneId    SceneId
}

func NewOptionScene(previous SceneId) *OptionScene {
	return &OptionScene{
		currentMenu:     nil,
		mainMenu:        nil,
		gameMenu:        nil,
		previousSceneId: previous,
	}
}

func (o *OptionScene) generateGameMenuOptions() []string {
	return []string{
		"Text dimension: " + strconv.Itoa(int(config.GlobalConfig.TextDimension)),
		"Screen Height: " + strconv.Itoa(config.GlobalConfig.ScreenHeight),
		"Reset to default",
		"Back to options",
	}
}

func (o *OptionScene) Draw(screen *ebiten.Image) {

	//If selected menu is main menu print relative options
	//When changing ball and paddle dimension, print relative on the screen

	o.currentMenu.Draw(screen)
}

func (o *OptionScene) FirstLoad() {
	o.mainMenu = &menu.RegularMenu{
		Options: []menu.MenuOption{
			{Label: "GAME"},
			{Label: "SCREEN"},
			{Label: "GENERAL"},
			{Label: "BACK"},
		},
		Selected:     0,
		LastMoveTime: time.Now(),
	}
	o.gameMenu = &menu.OptionMenu{
		Options:      o.generateGameMenuOptions(),
		Selected:     0,
		LastMoveTime: time.Now(),
		MenuName:     "GAME OPTIONS",
	}
	o.currentMenu = o.mainMenu
	o.lastEnterPressTime = time.Now()
	o.actionExecuted = false
}

func (o *OptionScene) OnEnter() {}

func (o *OptionScene) OnExit() {}

func (o *OptionScene) ShouldPreserveState(reason SceneChangeReason) bool {
	return false
}

func (o *OptionScene) Update() SceneId {
	// Updates the current menu to print correctly the options
	o.gameMenu.Options = o.generateGameMenuOptions()
	nextMenu := o.currentMenu.Update()
	if nextMenu != nil {
		o.currentMenu = nextMenu
		o.lastEnterPressTime = time.Now() // Resetta il tempo per evitare input immediati
		o.actionExecuted = false
	} else {
		// Evita l'esecuzione immediata dopo il cambio menu
		if time.Since(o.lastEnterPressTime) > 200*time.Millisecond {
			// Controlla se il menu corrente è un OptionMenu
			if _, ok := o.currentMenu.(*menu.OptionMenu); ok {

				// Effettua un'asserzione di tipo per accedere a OptionMenu
				optionMenu, ok := o.currentMenu.(*menu.OptionMenu)
				if !ok {
					fmt.Println("Errore: currentMenu non è un OptionMenu")
				}

				// Ottieni l'opzione selezionata
				selectedOption := optionMenu.Selected

				// Modifica l'opzione selezionata
				if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
					handleOptionSelection(selectedOption, true)
				}

				if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
					handleOptionSelection(selectedOption, false)
				}

				// Torna al menu principale

				//If enter is pressed AND label = Back? universal
				if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && selectedOption == 3 {
					o.currentMenu = o.mainMenu
				}
			} else {
				// Controlla se Enter è stato premuto e non abbiamo già eseguito l'azione
				if (inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)) && !o.actionExecuted {
					o.actionExecuted = true // Evita che venga eseguito più volte

					// Gestisci il passaggio dal mainMenu a un OptionMenu
					if o.currentMenu == o.mainMenu {
						switch o.mainMenu.Selected {
						case 0: // Prima opzione del mainMenu
							o.currentMenu = o.gameMenu
							o.gameMenu.Selected = 0
						// Puoi aggiungere altri case per altre opzioni del mainMenu
						case 3:
							return o.previousSceneId
						default:
							// Gestione di default (se necessario)
						}
					}
				}
			}
		}
	}

	// Se Enter viene rilasciato, permetti nuove azioni
	if inpututil.KeyPressDuration(ebiten.KeyEnter) == 0 && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		o.actionExecuted = false
	}

	return OptionsSceneId
}

func handleOptionSelection(selectedOption int, mode bool) {
	// If mode is true = +, if false = -
	switch selectedOption {
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
