package main

type Mesh struct {
	Heights        [][]float64
	XScale, ZScale float64
	Zero           Vec3d
}

func NewMesh(xscale, zscale float64, xsize, zsize int) *Mesh {
	m := new(Mesh)

	m.XScale = xscale
	m.ZScale = zscale
	m.Zero = Vec3d{0, 0, 0}

	m.Heights = make([][]float64, xsize)
	for i := range m.Heights {
		m.Heights[i] = make([]float64, zsize)
	}

	return m
}
