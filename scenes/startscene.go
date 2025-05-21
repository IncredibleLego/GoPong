package scenes

import (
	"fmt"
	"goPong/menu"
	"goPong/utils"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type StartScene struct { // is the scene loaded now
	currentMenu        *menu.RegularMenu
	mainMenu           *menu.RegularMenu
	playMenu           *menu.RegularMenu
	exitPopup          *utils.Popup
	lastEnterPressTime time.Time
	actionExecuted     bool
	selectedMode       int
}

func (s *StartScene) ShouldPreserveState(reason SceneChangeReason) bool {
	return false
}

func NewStartScene() *StartScene {
	return &StartScene{
		currentMenu: nil,
		mainMenu:    nil,
		playMenu:    nil,
	}
}

func (s *StartScene) Draw(screen *ebiten.Image) {
	//Letter 82 space 21 dimension 14

	utils.TitleDraw(screen)

	s.currentMenu.Draw(screen)

	if s.exitPopup.Active {
		s.exitPopup.Draw(screen)
	}
}

func (s *StartScene) FirstLoad() {
	s.mainMenu = &menu.RegularMenu{
		Options: []menu.MenuOption{
			{Label: "PLAY"},
			{Label: "OPTIONS"},
			{Label: "HIGHSCORES"},
			{Label: "CREDITS"},
			{Label: "QUIT"},
		},
		Selected:     0,
		LastMoveTime: time.Now(),
		Offset:       100,
	}
	s.playMenu = &menu.RegularMenu{
		Options: []menu.MenuOption{
			{Label: "SOLO MODE"},
			{Label: "COMPUTER MODE"},
			{Label: "MULTIPLAYER MODE"},
			{Label: "BACK"},
		},
		Selected:     0,
		LastMoveTime: time.Now(),
		Offset:       100,
	}
	s.exitPopup = &utils.Popup{
		Active:  false,
		Text:    "Are you sure you want to quit?",
		Options: []string{"YES", "NO"},
	}
	s.currentMenu = s.mainMenu
	s.lastEnterPressTime = time.Now()
	s.actionExecuted = false
}

/*
popupWidth := int(float64(config.GlobalConfig.ScreenWidth) * 0.4)  // 40% della larghezza
popupHeight := int(float64(config.GlobalConfig.ScreenHeight) * 0.2) // 20% dell'altezza
*/

func (s *StartScene) OnEnter() {

}

func (s *StartScene) OnExit() {

}

func (s *StartScene) Update() SceneId {
	if s.exitPopup.Active {
		//fmt.Println("ExitPopup Selected:", s.exitPopup.Selected)
		s.exitPopup.Update()
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			id := s.handleExitPopup()
			return id
		}
	} else {
		nextMenu := s.currentMenu.Update()
		if nextMenu != nil {
			if regularMenu, ok := nextMenu.(*menu.RegularMenu); ok {
				s.currentMenu = regularMenu
			} else {
				fmt.Println("Error: nextMenu is not of type *menu.RegularMenu")
			}
			s.lastEnterPressTime = time.Now() // Resetta il tempo per evitare input immediati
			s.actionExecuted = false
		} else {
			// Evita l'esecuzione immediata dopo il cambio menu
			if time.Since(s.lastEnterPressTime) > 200*time.Millisecond {
				// Controlla se Enter è stato premuto e non abbiamo già eseguito l'azione
				if (inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)) && !s.actionExecuted {
					id := s.handleMenuSelection()
					s.actionExecuted = true // Evita che venga eseguito più volte
					if id != StartSceneId {
						s.currentMenu = s.mainMenu
						return id
					}
				}
			}
		}
		// Se Enter viene rilasciato, permetti nuove azioni
		if inpututil.KeyPressDuration(ebiten.KeyEnter) == 0 && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			s.actionExecuted = false
		}
	}
	return StartSceneId
}

var _ Scene = (*StartScene)(nil)

func (s *StartScene) GetSelectedMode() int {
	return s.selectedMode
}

func (s *StartScene) handleMenuSelection() SceneId {
	selectedOption := s.currentMenu.Options[s.currentMenu.Selected].Label

	switch selectedOption {
	case "PLAY":
		s.currentMenu = s.playMenu
		s.playMenu.Selected = 0
	case "OPTIONS":
		return OptionsSceneId
	case "HIGHSCORES":
		fmt.Println("HIGHSCORES NOT YET IMPLEMENTED")
	case "CREDITS":
		fmt.Println("CREDITS NOT YET IMPLEMENTED")
	case "QUIT":
		s.exitPopup.Active = true
		s.exitPopup.Selected = 0
		//return ExitSceneId
	case "SOLO MODE":
		s.selectedMode = 1
		return NameInputSceneId
	case "COMPUTER MODE":
		s.selectedMode = 3
		return NameInputSceneId
	case "MULTIPLAYER MODE":
		s.selectedMode = 2
		return NameInputSceneId
	case "BACK":
		s.currentMenu = s.mainMenu
	}
	return StartSceneId
}

func (s *StartScene) handleExitPopup() SceneId {
	//fmt.Println("ExitPopup Selected:", s.exitPopup.Selected)
	if s.exitPopup.Selected == 0 {
		return ExitSceneId
	} else {
		s.exitPopup.Active = false
	}
	return StartSceneId
}
