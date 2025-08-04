package game_scene

import (
	"github.com/Alecktos/clg/assets/clg_json"
	"github.com/Alecktos/clg/game/view_model"
	"github.com/Alecktos/clg/game/view_providers"
	"image/color"

	"github.com/Alecktos/clg/assets/images"
	"github.com/Alecktos/clg/game/common"
	"github.com/hajimehoshi/ebiten/v2"
)

type gameBrick struct {
	common.Rectangle
	view_providers.VisibilityProvider
	buttonModel view_model.ButtonModel
	img         *ebiten.Image
	Challenge   clg_json.Challenge
}

const gameBrickWidth = 400
const gameBrickHeight = 400

func newGameBrick(challenge clg_json.Challenge) (*gameBrick, error) {
	brick := &gameBrick{
		buttonModel: view_model.NewButtonModel(),
		Rectangle: common.Rectangle{
			Position: common.Position{X: 0, Y: 0},
			Width:    gameBrickWidth,
			Height:   gameBrickHeight,
		},
		VisibilityProvider: view_providers.NewVisibilityProvider(),
		Challenge:          challenge,
	}
	err := brick.loadImage()
	return brick, err
}

func (g *gameBrick) update() {
	g.buttonModel.Update(g.Rectangle)
}

func (g *gameBrick) draw(screen *ebiten.Image) {
	if !g.IsVisible() {
		return
	}
	op := &ebiten.DrawImageOptions{}
	imgSourceSize := g.img.Bounds().Size()
	scaleX := float64(g.Width) / float64(imgSourceSize.X)
	scaleY := float64(g.Height) / float64(imgSourceSize.Y)
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(float64(g.Position.X), float64(g.Position.Y))

	if g.buttonModel.IsPressed() {
		op.ColorScale.ScaleWithColor(color.RGBA{R: 100, G: 100, B: 100, A: 255})
	}

	screen.DrawImage(g.img, op)
}

func (g *gameBrick) loadImage() error {
	var err error
	image, err := images.LoadBrickImage()
	g.img = image
	return err
}
