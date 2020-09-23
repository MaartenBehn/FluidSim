package Simulation

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"math/rand"
	"os"
	"strconv"
)

const (
	gas         = 2000
	restDensity = 998.29
)

var outFilePath string
var file os.File

var particles []particle
var particleCount int
var postionBounds float32
var startVelocityBounds float32
var g float32
var collisionDistance float32
var collisionElasticEnergy float32
var smoothingDistance float64

var frameCount int

func SetUpSimulation(_frameCount int, absPath string) {

	particleCount = 100
	postionBounds = 100
	startVelocityBounds = 0.5
	g = 1
	collisionDistance = 1
	collisionElasticEnergy = 1
	smoothingDistance = 1

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

	for i, newParticle := range newParticles {
		fmt.Printf("Calculating Particle %d of %d in Frame %d of %d \r", i, len(particles), frame, frameCount)

		// Set Up
		newParticle.density = 0
		newParticle.forces = mgl32.Vec3{}

		// Gravity
		for _, neigborParticle := range particles {
			if neigborParticle.position == newParticle.position {
				continue
			}

			relativePosition := neigborParticle.position.Sub(newParticle.position)
			distance := relativePosition.Len()

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

	for i, newParticle := range newParticles {

		// Colliding Neigbors
		newParticle.collidingParticles = make([]particle, 0)
		for _, neigborParticle := range particles {
			if neigborParticle.position == newParticle.position {
				continue
			}
			distance := neigborParticle.position.Add(neigborParticle.velocity).Sub(
				newParticle.position.Add(neigborParticle.velocity)).Len()
			if distance < collisionDistance*2 {
				newParticle.collidingParticles = append(newParticle.collidingParticles, neigborParticle)
			}
		}

		// Collision
		newParticle.lastFrameCollision = false
		collidingParticleAmmount := float32(len(newParticle.collidingParticles))
		if collidingParticleAmmount > 0 {

			newParticle.lastFrameCollision = true

			velocityElastic := newParticle.velocity
			velocityPastic := newParticle.velocity
			for _, collidingPraticle := range newParticle.collidingParticles {

				normVector := collidingPraticle.position.Sub(newParticle.position).Normalize()

				// r = d - n * 2 * dot(d, n)
				// velocity = velocity - norm(pos1 - pos2) * 2 * dot(velocity, norm(pos1 - pos2))
				velocityElastic = velocityElastic.Sub(normVector.Mul(
					2 * velocityElastic.Dot(normVector)))

				velocityPastic = velocityPastic.Add(collidingPraticle.velocity)

			}

			elasticEnergyConversiom := collisionElasticEnergy
			if particles[i].lastFrameCollision && elasticEnergyConversiom < 1 {
				elasticEnergyConversiom = 0
			}

			velocityElastic = velocityElastic.Mul(elasticEnergyConversiom)
			velocityPastic = velocityPastic.Mul((1 / (collidingParticleAmmount + 1)) * (1 - elasticEnergyConversiom))
			newParticle.velocity = velocityElastic.Add(velocityPastic)

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
