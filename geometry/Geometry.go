package geometry

import (
	"github.com/Steniaz/go-tracer/raytracer"
)

type Geometry interface {
	Hit(ray *raytracer.Ray) bool
}