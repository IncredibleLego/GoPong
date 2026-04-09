package main

import (
	"goPong/audio"
	"goPong/config"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	config.InitConfig()
	config.ApplyScaleToConfig(config.GlobalConfig, config.GlobalConfig.Scale)
	audio.Init()

	ebiten.SetWindowTitle("GoPong")
	ebiten.SetWindowSize(config.GlobalConfig.ScreenWidth, config.GlobalConfig.ScreenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	config.IsFullscreen = config.GlobalConfig.Fullscreen
	ebiten.SetFullscreen(config.IsFullscreen)

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
