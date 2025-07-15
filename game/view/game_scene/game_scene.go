package game_scene

import (
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view/game_scene/challenge"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	bricks         [9]*gameBrick
	challengeModal challenge.Modal
}

func NewGameScene() (*GameScene, error) {
	gameScene := &GameScene{}
	loadError := gameScene.load()
	return gameScene, loadError
}

func (s *GameScene) load() error {
	var err error
	s.challengeModal, err = challenge.NewChallengeModal()
	if err != nil {
		return err
	}
	loadError := s.loadGameBricks()
	return loadError
}

func (s *GameScene) loadGameBricks() error {
	var loadError error

	columns := 3
	spacing := 25

	// Beräkna total bredd för kolumner och startposition för centrerad layout
	totalWidth := columns*GAME_BRICK_WIDTH + (columns-1)*spacing
	startX := (config.WindowWidth - totalWidth) / 2
	x, y := startX, 75

	for index, _ := range s.bricks {

		brick, err := newGameBrick()
		if err != nil {
			loadError = err
			break
		}

		brick.Position.X = float64(x)
		brick.Position.Y = float64(y)
		x += GAME_BRICK_WIDTH + spacing

		// Gå till nästa rad efter varje tredje bricka
		if (index+1)%columns == 0 {
			x = startX
			y += GAME_BRICK_HEIGHT + spacing
		}
		s.bricks[index] = brick
	}
	return loadError
}

func (s *GameScene) Draw(screen *ebiten.Image) {

	for _, brick := range s.bricks {
		brick.draw(screen)
	}

	s.challengeModal.Draw(screen)
}

func (s *GameScene) Update() {
	s.challengeModal.Update()

	for _, brick := range s.bricks {
		brick.update()
		if brick.buttonModel.IsClicked() {
			s.challengeModal.Open()
		}
	}
}
