package main

type Mesh struct {
	Heights      [][]float64
	XStep, ZStep float64
	Zero         Vec3d
}

func NewMesh(xstep, zstep float64, xsize, zsize int) *Mesh {
	m := new(Mesh)

	m.XStep = xstep
	m.ZStep = zstep
	m.Zero = Vec3d{0, 0, 0}

	m.Heights = make([][]float64, xsize)
	for i := range m.Heights {
		m.Heights[i] = make([]float64, zsize)
	}

	return m
}

type HitRecord struct {
	Where Vec3d
	T     float64
}
