package buttons

import (
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view"
	"github.com/Alecktos/clg/game/view/text"
	"github.com/Alecktos/clg/game/view_models"
	"github.com/hajimehoshi/ebiten/v2"
	text2 "github.com/hajimehoshi/ebiten/v2/text/v2"
)

type PrimaryButton interface {
	Draw(screen *ebiten.Image)
	Update()
	IsClicked() bool
}

type primaryButton struct {
	backgroundRectangle view.RoundedRectangle
	buttonModel         view_models.ButtonModel
	buttonLabel         text.Text
}

func NewPrimaryButton(parentRectangle common.Rectangle, label string) (PrimaryButton, error) {
	width := config.WindowWidth - 500.0
	height := 200.0

	position := common.Position{X: parentRectangle.HorizontalCenter() - width/2, Y: parentRectangle.BottomY() - height - 75}
	rectangle := common.Rectangle{Position: position, Width: width, Height: height}

	backgroundRectangle, err := view.NewRoundedRectangle(rectangle, 100, config.PassionRed())
	if err != nil {
		return nil, err
	}

	buttonLabel, err := text.NewText()
	if err != nil {
		return nil, err
	}

	textLayout := text.NewTextLayout(common.Position{X: config.WindowWidth / 2, Y: rectangle.Position.Y + rectangle.Height/2})
	textLayout.VerticalAlign = text2.AlignCenter
	buttonLabel.SetText(label, textLayout)

	return &primaryButton{
		backgroundRectangle: backgroundRectangle,
		buttonLabel:         buttonLabel,
		buttonModel:         view_models.NewButtonModel(),
	}, nil
}

func (c *primaryButton) Update() {
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

func (c *primaryButton) Draw(screen *ebiten.Image) {
	c.backgroundRectangle.Draw(screen)
	c.buttonLabel.Draw(screen)
}

func (c *primaryButton) IsClicked() bool {
	return c.buttonModel.IsClicked()
}
