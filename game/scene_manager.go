package game

import (
	"github.com/Alecktos/clg/game/view/endgame_scene"
	"github.com/Alecktos/clg/game/view/error_scene"
	"github.com/Alecktos/clg/game/view/game_scene"
	"github.com/hajimehoshi/ebiten/v2"
)

type sceneState int

const (
	errorSceneState sceneState = iota
	gameSceneState
	endSceneState
)

type SceneManager interface {
	Load(loadError error)
	Draw(screen *ebiten.Image)
	Update()
}

type sceneManager struct {
	current      sceneState
	gameScene    *game_scene.GameScene
	endGameScene endgame_scene.EndGameScene
	errorScene   *error_scene.ErrorScene
}

func NewSceneManager() SceneManager {
	return &sceneManager{
		current: gameSceneState,
	}
}

func (sm *sceneManager) Load(predefinedError error) {
	sm.errorScene = error_scene.NewErrorScene() //always load error scene
	sm.errorScene.SetError(predefinedError)
	if predefinedError != nil {
		sm.current = errorSceneState
		return
	}

	var loadSceneError error
	sm.gameScene, loadSceneError = game_scene.NewGameScene(func() {
		sm.current = endSceneState
	})
	if loadSceneError != nil {
		sm.errorScene.SetError(loadSceneError)
		sm.current = errorSceneState
		return
	}

	sm.endGameScene, loadSceneError = endgame_scene.NewEndGameScene(func() {
		sm.current = gameSceneState
	})
	if loadSceneError != nil {
		sm.errorScene.SetError(loadSceneError)
		sm.current = errorSceneState
		return
	}
}

func (sm *sceneManager) Draw(screen *ebiten.Image) {
	switch sm.current {
	case errorSceneState:
		sm.errorScene.Draw(screen)
		break
	case endSceneState:
		sm.endGameScene.Draw(screen)
		break
	case gameSceneState:
		sm.gameScene.Draw(screen)
		break
	}

}

func (sm *sceneManager) Update() {
	switch sm.current {
	case gameSceneState:
		sm.gameScene.Update()
		break
	case endSceneState:
		sm.endGameScene.Update()
		break
	}
}
