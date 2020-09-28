package Renderer

import (
	of "OctaForceEngineGo"
	"bufio"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

func SelectDataFile(absPath string) {

	content, err := ioutil.ReadFile(absPath + "/builds/index.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\r\n")

	fmt.Print("Found data files are: \n")
	for i, line := range lines {
		fmt.Printf("%d: %s \n", i, line)
	}
	fmt.Print("Please type in the desired data file number to play the file: \n")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", 1)
	index := of.ParseInt(input)
	if index < 0 || index > len(lines) {
		index = 0
	}

	inFilePath = absPath + "/builds/" + lines[index] + ".txt"
}

var inFilePath string
var particleCount int
var FrameCount int
var particles []particle

func SetUpRenderer(absPath string) {
	mesh := of.LoadOBJ(absPath+"/mesh/Sphere.obj", false)

	content, err := ioutil.ReadFile(inFilePath)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")

	var currentFrame int

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

func UpdateRenderer(frame int) {
	for _, particle := range particles {
		transform := of.GetComponent(particle.entityId, of.ComponentTransform).(of.Transform)
		transform.SetPosition(particle.postions[frame].Mul(100))
		of.SetComponent(particle.entityId, of.ComponentTransform, transform)
	}
}
