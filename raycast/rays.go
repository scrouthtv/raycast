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
	Rays     []Ray
	Ts       []float64
	Absorbed bool
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

	if hit.Absorb {
		p.Absorbed = true
	}
}

func (p *RayPath) EndPoint() Vec3d {
	i := len(p.Rays) - 1
	return p.Rays[i].At(p.Ts[i])
}

type RayTracer struct {
	T               TraceConsumer
	S               *Scene
	Total, Absorbed int
}

func (t *RayTracer) Consume(r *Ray) {
	t.Total++
	trace := t.S.Trace(r)
	if !trace.Absorbed {
		return
	}

	t.Absorbed++

	t.T.Consume(trace)
}

type RayConsumer interface {
	Consume(r *Ray)
}

type TraceConsumer interface {
	Consume(t *RayPath)
}

type LineConsumer interface {
	Consume(r *Ray, tmin, tmax float64)
}
