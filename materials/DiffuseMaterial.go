package materials

import (
	"github.com/stijndehaes/go-tracer/math"
	"github.com/stijndehaes/go-tracer/raytracer"
)

type DiffuseMaterial struct {
	Reflectance *math.Color3
}

func (mat *DiffuseMaterial) Shade(in *math.Vector3, out *math.Vector3, ray *raytracer.Ray) *math.Color3 {
	return mat.Reflectance
}
