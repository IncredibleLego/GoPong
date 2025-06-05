package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneId uint

const (
	GameSceneId SceneId = iota
	StartSceneId
	ExitSceneId
	PauseSceneId
	ComputerSceneId
	MultiplayerSceneId
	OptionsSceneId
	NameInputSceneId
	HighScoresSceneId
)

type Scene interface {
	Update() SceneId
	Draw(screen *ebiten.Image)
	FirstLoad()
	OnEnter()
	OnExit()
	ShouldPreserveState(reason SceneChangeReason) bool
}

type SceneChangeReason string

const (
	Unpause SceneChangeReason = "unpause"
	Exit    SceneChangeReason = "exit"
	Other   SceneChangeReason = "other"
)
