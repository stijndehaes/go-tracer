package accelerators

import (
	"github.com/Steniaz/go-tracer/geometry"
	"github.com/Steniaz/go-tracer/raytracer"
)

type DumbAccelerator struct {
	geometries []geometry.Geometry
}

func (acc *DumbAccelerator) IntersectGeometry(ray *raytracer.Ray) {
	for _, geometry := range(acc.geometries) {
		geometry.Hit(ray)
	}
}

func (acc *DumbAccelerator) AddGeometry(geometry geometry.Geometry) {
	acc.geometries = append(acc.geometries, geometry)
}

func (acc *DumbAccelerator) build() {}