package scenes

import (
	"fmt"
	"goPong/config"
	"goPong/menu"
	"image/color"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type OptionScene struct {
	currentMenu        menu.Menu
	mainMenu           *menu.RegularMenu
	gameMenu           *menu.OptionMenu
	lastEnterPressTime time.Time
	actionExecuted     bool
	previousSceneId    SceneId
	showOption         int
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
		"Ball Speed: " + strconv.Itoa(config.GlobalConfig.BallSpeed),
		"Ball Size: " + strconv.Itoa(config.GlobalConfig.BallSize),
		"Paddle Speed: " + strconv.Itoa(config.GlobalConfig.PaddleSpeed),
		"Paddle Height: " + strconv.Itoa(config.GlobalConfig.PaddleHeight),
		"Paddle Distance: " + strconv.Itoa(config.GlobalConfig.PaddleDistanceFromWall),
		"Enemy Difficulty: " + fmt.Sprintf("%.2f", config.GlobalConfig.Difficulty),
		"Reset to default",
		"Back to options",
	}
}

func (o *OptionScene) Draw(screen *ebiten.Image) {

	//If selected menu is main menu print relative options
	//When changing ball and paddle dimension, print relative on the screen

	switch o.showOption {
	case 1, 2:
		x := config.GlobalConfig.ScreenWidth/2 - config.GlobalConfig.BallSize/2
		vector.DrawFilledRect(screen,
			float32(x), float32(400-config.GlobalConfig.BallSize/2),
			float32(config.GlobalConfig.BallSize), float32(config.GlobalConfig.BallSize),
			color.White, false,
		)
	case 3, 4, 5:
		x := config.GlobalConfig.ScreenWidth - config.GlobalConfig.PaddleDistanceFromWall
		y := config.GlobalConfig.ScreenHeight/2 - config.GlobalConfig.PaddleHeight/2
		vector.DrawFilledRect(screen,
			float32(x), float32(y),
			float32(15), float32(config.GlobalConfig.PaddleHeight),
			color.White, false,
		)
	default:
		_ = 0
	}

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

				moveInterval := time.Duration(time.Second / config.GlobalConfig.OptionsPerSecond)

				arrowRight := inpututil.KeyPressDuration(ebiten.KeyArrowRight)
				keyD := inpututil.KeyPressDuration(ebiten.KeyD)

				arrowLeft := inpututil.KeyPressDuration(ebiten.KeyArrowLeft)
				keyA := inpututil.KeyPressDuration(ebiten.KeyA)

				// Effettua un'asserzione di tipo per accedere a OptionMenu

				optionMenu, ok := o.currentMenu.(*menu.OptionMenu)
				if !ok {
					fmt.Println("Errore: currentMenu non è un OptionMenu")
				}

				// Make it print only if the menu is the gameMenu
				if optionMenu.MenuName == "GAME OPTIONS" {
					o.showOption = optionMenu.Selected + 1
				}

				// Modifica l'opzione selezionata
				/*
					if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
						handleOptionSelection(o, true)
					}

					if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
						handleOptionSelection(o, false)
					} */

				if (arrowRight > 0 || keyD > 0) && time.Since(o.lastEnterPressTime) >= moveInterval {
					handleOptionSelection(o, true)
					o.lastEnterPressTime = time.Now()
				}
				if (arrowLeft > 0 || keyA > 0) && time.Since(o.lastEnterPressTime) >= moveInterval {
					handleOptionSelection(o, false)
					o.lastEnterPressTime = time.Now()
				}

				// Torna al menu principale

				//If enter is pressed AND label = Back? universal
				if inpututil.IsKeyJustPressed(ebiten.KeyEnter) && optionMenu.Selected == len(optionMenu.Options)-1 {
					o.currentMenu = o.mainMenu
					o.showOption = 0
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

func handleOptionSelection(o *OptionScene, mode bool) {

	optionMenu, ok := o.currentMenu.(*menu.OptionMenu)
	if !ok {
		fmt.Println("Errore: currentMenu non è un OptionMenu")
	}

	// Gestisci le opzioni in base al menu corrente
	switch optionMenu.MenuName {
	case "GAME OPTIONS":
		handleGameMenuOptions(o, optionMenu.Selected, mode)
	// Aggiungi altri case per altri menu
	default:
		fmt.Println("Menu non riconosciuto:", optionMenu.MenuName)
	}
}

func handleGameMenuOptions(o *OptionScene, selectedOption int, mode bool) {
	// If mode is true = +, if false = -
	var err error

	switch selectedOption {
	case 0:
		err = config.UpdateConfig(func(cfg *config.Config) {
			if mode {
				cfg.BallSpeed += 1
			} else {
				cfg.BallSpeed -= 1
			}
		})
	case 1:
		err = config.UpdateConfig(func(cfg *config.Config) {
			if mode && cfg.BallSize < 200 {
				cfg.BallSize += 5
			} else if !mode && cfg.BallSize > 5 {
				cfg.BallSize -= 5
			}
		})
	case 2:
		err = config.UpdateConfig(func(cfg *config.Config) {
			if mode {
				cfg.PaddleSpeed += 1
			} else {
				cfg.PaddleSpeed -= 1
			}
		})
	case 3:
		err = config.UpdateConfig(func(cfg *config.Config) {
			if mode && cfg.PaddleHeight < 470 {
				cfg.PaddleHeight += 10
			} else if !mode && cfg.PaddleHeight > 10 {
				cfg.PaddleHeight -= 10
			}
		})
	case 4:
		err = config.UpdateConfig(func(cfg *config.Config) {
			if mode {
				cfg.PaddleDistanceFromWall += 5
			} else {
				cfg.PaddleDistanceFromWall -= 5
			}
		})
	case 5:
		err = config.UpdateConfig(func(cfg *config.Config) {
			if mode {
				cfg.Difficulty += 0.1
			} else {
				cfg.Difficulty -= 0.1
			}
		})
	case 6:
		config.SaveConfig(config.DefaultConfig)
		config.InitConfig()
	}

	if err != nil {
		fmt.Println("Error during option saving", err)
	}
}

var _ Scene = (*OptionScene)(nil)
