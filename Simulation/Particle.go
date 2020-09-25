package Simulation

import (
	"github.com/go-gl/mathgl/mgl32"
)

type particle struct {
	position mgl32.Vec3
	velocity mgl32.Vec3

	mass                   float32
	neigborParticleAmmount float32
	neigborRadius          float32

	lastFrameCollision bool
}

func (currentParticle *particle) applyPressureVelocity() {

	neigborCount := float32(0)
	toOtherParticleVector := mgl32.Vec3{}
	for _, neigborParticle := range particles {
		if neigborParticle.position == currentParticle.position {
			continue
		}

		relativePosition := neigborParticle.position.Sub(currentParticle.position)
		distance := relativePosition.Len()

		toOtherParticleVector = toOtherParticleVector.Add(relativePosition.Normalize())

		if distance < currentParticle.neigborRadius*1.5 {
			neigborCount++
		}
	}
	toOtherParticleVector.Normalize()

	countForce := mgl32.Vec3{}
	if neigborCount < currentParticle.neigborParticleAmmount {
		countForce = toOtherParticleVector.Mul((1 - (neigborCount / currentParticle.neigborParticleAmmount)) * pressureCountScale)
	}
	if neigborCount > currentParticle.neigborParticleAmmount {
		countForce = toOtherParticleVector.Mul((1 - (currentParticle.neigborParticleAmmount / neigborCount)) * -pressureCountScale)
	}

	currentParticle.velocity = currentParticle.velocity.Add(countForce.Mul(1 / currentParticle.mass))
}

func (currentParticle *particle) applyGravityVelocity() {

	for _, neigborParticle := range particles {
		if neigborParticle.position == currentParticle.position {
			continue
		}

		relativePosition := neigborParticle.position.Sub(currentParticle.position)
		distance := relativePosition.Len()

		if distance > collisionDistance*2 {

			// Gravity Force
			// force = gravityConst * mass1 * mass2 / |pos2 - pos1|^3 * pos2 - pos1
			force := relativePosition.Mul(gravityScale * currentParticle.mass * (neigborParticle.mass /
				(distance * distance * distance)))

			currentParticle.velocity = currentParticle.velocity.Add(force.Mul(1 / currentParticle.mass))

		}
	}
}

func (currentParticle *particle) applyVelocityToPosition() {
	// position = position + velocity
	currentParticle.position = currentParticle.position.Add(currentParticle.velocity)
}

const (
	collisionDistance              = 1
	doCollision                    = false
	collisionElasticEnergy float32 = 0.5
)

func (currentParticle *particle) applyColisionToVelocity() {

	if doCollision {
		currentParticle.lastFrameCollision = false

		normalVectorSum := mgl32.Vec3{}         // Sum of all relative Vectors for Elastic Collision
		velocitySum := currentParticle.velocity // Sum of all Collider Velocities for Plastic Collision
		collidingParticleAmmount := 0           // Number of Particles how are Colliding

		for _, neigborParticle := range particles {
			if neigborParticle.position == currentParticle.position {
				continue
			}

			// Public Vars
			relativePositionV := neigborParticle.position.Add(neigborParticle.velocity).Sub(
				currentParticle.position.Add(neigborParticle.velocity))
			distanceV := relativePositionV.Len()

			// Collision
			if distanceV < collisionDistance*2 {

				currentParticle.lastFrameCollision = true
				collidingParticleAmmount++

				normalVectorSum = normalVectorSum.Add(neigborParticle.position.Sub(currentParticle.position).Normalize())

				velocitySum = velocitySum.Add(neigborParticle.velocity)
			}
		}

		// If there are colliding Particles calculate Collision
		if collidingParticleAmmount > 0 {

			elasticEnergyConversiom := collisionElasticEnergy

			// Ensures that Particles stick to getter if they did not bounce high enough
			if currentParticle.lastFrameCollision && elasticEnergyConversiom < 1 {
				elasticEnergyConversiom = 0
			}

			// r = d - n * 2 * dot(d, n)
			// velocity = velocity - norm(pos1 - pos2) * 2 * dot(velocity, norm(pos1 - pos2))
			normalVectorSum = normalVectorSum.Normalize()
			velocityElastic := (currentParticle.velocity.Sub(normalVectorSum.Mul(
				2 * currentParticle.velocity.Dot(normalVectorSum)))).Mul(elasticEnergyConversiom)

			velocityPastic := velocitySum.Mul((1 / float32(collidingParticleAmmount+1)) * (1 - elasticEnergyConversiom))

			currentParticle.velocity = velocityElastic.Add(velocityPastic)
		}
	}
}
