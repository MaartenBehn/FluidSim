package main

import (
	of "OctaForceEngineGo"
	"github.com/go-gl/mathgl/mgl32"
)

type particle struct {
	postions  []mgl32.Vec3
	transform *of.Transform
}
