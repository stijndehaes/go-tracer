package raytracer

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"math"
)

func TestRay_IsHit_no_previous(t *testing.T) {
	ray := Ray{
		HitDistance: math.MaxFloat64,
	}
	result := ray.IsHit(1.0)
	assert.True(t, result)
	assert.True(t, ray.Hit)
	assert.Equal(t, 1.0, ray.HitDistance)
}

func TestRay_IsHit_with_previous(t *testing.T) {
	ray := Ray{
		HitDistance: 0.5,
	}
	result := ray.IsHit(1.0)
	assert.False(t, result)
	assert.False(t, ray.Hit)
}

func TestRay_IsHit_MinDistance(t *testing.T) {
	ray := Ray{
		HitDistance: math.MaxFloat64,
	}
	result := ray.IsHit(0.00000000000000001)
	assert.False(t, result)
	assert.False(t, ray.Hit)
}