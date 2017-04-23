package math

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestVector3_Add(t *testing.T) {
	vector1 := Vector3{ X: 1, Y: 1, Z: 1}
	vector2 := Vector3{ X: 1, Y: 1, Z: 1}
	vector3 := vector1.Add(&vector2)
	assert.Equal(t, 2.0, vector3.X)
	assert.Equal(t, 2.0, vector3.Y)
	assert.Equal(t, 2.0, vector3.Z)
}

func TestVector3_Multiply(t *testing.T) {
	vector1 := Vector3{ X: 1, Y: 1, Z: 1}
	vector2 := Vector3{ X: 1, Y: 2, Z: 3}
	vector3 := vector1.Multiply(&vector2)
	assert.Equal(t, 1.0, vector3.X)
	assert.Equal(t, 2.0, vector3.Y)
	assert.Equal(t, 3.0, vector3.Z)
}

func TestVector3_Subtract(t *testing.T) {
	vector1 := Vector3{ X: 1, Y: 1, Z: 1}
	vector2 := Vector3{ X: 1, Y: 1, Z: 1}
	vector3 := vector1.Subtract(&vector2)
	assert.Equal(t, 0.0, vector3.X)
	assert.Equal(t, 0.0, vector3.Y)
	assert.Equal(t, 0.0, vector3.Z)
}

func TestVector3_DotProduct(t *testing.T) {
	vector1 := Vector3{ X: 1, Y: 0, Z: 1}
	vector2 := Vector3{ X: 1, Y: 2, Z: 3}
	value := vector1.DotProduct(&vector2)
	assert.Equal(t, 4.0, value)
}

func TestVector3_MultiplyFloat(t *testing.T) {
	vector1 := Vector3{ X: 1, Y: 2, Z: 3}
	vector2 := vector1.MultiplyFloat(2.0)
	assert.Equal(t, 2.0, vector2.X)
	assert.Equal(t, 4.0, vector2.Y)
	assert.Equal(t, 6.0, vector2.Z)
}

func TestVector3_Negative(t *testing.T) {
	vector1 := Vector3{ X: 1, Y: 2, Z: 3}
	vector2 := vector1.Negative()
	assert.Equal(t, -1.0, vector2.X)
	assert.Equal(t, -2.0, vector2.Y)
	assert.Equal(t, -3.0, vector2.Z)
}

func TestVector3_Norm(t *testing.T) {
	vector1 := Vector3{ X: 1, Y: 2, Z: 3}
	assert.Equal(t, 3.7416573867739413, vector1.Norm())
}

func TestVector3_Normalize(t *testing.T) {
	vector1 := Vector3{ X: 1, Y: 2, Z: 3}
	vector2 := vector1.Normalize()
	assert.Equal(t, 1.0, vector2.Norm())
}

func TestVector3_CrossProduct(t *testing.T) {
	vector1 := Vector3{ X: 1}
	vector2 := Vector3{ Y: 1}
	vector3 := &Vector3{ Z: 1}
	assert.Equal(t, vector3, vector1.CrossProduct(&vector2))
}