package main

import (
	"image"
	"image/color"
	"math"
	"strconv"

	"github.com/golang/freetype"
	"github.com/scrouthtv/raytracing/raycast"
)

var (
	Radius = 5
)

type TargetDrawer struct {
	Img        image.RGBA
	midx, midy int

	S *raycast.Scene

	PPU float64

	text *freetype.Context
}

func NewTargetDrawer(s *raycast.Scene, w, h int, ppu float64) (*TargetDrawer, error) {
	t := &TargetDrawer{
		Img:  *image.NewRGBA(image.Rect(0, 0, w, h)),
		midx: w / 2,
		midy: h / 2,
		S:    s,
		PPU:  ppu,
	}

	err := t.initFont()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *TargetDrawer) Prepare() {
	for x := t.Img.Bounds().Min.X; x <= t.Img.Bounds().Max.X; x++ {
		for y := t.Img.Bounds().Min.Y; y <= t.Img.Bounds().Max.Y; y++ {
			t.Img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}

	// Draw the center mark:
	t.vline(t.midx, t.midy-5, t.midy+5, color.RGBA{255, 0, 0, 255})
	t.hline(t.midy, t.midx-5, t.midx+5, color.RGBA{255, 0, 0, 255})

	t.grid(int(t.PPU)/2, color.RGBA{200, 200, 200, 255})

	for _, angle := range []float64{5, 10, 20} {
		t.aoh(angle, color.RGBA{255, 0, 0, 255}, true)
	}

	for _, angle := range []float64{-5, 5, 10} {
		t.aov(angle, color.RGBA{255, 0, 0, 255}, false)
	}
}

func (t *TargetDrawer) aoh(angle float64, c color.Color, mirror bool) {
	x := int(t.S.Y * math.Tan(angle*math.Pi/180) * t.PPU)
	t.vline(t.midx+x, t.midy-5, t.midy+5, c)
	t.text.DrawString(strconv.FormatFloat(angle, 'f', 1, 64), freetype.Pt(t.midx+x, t.midy+30))

	if mirror {
		t.vline(t.midx-x, t.midy-5, t.midy+5, c)
		t.text.DrawString(strconv.FormatFloat(-angle, 'f', 1, 64), freetype.Pt(t.midx-x, t.midy+30))
	}
}

func (t *TargetDrawer) aov(angle float64, c color.Color, mirror bool) {
	y := int(t.S.Y * math.Tan(angle*math.Pi/180) * t.PPU)
	t.hline(t.midy+y, t.midx-5, t.midx+5, c)
	t.text.DrawString(strconv.FormatFloat(angle, 'f', 1, 64), freetype.Pt(t.midx+30, t.midy+y))

	if mirror {
		t.hline(t.midy-y, t.midx-5, t.midx+5, c)
		t.text.DrawString(strconv.FormatFloat(-angle, 'f', 1, 64), freetype.Pt(t.midx+30, t.midy-y))
	}
}

func (t *TargetDrawer) vline(x, y0, y1 int, c color.Color) {
	for y := y0; y < y1; y++ {
		t.Img.Set(x, y, c)
	}
}

func (t *TargetDrawer) hline(y, x0, x1 int, c color.Color) {
	for x := x0; x < x1; x++ {
		t.Img.Set(x, y, c)
	}
}

func (t *TargetDrawer) grid(step int, c color.Color) {
	for x := t.midx % step; x <= t.Img.Bounds().Max.X; x += step {
		t.vline(x, t.Img.Bounds().Min.Y, t.Img.Bounds().Max.Y, c)
	}

	for y := t.midy % step; y <= t.Img.Bounds().Max.Y; y += step {
		t.hline(y, t.Img.Bounds().Min.X, t.Img.Bounds().Max.X, c)
	}
}

func (t *TargetDrawer) dot(x0, y0 int, c color.Color) {
	for x := -Radius; x <= Radius; x++ {
		for y := -Radius; y <= Radius; y++ {
			if x*x+y*y <= Radius*Radius {
				t.Img.Set(x0+x, y0+y, c)
			}
		}
	}
}

func (t *TargetDrawer) Consume(r *raycast.RayPath) {
	if r.Absorbed {
		endp := r.EndPoint()
		if endp.Y-t.S.Y > 0.1 {
			println("ignore ray that was absorbed somewhere else:", endp.Y)
		}

		//println(endp.X, endp.Z)
		x := t.midx + int(endp.X*t.PPU)
		y := t.midy + int(endp.Z*t.PPU)
		println(x, y)

		//t.Img.Set(x, y, color.RGBA{0, 0, 255, 255})
		t.dot(x, y, color.RGBA{0, 0, 255, 255})
	}
}
