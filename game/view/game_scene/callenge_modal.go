package game_scene

import (
	"image/color"

	"github.com/Alecktos/clg/game/common"
	"github.com/Alecktos/clg/game/config"
	"github.com/Alecktos/clg/game/view"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type ChallengeModal struct {
	whiteImage    *ebiten.Image
	challengeText *view.Text
	rectangle     *common.Rectangle
}

func NewChallengeModal() (*ChallengeModal, error) {
	whiteImage := ebiten.NewImage(1, 1)
	whiteImage.Fill(color.White)

	width := float64(config.WindowWidth - 60)
	height := float64(300)
	x1 := 30.0
	y1 := config.WindowHeight/2 - height/2

	challengeText, err := view.NewCenterAlignedText("Here is challenge for you", y1)
	if err != nil {
		return nil, err
	}

	return &ChallengeModal{
		whiteImage:    whiteImage,
		challengeText: challengeText,
		rectangle:     &common.Rectangle{Position: common.Position{X: x1, Y: y1}, Width: width, Height: height},
	}, nil
}

func (cm *ChallengeModal) Draw(screen *ebiten.Image) {

	radius1 := float32(25)
	color1 := color.RGBA{R: 0, G: 128, B: 255, A: 255} // Blue color

	cm.drawRoundedRectManually(screen, radius1, color1)

	// Draw challengeText in the rectangle
	cm.challengeText.Draw(screen)
}

func (cm *ChallengeModal) drawRoundedRectManually(screen *ebiten.Image, radius float32, clr color.Color) {
	x := float32(cm.rectangle.Position.X)
	y := float32(cm.rectangle.Position.Y)
	width := float32(cm.rectangle.Width)
	height := float32(cm.rectangle.Height)

	// 1. Define Rounded Rectangle Path
	var p vector.Path
	p.MoveTo(x+radius, y)                                 // Starting point top edge, after the first corner
	p.LineTo(x+width-radius, y)                           // Top edge
	p.QuadTo(x+width, y, x+width, y+radius)               // Upper right corner
	p.LineTo(x+width, y+height-radius)                    // Right edge
	p.QuadTo(x+width, y+height, x+width-radius, y+height) // Lower right corner
	p.LineTo(x+radius, y+height)                          // Bottom edge
	p.QuadTo(x, y+height, x, y+height-radius)             // Lower left corner
	p.LineTo(x, y+radius)                                 // Left edge
	p.QuadTo(x, y, x+radius, y)                           // Upper left corner

	// 2. Generate vertices and indices
	vertices, indices := p.AppendVerticesAndIndicesForFilling(nil, nil)

	//Set rectangle color
	cr, cg, cb, ca := clr.RGBA()
	colorR := float32(cr) / 0xffff
	colorG := float32(cg) / 0xffff
	colorB := float32(cb) / 0xffff
	colorA := float32(ca) / 0xffff

	for i := range vertices {
		vertices[i].SrcX = 0 // point to  (0,0) (g.whiteImage)
		vertices[i].SrcY = 0
		vertices[i].ColorR = colorR // Set vertex color
		vertices[i].ColorG = colorG
		vertices[i].ColorB = colorB
		vertices[i].ColorA = colorA
	}

	//Mipmap might be expensive on mobile devices, so we disable it.
	screen.DrawTriangles(vertices, indices, cm.whiteImage, &ebiten.DrawTrianglesOptions{AntiAlias: true, DisableMipmaps: true})
}
