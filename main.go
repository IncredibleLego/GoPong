package main

import (
	"goPong/config"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var isFullscreen bool

func main() {
	config.InitConfig()
	//fmt.Printf("Loaded configuration: %+v\n", config.GlobalConfig)

	ebiten.SetWindowTitle("Pong in Go")
	ebiten.SetWindowSize(config.GlobalConfig.ScreenWidth, config.GlobalConfig.ScreenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	isFullscreen = config.GlobalConfig.Fullscreen
	ebiten.SetFullscreen(isFullscreen)

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
