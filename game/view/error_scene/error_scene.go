package error_scene

import (
	"github.com/Alecktos/clg/game/config"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func DrawErrorScene(screen *ebiten.Image, error string) {
	screen.Fill(config.MidnightBlue())
	DrawErrorMessage(screen)
	if config.DEV_MODE {
		ebitenutil.DebugPrint(screen, error)
	}
}
