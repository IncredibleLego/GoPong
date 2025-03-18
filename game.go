package main

import (
	"goPong/constants"
	"goPong/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sceneMap      map[scenes.SceneId]scenes.Scene
	activeSceneId scenes.SceneId
}

func NewGame() *Game {
	sceneMap := map[scenes.SceneId]scenes.Scene{
		scenes.GameSceneId:     scenes.NewGameScene(),
		scenes.StartSceneId:    scenes.NewStartScene(),
		scenes.PauseSceneId:    nil,
		scenes.ComputerSceneId: scenes.NewComputerScene(),
	}
	activeSceneId := scenes.StartSceneId
	sceneMap[activeSceneId].FirstLoad()
	return &Game{
		sceneMap,
		activeSceneId,
	}
}

func (g *Game) Update() error {
	nextSceneId := g.sceneMap[g.activeSceneId].Update() // updates the current scene
	if nextSceneId == scenes.ExitSceneId {
		g.sceneMap[g.activeSceneId].OnExit() // if the scene is the exit scene, the game is terminated
		return ebiten.Termination
	}
	if nextSceneId != g.activeSceneId {

		if nextSceneId == scenes.PauseSceneId {
			g.sceneMap[scenes.PauseSceneId] = scenes.NewPauseScene(g.activeSceneId)
		}

		nextScene := g.sceneMap[nextSceneId] // if the scene is different from the current scene, the current scene is exited and the new scene is entered
		if !nextScene.IsLoaded() {
			nextScene.FirstLoad() // if the scene is not loaded, it is loaded
		}
		nextScene.OnEnter()                  // the new scene is entered
		g.sceneMap[g.activeSceneId].OnExit() // the current scene is exited
	}
	g.activeSceneId = nextSceneId // the new scene is set as the current scene
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) { // draws the current scene
	g.sceneMap[g.activeSceneId].Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return constants.ScreenWidth, constants.ScreenHeight
}
