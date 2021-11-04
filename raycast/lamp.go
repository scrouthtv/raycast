package raycast

const (
	// RayStep indicates how many the distance between cast rays
	// on a plane 1 unit away.
	RayStep = 0.2
)

type Lamp struct {
	Pos Vec3d

	// Horizontal and Vertical describe how many units
	// the illuminated area of the lamp stretches out
	// to the left/right and up/down on a imaginary
	// plane 1 unit away. The entire illuminated area is
	// 2*Horizontal * 2*Vertical units^2.
	Horizontal, Vertical float64
}

func (l *Lamp) EachRay(c RayConsumer) {
	ray := Ray{l.Pos, Vec3d{0, 0, 0}}

	for x := -l.Horizontal; x <= l.Horizontal; x += RayStep {
		for y := -l.Vertical; y <= l.Vertical; y += RayStep {
			ray.Direction = Vec3d{x, y, -1}.Normalize()
			c.Consume(&ray)
		}
	}

}
