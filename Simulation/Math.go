package Simulation

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

func KernalFunction(pos1 mgl32.Vec3, pos2 mgl32.Vec3) float32 {
	H := pos2.Sub(pos1).Len() / smoothingRadius

	t1 := 1 - H
	if t1 < 0 {
		t1 = 0
	}

	t2 := 2 - H
	if t2 < 0 {
		t2 = 0
	}

	a := 1 / (6 * smoothingRadius)

	w := float32(0)
	if 0 <= H && H < 1 {
		w = a * (t2*t2*t2 - 4*t1*t1*t1)
	}
	if 1 <= H && H < 2 {
		w = a * (t2 * t2 * t2)
	}

	return w
}

func KernalFunction2(pos1 mgl32.Vec3, pos2 mgl32.Vec3) mgl32.Vec3 {
	H := pos2.Sub(pos1).Len() / smoothingRadius

	t1 := 1 - H
	if t1 < 0 {
		t1 = 0
	}

	t2 := 2 - H
	if t2 < 0 {
		t2 = 0
	}

	a := 1 / (6 * smoothingRadius)

	w := mgl32.Vec3{}
	if 0 <= H && H < 1 {
		w = pos1.Sub(pos2).Mul(1 / (H * smoothingRadius)).Mul(a * (-3*(t2*t2) + 12*t1*t1*t1))
	}
	if 1 <= H && H < 2 {
		w = pos1.Sub(pos2).Mul(1 / (H * smoothingRadius)).Mul(a * (-3 * (t2 * t2)))
	}

	return w
}

func CohesionSplineFunction(r float32) float32 {

	t1 := float64((smoothingRadius - r) * (smoothingRadius - r) * (smoothingRadius - r) * r * r * r)
	t2 := math.Pow(float64(smoothingRadius), 6) / 64
	a := 32 / (math.Pi * math.Pow(float64(smoothingRadius), 9))

	w := float64(0)

	if (2*r) > smoothingRadius && r <= smoothingRadius {
		w = a * t1
	} else if r > 0 && (2*r) <= smoothingRadius {
		w = a * (2*t1 - t2)
	}

	return float32(w)
}

func greatestCommonDivisor(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func isVec3NAN(vec3 mgl32.Vec3) bool {
	return math.IsNaN(float64(vec3[0])) || math.IsNaN(float64(vec3[1])) || math.IsNaN(float64(vec3[2]))
}
