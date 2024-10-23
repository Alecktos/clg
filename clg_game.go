package main

import (
	"github.com/Alecktos/couplesLoveGame/view"
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
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(50, 50)
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(clgSprite.Img, op)
}

func (g *ClgGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
