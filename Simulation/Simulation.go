package Simulation

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
	"math"
	"os"
	"strconv"
)

var outFilePath string
var file os.File

var particles []Particle
var frameCount int

const (
	dt                     = 1 / 2000.0 // s
	gravity                = 9.81 * m   // m/s^2
	m                      = 10.0
	cm                     = m / 100.0
	g                      = 1.0 // grams
	kg                     = 1000 * g
	newton                 = 1000 * g * m
	mu                     = 1.0e-3 * kg / m // Dynamic visosity, Ns/m^2
	pressureConstant       = 0.001 * m * m
	surfaceTensionConstant = 0.1

	spacing               = 0.5
	kernelRadius          = 1.5 * cm
	particleMass          = 1.0 * g
	particleRadius        = 0.3 * cm
	collisionDampingRatio = 0.8
)

var (
	rho           = 0.0 // Recalculated as needed
	particleCount = 0
)

func SetUpSimulation(_frameCount int, absPath string) {

	rho = 1.0 / math.Pow(spacing, 3) * g / (cm * cm * cm) // g/cm^3

	particles = make([]Particle, 0)

	createBlockofParticles(mgl64.Vec3{2, 2, 2}, mgl64.Vec3{}, mgl64.Vec3{})

	frameCount = _frameCount
	outFilePath = absPath + "/builds/simulationData.txt"
	newfile, err := os.Create(outFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	file = *newfile
	file.WriteString("info " + strconv.Itoa(particleCount) + " " + strconv.Itoa(frameCount+1) + "\n")
	file.WriteString("f " + strconv.Itoa(0) + "\n")

	for i, particle := range particles {
		file.WriteString("p " + strconv.FormatInt(int64(i), 10) + " " +
			strconv.FormatFloat(particle.position[0], 'f', -1, 64) + " " +
			strconv.FormatFloat(particle.position[1], 'f', -1, 64) + " " +
			strconv.FormatFloat(particle.position[2], 'f', -1, 64) + "\n")
	}

}

func createBlockofParticles(size mgl64.Vec3, position mgl64.Vec3, velocity mgl64.Vec3) {
	for x := 0.0; x <= size[0]; x += spacing {
		for y := 0.0; y <= size[1]; y += spacing {
			for z := 0.0; z <= size[2]; z += spacing {

				particle := Particle{
					position: mgl64.Vec3{(x + position[0] - (size[0] / 2)) * cm, (y + position[1] - (size[0] / 2)) * cm, (z + position[2] - (size[0] / 2)) * cm},
					velocity: mgl64.Vec3{velocity[0], velocity[1], velocity[2]},
					density:  rho,
				}

				particles = append(particles, particle)
				particleCount++
			}
		}
	}
}

func UpdateSimulation(frame int) {
	frame++
	file.WriteString("f " + strconv.Itoa(frame) + "\n")

	for i, particle := range particles {

		fmt.Printf("Calculating Particle %d of %d in Frame %d of %d \r", i, len(particles), frame, frameCount)

		particle.updateAcceleration()

		particles[i] = particle
	}

	for i, particle := range particles {

		particle.updateVelocity()
		particle.updatePosition()

		// Writing pos to file
		file.WriteString("p " + strconv.FormatInt(int64(i), 10) + " " +
			strconv.FormatFloat(float64(particle.position[0]), 'f', -1, 64) + " " +
			strconv.FormatFloat(float64(particle.position[1]), 'f', -1, 64) + " " +
			strconv.FormatFloat(float64(particle.position[2]), 'f', -1, 64) + "\n")

		particles[i] = particle
	}
}

func EndSimulation() {
	file.Close()
}
