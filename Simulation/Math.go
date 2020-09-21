package Simulation

import (
	"github.com/go-gl/mathgl/mgl32"
)

func normlizeVectorWithoutLength(v1 mgl32.Vec3, len float32) mgl32.Vec3 {
	l := 1.0 / len
	return mgl32.Vec3{v1[0] * l, v1[1] * l, v1[2] * l}
}
