package game

import (
	"fmt"
	"image/color"

	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/input"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	bricks    [9]*GameBrick
	loadError error
)

type ClgGame struct{}

func (g *ClgGame) Update() error {
	input.Update()
	return nil
}

func (g *ClgGame) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 40, G: 40, B: 40, A: 0}) // NOt sure alpha is correct
	if loadError != nil {
		//TODO: Print something pretty on screen
		ebitenutil.DebugPrint(screen, "Error loading image")
		return
	}
	// ebitenutil.DebugPrint(screen, "Hello, World!"+strconv.Itoa(len(input.GetTouchIDs())))

	//for _, v := range touch.GetTouchIDs() {
	// ebitenutil.DebugPrint(screen, "Hello, World!")
	// ebitenutil.DebugPrint(screen, "TouchID:"+strconv.Itoa(int(v)))
	//}
	for _, brick := range bricks {
		brick.Draw(screen)
	}
}

// Automatically scales.
func (g *ClgGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// fmt.Println(outsideWidth, outsideHeight)
	// 1290/4, 2796/4
	// return 320, 240

	// return outsideWidth, outsideHeight
	return config.WindowWidth, config.WindowHeight
}

func loadGameBricks() {
	x, y := 25, 75
	for index, _ := range bricks {
		bricks[index].Position.X = x
		bricks[index].Position.Y = x
		x += 25
		if x > 75 {
			y += 75
			x = 25
		}
	}
}

func init() {
	fmt.Println("init")
	for index, _ := range bricks {
		brick, err := NewGameBrick()
		if err != nil {
			loadError = err
			break
		}
		bricks[index] = brick
	}
}
