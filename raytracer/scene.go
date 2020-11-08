package raytracer

type Scene struct {
	Camera      *Camera
	Accelerator Accelerator
	Lights      []Light
}
