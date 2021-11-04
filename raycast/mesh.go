package raycast

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

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

// LoadMesh loads a hight profile from an .htb file.
// The file must have a trailing newline.
func LoadMesh(r *bufio.Reader) (*Mesh, error) {
	l, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	var xstep, zstep float64
	_, err = fmt.Sscanf(l, "%f %f", &xstep, &zstep)
	if err != nil {
		return nil, err
	}

	lines := make([]string, 0)

	l, err = r.ReadString('\n')
	for err == nil {
		lines = append(lines, strings.TrimRight(l, "\r\n"))
		l, err = r.ReadString('\n')
	}

	if len(lines) == 0 {
		return nil, errors.New("empty file")
	}

	xsize := strings.Count(lines[0], " ") + 1

	m := NewMesh(xstep, zstep, xsize, len(lines))

	for z, l := range lines {
		hs := strings.Split(l, " ")
		if len(hs) != xsize {
			return nil, &ErrMismatchingXsize{z, xsize, len(hs)}
		}

		for x, h := range hs {
			hf, err := strconv.ParseFloat(h, 64)
			if err != nil {
				return nil, &ErrBadNumber{x + 1, h, err}
			}

			m.Heights[x][z] = hf
		}
	}

	return m, nil
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

func (m *Mesh) Hit(r *Ray, tmin, tmax float64) (ok bool, rec *HitRecord) {
	ok = false
	rec = &HitRecord{T: tmax + 1}

	var myok bool
	var myrec *HitRecord
	for _, tri := range m.AllTris() {
		myok, myrec = tri.Hit(r, tmin, tmax)
		if !myok {
			continue
		}

		if myrec.T < rec.T {
			ok = true
			rec = myrec
		}
	}

	if !ok {
		return false, nil
	}

	return
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
	Where  Vec3d
	Normal Vec3d
	T      float64
}
