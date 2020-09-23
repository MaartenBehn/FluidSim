package Simulation

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

type particle struct {
	position mgl32.Vec3 // position = position + velocity
	velocity mgl32.Vec3 // velocity = velocity + forces / mass
	forces   mgl32.Vec3

	mass     float32
	density  float32 // pi = Element(j) * mj * Wij
	pressure float32

	collidingParticles []particle
}

func (currentParticle *particle) calculateDensity(otherParticle particle) {

	distance := otherParticle.position.Sub(currentParticle.position).Len()

	currentParticle.density += currentParticle.mass *
		float32((315.0/(64.0*math.Pi*math.Pow(smoothingDistance, 9)))*
			math.Pow(smoothingDistance-float64(distance), 3))
}

func (currentParticle *particle) calculatePressure() {
	currentParticle.pressure = gas * (currentParticle.density - restDensity)
}
