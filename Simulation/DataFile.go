package Simulation

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
)

func createFile(particleCount int, frameCount int, absPath string) {

	outFilePath = absPath + "/builds/simulationData.bin"
	newfile, err := os.Create(outFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	file = *newfile

	bytes := uint32ToByte(uint32(particleCount))
	bytes = append(bytes, uint32ToByte(uint32(frameCount))...)
	file.Write(bytes)
}

func writeParticle(particle Particle) {

	bytes := float32ToByte(float32(particle.position[0]))
	bytes = append(bytes, float32ToByte(float32(particle.position[1]))...)
	bytes = append(bytes, float32ToByte(float32(particle.position[2]))...)
	file.Write(bytes)
}

func uint32ToByte(i uint32) []byte {
	buffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(buffer, i)
	return buffer
}

func float32ToByte(f float32) []byte {
	buffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(buffer, math.Float32bits(f))
	return buffer
}
