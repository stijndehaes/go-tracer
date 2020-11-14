package math

import "math"

type Vector3 struct {
	X, Y, Z float64
}

func (first Vector3) Add(second Vector3) Vector3 {
	return Vector3{
		X: first.X + second.X,
		Y: first.Y + second.Y,
		Z: first.Z + second.Z,
	}
}

func (first Vector3) Subtract(second Vector3) Vector3 {
	return Vector3{
		X: first.X - second.X,
		Y: first.Y - second.Y,
		Z: first.Z - second.Z,
	}
}

func (first Vector3) Norm() float64 {
	return math.Sqrt(first.X*first.X + (first.Y * first.Y) + (first.Z * first.Z))
}

func (first Vector3) Normalize() Vector3 {
	normal := first.Norm()
	return Vector3{
		X: first.X / normal,
		Y: first.Y / normal,
		Z: first.Z / normal,
	}
}

func (first Vector3) MultiplyFloat(f float64) Vector3 {
	return Vector3{
		X: first.X * f,
		Y: first.Y * f,
		Z: first.Z * f,
	}
}

func (first Vector3) Negative() Vector3 {
	return Vector3{
		X: -first.X,
		Y: -first.Y,
		Z: -first.Z,
	}
}

func (first Vector3) DotProduct(second Vector3) float64 {
	return (first.X * second.X) + (first.Y * second.Y) + (first.Z * second.Z)
}

func (first Vector3) CrossProduct(second Vector3) Vector3 {
	return Vector3{
		X: (first.Y * second.Z) - (first.Z * second.Y),
		Y: (first.Z * second.X) - (first.X * second.Z),
		Z: (first.X * second.Y) - (first.Y * second.X),
	}
}
