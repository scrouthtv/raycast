package raycast

type HitRecord struct {
	Where  Vec3d
	Normal Vec3d
	T      float64
	Absorb bool
}

// RayPath describes a (multiple times) reflected ray.
// Each "subray" Rays[i] is used from T = 0 to T = Ts[i]
type RayPath struct {
	Rays []Ray
	Ts   []float64
}

func (p *RayPath) add(r *Ray, s *Scene, depth int) {
	ok, hit := s.Hit(r, 0.1, 10)
	if !ok {
		p.Rays = append(p.Rays, *r)
		p.Ts = append(p.Ts, 10)
		return
	}

	p.Rays = append(p.Rays, *r)
	p.Ts = append(p.Ts, hit.T)

	if !hit.Absorb && depth > 0 {
		ray := r.Reflect(hit)
		p.add(&ray, s, depth-1)
	}

	if depth == 0 {
		println("aborted ray")
	}
}

type RayConsumer interface {
	Consume(r *Ray)
}
