package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/fogleman/fauxgl"
	"github.com/scrouthtv/raytracing/raycast"
)

const (
	fov    = 30
	near   = .5
	far    = 30
	frames = 100
)

var (
	eye    = fauxgl.V(5, 0, 1)
	center = fauxgl.V(0, 0, 0.5)
	up     = fauxgl.V(0, 0, 1)
	light  = fauxgl.V(0, -1, 0).Normalize()
)

func main() {
	for t := 0.0; t < frames; t++ {
		eye = fauxgl.V(8*math.Cos(2*t*math.Pi/frames), 8*math.Sin(2*t*math.Pi/frames), 1.5)
		saveAs(fmt.Sprintf("out/%03.0f.png", t))
		fmt.Printf("%.0f/%d\n", t, frames)
	}
}

func saveAs(name string) {
	ctx := fauxgl.NewContext(1920, 1080)
	ctx.ClearColorBufferWith(fauxgl.HexColor("#87CEEB"))

	matrix := fauxgl.LookAt(eye, center, up).Perspective(fov, 1920.0/1080.0, near, far)
	shader := fauxgl.NewPhongShader(matrix, light, eye)
	shader.ObjectColor = fauxgl.HexColor("#cddc39")
	ctx.Shader = shader

	x := fauxgl.NewLineForPoints(fauxgl.V(0, 0, 0), fauxgl.V(1, 0, 0))
	y := fauxgl.NewLineForPoints(fauxgl.V(0, 0, 0), fauxgl.V(0, 1, 0))
	z := fauxgl.NewLineForPoints(fauxgl.V(0, 0, 0), fauxgl.V(0, 0, 1))

	ctx.DrawLine(x)
	ctx.DrawLine(y)
	ctx.DrawLine(z)

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

	mesh.Zero = raycast.Vec3d{X: -1, Y: .5, Z: .5}

	lamp := raycast.Lamp{Pos: raycast.Vec3d{X: 0, Y: 1.5, Z: 1.2}, Horizontal: 1, Vertical: 1}

	scene := raycast.NewScene(&lamp, mesh, 1920, 1080)

	rd := RayDrawer{ctx, scene, 0, 0}
	lamp.EachRay(&rd)
	fmt.Printf("%d/%d absorbed\n", rd.absorbed, rd.total)

	//rd.Consume(&raycast.Ray{Origin: lamp.Pos, Direction: raycast.Vec3d{X: 0, Y: -1, Z: 0}})

	fmt.Println("Loaded", len(mesh.AllTris()), "triangles.")
	fmesh := mesh.ToGL()
	//fmesh.SaveSTL("test01.stl")

	ctx.DrawMesh(fmesh)

	fauxgl.SavePNG(name, ctx.Image())
}

type RayDrawer struct {
	ctx             *fauxgl.Context
	scene           *raycast.Scene
	absorbed, total int
}

func (d *RayDrawer) Consume(r *raycast.Ray) {
	d.total++
	trace := d.scene.Trace(r)
	if !trace.Absorbed {
		return
	}
	d.absorbed++

	for i, r := range trace.Rays {
		d.draw(&r, trace.Ts[i])
	}
}

func (d *RayDrawer) draw(r *raycast.Ray, t float64) {
	a := r.At(0).ToGl()
	b := r.At(t).ToGl()
	l := fauxgl.NewLineForPoints(*a, *b)
	d.ctx.DrawLine(l)
}
