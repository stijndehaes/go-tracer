package raytracer

import (
	"github.com/stijndehaes/go-tracer/math"
)

type Integrator interface {
	RenderRay(ray *Ray) *math.Color3
}
