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
	postionBounds       = 4
	startVelocityBounds = 0

	volumePerPaeticle  float32 = 1
	densityPerParticle float32 = 1

	smoothingRadius float32 = 20

	cohesionScale   float32 = 1
	colorFieldScale float32 = 1
	curviatureScale float32 = 1

	pressureScale  float32 = 1
	viscosityScale float32 = 1

	gravityScale       float32 = 0.000001
	staticGravityScale float32 = 0 //0.001

)

var overAllDensity float32

func SetUpSimulation(_frameCount int, absPath string) {

	overAllDensity = densityPerParticle * particleCount

	frameCount = _frameCount
	outFilePath = absPath + "/simulationData.txt"

	newfile, err := os.Create(outFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	file = *newfile
	file.WriteString("info " + strconv.Itoa(particleCount) + " " + strconv.Itoa(frameCount+1) + "\n")
	particles = make([]particle, particleCount)

	file.WriteString("f " + strconv.Itoa(0) + "\n")
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

		currentParticle.volume = volumePerPaeticle
		currentParticle.mass = overAllDensity * currentParticle.volume

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
		// Console Print
		fmt.Printf("Calculating Particle %d of %d in Frame %d of %d \r", i, len(particles), frame, frameCount)

		currentParticle.calcDensityAndPressure()

		particles[i] = currentParticle
	}

	for i, currentParticle := range particles {

		currentParticle.calcColorFieldNormal()

		particles[i] = currentParticle
	}

	for i, currentParticle := range particles {

		currentParticle.applySurfaceTensionVelocity()
		//currentParticle.applyPressureVelocity()
		//currentParticle.applyViscosityVelocity()
		//currentParticle.applyStaticGravityVelocity()

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
