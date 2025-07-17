package challenge

import (
	"github.com/Alecktos/clg/assets/clg_json"
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view"
	"github.com/hajimehoshi/ebiten/v2"
)

type Modal interface {
	Update()
	Draw(screen *ebiten.Image)
	Open(challenge *clg_json.Challenge)
	IsVisible() bool
}

type challengeModal struct {
	header               view.Text
	description          view.Text
	roundedRectangle     view.RoundedRectangle
	closeChallengeButton *closeChallengeButton
	isVisible            bool
}

func NewChallengeModal() (Modal, error) {
	hMargin := 80.0
	width := config.WindowWidth - hMargin*2
	height := float64(config.WindowHeight - 800)
	x1 := hMargin
	y1 := config.WindowHeight/2 - height/2

	backgroundRectangle := common.Rectangle{Position: common.Position{X: x1, Y: y1}, Width: width, Height: height}

	header, err := view.NewCenterAlignedText()
	if err != nil {
		return nil, err
	}

	description, err := view.NewCenterAlignedText()
	if err != nil {
		return nil, err
	}

	roundedRectangle, err := view.NewRoundedRectangle(backgroundRectangle, 100, config.VelvetPlum())
	if err != nil {
		return nil, err
	}

	closeButton, err := newCloseChallengeButton(backgroundRectangle)
	if err != nil {
		return nil, err
	}

	return &challengeModal{
		header:               header,
		description:          description,
		roundedRectangle:     roundedRectangle,
		closeChallengeButton: closeButton,
		isVisible:            false,
	}, nil
}

func (c *challengeModal) Open(challenge *clg_json.Challenge) {
	y := c.roundedRectangle.GetRectangle().Position.Y + 85

	textLayout := view.NewTextLayout(common.Position{X: config.WindowWidth / 2, Y: y})
	textLayout.FontSize = config.HeaderFontSize
	textLayout.Color = config.BlushPink()
	textLayout.MaxWidth = int(c.roundedRectangle.GetRectangle().Width)

	c.header.SetText(challenge.Header, textLayout)

	textLayout.Position.Y += 200
	textLayout.Color = config.ChampagneGold()
	c.description.SetText(challenge.Description, textLayout)
	c.isVisible = true
}

func (c *challengeModal) IsVisible() bool {
	return c.isVisible
}

func (c *challengeModal) Update() {
	c.closeChallengeButton.update()
	if c.closeChallengeButton.buttonModel.IsClicked() {
		c.isVisible = false
	}
}

func (c *challengeModal) Draw(screen *ebiten.Image) {
	if !c.isVisible {
		return
	}
	c.roundedRectangle.Draw(screen)
	c.header.Draw(screen)
	c.description.Draw(screen)
	c.closeChallengeButton.draw(screen)
}
