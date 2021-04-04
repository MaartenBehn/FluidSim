package Simulation

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
	"os"
	"sync"
)

var outFilePath string
var file os.File

var particles []*Particle
var frameCount int

const (
	m  = 1.0
	cm = m / 100.0
	g  = 1.0 // grams
	kg = 1000 * g
	s  = 1.0

	dt = 1.0 / 25 * s

	particleRadius = 1.0 * cm
	spacing        = particleRadius * 4.0

	springNeigborMaxDist = spacing * 2
	springNeigborMinDist = spacing * 1.1
	springRestingLenght  = spacing
	springStiffness      = 0.1
	springDampingFactor  = 0.001
)

var (
	particleCount = 0
)

func SetUpSimulation(_frameCount int, absPath string) {

	fmt.Println("Starting simulation...")

	frameCount = _frameCount

	particles = make([]*Particle, 0)

	createBlockofParticles(mgl64.Vec3{30 * cm, 30 * cm, 30 * cm}, mgl64.Vec3{}, mgl64.Vec3{0, 0, 0})

	createBlockofParticles(mgl64.Vec3{10 * cm, 10 * cm, 10 * cm}, mgl64.Vec3{50 * cm, 0, 0}, mgl64.Vec3{-10 * cm, 0, 0})

	createFile(particleCount, frameCount, absPath)

	for _, particle := range particles {
		particle.calcSpringNeigbors()
	}

	for _, particle := range particles {

		writeParticle(particle)
	}

	fmt.Printf("Simulation contains %d particles.\n", particleCount)

}

func createBlockofParticles(size mgl64.Vec3, position mgl64.Vec3, velocity mgl64.Vec3) {
	for x := 0.0; x < size[0]; x += spacing * 0.8 {
		for y := 0.0; y < size[1]; y += spacing * 0.8 {
			for z := 0.0; z < size[2]; z += spacing * 0.8 {

				fmt.Printf("Creating particle %d. \r", particleCount)

				particle := Particle{
					position: mgl64.Vec3{x + position[0] - (size[0] / 2), y + position[1] - (size[0] / 2), z + position[2] - (size[0] / 2)},
					velocity: mgl64.Vec3{velocity[0], velocity[1], velocity[2]},
					mass:     1,
				}

				particles = append(particles, &particle)
				particleCount++
			}
		}
	}
}

func UpdateSimulation(frame int) {
	frame++

	fmt.Printf("Calculating Frame %d of %d. \r", frame, frameCount)

	runInParallel((*Particle).calcSpringNeigbors)
	runInParallel((*Particle).calcForce)
	runInParallel((*Particle).calcVelocity)
	runInParallel((*Particle).calcPosition)

	// Writing Particle Pos to file
	for _, particle := range particles {
		writeParticle(particle)
	}
}

func runInParallel(function func(particle *Particle)) {

	wg := sync.WaitGroup{}
	wg.Add(particleCount)

	for i, p := range particles {

		go func(index int, particle *Particle) {

			function(particle)
			particles[index] = particle
			wg.Done()

		}(i, p)
	}

	wg.Wait()
}

func EndSimulation() {
	file.Close()
}
