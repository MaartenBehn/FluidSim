package main

import (
	of "OctaForceEngineGo"
	"fmt"
	math "github.com/go-gl/mathgl/mgl32"
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

var camera int

func start() {
	of.SetMaxFPS(60)
	of.SetMaxUPS(20)
	of.AddUpdateCallback(update)

	camera = of.CreateEntity()
	of.AddComponent(camera, of.ComponentCamera)
	transform := of.GetComponent(camera, of.ComponentTransform).(of.Transform)
	transform.SetPosition(math.Vec3{0, 0, 10})
	of.SetComponent(camera, of.ComponentTransform, transform)
	of.SetActiveCameraEntity(camera)

	e1 := of.CreateEntity()
	mesh := of.AddComponent(e1, of.ComponentMesh).(of.Mesh)
	mesh = of.LoadOBJ("mesh/Cube.obj", false)
	mesh.Material = of.Material{DiffuseColor: math.Vec3{0.5, 0.2, 1}}
	of.SetComponent(e1, of.ComponentMesh, mesh)

	transform = of.GetComponent(e1, of.ComponentTransform).(of.Transform)
	transform.SetPosition(math.Vec3{-5, -5, -10})
	transform.SetRotaion(math.Vec3{0, 45, 45})
	of.SetComponent(e1, of.ComponentTransform, transform)

	e1 = of.CreateEntity()
	mesh = of.AddComponent(e1, of.ComponentMesh).(of.Mesh)
	mesh = of.LoadOBJ("mesh/Sphere.obj", false)
	mesh.Material = of.Material{DiffuseColor: math.Vec3{0.8, 0.2, 0.3}}
	of.SetComponent(e1, of.ComponentMesh, mesh)

	transform = of.GetComponent(e1, of.ComponentTransform).(of.Transform)
	transform.SetPosition(math.Vec3{5, 5, -10})
	of.SetComponent(e1, of.ComponentTransform, transform)
}

func update() {
	fmt.Println(of.GetFPS())

	transform := of.GetComponent(camera, of.ComponentTransform).(of.Transform)
	if of.KeyPressed(of.KeyW) {
		transform.MoveRelative(math.Vec3{0, 0, -1})
	}
	if of.KeyPressed(of.KeyS) {
		transform.MoveRelative(math.Vec3{0, 0, 1})
	}
	if of.KeyPressed(of.KeyA) {
		transform.MoveRelative(math.Vec3{-1, 0, 0})
	}
	if of.KeyPressed(of.KeyD) {
		transform.MoveRelative(math.Vec3{1, 0, 0})
	}
	if of.KeyPressed(of.KeyQ) {
		transform.Rotate(math.Vec3{1, 0, 0})
	}
	if of.KeyPressed(of.KeyE) {
		transform.Rotate(math.Vec3{-1, 0, 0})
	}
	of.SetComponent(camera, of.ComponentTransform, transform)
}

func stop() {

}
