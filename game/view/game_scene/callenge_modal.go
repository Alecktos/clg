package game_scene

import (
	"image/color"

	"github.com/Alecktos/clg/game/view"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type ChallengeModal struct {
	whiteImage    *ebiten.Image
	challengeText *view.Text
}

func NewChallengeModal() *ChallengeModal {
	whiteImage := ebiten.NewImage(1, 1)
	whiteImage.Fill(color.White)
	return &ChallengeModal{
		whiteImage:    whiteImage,
		challengeText: view.NewText("Here is challenge for you"),
	}
}

func (cm *ChallengeModal) Draw(screen *ebiten.Image) {

	x1 := float32(30)
	y1 := float32(30)
	width1 := float32(300)
	height1 := float32(500)
	radius1 := float32(25)
	color1 := color.RGBA{R: 0, G: 128, B: 255, A: 255} // Blue color

	cm.drawRoundedRectManually(screen, x1, y1, width1, height1, radius1, color1)

	// Draw challengeText in the rectangle
	cm.challengeText.Draw(screen)
}

func (cm *ChallengeModal) drawRoundedRectManually(screen *ebiten.Image, x, y, width, height, radius float32, clr color.Color) {
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
