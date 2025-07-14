package game_scene

import (
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view"
	"github.com/Alecktos/clg/game/view_model"
	"github.com/hajimehoshi/ebiten/v2"
)

type closeChallengeButton struct {
	backgroundRectangle view.RoundedRectangle
	buttonModel         view_model.ButtonModel
}

func newCloseChallengeButton(parentRectangle common.Rectangle) (*closeChallengeButton, error) {
	width := 100.0
	height := 50.0

	position := common.Position{X: parentRectangle.HorizontalCenter() - width/2, Y: parentRectangle.BottomY() - height - 10}
	rectangle := common.Rectangle{Position: position, Width: width, Height: height}

	backgroundRectangle, err := view.NewRoundedRectangle(rectangle, 10, config.PassionRed())
	if err != nil {
		return nil, err
	}

	return &closeChallengeButton{
		backgroundRectangle: backgroundRectangle,
	}, nil
}

func (c *closeChallengeButton) update() {
	c.buttonModel.Update(*c.backgroundRectangle.GetRectangle())
}

func (c *closeChallengeButton) draw(screen *ebiten.Image) {
	c.backgroundRectangle.Draw(screen)
}
