package accelerators

import (
	"github.com/Steniaz/go-tracer/raytracer"
	"github.com/Steniaz/go-tracer/geometry"
)

type Accelerator interface {
	IntersectGeometry(ray raytracer.Ray)

	AddGeometry(geometry geometry.Geometry)

	build()
}