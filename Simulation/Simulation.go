package Simulation

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
	"math"
	"os"
	"sync"
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
	particleRadius        = 0.1 * cm
	collisionDampingRatio = 1
)

var (
	rho           = 0.0 // Recalculated as needed
	particleCount = 0
)

func SetUpSimulation(_frameCount int, absPath string) {

	fmt.Println("Starting simulation...")

	rho = 1.0 / math.Pow(spacing, 3) * g / (cm * cm * cm) // g/cm^3
	frameCount = _frameCount

	particles = make([]Particle, 0)

	createBlockofParticles(mgl64.Vec3{8, 8, 8}, mgl64.Vec3{}, mgl64.Vec3{})

	createBlockofParticles(mgl64.Vec3{4, 4, 4}, mgl64.Vec3{-30, 4, 4}, mgl64.Vec3{10, 0, 0})

	createFile(particleCount, frameCount, absPath)

	for _, particle := range particles {
		writeParticle(particle)
	}

	fmt.Printf("Simulation contains %d particles.\n", particleCount)

}

func createBlockofParticles(size mgl64.Vec3, position mgl64.Vec3, velocity mgl64.Vec3) {
	for x := 0.0; x <= size[0]; x += spacing {
		for y := 0.0; y <= size[1]; y += spacing {
			for z := 0.0; z <= size[2]; z += spacing {

				fmt.Printf("Creating particle %d. \r", particleCount)

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

	fmt.Printf("Calculating Frame %d of %d. \r", frame, frameCount)

	runInParallel(updateAcceleration)
	runInParallel(updateVelocity)
	runInParallel(updatePosition)

	// Writing Particle Pos to file
	for _, particle := range particles {
		writeParticle(particle)
	}
}

func runInParallel(function func(particle *Particle)) {

	wg := sync.WaitGroup{}
	wg.Add(particleCount)

	for i, p := range particles {

		go func(index int, particle Particle) {

			function(&particle)
			particles[index] = particle
			wg.Done()

		}(i, p)
	}

	wg.Wait()
}

func EndSimulation() {
	file.Close()
}
