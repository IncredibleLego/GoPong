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
		scenes.NameInputSceneId:   nil,
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
	// Updates the current scene and checks if the scene has changed if nextSceneId is different from the current scene
	nextSceneId := g.sceneMap[g.activeSceneId].Update()
	// If the next scene is the exit scene, the game is terminated
	if nextSceneId == scenes.ExitSceneId {
		g.sceneMap[g.activeSceneId].OnExit()
		return ebiten.Termination
	}
	// Instead, if the scene has changed:
	if nextSceneId != g.activeSceneId {

		var reason scenes.SceneChangeReason
		if nextSceneId == scenes.PauseSceneId {
			// If the next scene is the pause scene, it is created and the reason is set to "other"
			g.sceneMap[scenes.PauseSceneId] = scenes.NewPauseScene(g.activeSceneId)
			reason = scenes.Other
		} else if g.activeSceneId == scenes.PauseSceneId && nextSceneId != scenes.ExitSceneId {
			// If the current scene is the pause scene and the next scene is not the exit scene, the reason is set to "unpause"
			reason = scenes.Unpause
		} else if nextSceneId == scenes.NameInputSceneId {
			// If the next scene is the name input scene, it is created and the reason is set to "other"
			startScene := g.sceneMap[scenes.StartSceneId].(*scenes.StartScene)
			selectedMode := startScene.GetSelectedMode()
			// A new name input scene is created with the selected mode from the start scene
			g.sceneMap[scenes.NameInputSceneId] = scenes.NewNameInputScene(selectedMode)
			reason = scenes.Other
		} else {
			// If the next scene is not the pause scene, the name input scene or the exit scene, the reason is set to "exit"
			reason = scenes.Exit
		}

		// nextScene is the new scene to be loaded
		nextScene := g.sceneMap[nextSceneId]
		// If the scene is not loaded or should not preserve state, it is initialized wirh FirstLoad
		if !g.loadedScenes[nextSceneId] || !nextScene.ShouldPreserveState(reason) {
			nextScene.FirstLoad()
			g.loadedScenes[nextSceneId] = true
		}
		// The current scene is exited
		g.sceneMap[g.activeSceneId].OnExit()
		// The new scene is entered
		nextScene.OnEnter()
	}
	// Sets the new scene as the current scene (if nextSceneId is the same no action is taken)
	g.activeSceneId = nextSceneId
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) { // draws the current scene
	g.sceneMap[g.activeSceneId].Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.GlobalConfig.ScreenWidth, config.GlobalConfig.ScreenHeight
}
