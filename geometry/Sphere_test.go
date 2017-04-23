package geometry

import (
	"testing"
	"github.com/Steniaz/go-tracer/math"
	"github.com/stretchr/testify/assert"
	"github.com/Steniaz/go-tracer/raytracer"
	math2 "math"
)

func TestSphere_Hit(t *testing.T) {
	sphere := Sphere{
		Radius: 1,
	}
	ray := raytracer.Ray{
		Eye: &math.Vector3{ X: -10 },
		Direction: &math.Vector3{X: 1 },
		HitDistance: math2.MaxFloat64,
	}
	sphere.Hit(&ray)
	assert.True(t, ray.Hit)
	assert.Equal(t, 9.0, ray.HitDistance)
}

func TestSphere_Hit2(t *testing.T) {
	sphere := Sphere{
		Radius: 1,
	}
	ray := raytracer.Ray{
		Eye: &math.Vector3{ X: -10 },
		Direction: &math.Vector3{X: 1 , Y: 1},
		HitDistance: math2.MaxFloat64,
	}
	sphere.Hit(&ray)
	assert.False(t, ray.Hit)
}