package materials

import (
	"github.com/Steniaz/go-tracer/raytracer"
	"github.com/Steniaz/go-tracer/math"
)

type DiffuseMaterial struct {
	Reflectance *math.Color3
}

func (mat *DiffuseMaterial) Shade(in *math.Vector3, out *math.Vector3, ray *raytracer.Ray) *math.Color3 {
	return mat.Reflectance
}