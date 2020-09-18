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
	of.SetMaxUPS(30)
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
	transform.SetPosition(math.Vec3{0, 0, -10})
	transform.SetRotaion(math.Vec3{0, 45, 45})
	of.SetComponent(e1, of.ComponentTransform, transform)

	e1 = of.CreateEntity()
	mesh = of.AddComponent(e1, of.ComponentMesh).(of.Mesh)
	mesh = of.LoadOBJ("mesh/Sphere.obj", false)
	mesh.Material = of.Material{DiffuseColor: math.Vec3{0.8, 0.2, 0.3}}
	of.SetComponent(e1, of.ComponentMesh, mesh)

	transform = of.GetComponent(e1, of.ComponentTransform).(of.Transform)
	transform.SetPosition(math.Vec3{1, 0, -20})
	of.SetComponent(e1, of.ComponentTransform, transform)
}

const (
	movementSpeed float32 = 10
	mouseSpeed    float32 = 3
)

func update() {
	fmt.Println(of.GetFPS())

	deltaTime := float32(of.GetDeltaTime())

	transform := of.GetComponent(camera, of.ComponentTransform).(of.Transform)
	if of.KeyPressed(of.KeyW) {
		transform.MoveRelative(math.Vec3{0, 0, -1}.Mul(deltaTime * movementSpeed))
	}
	if of.KeyPressed(of.KeyS) {
		transform.MoveRelative(math.Vec3{0, 0, 1}.Mul(deltaTime * movementSpeed))
	}
	if of.KeyPressed(of.KeyA) {
		transform.MoveRelative(math.Vec3{-1, 0, 0}.Mul(deltaTime * movementSpeed))
	}
	if of.KeyPressed(of.KeyD) {
		transform.MoveRelative(math.Vec3{1, 0, 0}.Mul(deltaTime * movementSpeed))
	}
	if of.MouseButtonPressed(of.MouseButtonLeft) {
		mouseMovement := of.GetMouseMovement()
		transform.Rotate(math.Vec3{-1, 0, 0}.Mul(mouseMovement.Y() * deltaTime * mouseSpeed))
		transform.Rotate(math.Vec3{0, 1, 0}.Mul(mouseMovement.X() * deltaTime * mouseSpeed))
	}
	of.SetComponent(camera, of.ComponentTransform, transform)
}

func stop() {

}
