package raytracer

type Geometry interface {
	Hit(ray *Ray)
	Material() Material
}
