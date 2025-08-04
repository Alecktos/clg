package game

import (
	"github.com/Alecktos/clg/assets/fonts"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/input"
	"github.com/hajimehoshi/ebiten/v2"
)

type ClgGame struct {
	sceneManager SceneManager
}

func NewClgGame() *ClgGame {
	clgGame := &ClgGame{
		sceneManager: NewSceneManager(),
	}
	clgGame.load()
	return clgGame
}

func (g *ClgGame) load() {
	loadError := fonts.LoadFonts()

	g.sceneManager.Load(loadError)

}

func (g *ClgGame) Update() error {
	input.Update()
	g.sceneManager.Update()
	return nil
}

func (g *ClgGame) Draw(screen *ebiten.Image) {
	g.sceneManager.Draw(screen)
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
