package geometry

import (
	"github.com/Steniaz/go-tracer/math"
	"github.com/Steniaz/go-tracer/raytracer"
	gomath "math"
)

type Sphere struct {
	Center math.Vector3
	Radius float64
	SphereMaterial raytracer.Material
}

func (sphere *Sphere) radius2() float64 {
	return sphere.Radius * sphere.Radius
}

func (sphere *Sphere) Material() raytracer.Material {
	return sphere.SphereMaterial
}


func (sphere *Sphere) Hit(ray *raytracer.Ray) {
	a := ray.Direction.DotProduct(ray.Direction)

	eMinC := ray.Eye.Subtract(&sphere.Center)
	b :=  ray.Direction.DotProduct(eMinC) * 2.0
	c := eMinC.DotProduct(eMinC) - sphere.radius2()
	discriminant := b * b - (4 * a * c)
	if discriminant < 0 {
		return
	}

	t1 := (-b + gomath.Sqrt(discriminant)) / (2 * a)
	t2 := (-b - gomath.Sqrt(discriminant)) / (2 * a)
	if t1 > t2 {
		temp := t1
		t1 = t2
		t2 = temp
	}

	if ray.IsHit(t1, sphere) || ray.IsHit(t2, sphere){
		ray.Normal = ray.HitPoint().Subtract(&sphere.Center).Normalize()
	}
}