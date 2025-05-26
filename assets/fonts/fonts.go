package fonts

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	GameFont *text.GoTextFaceSource
)

func LoadFonts() error {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(font1_ttf))
	if err != nil {
		return err
	}
	GameFont = s
	return nil
}
