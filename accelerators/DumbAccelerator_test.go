package accelerators

import (
	"github.com/stijndehaes/go-tracer/geometry"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDumbAccelerator_AddGeometry(t *testing.T) {
	acc := DumbAccelerator{}

	sphere := geometry.Sphere{}
	acc.AddGeometry(&sphere)
	assert.Equal(t, 1, len(acc.geometries))
}
