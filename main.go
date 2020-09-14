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
	OF.StartUp(start, stop)
}

func start() {
	OF.SetMaxFPS(60)
	OF.SetMaxUPS(20)
	OF.AddUpdateCallback(update)

	e0 := OF.CreateEntity()
	OF.AddComponent(e0, OF.COMPONENT_Camera)
	transform := OF.GetComponent(e0, OF.COMPONENT_Transform).(OF.Transform)
	transform.Position = mgl32.Vec3{5, 0, 10}
	transform.Rotation = mgl32.Vec3{0, 0, 0}
	OF.SetComponent(e0, OF.COMPONENT_Transform, transform)
	OF.SetActiveCameraEntity(e0)

	e1 := OF.CreateEntity()
	mesh := OF.AddComponent(e1, OF.COMPONENT_Mesh).(OF.Mesh)
	mesh = OF.LoadOBJ(absPath + "/mesh/Cube.obj")
	OF.SetComponent(e1, OF.COMPONENT_Mesh, mesh)

	transform = OF.GetComponent(e1, OF.COMPONENT_Transform).(OF.Transform)
	transform.Position = mgl32.Vec3{-5, -5, -10}
	OF.SetComponent(e1, OF.COMPONENT_Transform, transform)

	e1 = OF.CreateEntity()
	mesh = OF.AddComponent(e1, OF.COMPONENT_Mesh).(OF.Mesh)
	mesh = OF.LoadOBJ(absPath + "/mesh/Sphere.obj")
	OF.SetComponent(e1, OF.COMPONENT_Mesh, mesh)

	transform = OF.GetComponent(e1, OF.COMPONENT_Transform).(OF.Transform)
	transform.Position = mgl32.Vec3{5, 5, -10}
	OF.SetComponent(e1, OF.COMPONENT_Transform, transform)
}

func update() {
	log.Print(OF.GetFPS())
}

func stop() {

}
