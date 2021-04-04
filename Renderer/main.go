package main

import (
	of "OctaForceEngineGo"
	"github.com/go-gl/mathgl/mgl32"
	"os"
	"path"
	"runtime"
)

var absPath string

func main() {
	//defer profile.Start().Stop()

	_, b, _, _ := runtime.Caller(0)
	dir, err := os.Open(path.Join(b, "../../"))
	if err != nil {
		panic(err)
	}
	absPath = dir.Name()

	SelectDataFile(absPath)
	of.Init(start)
}

var camera int

func start() {
	camera := of.NewCamera()

	camera.Transform = of.NewTransform()
	camera.Transform.SetPosition(mgl32.Vec3{0, 0, 200})

	of.SetActiveCamera(camera)

	task := of.NewTask(func() {

		deltaTime := float32(of.GetDeltaTime())
		if of.KeyPressed(of.KeyW) {
			camera.Transform.MoveRelative(mgl32.Vec3{0, 0, -1}.Mul(deltaTime * movementSpeed))
		}
		if of.KeyPressed(of.KeyS) {
			camera.Transform.MoveRelative(mgl32.Vec3{0, 0, 1}.Mul(deltaTime * movementSpeed))
		}
		if of.KeyPressed(of.KeyA) {
			camera.Transform.MoveRelative(mgl32.Vec3{-1, 0, 0}.Mul(deltaTime * movementSpeed))
		}
		if of.KeyPressed(of.KeyD) {
			camera.Transform.MoveRelative(mgl32.Vec3{1, 0, 0}.Mul(deltaTime * movementSpeed))
		}
		if of.MouseButtonPressed(of.MouseButtonLeft) {
			mouseMovement := of.GetMouseMovement()
			camera.Transform.Rotate(mgl32.Vec3{-1, 0, 0}.Mul(mouseMovement.Y() * deltaTime * mouseSpeed))
			camera.Transform.Rotate(mgl32.Vec3{0, -1, 0}.Mul(mouseMovement.X() * deltaTime * mouseSpeed))
		}
	})
	task.SetRepeating(true)
	task.SetRaceTask(of.GetEngineTask(of.WindowUpdateTask), of.GetEngineTask(of.RenderTask))
	of.AddTask(task)

	SetUpRenderer(absPath)
	UpdateRenderer(0)

	task = of.NewTask(update)
	task.SetRepeating(true)
	of.AddTask(task)
}

const (
	movementSpeed float32 = 100
	mouseSpeed    float32 = 3
)

var currentFrame int

func update() {

	if of.KeyPressed(of.KeyQ) {
		UpdateRenderer(currentFrame)
		if currentFrame < FrameCount-1 {
			currentFrame++
		}
	}
	if of.KeyPressed(of.KeyE) {
		UpdateRenderer(currentFrame)
		if currentFrame > 0 {
			currentFrame--
		}
	}
}
