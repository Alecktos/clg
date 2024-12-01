package game

import (
	"fmt"
	"github.com/Alecktos/clg/game/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	clgSprite, err = NewClgSprite()
)

type ClgGame struct{}

func (g *ClgGame) Update() error {
	input.Update()
	return nil
}

func (g *ClgGame) Draw(screen *ebiten.Image) {
	if err != nil {
		ebitenutil.DebugPrint(screen, "Error loading image")
		return
	}
	// ebitenutil.DebugPrint(screen, "Hello, World!"+strconv.Itoa(len(input.GetTouchIDs())))

	//for _, v := range touch.GetTouchIDs() {
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	// ebitenutil.DebugPrint(screen, "TouchID:"+strconv.Itoa(int(v)))
	//}

	clgSprite.Draw(screen)
}

// Automatically scales.
func (g *ClgGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	fmt.Println(outsideWidth, outsideHeight)
	// 1290/4, 2796/4
	// return 320, 240
	return outsideWidth, outsideHeight
}
