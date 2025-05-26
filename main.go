package main

import (
	"log"

	"github.com/Alecktos/clg/game"
	"github.com/Alecktos/clg/game/config"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// does nothing on phone
	ebiten.SetWindowSize(config.WindowWidth, config.WindowHeight)

	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(game.NewClgGame()); err != nil {
		log.Fatal(err)
	}
}
