package geometry

import (
	"github.com/Steniaz/go-tracer/math"
	"github.com/Steniaz/go-tracer/raytracer"
	gomath "math"
)

type Sphere struct {
	Center math.Vector3
	Radius float64
}

func (sphere *Sphere) radius2() float64 {
	return sphere.Radius * sphere.Radius
}

func (sphere *Sphere) Hit(ray *raytracer.Ray) bool {

	l := sphere.Center.Subtract(ray.Eye);
	tca := l.DotProduct(ray.Direction)
	d2 := l.DotProduct(&l) - tca * tca;
	if d2 > sphere.radius2() {
		return false
	}
	thc := gomath.Sqrt(sphere.radius2() - d2)
	t0 := tca - thc
	t1 := tca + thc
	if (t0 > t1) {
		temp := t0
		t0 = t1
		t1 = temp
	}
	if t0 < 0 {
		t0 = t1; // if t0 is negative, let's use t1 instead
		if t0 < 0 {
			return false
		} // both t0 and t1 are negative
	}
	return ray.IsHit(t0)
}