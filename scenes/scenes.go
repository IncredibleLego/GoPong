package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneId uint

const (
	StartSceneId SceneId = iota
	ExitSceneId
	PauseSceneId
	GameSceneId
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
	ShouldPreserveState(reason SceneChangeReason) bool
}

type SceneChangeReason string

const (
	Unpause SceneChangeReason = "unpause"
	Exit    SceneChangeReason = "exit"
	Other   SceneChangeReason = "other"
)
