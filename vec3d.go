package main

import "math"

type Vec3d struct {
	X, Y, Z float64
}

func (v Vec3d) Add(v2 Vec3d) Vec3d {
	return Vec3d{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vec3d) Sub(v2 Vec3d) Vec3d {
	return Vec3d{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func (v Vec3d) Mul(s float64) Vec3d {
	return Vec3d{v.X * s, v.Y * s, v.Z * s}
}

func (v Vec3d) Dot(v2 Vec3d) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func (v Vec3d) LenSquared() float64 {
	return v.Dot(v)
}

func (v Vec3d) Len() float64 {
	return math.Sqrt(v.LenSquared())
}

func (v Vec3d) Normalize() Vec3d {
	return v.Mul(1 / v.Len())
}

func (v Vec3d) Cross(v2 Vec3d) Vec3d {
	return Vec3d{v.Y*v2.Z - v.Z*v2.Y, v.Z*v2.X - v.X*v2.Z, v.X*v2.Y - v.Y*v2.X}
}

type Ray struct {
	Origin, Direction Vec3d
}

func (r *Ray) At(t float64) Vec3d {
	return r.Origin.Add(r.Direction.Mul(t))
}
