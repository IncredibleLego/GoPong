package main

import (
	"goPong/config"
	"goPong/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	sceneMap      map[scenes.SceneId]scenes.Scene
	activeSceneId scenes.SceneId
	loadedScenes  map[scenes.SceneId]bool
}

func NewGame() *Game {
	sceneMap := map[scenes.SceneId]scenes.Scene{
		scenes.GameSceneId:        scenes.NewGameScene(),
		scenes.StartSceneId:       scenes.NewStartScene(),
		scenes.PauseSceneId:       nil,
		scenes.ComputerSceneId:    scenes.NewComputerScene(),
		scenes.MultiplayerSceneId: scenes.NewMultiplayerScene(),
		scenes.OptionsSceneId:     scenes.NewOptionScene(),
	}
	activeSceneId := scenes.StartSceneId
	sceneMap[activeSceneId].FirstLoad()
	return &Game{
		sceneMap:      sceneMap,
		activeSceneId: activeSceneId,
		loadedScenes:  map[scenes.SceneId]bool{activeSceneId: true},
	}
}

func (g *Game) Update() error {
	nextSceneId := g.sceneMap[g.activeSceneId].Update() // updates the current scene
	if nextSceneId == scenes.ExitSceneId {
		g.sceneMap[g.activeSceneId].OnExit() // if the scene is the exit scene, the game is terminated
		return ebiten.Termination
	}
	if nextSceneId != g.activeSceneId {

		var reason scenes.SceneChangeReason
		if nextSceneId == scenes.PauseSceneId {
			g.sceneMap[scenes.PauseSceneId] = scenes.NewPauseScene(g.activeSceneId)
			reason = scenes.Other
		} else if g.activeSceneId == scenes.PauseSceneId && nextSceneId != scenes.ExitSceneId {
			reason = scenes.Unpause
		} else {
			reason = scenes.Exit
		}

		nextScene := g.sceneMap[nextSceneId] // if the scene is different from the current scene, the current scene is exited and the new scene is entered
		if !g.loadedScenes[nextSceneId] || !nextScene.ShouldPreserveState(reason) {
			nextScene.FirstLoad() // if the scene is not loaded or should not preserve state, it is loaded
			g.loadedScenes[nextSceneId] = true
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
	return config.GlobalConfig.ScreenWidth, config.GlobalConfig.ScreenHeight
}
