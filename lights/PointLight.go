package lights

import (
	"github.com/Steniaz/go-tracer/math"
)

type PointLight struct {
	Color *math.Color3
	Origin *math.Vector3
}

func (light PointLight) SampleLight(hit *math.Vector3) (*math.Vector3, float64, *math.Color3) {
	temp := light.Origin.Subtract(hit)
	direction := temp.Normalize()
	distance := temp.Norm()
	return direction, distance, light.Color.DivideFloat(distance * distance)
}