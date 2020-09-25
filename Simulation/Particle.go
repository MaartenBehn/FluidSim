package Simulation

import (
	"github.com/go-gl/mathgl/mgl32"
)

type particle struct {
	position mgl32.Vec3
	velocity mgl32.Vec3

	mass     float32
	density  float32
	pressure float32
	volume   float32

	lastFrameCollision bool
}

func (currentParticle *particle) calcDensityAndPressure() {

	currentParticle.density = 0

	for _, neigborParticle := range particles {
		if neigborParticle.position == currentParticle.position {
			continue
		}

		currentParticle.density += neigborParticle.mass *
			KernalFunction(currentParticle.position, neigborParticle.position)
	}

	currentParticle.pressure = pressureScale * ((currentParticle.density / overAllDensity) - 1)
	if currentParticle.pressure < 0 {
		currentParticle.pressure = 0
	}
}

func (currentParticle *particle) applyPressureVelocity() {
	for _, neigborParticle := range particles {
		if neigborParticle.position == currentParticle.position {
			continue
		}

		velocity := KernalFunction2(currentParticle.position, neigborParticle.position).Mul(-1 * neigborParticle.mass *
			((currentParticle.pressure / (currentParticle.density * currentParticle.density)) +
				(neigborParticle.pressure / (neigborParticle.density * neigborParticle.density))))

		if !isVec3NAN(velocity) {
			currentParticle.velocity = currentParticle.velocity.Add(velocity)
		}
	}
}

func (currentParticle *particle) applyViscosityVelocity() {

	velocity := mgl32.Vec3{}
	for _, neigborParticle := range particles {
		if neigborParticle.position == currentParticle.position {
			continue
		}

		velocity = velocity.Add(
			KernalFunction2(currentParticle.position, neigborParticle.position).Mul(
				(neigborParticle.mass / neigborParticle.density) *
					neigborParticle.velocity.Dot(neigborParticle.position) /
					(neigborParticle.position.Dot(neigborParticle.position) +
						(0.01 * kernelSmoothingRadius * kernelSmoothingRadius))))

	}
	velocity.Mul(
		2 * viscosityScale * currentParticle.velocity.Dot(currentParticle.position) /
			(currentParticle.position.Dot(currentParticle.position) +
				(0.01 * kernelSmoothingRadius * kernelSmoothingRadius)))

	if !isVec3NAN(velocity) {
		currentParticle.velocity = currentParticle.velocity.Add(velocity)
	}
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
