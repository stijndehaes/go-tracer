package raytracer

import (
	"testing"
	"github.com/Steniaz/go-tracer/math"
	"github.com/stretchr/testify/assert"
	gomath "math"
)

func TestCamera_GenerateRay(t *testing.T) {
	camera := GetCamera(
		&math.Vector3{},
		&math.Vector3{Y: 1.0},
		&math.Vector3{X: 1.0},
		90.0,
		1.0,
		3,
		3,
	)
	ray := camera.GenerateRay(1, 1)
	assert.Equal(t, math.Vector3{}, *ray.Eye)
	assert.Equal(t, gomath.MaxFloat64 , ray.HitDistance)
}

func TestCamera_getRayDirection(t *testing.T) {
	camera := GetCamera(
		&math.Vector3{},
		&math.Vector3{Y: 1.0},
		&math.Vector3{X: 1.0},
		90.0,
		1.0,
		3,
		3,
	)
	direction := camera.getRayDirection(1, 1)
	assert.Equal(t, math.Vector3{X:1.0}, *direction)
}