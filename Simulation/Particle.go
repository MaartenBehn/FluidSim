package Simulation

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Particle struct {
	position     mgl64.Vec3
	velocity     mgl64.Vec3
	acceleration mgl64.Vec3

	density, nextDensity float64
}

func updateAcceleration(particle *Particle) {
	// Get the neighbors and distances to them
	neighbors, distances := particle.findNeigbors(kernelRadius)

	// Initialize all accumulated quantities to zero
	density := 0.0
	colorLaplace := 0.0
	colorGradient := mgl64.Vec3{}
	force := mgl64.Vec3{}

	// Compute pressure due to the Particle itself. This is used for computing the
	// actual pressure by averaging this with the pressures due to other particles.
	selfPressure := pressureConstant * (particle.density - rho)

	// For all particles within the kernel radius
	for i, neighbor := range neighbors {
		// Density. Note that unlike other properties, the Particle contributes
		// to its own density.  This avoids having particles with a zero density.
		density += particleMass * smoothingKernel(distances[i])

		if neighbor != *particle {
			// Pressure
			neighborPressure := pressureConstant * (neighbor.density - rho)
			pressureScale := particleMass * (selfPressure + neighborPressure) / (2 * neighbor.density) * derivPressureKernel(distances[i])

			neighborDirection := neighbor.position.Sub(particle.position)
			force = force.Add(neighborDirection.Mul(pressureScale / distances[i]))

			// Surface tension (color!)
			colorScaling := particleMass / neighbor.density * derivSmoothingKernel(distances[i])
			colorLaplace += particleMass / neighbor.density * laplaceSmoothingKernel(distances[i])
			colorGradient = colorGradient.Add(neighborDirection.Mul(colorScaling))

			// Viscosity
			viscosityScale := mu * particleMass / neighbor.density * laplaceViscosityKernel(distances[i])
			viscosityDirection := neighbor.velocity.Sub(particle.velocity)
			force = force.Add(viscosityDirection.Mul(viscosityScale))
		}
	}

	// Only count color when it exists Otherwise, this causes the normalization
	// to fail and then particles become NaN.
	if colorGradient.X() > 0 || colorGradient.Y() > 0 || colorGradient.Z() > 0 {
		force = force.Add(colorGradient.Normalize().Mul(colorLaplace * surfaceTensionConstant))
	}

	// Gravity, pointing down
	//force.Z += -gravity

	// Set acceleration based on force, via F = ma
	particle.acceleration = force.Mul(1.0 / particleMass)

	// Set the to-be density
	particle.nextDensity = density
}

func (particle *Particle) findNeigbors(radius float64) ([]Particle, []float64) {

	distances := make([]float64, 0)
	neigbors := make([]Particle, 0)
	for _, otherParticle := range particles {
		distance := otherParticle.position.Sub(particle.position).Len()

		if distance < radius {
			distances = append(distances, distance)
			neigbors = append(neigbors, otherParticle)
		}
	}

	return neigbors, distances
}

// Update Particle state and display properties.
func updateVelocity(particle *Particle) {
	// Update velocity
	particle.velocity = particle.velocity.Add(particle.acceleration.Mul(dt))
}

func updateCollisions(particle *Particle) {

	// Get the neighbors and distances to them
	neighbors, distances := particle.findNeigbors(kernelRadius)

	collisionVelocities := mgl64.Vec3{}
	collisionCounter := 0.0
	// For all particles within the kernel radius
	for i, neighbor := range neighbors {

		if neighbor != *particle && distances[i] < particleRadius*2 {

			collisionCounter++

			// r = v - 2<v, n> n
			// velocity = velocity - 2 * dot(velocity, norm(pos1 - pos2)) * norm(pos1 - pos2)
			normal := particle.position.Sub(neighbor.position).Normalize()
			collisionVelocities = collisionVelocities.Add(
				(particle.velocity.Sub(normal.Mul(
					2 * particle.velocity.Dot(normal)))).Mul(collisionDampingRatio))

		}
	}

	if collisionCounter > 0 {
		particle.velocity = collisionVelocities.Mul(1 / collisionCounter)
	}
}

func updatePosition(particle *Particle) {
	// Update velocity, position
	particle.position = particle.position.Add(particle.velocity.Mul(dt))

	// Update density
	particle.density = particle.nextDensity
}
