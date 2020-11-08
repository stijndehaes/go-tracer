package raytracer

import (
	"github.com/stijndehaes/go-tracer/math"
	gomath "math"
)

type Camera struct {
	eye, up, direction          math.Vector3
	fov, distance               float64
	nx, ny                      int
	l, r, b, t, tanfov2, nxopny float64
	w, u, v                     math.Vector3
}

func (camera *Camera) getRayDirection(x, y int) math.Vector3 {
	uVal := camera.l + (camera.r-camera.l)*(float64(x)+0.5)/float64(camera.nx)
	vVal := camera.b + (camera.t-camera.b)*(float64(y)+0.5)/float64(camera.ny)
	utimes := camera.u.MultiplyFloat(uVal)
	vtimes := camera.v.MultiplyFloat(vVal)
	dtimes := camera.w.MultiplyFloat(-1 * camera.distance)
	uvtimes := utimes.Add(vtimes)
	uvdtimes := uvtimes.Add(dtimes)
	return uvdtimes.Normalize()
}

func (camera *Camera) GenerateRay(x, y int) Ray {
	rayDir := camera.getRayDirection(x, y)
	return NewRay(camera.eye, rayDir)
}

func GetCamera(eye, up, direction math.Vector3, fov, distance float64, nx, ny int) Camera {
	normalUp := up.Normalize()
	normalDir := direction.Normalize()
	w := normalDir.Negative()
	u := normalDir.CrossProduct(normalUp)
	v := w.CrossProduct(u)
	tanfov2 := gomath.Tan(fov / 360 * gomath.Pi)
	t := gomath.Abs(distance) * tanfov2
	b := -t
	nxopny := float64(nx) / float64(ny)
	r := t * nxopny
	l := -r
	return Camera{
		eye:       eye,
		up:        normalUp,
		direction: direction,
		fov:       fov,
		distance:  distance,
		nx:        nx,
		ny:        ny,
		w:         w,
		u:         u,
		v:         v,
		t:         t,
		b:         b,
		tanfov2:   tanfov2,
		nxopny:    nxopny,
		r:         r,
		l:         l,
	}
}
