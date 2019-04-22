package lights

import (
	"github.com/stijndehaes/go-tracer/math"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointLight_SampleLight(t *testing.T) {
	light := PointLight{
		Color:  &math.Color3{R: 1.0},
		Origin: &math.Vector3{0, 10, 0},
	}
	hit := &math.Vector3{0, 1, 0}

	direction, distance, color := light.SampleLight(hit)

	assert.Equal(t, &math.Vector3{Y: 1.0}, direction)
	assert.Equal(t, distance, 9.0)
	assert.Equal(t, &math.Color3{R: 1.0 / (9.0 * 9.0)}, color)
}
