package main

import (
	"goPong/config"
	"goPong/scenes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	sceneMap      map[scenes.SceneId]scenes.Scene
	activeSceneId scenes.SceneId
	lastSceneId   scenes.SceneId // ID of the last played scene between solo, computer and multiplayer
	loadedScenes  map[scenes.SceneId]bool
}

func NewGame() *Game {
	sceneMap := map[scenes.SceneId]scenes.Scene{
		scenes.GameSceneId:        scenes.NewGameScene(),
		scenes.StartSceneId:       scenes.NewStartScene(),
		scenes.PauseSceneId:       nil,
		scenes.ComputerSceneId:    scenes.NewComputerScene(),
		scenes.MultiplayerSceneId: scenes.NewMultiplayerScene(),
		scenes.OptionsSceneId:     nil,
		scenes.NameInputSceneId:   nil,
		scenes.HighScoresSceneId:  nil,
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

	// Using "Esc" you can enter or exit fullscreen
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		config.IsFullscreen = !config.IsFullscreen
		ebiten.SetFullscreen(config.IsFullscreen)
	}

	// Updates the current scene and checks if the scene has changed if nextSceneId is different from the current scene
	nextSceneId := g.sceneMap[g.activeSceneId].Update()
	// If the next scene is the exit scene, the game is terminated
	if nextSceneId == scenes.ExitSceneId {
		return ebiten.Termination
	}
	// Instead, if the scene has changed:
	if nextSceneId != g.activeSceneId {
		// Saves Highscores if you're returing to menu after playing a game
		if g.lastSceneId != 0 && nextSceneId == scenes.StartSceneId {
			if g.lastSceneId == scenes.GameSceneId && scenes.DirtySoloScore.Score > 0 {
				scenes.AddSoloScore(scenes.DirtySoloScore)
				scenes.DirtySoloScore = scenes.SoloScore{}
			} else if g.lastSceneId == scenes.ComputerSceneId && scenes.DirtyComputerScore.Score > 0 {
				scenes.AddComputerScore(scenes.DirtyComputerScore)
				scenes.DirtyComputerScore = scenes.ComputerScore{}
			} else if g.lastSceneId == scenes.MultiplayerSceneId && scenes.DirtyMultiplayerScore.Score > 0 {
				scenes.AddMultiplayerScore(scenes.DirtyMultiplayerScore)
				scenes.DirtyMultiplayerScore = scenes.MultiplayerScore{}
			}
			g.lastSceneId = 0
		}
		var reason scenes.SceneChangeReason
		if nextSceneId == scenes.PauseSceneId {
			// If the next scene is the pause scene, it is created and the reason is set to "other"
			if g.activeSceneId == scenes.OptionsSceneId {
				// If the current scene is the options scene, the last scene id is set to the previous scene
				g.sceneMap[scenes.PauseSceneId] = scenes.NewPauseScene(g.lastSceneId, true)
			} else {
				// If the current scene is not the options scene, the last scene id is set to the active scene
				g.sceneMap[scenes.PauseSceneId] = scenes.NewPauseScene(g.activeSceneId, false)
			}
			reason = scenes.Other
		} else if nextSceneId == scenes.OptionsSceneId {
			// If the next scene is the pause scene, it is created and the reason is set to "other"
			g.sceneMap[scenes.OptionsSceneId] = scenes.NewOptionScene(g.activeSceneId)
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
		} else if nextSceneId == scenes.HighScoresSceneId {
			// If the next scene is the high scores scene, it is created and the reason is set to "other"
			g.sceneMap[scenes.HighScoresSceneId] = scenes.NewHighScoresScene()
			reason = scenes.Other
		} else {
			// If the next scene is not the pause scene, the name input scene, the highscoresscene or the exit scene, the reason is set to "exit"
			reason = scenes.Exit
		}

		// nextScene is the new scene to be loaded
		nextScene := g.sceneMap[nextSceneId]
		// If the scene is not loaded or should not preserve state, it is initialized wirh FirstLoad
		if !g.loadedScenes[nextSceneId] || !nextScene.ShouldPreserveState(reason) {
			nextScene.FirstLoad()
			g.loadedScenes[nextSceneId] = true
		}
		// The current scene is exited. If the current scene is not the pause scene or the options scene, the last scene id is saved
		if g.activeSceneId == scenes.GameSceneId || g.activeSceneId == scenes.ComputerSceneId || g.activeSceneId == scenes.MultiplayerSceneId {
			g.lastSceneId = g.activeSceneId
		}
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
