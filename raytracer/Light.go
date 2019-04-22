package raytracer

import (
	"github.com/stijndehaes/go-tracer/math"
)

type Light interface {
	SampleLight(hit *math.Vector3) (*math.Vector3, float64, *math.Color3)
}
