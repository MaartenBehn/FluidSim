package main

import (
	render "FluidSim/Renderer"
	of "OctaForceEngineGo"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"path/filepath"
	"runtime"
)

func main() {
	//defer profile.Start().Stop()
	of.StartUp(start, update, stop)
}

var camera int

func start() {
	of.SetMaxFPS(60)
	of.SetMaxUPS(30)

	camera = of.CreateEntity()
	of.AddComponent(camera, of.ComponentCamera)
	transform := of.GetComponent(camera, of.ComponentTransform).(of.Transform)
	transform.SetPosition(mgl32.Vec3{0, 0, 500})
	of.SetComponent(camera, of.ComponentTransform, transform)
	of.SetActiveCameraEntity(camera)

	_, b, _, _ := runtime.Caller(0)
	absPath := filepath.Dir(b)

	render.SetUpRenderer(absPath)
	render.UpdateRenderer(0)
}

const (
	movementSpeed float32 = 100
	mouseSpeed    float32 = 3
)

var currentFrame uint32

func update() {
	fmt.Printf("FPS: %f UPS: %f \r", of.GetFPS(), of.GetUPS())

	deltaTime := float32(of.GetDeltaTime())

	transform := of.GetComponent(camera, of.ComponentTransform).(of.Transform)
	if of.KeyPressed(of.KeyW) {
		transform.MoveRelative(mgl32.Vec3{0, 0, -1}.Mul(deltaTime * movementSpeed))
	}
	if of.KeyPressed(of.KeyS) {
		transform.MoveRelative(mgl32.Vec3{0, 0, 1}.Mul(deltaTime * movementSpeed))
	}
	if of.KeyPressed(of.KeyA) {
		transform.MoveRelative(mgl32.Vec3{-1, 0, 0}.Mul(deltaTime * movementSpeed))
	}
	if of.KeyPressed(of.KeyD) {
		transform.MoveRelative(mgl32.Vec3{1, 0, 0}.Mul(deltaTime * movementSpeed))
	}
	if of.MouseButtonPressed(of.MouseButtonLeft) {
		mouseMovement := of.GetMouseMovement()
		transform.Rotate(mgl32.Vec3{-1, 0, 0}.Mul(mouseMovement.Y() * deltaTime * mouseSpeed))
		transform.Rotate(mgl32.Vec3{0, -1, 0}.Mul(mouseMovement.X() * deltaTime * mouseSpeed))
	}
	of.SetComponent(camera, of.ComponentTransform, transform)

	if of.KeyPressed(of.Key1) {
		render.UpdateRenderer(currentFrame)
		if currentFrame < render.FrameCount-1 {
			currentFrame++
		}
	}
	if of.KeyPressed(of.Key2) {
		render.UpdateRenderer(currentFrame)
		if currentFrame > 0 {
			currentFrame--
		}
	}
}

func stop() {

}
