package main

type World struct {
	particles []*Particle
	drops     []*Drop
}

var readWorld *World
var writeWorld *World

func swapWorlds() {
	readWorldBuffer := readWorld
	readWorld = writeWorld
	writeWorld = readWorldBuffer
}
