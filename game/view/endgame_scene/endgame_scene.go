package endgame_scene

import (
	"github.com/Alecktos/clg/assets/clg_json"
	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view"
	"github.com/Alecktos/clg/game/view/buttons"
	"github.com/Alecktos/clg/game/view/text"
	"github.com/hajimehoshi/ebiten/v2"
)

type EndGameScene interface {
	Draw(screen *ebiten.Image)
	Update()
}

type endGameScene struct {
	backgroundRectangle view.RoundedRectangle
	header              text.Text
	description         text.Text
	startGameButton     buttons.PrimaryButton
	onDone              func()
}

func NewEndGameScene(onDone func()) (EndGameScene, error) {
	hMargin := 40.0
	width := config.WindowWidth - hMargin*2
	height := float64(config.WindowHeight - hMargin*4)
	x1 := hMargin
	y1 := config.WindowHeight/2 - height/2

	background, err := view.NewRoundedRectangle(common.NewRectangle(x1, y1, width, height), 100, config.VelvetPlum())
	if err != nil {
		return nil, err
	}

	endGameData, err := clg_json.LoadEndGame()
	if err != nil {
		return nil, err
	}

	header, err := text.NewText()
	if err != nil {
		return nil, err
	}

	textLayout := text.NewTextLayout(common.Position{X: config.WindowWidth / 2, Y: y1 + 400})
	textLayout.FontSize = config.HeaderFontSize
	textLayout.Color = config.PassionRed()

	header.SetText(endGameData.Header, textLayout)

	description, err := text.NewText()
	if err != nil {
		return nil, err
	}

	textLayout = text.NewTextLayout(common.Position{X: config.WindowWidth / 2, Y: y1 + 700})
	textLayout.Color = config.ChampagneGold()
	description.SetText(endGameData.Description, textLayout)

	actionButton, err := buttons.NewPrimaryButton(*background.GetRectangle(), endGameData.ButtonText)
	if err != nil {
		return nil, err
	}

	return &endGameScene{
		header:              header,
		description:         description,
		backgroundRectangle: background,
		startGameButton:     actionButton,
		onDone:              onDone,
	}, nil
}

func (e *endGameScene) Draw(screen *ebiten.Image) {
	e.backgroundRectangle.Draw(screen)
	e.header.Draw(screen)
	e.description.Draw(screen)
	e.startGameButton.Draw(screen)
}
func (e *endGameScene) Update() {
	e.startGameButton.Update()
	if e.startGameButton.IsClicked() {
		e.onDone()
	}
}
