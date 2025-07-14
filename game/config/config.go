package config

import "image/color"

type DeviceType int64

const (
	DeviceTypeMobile DeviceType = iota
	DeviceTypeDesktop
)

const (
	WindowWidth  = 1290 / 4
	WindowHeight = 2532 / 4

	DEV_MODE = true
)

const (
	STANDARD_FONT_SIZE = 20
)

// -- Primary Colors --

// Main buttons, highlights, challenges
func PassionRed() color.Color {
	return color.RGBA{R: 215, G: 38, B: 56, A: 255} // #D72638
}

// Backgrounds, overlays, cards
func VelvetPlum() color.Color {
	return color.RGBA{R: 94, G: 42, B: 93, A: 255} // #5E2A5D
}

// -- Secondary Colors --

// Hover effects, soft UI accents
func BlushPink() color.Color {
	return color.RGBA{R: 247, G: 141, B: 167, A: 255} // #F78DA7
}

// Deep backgrounds, contrast zones
func MidnightBlue() color.Color {
	return color.RGBA{R: 28, G: 28, B: 58, A: 255} // #1C1C3A
}

// Decorative lines, elegant text
func ChampagneGold() color.Color {
	return color.RGBA{R: 244, G: 226, B: 216, A: 255} // #F4E2D8
}

// -- Accent / Feedback Colors --

// Progress bars, timer, glow effects
func TeasePurple() color.Color {
	return color.RGBA{R: 166, G: 120, B: 180, A: 255} // #A678B4
}

// Background gradients, intimacy
func SkinToneBeige() color.Color {
	return color.RGBA{R: 255, G: 214, B: 192, A: 255} // #FFD6C0
}

// Text, icons, subtle shadows
func ForbiddenBlack() color.Color {
	return color.RGBA{R: 17, G: 17, B: 17, A: 255} // #111111
}

const CurrentDevice DeviceType = DeviceTypeDesktop
