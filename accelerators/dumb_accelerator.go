package accelerators

import (
	"github.com/stijndehaes/go-tracer/raytracer"
)

type DumbAccelerator struct {
	geometries []raytracer.Geometry
}

func (acc *DumbAccelerator) IntersectGeometry(ray *raytracer.Ray) {
	for _, geometry := range acc.geometries {
		geometry.Hit(ray)
	}
}

func (acc *DumbAccelerator) AddGeometry(geometry raytracer.Geometry) {
	acc.geometries = append(acc.geometries, geometry)
}

func (acc *DumbAccelerator) Build() {}
