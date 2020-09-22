package Simulation

import (
	"github.com/go-gl/mathgl/mgl32"
)

type particle struct {
	position mgl32.Vec3
	mass     float32
	velocity mgl32.Vec3
}

func updateParticle(particle particle) particle {
	for _, neigborParticle := range particles {
		if neigborParticle.position == particle.position {
			continue
		}

		relativePostion := particle.position.Sub(neigborParticle.position)
		distance := relativePostion.Len()
		if distance > collisionDistance*2 {
			particle.velocity = particle.velocity.Sub(
				normlizeVectorWithoutLength(relativePostion, distance).Mul(g / (distance * distance)))
		}
	}
	particle.position = particle.position.Add(particle.velocity)

	return particle
}
