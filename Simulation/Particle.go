package Simulation

import (
	"github.com/go-gl/mathgl/mgl32"
)

type particle struct {
	position           mgl32.Vec3
	mass               float32
	velocity           mgl32.Vec3
	collidingParticles []particle
}

func updateParticle(currentParticle particle) particle {

	currentParticle.collidingParticles = make([]particle, 0)

	for _, neigborParticle := range particles {
		if neigborParticle.position == currentParticle.position {
			continue
		}

		relativePostion := neigborParticle.position.Sub(currentParticle.position)
		distance := relativePostion.Len()

		if distance < collisionDistance*2 {
			currentParticle.collidingParticles = append(currentParticle.collidingParticles, neigborParticle)
		} else {
			currentParticle.velocity = currentParticle.velocity.Add(relativePostion.Mul(
				(g * currentParticle.mass * neigborParticle.mass) / (distance * distance * distance)))
		}
	}

	return currentParticle
}

// r = d - n * 2 * dot(d, n)
func updateColision(currentParticle particle) particle {
	if len(currentParticle.collidingParticles) > 0 {
		for _, collidingPraticle := range currentParticle.collidingParticles {

			normVector := currentParticle.position.Sub(collidingPraticle.position).Normalize()
			currentParticle.velocity = currentParticle.velocity.Sub(normVector.Mul(
				2 * currentParticle.velocity.Dot(normVector)))
		}
	}
	return currentParticle
}

func updatePosition(currentParticle particle) particle {
	currentParticle.position = currentParticle.position.Add(currentParticle.velocity)
	return currentParticle
}
