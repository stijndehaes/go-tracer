package geometry

import (
	"github.com/stijndehaes/go-tracer/math"
	"github.com/stijndehaes/go-tracer/raytracer"
	gomath "math"
)

type Triangle struct {
	Vertex0, Vertex1, Vertex2 math.Vector3
	Edge1, Edge2              math.Vector3
	TriangleMaterial          raytracer.Material
	normal                    math.Vector3
}

func NewTriangle(vertex0, vertex1, vertex2 math.Vector3, material raytracer.Material) *Triangle {
	// compute plane's normal
	v0v1 := vertex1.Subtract(vertex0)
	v0v2 := vertex2.Subtract(vertex0)
	// no need to normalize
	normal := v0v1.CrossProduct(v0v2)
	return &Triangle{
		Vertex0:          vertex0,
		Vertex1:          vertex1,
		Vertex2:          vertex2,
		Edge1:            v0v1,
		Edge2:            v0v2,
		TriangleMaterial: material,
		normal:           normal,
	}
}

func (triangle *Triangle) Material() raytracer.Material {
	return triangle.TriangleMaterial
}

func (triangle *Triangle) Hit(ray *raytracer.Ray) {
	dirCrossE2 := ray.Direction.CrossProduct(triangle.Edge2)
	determinant := triangle.Edge1.DotProduct(dirCrossE2)
	if gomath.Abs(determinant) < raytracer.MinDistance {
		return
	}

	// Triangle misses over P1-P3 edge
	f := 1.0 / determinant

	p1ToOrigin := ray.Eye.Subtract(triangle.Vertex0)
	u := f * p1ToOrigin.DotProduct(dirCrossE2)
	if u < 0 || u > 1 {
		return
	}

	originCrossE1 := p1ToOrigin.CrossProduct(triangle.Edge1)
	v := f * ray.Direction.DotProduct(originCrossE1)
	if v < 0 || (u+v) > 1 {
		return
	}
	tdist := f * triangle.Edge2.DotProduct(originCrossE1)
	if ray.IsHit(tdist, triangle) {
		ray.Normal = triangle.normal
	}
}
