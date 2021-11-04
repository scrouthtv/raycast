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

// RayPath describes a (multiple times) reflected ray.
// Each "subray" Rays[i] is used from T = 0 to T = Ts[i]
type RayPath struct {
	Rays []Ray
	Ts   []float64
}

func (s *Scene) Trace(r *Ray) *RayPath {
	return nil
}

type RayConsumer interface {
	Consume(r *Ray)
}
