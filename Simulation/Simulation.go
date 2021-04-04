package main

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl64"
	"os"
	"sync"
)

var outFilePath string
var file os.File
var frameCount int

var particles []*Particle
var drops []*Drop

const (
	m  = 1.0
	cm = m / 100.0
	g  = 1.0 // grams
	kg = 1000 * g
	s  = 1.0

	dt = 1.0 / 25 * s

	particleRadius = 1.0 * cm
	spacing        = particleRadius * 4.0
	radiusFactor   = 0.01
	mergeFactor    = 0.7

	springStiffness     = 0.01
	springDampingFactor = 0.001
)

var (
	particleCount = 0
)

func main() {
	run()
}
func run() {
	frameCount = 500

	SetUpSimulation()

	for i := 0; i < frameCount; i++ {
		UpdateSimulation(i)
	}
	EndSimulation()
}

func SetUpSimulation() {

	fmt.Println("Starting simulation...")

	createBlockofParticles(mgl64.Vec3{30 * cm, 30 * cm, 30 * cm}, mgl64.Vec3{}, mgl64.Vec3{0, 0, 0})

	createBlockofParticles(mgl64.Vec3{10 * cm, 10 * cm, 10 * cm}, mgl64.Vec3{200 * cm, 0, 0}, mgl64.Vec3{-50 * cm, 0, 0})

	createFile(particleCount, frameCount)

	for _, particle := range particles {

		writeParticle(particle)
	}

	fmt.Printf("Simulation contains %d particles.\n", particleCount)

}

func createBlockofParticles(size mgl64.Vec3, position mgl64.Vec3, velocity mgl64.Vec3) {
	drop := &Drop{}

	for x := 0.0; x < size[0]; x += spacing * 0.8 {
		for y := 0.0; y < size[1]; y += spacing * 0.8 {
			for z := 0.0; z < size[2]; z += spacing * 0.8 {

				fmt.Printf("Creating particle %d. \r", particleCount)

				particle := &Particle{
					position: mgl64.Vec3{x + position[0] - (size[0] / 2), y + position[1] - (size[0] / 2), z + position[2] - (size[0] / 2)},
					velocity: mgl64.Vec3{velocity[0], velocity[1], velocity[2]},
					mass:     1,
					drop:     drop,
				}
				drop.particles = append(drop.particles, particle)

				particles = append(particles, particle)
				particleCount++
			}
		}
	}
	drops = append(drops, drop)
}

func UpdateSimulation(frame int) {
	frame++

	fmt.Printf("Calculating Frame %d of %d. \r", frame, frameCount)

	for _, drop := range drops {
		drop.calcPosition()
	}
	for i := len(drops) - 1; i >= 0; i-- {
		drops[i].checkDrop()
	}
	for _, drop := range drops {
		drop.calcDropRadius()
	}

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
