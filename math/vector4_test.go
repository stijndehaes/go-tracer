package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddVector4(t *testing.T) {
	vector1 := Vector4{X: 1, Y: 1, Z: 1, T: 1}
	vector2 := Vector4{X: 1, Y: 1, Z: 1, T: 1}
	vector3 := vector1.add(vector2)
	assert.Equal(t, 2.0, vector3.X)
	assert.Equal(t, 2.0, vector3.Y)
	assert.Equal(t, 2.0, vector3.Z)
	assert.Equal(t, 2.0, vector3.T)
}

func TestMultiplyVector4(t *testing.T) {
	vector1 := Vector4{X: 1, Y: 1, Z: 1, T: 1}
	vector2 := Vector4{X: 1, Y: 2, Z: 3, T: 4}
	vector3 := vector1.multiply(vector2)
	assert.Equal(t, 1.0, vector3.X)
	assert.Equal(t, 2.0, vector3.Y)
	assert.Equal(t, 3.0, vector3.Z)
	assert.Equal(t, 4.0, vector3.T)
}

func TestSubtractVector4(t *testing.T) {
	vector1 := Vector4{X: 1, Y: 1, Z: 1, T: 1}
	vector2 := Vector4{X: 1, Y: 1, Z: 1, T: 1}
	vector3 := vector1.subtract(vector2)
	assert.Equal(t, 0.0, vector3.X)
	assert.Equal(t, 0.0, vector3.Y)
	assert.Equal(t, 0.0, vector3.Z)
	assert.Equal(t, 0.0, vector3.T)

}
