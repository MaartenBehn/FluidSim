package main

import (
	"OctaForceEngineGo"
	"log"
	"path/filepath"
	"runtime"
)

var absPath string

func init() {
	_, b, _, _ := runtime.Caller(0)
	absPath = filepath.Dir(b)
}

func main() {
	OctaForce.StartUp(start, stop)
}

func start() {
	OctaForce.SetMaxFPS(60)
	OctaForce.SetMaxUPS(20)
	OctaForce.AddUpdateCallback(update)

	e1 := OctaForce.CreateEntity()
	mesh := OctaForce.AddComponent(e1, OctaForce.COMPONENT_Mesh).(OctaForce.Mesh)
	mesh = OctaForce.LoadOBJ(absPath + "/mesh/cube.obj")
	OctaForce.SetComponent(e1, OctaForce.COMPONENT_Mesh, mesh)
}

func update() {
	log.Print(OctaForce.GetFPS())
}

func stop() {

}
