package raycast

import (
	"testing"
)

func TestToV(t *testing.T) {
	m := NewMesh(1, 1, 5, 5)
	for x := 0; x < 3; x++ {
		for z := 0; z < 3; z++ {
			t.Logf("%d/%d: %v", x, z, m.toVec(x, z))
		}
	}
}

func TestAllTris(t *testing.T) {
	m := NewMesh(1, 1, 2, 7)
	m.Heights = [][]float64{
		{3, 5, 2, 7, 4, 6},
		{2, 4, 4, 6, 5, 4},
	}

	t.Log(" x1 z1 -- x2 z2 -- x3 z3 ")
	tris := m.AllTris()
	for _, tri := range tris {
		t.Logf("%3.1f %.1f -- %.1f %.1f -- %.1f %.1f", tri.Vs[0].X, tri.Vs[0].Z, tri.Vs[1].X, tri.Vs[1].Z, tri.Vs[2].X, tri.Vs[2].Z)
	}
}
