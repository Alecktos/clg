package game_scene

import (
	"github.com/Alecktos/clg/assets/clg_json"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view/game_scene/challenge"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	bricks         [12]*gameBrick
	challengeModal challenge.Modal
	challenges     []clg_json.Challenge
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

	challenges, err := clg_json.LoadChallenges()
	if err != nil {
		return err
	}
	s.challenges = challenges

	loadError := s.loadGameBricks()
	return loadError
}

func (s *GameScene) loadGameBricks() error {
	var loadError error

	columns := 3
	spacing := 25

	// Beräkna total bredd för kolumner och startposition för centrerad layout
	totalWidth := columns*GameBrickWidth + (columns-1)*spacing
	startX := (config.WindowWidth - totalWidth) / 2
	x, y := startX, 75

	for index := range s.bricks {

		brick, err := newGameBrick()
		if err != nil {
			loadError = err
			break
		}

		brick.Position.X = float64(x)
		brick.Position.Y = float64(y)
		x += GameBrickWidth + spacing

		// Gå till nästa rad efter varje tredje bricka
		if (index+1)%columns == 0 {
			x = startX
			y += GameBrickHeight + spacing
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
		if !s.challengeModal.IsVisible() && brick.buttonModel.IsClicked() {
			s.challengeModal.Open()
		}
	}
}
