package raytracer

import (
	"github.com/stijndehaes/go-tracer/math"
)

type Material interface {
	Shade(in math.Vector3, out math.Vector3, ray *Ray) math.Color3
}
