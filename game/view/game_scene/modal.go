package game_scene

import (
	"github.com/Alecktos/clg/assets/clg_json"
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view"
	"github.com/Alecktos/clg/game/view/buttons"
	"github.com/Alecktos/clg/game/view/text"
	"github.com/hajimehoshi/ebiten/v2"
)

type modal struct {
	header               text.Text
	description          text.Text
	roundedRectangle     view.RoundedRectangle
	closeChallengeButton buttons.PrimaryButton
	isVisible            bool
	onClosed             func()
}

func newModal() (*modal, error) {
	hMargin := 80.0
	width := config.WindowWidth - hMargin*2
	height := float64(config.WindowHeight - 800)
	x1 := hMargin
	y1 := config.WindowHeight/2 - height/2

	backgroundRectangle := common.Rectangle{Position: common.Position{X: x1, Y: y1}, Width: width, Height: height}

	header, err := text.NewText()
	if err != nil {
		return nil, err
	}

	description, err := text.NewText()
	if err != nil {
		return nil, err
	}

	roundedRectangle, err := view.NewRoundedRectangle(backgroundRectangle, 100, config.VelvetPlum())
	if err != nil {
		return nil, err
	}

	closeButton, err := buttons.NewPrimaryButton(backgroundRectangle, "Close")
	if err != nil {
		return nil, err
	}

	return &modal{
		header:               header,
		description:          description,
		roundedRectangle:     roundedRectangle,
		closeChallengeButton: closeButton,
		isVisible:            false,
		onClosed:             func() {},
	}, nil
}

func (c *modal) Open(challenge *clg_json.Challenge, onClosed func()) {
	y := c.roundedRectangle.GetRectangle().Position.Y + 85

	textLayout := text.NewTextLayout(common.Position{X: config.WindowWidth / 2, Y: y})
	textLayout.FontSize = config.HeaderFontSize
	textLayout.Color = config.BlushPink()
	textLayout.MaxWidth = int(c.roundedRectangle.GetRectangle().Width)

	c.header.SetText(challenge.Header, textLayout)

	textLayout.Position.Y += c.header.MeasureHeight() + 50
	textLayout.FontSize = config.StandardFontSize
	textLayout.Color = config.ChampagneGold()
	c.description.SetText(challenge.Description, textLayout)
	c.isVisible = true
	c.onClosed = onClosed
}

func (c *modal) IsVisible() bool {
	return c.isVisible
}

func (c *modal) Update() {
	c.closeChallengeButton.Update()
	if c.closeChallengeButton.IsClicked() {
		c.isVisible = false
		c.onClosed()
	}
}

func (c *modal) Draw(screen *ebiten.Image) {
	if !c.isVisible {
		return
	}
	c.roundedRectangle.Draw(screen)
	c.header.Draw(screen)
	c.description.Draw(screen)
	c.closeChallengeButton.Draw(screen)
}
