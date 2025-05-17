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
	screenMenu         *menu.OptionMenu
	generalMenu        *menu.OptionMenu
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
		screenMenu:      nil,
		generalMenu:     nil,
		previousSceneId: previous,
	}
}

func (o *OptionScene) generateGameMenuOptions() []string {
	return []string{
		"Ball Speed: " + strconv.Itoa(config.GlobalConfig.BallSpeed),
		"Ball Size: " + strconv.Itoa(config.GlobalConfig.BallSize),
		"Paddle Speed: " + strconv.Itoa(config.GlobalConfig.PaddleSpeed),
		"Paddle Height: " + strconv.Itoa(config.GlobalConfig.PaddleHeight),
		"Paddle Width: " + strconv.Itoa(config.GlobalConfig.PaddleWidth),
		"Paddle Distance: " + strconv.Itoa(config.GlobalConfig.PaddleDistanceFromWall),
		"Enemy Difficulty: " + fmt.Sprintf("%.2f", config.GlobalConfig.Difficulty),
		"Reset to default",
		"Back to options",
	}
}

func (o *OptionScene) generateScreenMenuOptions() []string {
	return []string{
		"Text Dimension: " + strconv.Itoa(int(config.GlobalConfig.TextDimension)),
		"Screen Size: " + strconv.Itoa(config.GlobalConfig.ScreenWidth) + " x " + strconv.Itoa(config.GlobalConfig.ScreenHeight),
		"FullScreen: " + strconv.FormatBool(config.GlobalConfig.Fullscreen),
		"Reset to default",
		"Back to options",
	}
}

func (o *OptionScene) generateGeneralMenuOptions() []string {
	return []string{
		"Menu opt. per second: " + strconv.Itoa(int(config.GlobalConfig.MenuOptionsPerSecond)),
		"Options per second: " + strconv.Itoa(int(config.GlobalConfig.OptionsPerSecond)),
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
	case 3, 4, 5, 6:
		x := config.GlobalConfig.ScreenWidth - config.GlobalConfig.PaddleDistanceFromWall
		y := config.GlobalConfig.ScreenHeight/2 - config.GlobalConfig.PaddleHeight/2
		vector.DrawFilledRect(screen,
			float32(x), float32(y),
			float32(config.GlobalConfig.PaddleWidth), float32(config.GlobalConfig.PaddleHeight),
			color.White, false,
		)
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
		Position:     100,
	}
	o.screenMenu = &menu.OptionMenu{
		Options:      o.generateScreenMenuOptions(),
		Selected:     0,
		LastMoveTime: time.Now(),
		MenuName:     "SCREEN OPTIONS",
		Position:     150,
	}
	o.generalMenu = &menu.OptionMenu{
		Options:      o.generateGeneralMenuOptions(),
		Selected:     0,
		LastMoveTime: time.Now(),
		MenuName:     "GENERAL OPTIONS",
		Position:     200,
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
	o.screenMenu.Options = o.generateScreenMenuOptions()
	o.generalMenu.Options = o.generateGeneralMenuOptions()

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
						case 1:
							o.currentMenu = o.screenMenu
							o.screenMenu.Selected = 0
						case 2:
							o.currentMenu = o.generalMenu
							o.generalMenu.Selected = 0
						case 3:
							return o.previousSceneId
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
	case "SCREEN OPTIONS":
		handleScreenMenuOptions(o, optionMenu.Selected, mode)
	case "GENERAL OPTIONS":
		handleGeneralMenuOptions(o, optionMenu.Selected, mode)
	default:
		fmt.Println("Menu non riconosciuto:", optionMenu.MenuName)
	}
}

func updateConfigValue(option *int, min, max, step int, mode bool) {
	err := config.UpdateConfig(func(cfg *config.Config) {
		if mode && *option < max {
			*option += step
		} else if !mode && *option > min {
			*option -= step
		}
	})
	if err != nil {
		fmt.Println("Error during option saving", err)
	}
}

func updateConfigValueFloat(option *float64, min, max, step float64, mode bool) bool {
	originalValue := *option
	err := config.UpdateConfig(func(cfg *config.Config) {
		if mode && *option < max {
			*option += step
		} else if !mode && *option > min {
			*option -= step
		}
	})
	if err != nil {
		fmt.Println("Error during option saving", err)
	}
	return originalValue != *option
}

