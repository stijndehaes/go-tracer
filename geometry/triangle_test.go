package geometry

import (
	"github.com/stijndehaes/go-tracer/materials"
	"github.com/stijndehaes/go-tracer/math"
	"github.com/stijndehaes/go-tracer/raytracer"
	"github.com/stretchr/testify/assert"
	math2 "math"
	"testing"
)

func TestTriangle_Hit(t *testing.T) {
	triangle := NewTriangle(
		math.Vector3{
			X: 0,
			Y: 100,
			Z: -100,
		},
		math.Vector3{
			X: 0,
			Y: 100,
			Z: 100,
		},
		math.Vector3{
			X: 0,
			Y: -100,
			Z: -100,
		},
		&materials.DiffuseMaterial{},
	)
	ray := &raytracer.Ray{
		Eye:         math.Vector3{X: -10},
		Direction:   math.Vector3{X: 1},
		HitDistance: math2.MaxFloat64,
	}
	triangle.Hit(ray)
	assert.True(t, ray.Hit)
	assert.Equal(t, 10.0, ray.HitDistance)
}
