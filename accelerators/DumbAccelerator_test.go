package accelerators

import (
	"testing"
	"github.com/Steniaz/go-tracer/geometry"
	"github.com/stretchr/testify/assert"
)

func TestDumbAccelerator_AddGeometry(t *testing.T) {
	acc := DumbAccelerator{}

	sphere := geometry.Sphere{}
	acc.AddGeometry(&sphere)
	assert.Equal(t, 1, len(acc.geometries))
}
