package raytracer

import (
	"github.com/Steniaz/go-tracer/math"
)

type Integrator interface {
	RenderRay(ray Ray) *math.Color3
}