package main

import (
	"OctaForceEngineGo"
	"github.com/go-gl/mathgl/mgl32"
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
	mesh = OctaForce.LoadOBJ(absPath + "/mesh/Cube.obj")
	OctaForce.SetComponent(e1, OctaForce.COMPONENT_Mesh, mesh)

	transform := OctaForce.GetComponent(e1, OctaForce.COMPONENT_Transform).(OctaForce.Transform)
	transform.Position = mgl32.Vec3{-5, -5, -10}
	OctaForce.SetComponent(e1, OctaForce.COMPONENT_Transform, transform)

	e1 = OctaForce.CreateEntity()
	mesh = OctaForce.AddComponent(e1, OctaForce.COMPONENT_Mesh).(OctaForce.Mesh)
	mesh = OctaForce.LoadOBJ(absPath + "/mesh/Sphere.obj")
	OctaForce.SetComponent(e1, OctaForce.COMPONENT_Mesh, mesh)

	transform = OctaForce.GetComponent(e1, OctaForce.COMPONENT_Transform).(OctaForce.Transform)
	transform.Position = mgl32.Vec3{5, 5, -10}
	OctaForce.SetComponent(e1, OctaForce.COMPONENT_Transform, transform)
}

func update() {
	log.Print(OctaForce.GetFPS())
}

func stop() {

}
