package main

import (
	"FluidSim/Simulation"
	"path/filepath"
	"runtime"
)

func main() {

	frames := 1000
	_, b, _, _ := runtime.Caller(0)
	absPath := filepath.Dir(b)

	Simulation.SetUpSimulation(frames, absPath)

	for i := 0; i < frames; i++ {
		Simulation.UpdateSimulation(i)
	}
	Simulation.EndSimulation()
}
