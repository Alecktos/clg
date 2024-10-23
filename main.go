package main

import (
	"github.com/Alecktos/clg/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&game.ClgGame{}); err != nil {
		log.Fatal(err)
	}
}
