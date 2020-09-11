package main

import (
	"OctaForceEngineGo"
	"github.com/go-gl/mathgl/mgl32"
	"log"
)

func main() {
	OctaForce.StartUp(start, stop)
}

func start() {
	OctaForce.SetMaxFPS(60)
	OctaForce.SetMaxUPS(20)
	OctaForce.AddUpdateCallback(update)

	e1 := OctaForce.CreateEntity()
	OctaForce.AddComponent(e1, OctaForce.COMPONENT_Transform)
	transform1 := OctaForce.GetComponent(e1, OctaForce.COMPONENT_Transform).(OctaForce.Transform)
	transform1.Position = mgl32.Vec3{1, 1, 1}

	e2 := OctaForce.CreateEntity()
	OctaForce.AddComponent(e2, OctaForce.COMPONENT_Transform)
	transform2 := OctaForce.GetComponent(e2, OctaForce.COMPONENT_Transform).(OctaForce.Transform)
	transform2.Position = mgl32.Vec3{2, 2, 2}

	OctaForce.RemoveComponent(e2, OctaForce.COMPONENT_Transform)
	OctaForce.DeleteEntity(e2)
}

func update() {
	log.Print(OctaForce.GetFPS())
}

func stop() {

}
