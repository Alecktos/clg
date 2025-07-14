package game_scene

import (
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view"
	"github.com/hajimehoshi/ebiten/v2"
)

type challengeModal struct {
	challengeText    *view.Text
	roundedRectangle view.RoundedRectangle
}

func newChallengeModal() (*challengeModal, error) {
	width := float64(config.WindowWidth - 60)
	height := float64(300)
	x1 := 30.0
	y1 := config.WindowHeight/2 - height/2

	challengeText, err := view.NewCenterAlignedText("Here is challenge for you", y1+20)
	if err != nil {
		return nil, err
	}

	roundedRectangle, err := view.NewRoundedRectangle(common.Rectangle{Position: common.Position{X: x1, Y: y1}, Width: width, Height: height}, 25.0, config.VelvetPlum())
	if err != nil {
		return nil, err
	}

	return &challengeModal{
		challengeText:    challengeText,
		roundedRectangle: roundedRectangle,
	}, nil
}

func (cm *challengeModal) draw(screen *ebiten.Image) {
	cm.roundedRectangle.Draw(screen)
	cm.challengeText.Draw(screen)
}
