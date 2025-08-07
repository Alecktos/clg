package game_scene

import (
	"math/rand"
	"time"

	"github.com/Alecktos/clg/assets/clg_json"
	"github.com/Alecktos/clg/game/config"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameScene struct {
	bricks         [15]*gameBrick
	challengeModal *modal
	onDone         func()
}

func NewGameScene(onDone func()) (*GameScene, error) {
	gameScene := &GameScene{
		onDone: onDone,
	}
	loadError := gameScene.load()
	return gameScene, loadError
}

func (s *GameScene) load() error {
	var err error
	s.challengeModal, err = newModal()
	if err != nil {
		return err
	}

	loadError := s.loadGameBricks()
	return loadError
}

func (s *GameScene) loadGameBricks() error {
	challenges, err := clg_json.LoadChallenges()
	if err != nil {
		return err
	}

	var loadError error

	columns := 3
	spacing := 25

	// Beräkna total bredd för kolumner och startposition för centrerad layout
	totalWidth := columns*gameBrickWidth + (columns-1)*spacing
	startX := (config.WindowWidth - totalWidth) / 2
	x, y := startX, 75

	for index := range s.bricks {

		brick, err := newGameBrick(challenges[index])
		if err != nil {
			loadError = err
			break
		}

		brick.Position.X = float64(x)
		brick.Position.Y = float64(y)
		x += gameBrickWidth + spacing

		// Gå till nästa rad efter varje tredje bricka
		if (index+1)%columns == 0 {
			x = startX
			y += gameBrickHeight + spacing
		}
		s.bricks[index] = brick
	}

	s.shuffleBricks()

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
			s.disableAllBricks()
			s.challengeModal.Open(&brick.Challenge, func() {
				s.enableAllBricks()
				brick.buttonModel.SetVisibility(false)
				if s.allChallengesCompleted() {
					s.onDone()
				}
			})
		}
	}
}

func (s *GameScene) allChallengesCompleted() bool {
	for _, brick := range s.bricks {
		if brick.buttonModel.IsVisible() {
			return false
		}
	}
	return true
}

func (s *GameScene) disableAllBricks() {
	for _, brick := range s.bricks {
		brick.buttonModel.SetDisabled(true)
	}
}

func (s *GameScene) enableAllBricks() {
	for _, brick := range s.bricks {
		brick.buttonModel.SetDisabled(false)
	}
}

func (s *GameScene) Reset() {
	for _, brick := range s.bricks {
		if !brick.buttonModel.IsVisible() {
			brick.buttonModel.SetVisibility(true)
		}
	}
	s.shuffleBricks()
}

func (s *GameScene) shuffleBricks() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(s.bricks), func(i, j int) {
		tmp := s.bricks[i].Challenge
		s.bricks[i].Challenge = s.bricks[j].Challenge
		s.bricks[j].Challenge = tmp
	})
}
