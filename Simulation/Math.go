package Simulation

import "github.com/go-gl/mathgl/mgl32"

const (
	h float32 = 10
)

func KernalFunction(pos1 mgl32.Vec3, pos2 mgl32.Vec3) float32 {
	H := pos2.Sub(pos1).Len() / h

	t1 := 1 - H
	if t1 < 0 {
		t1 = 0
	}

	t2 := 2 - H
	if t2 < 0 {
		t2 = 0
	}

	a := 1 / (6 * h)

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
	H := pos2.Sub(pos1).Len() / h

	t1 := 1 - H
	if t1 < 0 {
		t1 = 0
	}

	t2 := 2 - H
	if t2 < 0 {
		t2 = 0
	}

	a := 1 / (6 * h)

	w := mgl32.Vec3{}
	if 0 <= H && H < 1 {
		w = pos1.Sub(pos2).Mul(1 / (H * h)).Mul(a * (-3*(t2*t2) + 12*t1*t1*t1))
	}
	if 1 <= H && H < 2 {
		w = pos1.Sub(pos2).Mul(1 / (H * h)).Mul(a * (-3 * (t2 * t2)))
	}

	return w
}
