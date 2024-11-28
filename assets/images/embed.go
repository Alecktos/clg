package images

import (
	_ "embed"
)

var (
	//go:embed brick.png
	brick_png []byte

	//go:embed brick_colors.png
	brick_colors_png []byte
)
