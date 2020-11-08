package materials

import (
	"github.com/stijndehaes/go-tracer/math"
	"github.com/stijndehaes/go-tracer/raytracer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiffuseMaterial_Shade(t *testing.T) {
	mat := DiffuseMaterial{&math.Color3{R: 1.0}}
	in := &math.Vector3{}
	out := &math.Vector3{}
	ray := &raytracer.Ray{}
	shade := mat.Shade(in, out, ray)
	assert.Equal(t, &math.Color3{R: 1.0}, shade)
}
