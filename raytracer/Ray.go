package raytracer

import (
	"github.com/Steniaz/go-tracer/math"
	gomath "math"
)



type Ray struct {
	Eye *math.Vector3
	Direction *math.Vector3
	HitDistance float64
	Hit bool
	Geometry Geometry
	Normal *math.Vector3
}

func (ray *Ray) IsHit(f float64, geometry Geometry) bool {
	if MinDistance <= f && f <= ray.HitDistance{
		ray.HitDistance = f
		ray.Hit = true
		ray.Geometry = geometry
		return true
	}
	return false
}

func (ray *Ray) HitPoint() *math.Vector3 {
	return ray.Direction.MultiplyFloat(ray.HitDistance).Add(ray.Eye)
}

func NewRay(eye *math.Vector3, direction *math.Vector3) Ray {
	return Ray{
		Eye: eye,
		Direction: direction,
		HitDistance: gomath.MaxFloat64,
	}
}