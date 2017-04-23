package math

import "math"

type Vector3 struct {
	X, Y, Z float64
}

func (first *Vector3) Multiply(second *Vector3) *Vector3 {
	return &Vector3{
		X: first.X * second.X,
	 	Y: first.Y * second.Y,
		Z: first.Z * second.Z,
	}
}

func (first *Vector3) Add(second *Vector3) *Vector3 {
	return &Vector3{
		X: first.X + second.X,
		Y: first.Y + second.Y,
		Z: first.Z + second.Z,
	}
}

func (first *Vector3) Subtract(second *Vector3) *Vector3 {
	return &Vector3{
		X: first.X - second.X,
		Y: first.Y - second.Y,
		Z: first.Z - second.Z,
	}
}

func (vector *Vector3) Norm() float64 {
	return math.Sqrt( vector.X * vector.X + (vector.Y * vector.Y) + (vector.Z * vector.Z))
}

func (vector *Vector3) Normalize() *Vector3 {
	normal := vector.Norm()
	return &Vector3{
		X: vector.X / normal,
		Y: vector.Y / normal,
		Z: vector.Z / normal,
	}
}

func (vector *Vector3) MultiplyFloat(f float64) *Vector3 {
	return &Vector3{
		X: vector.X * f,
		Y: vector.Y * f,
		Z: vector.Z * f,
	}
}

func (vector *Vector3) Negative() *Vector3 {
	return &Vector3{
		X: -vector.X,
		Y: -vector.Y,
		Z: -vector.Z,
	}
}

func (first *Vector3) DotProduct(second *Vector3) float64 {
	return first.X * second.X + (first.Y * second.Y) + (first.Z * second.Z)
}

func (first *Vector3) CrossProduct(second *Vector3) *Vector3 {
	return &Vector3{
		X: (first.Y * second.Z) - (first.Z * second.Y),
		Y: (first.Z * second.X) - (first.X * second.Z),
		Z: (first.X * second.Y) - (first.Y * second.X),
	}
}