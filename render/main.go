package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fogleman/fauxgl"
	"github.com/scrouthtv/raytracing/raycast"
)

const (
	fov  = 30
	near = .5
	far  = 10
)

var (
	eye    = fauxgl.V(3, 1, .75)
	center = fauxgl.V(0, 0, 0)
	up     = fauxgl.V(0, 0, 1)
	light  = fauxgl.V(-.75, 1, .25).Normalize()
)

func main() {
	ctx := fauxgl.NewContext(1920, 1080)
	ctx.ClearColorBufferWith(fauxgl.HexColor("#FFF8E3"))

	matrix := fauxgl.LookAt(eye, center, up).Perspective(fov, 1920.0/1080.0, near, far)
	shader := fauxgl.NewPhongShader(matrix, light, eye)
	shader.ObjectColor = fauxgl.HexColor("#468966")
	ctx.Shader = shader

	/*mesh, _ := fauxgl.LoadSTL("taste01niwwer.stl")
	mesh.BiUnitCube()
	ctx.DrawMesh(mesh)*/

	x := fauxgl.NewLineForPoints(fauxgl.V(0, 0, 0), fauxgl.V(1, 0, 0))
	y := fauxgl.NewLineForPoints(fauxgl.V(0, 0, 0), fauxgl.V(0, 1, 0))
	z := fauxgl.NewLineForPoints(fauxgl.V(0, 0, 0), fauxgl.V(0, 0, 1))
	marks := []*fauxgl.Line{
		fauxgl.NewLineForPoints(fauxgl.V(.2, -.1, -.1), fauxgl.V(.2, .1, .1)),
		fauxgl.NewLineForPoints(fauxgl.V(-.1, .2, -.1), fauxgl.V(.1, .2, .1)),
		fauxgl.NewLineForPoints(fauxgl.V(-.1, .3, -.1), fauxgl.V(.1, .3, .1)),
		fauxgl.NewLineForPoints(fauxgl.V(-.1, -.1, .2), fauxgl.V(.1, .1, .2)),
		fauxgl.NewLineForPoints(fauxgl.V(-.1, -.1, .3), fauxgl.V(.1, .1, .3)),
		fauxgl.NewLineForPoints(fauxgl.V(-.1, -.1, .4), fauxgl.V(.1, .1, .4)),
	}

	ctx.DrawLine(x)
	ctx.DrawLine(y)
	ctx.DrawLine(z)
	for _, l := range marks {
		ctx.DrawLine(l)
	}

	f, err := os.Open("test01.htb")
	if err != nil {
		fmt.Println(err)
		return
	}

	mesh, err := raycast.LoadMesh(bufio.NewReader(f))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Loaded", len(mesh.AllTris()), "triangles.")
	//fmt.Println(mesh.AllTris())
	fmesh := mesh.ToGL()
	fmesh.SaveSTL("test01.stl")

	//mesh.Zero = raycast.Vec3d{}
	info := ctx.DrawMesh(fmesh)
	fmt.Println(info)

	fauxgl.SavePNG("out.png", ctx.Image())
}
