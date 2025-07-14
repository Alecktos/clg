package game

import (
	"github.com/Alecktos/clg/assets/fonts"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/input"
	"github.com/Alecktos/clg/game/view/error_scene"
	"github.com/Alecktos/clg/game/view/game_scene"
	"github.com/hajimehoshi/ebiten/v2"
)

type ClgGame struct {
	gameScene  *game_scene.GameScene
	errorScene *error_scene.ErrorScene
	loadError  error
}

func NewClgGame() *ClgGame {
	clgGame := &ClgGame{}
	clgGame.load()
	return clgGame
}

func (g *ClgGame) load() {
	g.loadError = fonts.LoadFonts()
	g.errorScene = error_scene.NewErrorScene()

	if g.loadError != nil {
		return
	}
	g.gameScene, g.loadError = game_scene.NewGameScene()
	if g.loadError != nil {
		return
	}
}

func (g *ClgGame) Update() error {
	input.Update()
	g.gameScene.Update()
	return nil
}

func (g *ClgGame) Draw(screen *ebiten.Image) {
	// screen.Fill(color.RGBA{R: 40, G: 40, B: 40, A: 0}) // NOt sure alpha is correct
	if g.loadError != nil {
		errorMessage := ""
		if g.loadError != nil {
			errorMessage = g.loadError.Error()
		}
		g.errorScene.Draw(screen, errorMessage)
		return
	}

	g.gameScene.Draw(screen)
}

// Automatically scales.
func (g *ClgGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// fmt.Println(outsideWidth, outsideHeight)
	// 1290/4, 2796/4
	// return 320, 240

	// return outsideWidth, outsideHeight
	return config.WindowWidth, config.WindowHeight
}

func init() {
}
