package game

import (
	"github.com/Alecktos/clg/game/input"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	clgSprite *ClgSprite = NewClgSprite()
)

type ClgGame struct{}

func (g *ClgGame) Update() error {
	input.Update()
	return nil
}

func (g *ClgGame) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!"+strconv.Itoa(len(input.GetTouchIDs())))

	//for _, v := range touch.GetTouchIDs() {
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	// ebitenutil.DebugPrint(screen, "TouchID:"+strconv.Itoa(int(v)))
	//}

	clgSprite.Draw(screen)
}

func (g *ClgGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}
