package Renderer

import (
	of "OctaForceEngineGo"
	"bufio"
	"encoding/binary"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"io/ioutil"
	"log"
	"math"
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

	var index = 0
	if len(lines) > 1 {
		fmt.Print("Found data files are: \n")
		for i, line := range lines {
			fmt.Printf("%d: %s \n", i, line)
		}
		fmt.Print("Please type in the desired data file number to play the file: \n")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", 1)
		index = of.ParseInt(input)
		if index < 0 || index > len(lines) {
			index = 0
		}
	}

	inFilePath = absPath + "/builds/" + lines[index] + ".bin"
}

var inFilePath string
var particleCount int
var FrameCount int
var particles []particle

func SetUpRenderer(absPath string) {
	mesh := of.NewMesh()
	mesh.LoadOBJ(absPath+"/mesh/Sphere.obj", false)

	content, err := ioutil.ReadFile(inFilePath)
	if err != nil {
		log.Fatal(err)
	}

	counter := 0
	particleCount = int(byteToInt32(content[counter*4 : (counter+1)*4]))
	counter++

	FrameCount = int(byteToInt32(content[counter*4 : (counter+1)*4]))
	counter++

	particles = make([]particle, particleCount)
	particles[0] = particle{
		postions:  make([]mgl32.Vec3, FrameCount),
		transform: mesh.Transform,
	}

	mesh.Material = of.Material{DiffuseColor: mgl32.Vec3{
		rand.Float32(),
		rand.Float32(),
		rand.Float32(),
	}}

	for i := range particles {
		if i == 0 {
			continue
		}

		material := of.Material{DiffuseColor: mgl32.Vec3{
			rand.Float32(),
			rand.Float32(),
			rand.Float32(),
		}}
		meshInstant := of.NewMeshInstant(mesh, &material)
		particles[i] = particle{
			postions:  make([]mgl32.Vec3, FrameCount),
			transform: meshInstant.Transform,
		}
	}
	of.GetActiveMeshes().AddMesh(mesh)

	var frame int
	var index int

	for frame < FrameCount {

		pos := mgl32.Vec3{}

		pos[0] = byteToFloat32(content[counter*4 : (counter+1)*4])
		counter++

		pos[1] = byteToFloat32(content[counter*4 : (counter+1)*4])
		counter++

		pos[2] = byteToFloat32(content[counter*4 : (counter+1)*4])
		counter++

		particles[index].postions[frame] = pos

		frame = counter / particleCount / 3
		index = (counter - frame*particleCount*3) / 3
	}
}

func byteToInt32(buffer []byte) uint32 {
	return binary.LittleEndian.Uint32(buffer)
}

func byteToFloat32(buffer []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(buffer))
}

func UpdateRenderer(frame int) {
	for _, particle := range particles {
		particle.transform.SetPosition(particle.postions[frame].Mul(100))
	}
}
