package main

import (
	"log"

	"github.com/Alecktos/clg/game"
	"github.com/Alecktos/clg/game/config"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// does nothing on phone
	ebiten.SetWindowSize(config.WindowWidth/4, config.WindowHeight/4)

	if config.DevMode {
		ebiten.SetWindowTitle("Clg Game - Dev Mode")
	}

	if err := ebiten.RunGame(game.NewClgGame()); err != nil {
		log.Fatal(err)
	}
}
