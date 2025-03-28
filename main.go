package main

import (
	"fmt"
	"goPong/config"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Pong in Go")
	ebiten.SetWindowSize(config.GlobalConfig.ScreenWidth, config.GlobalConfig.ScreenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	config.InitConfig()
	fmt.Printf("Configurazione caricata: %+v\n", config.GlobalConfig)

	game := NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
