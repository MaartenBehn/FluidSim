package main

import (
	"OctaForce"
	"log"
)

func main() {
	OctaForce.StartUp(start, stop)
}

func start() {
	OctaForce.SetMaxFPS(60)
	OctaForce.SetMaxUPS(20)
	OctaForce.AddUpdateCallback(update)
}

func update() {
	log.Print(OctaForce.GetFPS())
}

func stop() {

}
