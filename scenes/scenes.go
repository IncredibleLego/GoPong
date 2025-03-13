package scenes

import (
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type SceneId uint

const (
	GameSceneId SceneId = iota
	StartSceneId
	ExitSceneId
	PauseSceneId
)

//go:embed PressStart2P-Regular.ttf
var pressStart2P []byte
var pressStart2PFaceSource *text.GoTextFaceSource

type Scene interface {
	Update() SceneId
	Draw(screen *ebiten.Image)
	FirstLoad()
	OnEnter()
	OnExit()
	IsLoaded() bool
}
