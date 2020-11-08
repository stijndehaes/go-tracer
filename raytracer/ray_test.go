package raytracer

import (
	math2 "github.com/stijndehaes/go-tracer/math"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestRay_IsHit_no_previous(t *testing.T) {
	ray := Ray{
		HitDistance: math.MaxFloat64,
	}
	ray.IsHit(1.0, nil)
	assert.True(t, ray.Hit)
	assert.Equal(t, 1.0, ray.HitDistance)
}

func TestRay_IsHit_with_previous(t *testing.T) {
	ray := Ray{
		HitDistance: 0.5,
	}
	ray.IsHit(1.0, nil)
	assert.False(t, ray.Hit)
}

func TestRay_IsHit_MinDistance(t *testing.T) {
	ray := Ray{
		HitDistance: math.MaxFloat64,
	}
	ray.IsHit(0.00000000000000001, nil)
	assert.False(t, ray.Hit)
}

func TestRay_HitPoint(t *testing.T) {
	ray := Ray{
		Eye:         &math2.Vector3{},
		Direction:   &math2.Vector3{X: 1},
		HitDistance: 1,
	}
	hitPoint := ray.HitPoint()
	assert.Equal(t, &math2.Vector3{X: 1}, hitPoint)
}
