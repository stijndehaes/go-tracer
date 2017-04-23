package raytracer

import (
	"github.com/Steniaz/go-tracer/math"
)

type Material interface {
	Shade(in *math.Vector3, out *math.Vector3, ray *Ray) *math.Color3
}