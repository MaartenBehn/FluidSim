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
	particleCount       = 5
	postionBounds       = 5
	startVelocityBounds = 0

	gravityScale       float32 = 0
	pressureCountScale float32 = 0.1
	pressureScale      float32 = 0.1
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
	file.WriteString("info " + strconv.Itoa(particleCount) + " " + strconv.Itoa(frameCount+1) + "\n")
	file.WriteString("f " + strconv.Itoa(0) + "\n")

	particles = make([]particle, particleCount)
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
			mass:                   1,
			neigborRadius:          10,
			neigborParticleAmmount: 1,
		}

		file.WriteString("p " + strconv.FormatInt(int64(i), 10) + " " +
			strconv.FormatFloat(float64(currentParticle.position[0]), 'f', -1, 64) + " " +
			strconv.FormatFloat(float64(currentParticle.position[1]), 'f', -1, 64) + " " +
			strconv.FormatFloat(float64(currentParticle.position[2]), 'f', -1, 64) + "\n")

		particles[i] = currentParticle
	}

}

func UpdateSimulation(frame int) {
	frame++

	file.WriteString("f " + strconv.Itoa(frame) + "\n")

	for i, currentParticle := range particles {

		fmt.Printf("Calculating Particle %d of %d in Frame %d of %d \r", i, len(particles), frame, frameCount)

		particles[i] = currentParticle
	}

	for i, currentParticle := range particles {

		currentParticle.applyPressureVelocity()
		currentParticle.applyGravityVelocity()
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
