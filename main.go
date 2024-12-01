package main

import (
	"github.com/Alecktos/clg/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	// does nothing on phone
	ebiten.SetWindowSize(1290/4, 2796/4)

	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&game.ClgGame{}); err != nil {
		log.Fatal(err)
	}
}
