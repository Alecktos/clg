package view

import (
	"github.com/Alecktos/clg/game/common"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type RoundedRectangle interface {
	Draw(image *ebiten.Image)
}

type roundedRectangle struct {
	common.Rectangle
	radius     float32
	color      color.Color
	whiteImage *ebiten.Image
}

func NewRoundedRectangle(rectangle common.Rectangle, radius float32, backgroundColor color.Color) (RoundedRectangle, error) {
	whiteImage := ebiten.NewImage(1, 1)
	whiteImage.Fill(color.White)

	return &roundedRectangle{
		Rectangle:  rectangle,
		radius:     radius,
		color:      backgroundColor,
		whiteImage: whiteImage,
	}, nil
}

func (r *roundedRectangle) Draw(image *ebiten.Image) {
	r.drawRoundedRectManually(image)
}

func (r *roundedRectangle) drawRoundedRectManually(screen *ebiten.Image) {
	x := float32(r.Position.X)
	y := float32(r.Position.Y)
	width := float32(r.Width)
	height := float32(r.Height)

	// 1. Define Rounded Rectangle Path
	var p vector.Path
	p.MoveTo(x+r.radius, y)                                 // Starting point top edge, after the first corner
	p.LineTo(x+width-r.radius, y)                           // Top edge
	p.QuadTo(x+width, y, x+width, y+r.radius)               // Upper right corner
	p.LineTo(x+width, y+height-r.radius)                    // Right edge
	p.QuadTo(x+width, y+height, x+width-r.radius, y+height) // Lower right corner
	p.LineTo(x+r.radius, y+height)                          // Bottom edge
	p.QuadTo(x, y+height, x, y+height-r.radius)             // Lower left corner
	p.LineTo(x, y+r.radius)                                 // Left edge
	p.QuadTo(x, y, x+r.radius, y)                           // Upper left corner

	// 2. Generate vertices and indices
	vertices, indices := p.AppendVerticesAndIndicesForFilling(nil, nil)

	//Set rectangle color
	cr, cg, cb, ca := r.color.RGBA()
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
	screen.DrawTriangles(vertices, indices, r.whiteImage, &ebiten.DrawTrianglesOptions{AntiAlias: true, DisableMipmaps: true})
}
