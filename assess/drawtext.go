package main

import (
	_ "embed"
	"image"
	"image/color"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

func (t *TargetDrawer) initFont() error {
	fdata, err := os.ReadFile("osifont.ttf")
	if err != nil {
		return err
	}

	f, err := freetype.ParseFont(fdata)
	if err != nil {
		return err
	}

	t.text = freetype.NewContext()
	t.text.SetClip(t.Img.Bounds())
	t.text.SetDPI(72)
	t.text.SetDst(&t.Img)
	t.text.SetFont(f)
	t.text.SetFontSize(16)
	t.text.SetHinting(font.HintingFull)
	t.text.SetSrc(image.NewUniform(color.RGBA{255, 0, 0, 255}))

	return nil
}
