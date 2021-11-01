package main

// Polygon is a triangle in three-dimensional space.
type Polygon struct {
	Vs [3]Vec3d
}

// Hit tests whether the specified ray hits the polygon.
func (p *Polygon) Hit(r *Ray, tmin, tmax float64) (bool, *HitRecord) {
	edge1 := p.Vs[1].Sub(p.Vs[0])
	edge2 := p.Vs[2].Sub(p.Vs[0])

	h := r.Direction.Cross(edge2)
	a := edge1.Dot(h)

	if a > -Epsilon && a < Epsilon {
		// ray is parallel to the plane of the triangle
		return false, nil
	}

	f := 1.0 / a
	s := r.Origin.Sub(p.Vs[0])
	u := f * s.Dot(h)

	if u < 0.0 || u > 1.0 {
		return false, nil
	}

	q := s.Cross(edge1)
	v := f * r.Direction.Dot(q)
	if v < 0.0 || u+v > 1.0 {
		return false, nil
	}

	// Ray intersects
	t := f * edge2.Dot(q)
	return true, &HitRecord{
		Where: r.At(t),
		T:     t,
	}
}
