package raytracer

import (
	"github.com/Steniaz/go-tracer/math"
	gomath "math"
)


const minDistance = 0.00001

type Ray struct {
	Eye *math.Vector3
	Direction *math.Vector3
	HitDistance float64
	Hit bool
}

func (ray *Ray) IsHit(f float64) bool {
	if ray.HitDistance > f && f > minDistance{
		ray.HitDistance = f
		ray.Hit = true
		return true
	}
	return false
}

func NewRay(eye *math.Vector3, direction *math.Vector3) Ray {
	return Ray{
		Eye: eye,
		Direction: direction,
		HitDistance: gomath.MaxFloat64,
	}
}