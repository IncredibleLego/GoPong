package scenes

import (
	"fmt"
	"goPong/menu"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type StartScene struct { // is the scene loaded now
	currentMenu        *menu.Menu
	mainMenu           *menu.Menu
	playMenu           *menu.Menu
	lastEnterPressTime time.Time
	actionExecuted     bool
}

func (s *StartScene) ShouldPreserveState() bool {
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

	//screen.Fill(color.RGBA{0, 0, 0, 1})

	menu.ScreenDraw(13, 250, 60, 1, 1, 1, 1, 1.5, screen, "Pong in Go")
	menu.ScreenDraw(13, 210, 80, 1, 1, 1, 1, 1.5, screen, "by IncredibleLego")

	s.currentMenu.Draw(screen)

}

func (s *StartScene) FirstLoad() {
	s.mainMenu = &menu.Menu{
		Options: []menu.MenuOption{
			{Label: "PLAY"},
			{Label: "OPTIONS"},
			{Label: "CREDITS"},
			{Label: "QUIT"},
		},
		Selected:     0,
		LastMoveTime: time.Now(),
	}
	s.playMenu = &menu.Menu{
		Options: []menu.MenuOption{
			{Label: "SOLO MODE"},
			{Label: "COMPUTER MODE"},
			{Label: "MULTIPLAYER MODE"},
			{Label: "BACK"},
		},
		Selected:     0,
		LastMoveTime: time.Now(),
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
		s.currentMenu = nextMenu
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

func (s *StartScene) handleMenuSelection() SceneId {
	selectedOption := s.currentMenu.Options[s.currentMenu.Selected].Label

	switch selectedOption {
	case "PLAY":
		s.currentMenu = s.playMenu
		s.playMenu.Selected = 0
	case "OPTIONS":
		fmt.Println("OPTIONS NOT YET IMPLEMENTED")
	case "CREDITS":
		fmt.Println("CREDITS NOT YET IMPLEMENTED")
	case "QUIT":
		return ExitSceneId
	case "SOLO MODE":
		return GameSceneId
	case "COMPUTER MODE":
		return ComputerSceneId
	case "MULTIPLAYER MODE":
		return MultiplayerSceneId
	case "BACK":
		s.currentMenu = s.mainMenu
	}

	return StartSceneId
}
