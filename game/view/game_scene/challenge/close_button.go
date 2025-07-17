package challenge

import (
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view"
	"github.com/Alecktos/clg/game/view_model"
	"github.com/hajimehoshi/ebiten/v2"
	text2 "github.com/hajimehoshi/ebiten/v2/text/v2"
)

type closeChallengeButton struct {
	backgroundRectangle view.RoundedRectangle
	buttonModel         view_model.ButtonModel
	buttonLabel         view.Text
}

func newCloseChallengeButton(parentRectangle common.Rectangle) (*closeChallengeButton, error) {
	width := config.WindowWidth - 500.0
	height := 200.0

	position := common.Position{X: parentRectangle.HorizontalCenter() - width/2, Y: parentRectangle.BottomY() - height - 75}
	rectangle := common.Rectangle{Position: position, Width: width, Height: height}

	backgroundRectangle, err := view.NewRoundedRectangle(rectangle, 100, config.PassionRed())
	if err != nil {
		return nil, err
	}

	buttonLabel, err := view.NewCenterAlignedText()
	if err != nil {
		return nil, err
	}

	textLayout := view.NewTextLayout(common.Position{X: config.WindowWidth / 2, Y: rectangle.Position.Y + rectangle.Height/2})
	textLayout.VerticalAlign = text2.AlignCenter
	buttonLabel.SetText("Close", textLayout)

	return &closeChallengeButton{
		backgroundRectangle: backgroundRectangle,
		buttonLabel:         buttonLabel,
		buttonModel:         view_model.NewButtonModel(),
	}, nil
}

func (c *closeChallengeButton) update() {
	c.buttonModel.Update(*c.backgroundRectangle.GetRectangle())
	if c.buttonModel.FirstPressedTick() {
		var darkenFactor float32 = 0.8
		c.buttonLabel.DarkenColor(darkenFactor)
		c.backgroundRectangle.DarkenColor(darkenFactor)
	} else if !c.buttonModel.IsPressed() {
		c.buttonLabel.ResetToInitialColor()
		c.backgroundRectangle.ResetToInitialColor()
	}
}

func (c *closeChallengeButton) draw(screen *ebiten.Image) {
	c.backgroundRectangle.Draw(screen)
	c.buttonLabel.Draw(screen)
}
