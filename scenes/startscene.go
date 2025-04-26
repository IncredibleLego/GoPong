package scenes

import (
	"fmt"
	"goPong/menu"
	"goPong/utils"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type StartScene struct { // is the scene loaded now
	currentMenu        *menu.RegularMenu
	mainMenu           *menu.RegularMenu
	playMenu           *menu.RegularMenu
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

	//utils.ScreenDraw(-7, 250, 60, "white", screen, "Pong in Go")

	//Lettere 82 spazio 21 dimensione 14
	// Create file to print this text

	var dimension float32 = 9
	var firstY float32 = 30
	var space float32 = 21

	// Draw "G"
	vector.DrawFilledRect(screen,
		21, firstY, dimension, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21, firstY, 82, dimension,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21, firstY+100-dimension, 82, dimension,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+41, firstY+50-dimension/2, 41, dimension,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82-dimension, firstY+50-dimension/2, dimension, 50,
		color.White, false,
	)
	// Draw "O"
	vector.DrawFilledRect(screen,
		21+82+space, firstY, dimension, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space, firstY, 82, dimension,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82-dimension, firstY, dimension, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space, firstY+100-dimension, 82, dimension,
		color.White, false,
	)
	// Draw "P"
	vector.DrawFilledRect(screen,
		21+82+space+82+space, firstY, dimension, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space, firstY, 82, dimension,
		color.White, false,
	)

	vector.DrawFilledRect(screen,
		21+82+space+82+space, firstY+45, 82, dimension,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82-dimension, firstY, dimension, 45,
		color.White, false,
	)
	// Draw "O"
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space, firstY, dimension, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space, firstY, 82, dimension,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82-dimension, firstY, dimension, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space, firstY+100-dimension, 82, dimension,
		color.White, false,
	)
	// Draw "N"
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82+space, firstY, dimension, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82+space+82-dimension, firstY, dimension, 100,
		color.White, false,
	)
	// Draw "G"
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82+space+82+space, firstY, dimension, 100,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82+space+82+space, firstY, 82, dimension,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82+space+82+space, firstY+100-dimension, 82, dimension,
		color.White, false,
	)

	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82+space+82+space+41, firstY+50-dimension/2, 41, dimension,
		color.White, false,
	)
	vector.DrawFilledRect(screen,
		21+82+space+82+space+82+space+82+space+82+space+82-dimension, firstY+50-dimension/2, dimension, 50,
		color.White, false,
	)

	utils.ScreenDraw(-5, 195, 200, "white", screen, "by IncredibleLego")

	s.currentMenu.Draw(screen)

}

func (s *StartScene) FirstLoad() {
	s.mainMenu = &menu.RegularMenu{
		Options: []menu.MenuOption{
			{Label: "PLAY"},
			{Label: "OPTIONS"},
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
		Offset:       40,
	}
	s.currentMenu = s.mainMenu
	s.lastEnterPressTime = time.Now()
	s.actionExecuted = false
}

func (s *StartScene) OnEnter() {

}

func (s *StartScene) OnExit() {

}

func (s *StartScene) Update() SceneId {
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
	case "CREDITS":
		fmt.Println("CREDITS NOT YET IMPLEMENTED")
	case "QUIT":
		return ExitSceneId
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