func handleGameMenuOptions(o *OptionScene, selectedOption int, mode bool) {
	// If mode is true = +, if false = -
	switch selectedOption {
	case 0:
		updateConfigValue(&config.GlobalConfig.BallSpeed, 1, 200, 1, mode)
	case 1:
		updateConfigValue(&config.GlobalConfig.BallSize, 5, 200, 5, mode)
	case 2:
		updateConfigValue(&config.GlobalConfig.PaddleSpeed, 1, 200, 1, mode)
	case 3:
		updateConfigValue(&config.GlobalConfig.PaddleHeight, 10, 470, 10, mode)
	case 4:
		updateConfigValue(&config.GlobalConfig.PaddleWidth, 5, config.GlobalConfig.PaddleDistanceFromWall-5, 5, mode)
	case 5:
		updateConfigValue(&config.GlobalConfig.PaddleDistanceFromWall, 15, config.GlobalConfig.ScreenWidth/2, 5, mode)
	case 6:
		updateConfigValueFloat(&config.GlobalConfig.Difficulty, 0.2, 0.9, 0.1, mode)
	case 7:
		err := config.UpdateConfig(func(cfg *config.Config) {
			scale := cfg.Scale
			cfg.BallSpeed = int(float64(config.DefaultConfig.BallSpeed) * scale)
			cfg.BallSize = int(float64(config.DefaultConfig.BallSize) * scale)
			cfg.PaddleSpeed = int(float64(config.DefaultConfig.PaddleSpeed) * scale)
			cfg.PaddleHeight = int(float64(config.DefaultConfig.PaddleHeight) * scale)
			cfg.PaddleWidth = int(float64(config.DefaultConfig.PaddleWidth) * scale)
			cfg.PaddleDistanceFromWall = int(float64(config.DefaultConfig.PaddleDistanceFromWall) * scale)
			cfg.Difficulty = config.DefaultConfig.Difficulty
		})
		if err != nil {
			fmt.Println("Error during option saving", err)
		}
	}
}

func handleScreenMenuOptions(o *OptionScene, selectedOption int, mode bool) {
	// If mode is true = +, if false = -
	switch selectedOption {
	case 0:
		updateConfigValueFloat(&config.GlobalConfig.TextDimension, 1, 35, 1, mode)
	case 1:
		if updateConfigValueFloat(&config.GlobalConfig.Scale, 0.67, 1.33, 0.33, mode) {
			err := config.ChangeScale(config.GlobalConfig.Scale)
			if err != nil {
				fmt.Println("Error during option saving", err)
			} else {
				// Show a message to restart the game or confirm
			}
		}
	case 2:
		err := config.UpdateConfig(func(cfg *config.Config) {
			cfg.Fullscreen = !cfg.Fullscreen
		})
		if err != nil {
			fmt.Println("Error during option saving", err)
		}
	case 3:
		err := config.UpdateConfig(func(cfg *config.Config) {
			cfg.TextDimension = config.DefaultConfig.TextDimension * config.GlobalConfig.Scale
			cfg.Scale = config.DefaultConfig.Scale
			cfg.Fullscreen = config.DefaultConfig.Fullscreen

		})
		if err != nil {
			fmt.Println("Error during option saving", err)
		}
	}
}

func handleGeneralMenuOptions(o *OptionScene, selectedOption int, mode bool) {
	// If mode is true = +, if false = -
	var err error

	switch selectedOption {
	case 0:
		err = config.UpdateConfig(func(cfg *config.Config) {
			if mode && cfg.MenuOptionsPerSecond < 35 {
				cfg.MenuOptionsPerSecond += 1
			} else if !mode && cfg.MenuOptionsPerSecond > 1 {
				cfg.MenuOptionsPerSecond -= 1
			}
		})
	case 1:
		err = config.UpdateConfig(func(cfg *config.Config) {
			if mode && cfg.OptionsPerSecond < 200 { //Random value, change
				cfg.OptionsPerSecond += 1
			} else if !mode && cfg.OptionsPerSecond > 0 {
				cfg.OptionsPerSecond -= 1
			}
		})
	case 2:
		err = config.UpdateConfig(func(cfg *config.Config) {
			cfg.MenuOptionsPerSecond = config.DefaultConfig.MenuOptionsPerSecond
			cfg.OptionsPerSecond = config.DefaultConfig.OptionsPerSecond
		})
	}

	if err != nil {
		fmt.Println("Error during option saving", err)
	}
}

var _ Scene = (*OptionScene)(nil)
