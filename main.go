package main

import (
	of "OctaForceEngineGo"
	math "github.com/go-gl/mathgl/mgl32"
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
	of.StartUp(start, stop)
}

func start() {
	of.SetMaxFPS(60)
	of.SetMaxUPS(20)
	of.AddUpdateCallback(update)

	e0 := of.CreateEntity()
	of.AddComponent(e0, of.ComponentCamera)
	transform := of.GetComponent(e0, of.ComponentTransform).(of.Transform)
	transform.Position = math.Vec3{5, 0, 10}
	transform.Rotation = math.Vec3{0, 0, 0}
	of.SetComponent(e0, of.ComponentTransform, transform)
	of.SetActiveCameraEntity(e0)

	e1 := of.CreateEntity()
	mesh := of.AddComponent(e1, of.ComponentMesh).(of.Mesh)
	mesh = of.LoadOBJ(absPath + "/mesh/Cube.obj")
	of.SetComponent(e1, of.ComponentMesh, mesh)

	transform = of.GetComponent(e1, of.ComponentTransform).(of.Transform)
	transform.Position = math.Vec3{-5, -5, -10}
	of.SetComponent(e1, of.ComponentTransform, transform)

	e1 = of.CreateEntity()
	mesh = of.AddComponent(e1, of.ComponentMesh).(of.Mesh)
	mesh = of.LoadOBJ(absPath + "/mesh/Sphere.obj")
	of.SetComponent(e1, of.ComponentMesh, mesh)

	transform = of.GetComponent(e1, of.ComponentTransform).(of.Transform)
	transform.Position = math.Vec3{5, 5, -10}
	of.SetComponent(e1, of.ComponentTransform, transform)
}

func update() {
	log.Print(of.GetFPS())
}

func stop() {

}
