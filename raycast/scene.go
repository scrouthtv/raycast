package raycast

import "image"

const (
	Epsilon = 0.0000001
)

type Scene struct {
	L *Lamp
	M *Mesh

	// Y is the distance to the target plane on the y axis.
	Y float64

	Img *image.RGBA
}

func NewScene(l *Lamp, m *Mesh, w, h int) *Scene {
	s := &Scene{
		L:   l,
		M:   m,
		Y:   5,
		Img: image.NewRGBA(image.Rect(0, 0, w, h)),
	}

	return s
}

func (s *Scene) Draw() {

}

type RayConsumer interface {
	Consume(r *Ray)
}
