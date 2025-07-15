package game_scene

import (
	"github.com/Alecktos/clg/game/view_model"
	"image/color"

	"github.com/Alecktos/clg/assets/images"
	"github.com/Alecktos/clg/game/common"
	"github.com/hajimehoshi/ebiten/v2"
)

type gameBrick struct {
	common.Rectangle
	buttonModel view_model.ButtonModel
	Img         *ebiten.Image // private?
}

const GameBrickWidth = 400
const GameBrickHeight = 400

func newGameBrick() (*gameBrick, error) {
	clgImage := &gameBrick{
		buttonModel: view_model.NewButtonModel(),
		Rectangle: common.Rectangle{
			Position: common.Position{X: 0, Y: 0},
			Width:    GameBrickWidth,
			Height:   GameBrickHeight,
		},
	}
	err := clgImage.loadImage()
	return clgImage, err
}

func (g *gameBrick) update() {
	g.buttonModel.Update(g.Rectangle)
}

func (g *gameBrick) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	imgSourceSize := g.Img.Bounds().Size()
	scaleX := float64(g.Width) / float64(imgSourceSize.X)
	scaleY := float64(g.Height) / float64(imgSourceSize.Y)
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(float64(g.Position.X), float64(g.Position.Y))

	if g.buttonModel.IsPressed() {
		op.ColorScale.ScaleWithColor(color.RGBA{R: 100, G: 100, B: 100, A: 255})
	}

	screen.DrawImage(g.Img, op)
}

func (g *gameBrick) loadImage() error {
	var err error
	image, err := images.LoadBrickImage()
	g.Img = image
	return err
}
