package game

import (
	"github.com/Alecktos/clg/view"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	clgSprite *view.ClgSprite = view.NewClgSprite()
)

type ClgGame struct{}

func (g *ClgGame) Update() error {
	return nil
}

func (g *ClgGame) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	clgSprite.Draw(screen)
}

func (g *ClgGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
