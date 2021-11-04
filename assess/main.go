package main

import (
	"bufio"
	"fmt"
	"image/png"
	"os"

	"github.com/scrouthtv/raytracing/raycast"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s file.htb\n", os.Args[0])
		os.Exit(2)
	}

	lamp := raycast.Lamp{Pos: raycast.Vec3d{X: 0, Y: 1.5, Z: 1.2}, Horizontal: 1, Vertical: 1}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, "error opening", os.Args[1]+":", err.Error())
		os.Exit(2)
	}
	mesh, err := raycast.LoadMesh(bufio.NewReader(f))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	mesh.Zero = raycast.Vec3d{X: -1, Y: .5, Z: .5}

	s := raycast.NewScene(&lamp, mesh)

	t, err := NewTargetDrawer(s, 1920, 1080, 300)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	t.Prepare()

	s.EachTrace(t)

	//t.dot(1378, 1061, color.RGBA{0, 255, 0, 255})

	outf, err := os.Create(os.Args[1] + ".png")
	if err != nil {
		fmt.Fprint(os.Stderr, "error opening output file", os.Args[1]+".png:", err)
		os.Exit(2)
	}

	png.Encode(outf, &t.Img)

	stat, err := outf.Stat()
	if err == nil {
		fmt.Printf("%d kB saved to %s.\n", stat.Size()/1024, os.Args[1]+".png")
	}
}
