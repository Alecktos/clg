package config

import (
	"image/color"
	"runtime"
)

type DeviceType int64

const (
	DeviceTypeMobile DeviceType = iota
	DeviceTypeDesktop
)

const (
	WindowWidth  = 1290
	WindowHeight = 2532

	DevMode = true
)

const (
	StandardFontSize = 90
	HeaderFontSize   = 120
)

var cachedDeviceType *DeviceType

func DetectDevicetype() DeviceType {
	if cachedDeviceType != nil {
		return *cachedDeviceType
	}
	os := runtime.GOOS
	var deviceType DeviceType
	if os == "ios" {
		deviceType = DeviceTypeMobile
	} else {
		deviceType = DeviceTypeDesktop
	}
	cachedDeviceType = &deviceType
	return deviceType
}

type ClgColor struct {
	color.Color
}

// -- Primary Colors --

// Main buttons, highlights, challenges
func PassionRed() ClgColor {
	return ClgColor{color.RGBA{R: 215, G: 38, B: 56, A: 255}} // #D72638
}

// Backgrounds, overlays, cards
func VelvetPlum() ClgColor {
	return ClgColor{color.RGBA{R: 94, G: 42, B: 93, A: 255}} // #5E2A5D
}

// -- Secondary Colors --

// Hover effects, soft UI accents
func BlushPink() ClgColor {
	return ClgColor{color.RGBA{R: 247, G: 141, B: 167, A: 255}} // #F78DA7
}

// Deep backgrounds, contrast zones
func MidnightBlue() ClgColor {
	return ClgColor{color.RGBA{R: 28, G: 28, B: 58, A: 255}} // #1C1C3A
}

// Decorative lines, elegant text
func ChampagneGold() ClgColor {
	return ClgColor{color.RGBA{R: 244, G: 226, B: 216, A: 255}} // #F4E2D8
}

// -- Accent / Feedback Colors --

// Progress bars, timer, glow effects
func TeasePurple() ClgColor {
	return ClgColor{color.RGBA{R: 166, G: 120, B: 180, A: 255}} // #A678B4
}

// Background gradients, intimacy
func SkinToneBeige() ClgColor {
	return ClgColor{color.RGBA{R: 255, G: 214, B: 192, A: 255}} // #FFD6C0
}

// Text, icons, subtle shadows
func ForbiddenBlack() ClgColor {
	return ClgColor{color.RGBA{R: 17, G: 17, B: 17, A: 255}} // #111111
}
