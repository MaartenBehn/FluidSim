package Simulation

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"math/rand"
	"os"
	"strconv"
)

var outFilePath string
var file os.File

var particles []particle
var particleCount int
var postionBounds float32
var startVelocityBounds float32
var g float32
var collisionDistance float32
var collisionFriction float32
var frameCount int

func SetUpSimulation(_frameCount int, absPath string) {

	particleCount = 100
	postionBounds = 100
	startVelocityBounds = 0.1
	g = 1
	collisionDistance = 1
	collisionFriction = 0.1
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
	for i, particle := range particles {
		fmt.Printf("Calculating Particle %d of %d in Frame %d of %d \r", i, len(particles), frame, frameCount)

		newParticles[i] = updateParticle(particle)
	}

	for i, _ := range newParticles {
		newParticles[i] = updateColision(newParticles[i])
		newParticles[i] = updatePosition(newParticles[i])

		file.WriteString("p " + strconv.FormatInt(int64(i), 10) + " " +
			strconv.FormatFloat(float64(newParticles[i].position[0]), 'f', -1, 64) + " " +
			strconv.FormatFloat(float64(newParticles[i].position[1]), 'f', -1, 64) + " " +
			strconv.FormatFloat(float64(newParticles[i].position[2]), 'f', -1, 64) + "\n")
	}

	particles = newParticles
}

func EndSimulation() {
	file.Close()
}
