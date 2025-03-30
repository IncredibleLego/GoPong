package main

import (
	"goPong/config"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Pong in Go")
	ebiten.SetWindowSize(config.GlobalConfig.ScreenWidth, config.GlobalConfig.ScreenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	config.InitConfig()
	//fmt.Printf("Loaded configuration: %+v\n", config.GlobalConfig)

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
