package endgame_scene

import (
	"github.com/Alecktos/clg/assets/clg_json"
	"github.com/Alecktos/clg/game/view/modal"
	"github.com/hajimehoshi/ebiten/v2"
)

type EndGameScene interface {
	Draw(screen *ebiten.Image)
	Update()
}

type endGameScene struct {
	endGameModal modal.Modal
}

func NewEndGameScene() (EndGameScene, error) {
	endGameModal, err := modal.NewModal()
	if err != nil {
		return nil, err
	}

	endGameModal.Open(&clg_json.Challenge{Header: "", Description: ""}, func() {

	})

	return &endGameScene{
		endGameModal: endGameModal,
	}, nil
}

func (e *endGameScene) Draw(screen *ebiten.Image) {
	e.endGameModal.Draw(screen)
}
func (e *endGameScene) Update() {
	e.endGameModal.Update()
}
