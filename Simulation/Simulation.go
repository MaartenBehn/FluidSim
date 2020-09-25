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
var frameCount int

const (
	particleCount       = 30
	postionBounds       = 20
	startVelocityBounds = 0

	volumePerPaeticle  float32 = 100
	densityPerParticle float32 = 10

	pressureScale         float32 = 1
	gravityScale          float32 = 0.000001
	kernelSmoothingRadius float32 = 10
	viscosityScale        float32 = 1
)

var overAllVolume float32
var overAllDensity float32

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
	overAllDensity = densityPerParticle * particleCount
	overAllVolume = volumePerPaeticle * particleCount

	for i := 0; i < particleCount; i++ {
		currentParticle := particle{
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
		}

		currentParticle.volume = overAllVolume / particleCount
		currentParticle.mass = overAllDensity * currentParticle.volume

		particles[i] = currentParticle
	}
}

func UpdateSimulation(frame int) {
	file.WriteString("f " + strconv.Itoa(frame) + "\n")

	for i, currentParticle := range particles {
		// Console Print
		fmt.Printf("Calculating Particle %d of %d in Frame %d of %d \r", i, len(particles), frame, frameCount)

		currentParticle.calcDensityAndPressure()

		particles[i] = currentParticle
	}

	for i, currentParticle := range particles {

		currentParticle.applyPressureVelocity()
		currentParticle.applyViscosityVelocity()
		currentParticle.applyGravityVelocity()

		particles[i] = currentParticle
	}

	for i, currentParticle := range particles {

		currentParticle.applyVelocityToPosition()

		// Writing pos to file
		file.WriteString("p " + strconv.FormatInt(int64(i), 10) + " " +
			strconv.FormatFloat(float64(currentParticle.position[0]), 'f', -1, 64) + " " +
			strconv.FormatFloat(float64(currentParticle.position[1]), 'f', -1, 64) + " " +
			strconv.FormatFloat(float64(currentParticle.position[2]), 'f', -1, 64) + "\n")

		particles[i] = currentParticle
	}
}

func EndSimulation() {
	file.Close()
}
