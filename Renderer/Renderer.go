package Renderer

import (
	of "OctaForceEngineGo"
	"github.com/go-gl/mathgl/mgl32"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
)

var inFilePath string
var particleCount uint32
var FrameCount uint32
var particles []particle

func SetUpRenderer(absPath string) {

	mesh := of.LoadOBJ(absPath+"/mesh/Sphere.obj", false)

	inFilePath = absPath + "/simulationData.txt"
	content, err := ioutil.ReadFile(inFilePath)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	var currentFrame uint32

	for _, line := range lines {
		values := strings.Split(line, " ")
		values[len(values)-1] = strings.Replace(values[len(values)-1], "\r", "", 1)

		switch values[0] {
		case "info":
			particleCount = of.ParseInt(values[1])
			FrameCount = of.ParseInt(values[2])

			particles = make([]particle, particleCount)
			for i := range particles {
				particles[i] = particle{
					postions: make([]mgl32.Vec3, FrameCount),
					entityId: of.CreateEntity(),
				}
				of.AddComponent(particles[i].entityId, of.ComponentMesh)
				mesh.Material = of.Material{DiffuseColor: mgl32.Vec3{
					rand.Float32(),
					rand.Float32(),
					rand.Float32(),
				}}
				of.SetComponent(particles[i].entityId, of.ComponentMesh, mesh)
			}
			break
		case "f":
			currentFrame = of.ParseInt(values[1])
			break
		case "p":
			currentParticle := of.ParseInt(values[1])
			particles[currentParticle].postions[currentFrame] = mgl32.Vec3{
				of.ParseFloat(values[2]),
				of.ParseFloat(values[3]),
				of.ParseFloat(values[4])}
			break
		}
	}
}

func UpdateRenderer(frame uint32) {
	for _, particle := range particles {
		transform := of.GetComponent(particle.entityId, of.ComponentTransform).(of.Transform)
		transform.SetPosition(particle.postions[frame])
		of.SetComponent(particle.entityId, of.ComponentTransform, transform)
	}
}
