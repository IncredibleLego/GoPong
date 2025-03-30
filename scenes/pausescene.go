package scenes

import (
	"goPong/menu"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type PauseScene struct {
	pauseMenu       *menu.Menu
	actionExecuted  bool
	previousSceneId SceneId
	options         bool //true if the last scene was options
}

func (p *PauseScene) ShouldPreserveState(reason SceneChangeReason) bool {
	return false
}

func NewPauseScene(previous SceneId, opt bool) *PauseScene {
	return &PauseScene{
		pauseMenu:       nil,
		previousSceneId: previous,
		options:         opt,
	}
}

func (p *PauseScene) Draw(screen *ebiten.Image) {

	p.pauseMenu.Draw(screen)
}

func (p *PauseScene) FirstLoad() {
	p.pauseMenu = &menu.Menu{
		Options: []menu.MenuOption{
			{Label: "UNPAUSE"},
			{Label: "OPTIONS"},
			{Label: "EXIT"},
		},
		Selected:     0,
		LastMoveTime: time.Now(),
	}
}

func (p *PauseScene) OnEnter() {
	if p.options {
		p.pauseMenu.Selected = 1
	}
}

func (p *PauseScene) OnExit() {

}

func (p *PauseScene) Update() SceneId {

	p.pauseMenu.Update()

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && !p.actionExecuted {
		id := p.handleMenuSelection()
		p.actionExecuted = true
		if id != PauseSceneId {
			return id
		}
	}

	if inpututil.KeyPressDuration(ebiten.KeyEnter) == 0 && !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		p.actionExecuted = false
	}

	return PauseSceneId
}

var _ Scene = (*PauseScene)(nil)

func (p *PauseScene) handleMenuSelection() SceneId {
	selectedOption := p.pauseMenu.Options[p.pauseMenu.Selected].Label

	switch selectedOption {
	case "UNPAUSE":
		p.pauseMenu.Selected = 0
		return p.previousSceneId
	case "OPTIONS":
		return OptionsSceneId
	case "EXIT":
		p.pauseMenu.Selected = 0
		return StartSceneId
	}

	return PauseSceneId
}
