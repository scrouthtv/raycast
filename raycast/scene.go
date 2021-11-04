package raycast

const (
	Epsilon = 0.0000001
)

var (
	targetNormal = Vec3d{0, 1, 0}
)

type Scene struct {
	L *Lamp
	M *Mesh

	// Y is the distance to the target plane on the y axis.
	// The target plane is infinitely large in width and height.
	Y float64
}

func NewScene(l *Lamp, m *Mesh) *Scene {
	s := &Scene{
		L: l,
		M: m,
		Y: -5,
	}

	return s
}

func (s *Scene) Hit(r *Ray, tmin, tmax float64) (bool, *HitRecord) {
	ok1, hit1 := s.M.Hit(r, tmin, tmax)
	ok2, hit2 := s.hitTarget(r, tmin, tmax)
	if ok1 && ok2 {
		if hit1.T < hit2.T {
			return true, hit1
		} else {
			return true, hit2
		}
	} else if ok1 {
		return true, hit1
	} else if ok2 {
		return true, hit2
	} else {
		return false, nil
	}
}

func (s *Scene) hitTarget(r *Ray, tmin, tmax float64) (bool, *HitRecord) {
	// y = r.Origin.Y + t * r.Direction.Y
	// t = (y - r.Origin.Y) / r.Direction.Y
	t := (s.Y - r.Origin.Y) / r.Direction.Y
	if t < tmin || t > tmax {
		return false, nil
	}

	return true, &HitRecord{
		Where:  r.At(t),
		Normal: targetNormal,
		T:      t,
		Absorb: true,
	}
}

func (s *Scene) EachTrace(t TraceConsumer) *RayTracer {
	r := &RayTracer{t, s, 0, 0}
	s.L.eachRay(r)
	return r
}

func (s *Scene) Trace(r *Ray) *RayPath {
	if r == nil {
		println("nil ray")
		return &RayPath{}
	}

	path := RayPath{nil, nil, false}
	path.add(r, s, 30)
	return &path
}
