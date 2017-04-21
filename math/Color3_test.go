package math

import (
	gomath "math"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestColor3_RGBA64(t *testing.T) {
	color3 := Color3{0.0, 0.5, 1.0}.RGBA64()
	assert.Equal(t, uint16(0), color3.R)
	assert.Equal(t, uint16(gomath.MaxUint16) / 2, color3.G)
	assert.Equal(t, uint16(gomath.MaxUint16), color3.B)
	assert.Equal(t, uint16(gomath.MaxUint16), color3.A)
}
