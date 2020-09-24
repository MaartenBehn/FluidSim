package Simulation

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"math"
	"math/rand"
	"os"
	"strconv"
)

const (
	particleCount       = 4
	postionBounds       = 10
	startVelocityBounds = 0

	g                 = 1
	collisionDistance = 1

	doCollision                    = false
	collisionElasticEnergy float32 = 1.0

	gas               = 3.0
	restDensity       = 998.29
	smoothingDistance = 0.0457
)

var outFilePath string
var file os.File

var particles []particle
var frameCount int

func SetUpSimulation(_frameCount int, absPath string) {

	frameCount = _frameCount
	outFilePath = absPath + "/simulationData.txt"

	newfile, err := os.Create(outFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	file = *newfile
	file.WriteString("info " + strconv.Itoa(particleCount) + " " + strconv.Itoa(frameCount) + "\n")

	particles = make([]particle, particleCount)

	for i := 0; i < particleCount; i++ {
		particle := particle{
			position: mgl32.Vec3{
				(rand.Float32()*2 - 1) * postionBounds,
				(rand.Float32()*2 - 1) * postionBounds,
				(rand.Float32()*2 - 1) * postionBounds,
			},
			velocity: mgl32.Vec3{
				(rand.Float32()*2 - 1) * startVelocityBounds,
				(rand.Float32()*2 - 1) * startVelocityBounds,
				(rand.Float32()*2 - 1) * startVelocityBounds,
			},
			mass: 1,
		}

		particles[i] = particle
	}
}

func UpdateSimulation(frame int) {
	file.WriteString("f " + strconv.Itoa(frame) + "\n")

	newParticles := make([]particle, particleCount)
	copy(newParticles, particles)

	// First Loop: Setting up Particle Vars
	for i, newParticle := range newParticles {

		// Console Print
		fmt.Printf("Calculating Particle %d of %d in Frame %d of %d \r", i, len(particles), frame, frameCount)

		// Set Up Vars
		newParticle.density = 0
		newParticle.forces = mgl32.Vec3{}

		// For Neigbor
		for _, neigborParticle := range particles {
			if neigborParticle.position == newParticle.position {
				continue
			}

			// Public Vars
			relativePosition := neigborParticle.position.Sub(newParticle.position)
			distance := relativePosition.Len()

			// Particle density Var
			// pi = Element(j) * mj * Wij
			newParticle.density += Wpoly6(distance * distance)
		}

		newParticle.density *= newParticle.mass

		// Particle pressure Var
		newParticle.pressure = gas * (newParticle.density - restDensity)

		newParticles[i] = newParticle
	}

	// Second Loop: Forces -> Velocity
	for i, newParticle := range newParticles {

		// For Neigbor
		for _, neigborParticle := range newParticles {
			if neigborParticle.position == newParticle.position {
				continue
			}

			// Public Vars
			relativePosition := neigborParticle.position.Sub(newParticle.position)
			distance := relativePosition.Len()

			// Pressure Force

			newParticle.forces = newParticle.forces.Add(WspikyGradient(relativePosition, distance*distance).Mul(
				newParticle.pressure/float32(math.Pow(float64(newParticle.density), 2)) +
					neigborParticle.pressure/float32(math.Pow(float64(neigborParticle.density), 2))))

			// Gravity Force
			// force = g * mass1 * mass2 / |pos2 - pos1|^3 * pos2 - pos1
			if distance > collisionDistance*2 {
				newParticle.forces = newParticle.forces.Add(
					relativePosition.Mul(g * newParticle.mass * neigborParticle.mass /
						(distance * distance * distance)))
			}
		}

		// Velocity
		newParticle.velocity = newParticle.velocity.Add(
			newParticle.forces.Mul(1 / newParticle.mass))

		newParticles[i] = newParticle
	}

	// Third Loop: Velocity -> Position (with Collision)
	for i, newParticle := range newParticles {

		if doCollision {
			newParticle.lastFrameCollision = false

			normalVectorSum := mgl32.Vec3{}     // Sum of all relative Vectors for Elastic Collision
			velocitySum := newParticle.velocity // Sum of all Collider Velocities for Plastic Collision
			collidingParticleAmmount := 0       // Number of Particles how are Colliding

			for _, neigborParticle := range newParticles {
				if neigborParticle.position == newParticle.position {
					continue
				}

				// Public Vars
				relativePositionV := neigborParticle.position.Add(neigborParticle.velocity).Sub(
					newParticle.position.Add(neigborParticle.velocity))
				distanceV := relativePositionV.Len()

				// Collision
				if distanceV < collisionDistance*2 {

					newParticle.lastFrameCollision = true
					collidingParticleAmmount++

					normalVectorSum = normalVectorSum.Add(neigborParticle.position.Sub(newParticle.position).Normalize())

					velocitySum = velocitySum.Add(neigborParticle.velocity)
				}
			}

			// If there are colliding Particles calculate Collision
			if collidingParticleAmmount > 0 {

				elasticEnergyConversiom := collisionElasticEnergy

				// Ensures that Particles stick to getter if they did not bounce high enough
				if particles[i].lastFrameCollision && elasticEnergyConversiom < 1 {
					elasticEnergyConversiom = 0
				}

				// r = d - n * 2 * dot(d, n)
				// velocity = velocity - norm(pos1 - pos2) * 2 * dot(velocity, norm(pos1 - pos2))
				normalVectorSum = normalVectorSum.Normalize()
				velocityElastic := (newParticle.velocity.Sub(normalVectorSum.Mul(
					2 * newParticle.velocity.Dot(normalVectorSum)))).Mul(elasticEnergyConversiom)

				velocityPastic := velocitySum.Mul((1 / float32(collidingParticleAmmount+1)) * (1 - elasticEnergyConversiom))

				newParticle.velocity = velocityElastic.Add(velocityPastic)
			}
		}

		// Position
		newParticle.position = newParticle.position.Add(newParticle.velocity)

		// Writing pos to file
		file.WriteString("p " + strconv.FormatInt(int64(i), 10) + " " +
			strconv.FormatFloat(float64(newParticle.position[0]), 'f', -1, 64) + " " +
			strconv.FormatFloat(float64(newParticle.position[1]), 'f', -1, 64) + " " +
			strconv.FormatFloat(float64(newParticle.position[2]), 'f', -1, 64) + "\n")

		newParticles[i] = newParticle
	}

	particles = newParticles
}

func EndSimulation() {
	file.Close()
}
