package math

type Vector4 struct {
	X, Y, Z, T float64
}

func (first Vector4) multiply(second Vector4) Vector4 {
	return Vector4{
		X: first.X * second.X,
		Y: first.Y * second.Y,
		Z: first.Z * second.Z,
		T: first.T * second.T,
	}
}

func (first Vector4) add(second Vector4) Vector4 {
	return Vector4{
		X: first.X + second.X,
		Y: first.Y + second.Y,
		Z: first.Z + second.Z,
		T: first.T + second.T,
	}
}

func (first Vector4) subtract(second Vector4) Vector4 {
	return Vector4{
		X: first.X - second.X,
		Y: first.Y - second.Y,
		Z: first.Z - second.Z,
		T: first.T - second.T,
	}
}
