package Simulation

import (
	"github.com/go-gl/mathgl/mgl32"
)

type particle struct {
	position mgl32.Vec3 // position = position + velocity
	velocity mgl32.Vec3 // velocity = velocity + forces / mass
	forces   mgl32.Vec3

	mass     float32
	density  float32 // pi = Element(j) * mj * Wij
	pressure float32

	lastFrameCollision bool
}

func (currentParticle *particle) calculatePressure() {

}
