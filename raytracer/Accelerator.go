package raytracer

type Accelerator interface {
	IntersectGeometry(ray *Ray)

	AddGeometry(geometry Geometry)

	Build()
}