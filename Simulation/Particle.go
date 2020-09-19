package Simulation

import "github.com/go-gl/mathgl/mgl32"

type particle struct {
	postion  mgl32.Vec3
	mass     float32
	velocity mgl32.Vec3
}
