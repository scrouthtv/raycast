package raycast

import (
	"github.com/fogleman/fauxgl"
)

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

func (m *Mesh) toVec(x, z int) Vec3d {
	var xval float64 = float64(x) * m.XStep
	if z%2 == 1 {
		xval += 0.5 * m.XStep
	}

	return m.Zero.Add(Vec3d{xval, m.Heights[x][z], float64(z) * m.ZStep})
}

func (m *Mesh) AllTris() []Triangle {
	tris := make([]Triangle, 0, len(m.Heights)*len(m.Heights[0]))

	var v0, v1, v2 Vec3d

	for z := 0; z < len(m.Heights[0]); z++ {
		for x := 0; x < len(m.Heights); x++ {
			v0 = m.toVec(x, z)

			if z%2 == 0 {
				// lime triangles
				if z-1 >= 0 && z+1 < len(m.Heights[0]) {
					v1 = m.toVec(x, z-1)
					v2 = m.toVec(x, z+1)
					tris = append(tris, Triangle{[3]Vec3d{v0, v1, v2}})
				}

				// blue triangles
				if z+2 < len(m.Heights[0]) {
					v1 = m.toVec(x, z+1)
					v2 = m.toVec(x, z+2)
					tris = append(tris, Triangle{[3]Vec3d{v0, v1, v2}})
				}
			} else {
				// orange triangles
				if x+1 < len(m.Heights) && z-1 >= 0 && z+1 < len(m.Heights[0]) {
					v1 = m.toVec(x+1, z-1)
					v2 = m.toVec(x+1, z+1)
					tris = append(tris, Triangle{[3]Vec3d{v0, v1, v2}})
				}

				// blue triangles
				if x+1 < len(m.Heights) && z+2 < len(m.Heights[0]) {
					v1 = m.toVec(x+1, z+1)
					v2 = m.toVec(x, z+2)
					tris = append(tris, Triangle{[3]Vec3d{v0, v1, v2}})
				}
			}
		}
	}

	return tris
}

func (m *Mesh) ToGL() *fauxgl.Mesh {
	tris := m.AllTris()
	gltris := make([]*fauxgl.Triangle, len(tris))

	for i, t := range tris {
		gltris[i] = t.ToGL()
	}

	return fauxgl.NewTriangleMesh(gltris)
}

type HitRecord struct {
	Where Vec3d
	T     float64
}
