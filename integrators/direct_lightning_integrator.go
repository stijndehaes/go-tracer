package integrators

import (
	"github.com/stijndehaes/go-tracer/math"
	"github.com/stijndehaes/go-tracer/raytracer"
)

type DirectLightningIntegrator struct {
	Scene *raytracer.Scene
}

func (integrator *DirectLightningIntegrator) RenderRay(ray *raytracer.Ray) math.Color3 {
	integrator.Scene.Accelerator.IntersectGeometry(ray)

	color := math.Color3{}
	if !ray.Hit {
		return color
	}

	for _, light := range integrator.Scene.Lights {
		color = color.Add(integrator.shadeGeometryWithLight(light, ray))
	}
	return color
}

func (integrator *DirectLightningIntegrator) shadeGeometryWithLight(light raytracer.Light, ray *raytracer.Ray) math.Color3 {
	hit := ray.HitPoint()
	direction, distance, color := light.SampleLight(hit)
	lightRay := &raytracer.Ray{Eye: hit, Direction: direction, HitDistance: distance}
	integrator.Scene.Accelerator.IntersectGeometry(lightRay)
	dirMulNormal := direction.DotProduct(ray.Normal)
	if !lightRay.Hit && dirMulNormal > 0 {
		return ray.Geometry.Material().Shade(ray.Direction, direction, ray).Multiply(color).MultiplyFloat(dirMulNormal)
	} else {
		return math.Color3{}
	}
}
