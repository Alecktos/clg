package game_scene

import (
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view"
	"github.com/hajimehoshi/ebiten/v2"
)

type challengeModal struct {
	challengeText        *view.Text
	roundedRectangle     view.RoundedRectangle
	closeChallengeButton *closeChallengeButton
	isVisible            bool
}

func newChallengeModal() (*challengeModal, error) {
	width := float64(config.WindowWidth - 60)
	height := float64(300)
	x1 := 30.0
	y1 := config.WindowHeight/2 - height/2

	backgroundRectangle := common.Rectangle{Position: common.Position{X: x1, Y: y1}, Width: width, Height: height}

	challengeText, err := view.NewCenterAlignedText("Here is challenge for you", y1+20, config.STANDARD_FONT_SIZE)
	if err != nil {
		return nil, err
	}

	roundedRectangle, err := view.NewRoundedRectangle(backgroundRectangle, 25.0, config.VelvetPlum())
	if err != nil {
		return nil, err
	}

	closeButton, err := newCloseChallengeButton(backgroundRectangle)
	if err != nil {
		return nil, err
	}

	return &challengeModal{
		challengeText:        challengeText,
		roundedRectangle:     roundedRectangle,
		closeChallengeButton: closeButton,
		isVisible:            false,
	}, nil
}

func (c *challengeModal) update() {
	c.closeChallengeButton.update()
	if c.closeChallengeButton.buttonModel.IsClicked() {
		c.isVisible = false
	}
}

func (c *challengeModal) draw(screen *ebiten.Image) {
	if !c.isVisible {
		return
	}
	c.roundedRectangle.Draw(screen)
	c.challengeText.Draw(screen)
	c.closeChallengeButton.draw(screen)
}
