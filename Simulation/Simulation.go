package Simulation

import (
	of "OctaForceEngineGo"
	"github.com/go-gl/mathgl/mgl32"
	"math/rand"
)

var particles []particle
var particleCount int
var postionBounds float32
var startVelocityBounds float32
var g float32
var collisionDistance float32

func SetUpSimulation() {

	particleCount = 1000
	postionBounds = 100
	startVelocityBounds = 0
	g = 1
	collisionDistance = 0

	particles = make([]particle, particleCount)
	mesh := of.LoadOBJ("mesh/Sphere.obj", false)

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
			mass:     1,
			entityId: of.CreateEntity(),
		}

		of.AddComponent(particle.entityId, of.ComponentMesh)
		mesh.Material = of.Material{DiffuseColor: mgl32.Vec3{
			rand.Float32(),
			rand.Float32(),
			rand.Float32(),
		}}
		of.SetComponent(particle.entityId, of.ComponentMesh, mesh)

		particle.setTransform()
		particles[i] = particle
	}
}

func UpdateSimulation() {
	newParticles := make([]particle, particleCount)
	for i, particle := range particles {
		newParticles[i] = updateParticle(particle)
		newParticles[i].setTransform()
	}
	particles = newParticles
}
