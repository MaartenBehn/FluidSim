package Simulation

import (
	"github.com/go-gl/mathgl/mgl64"
	"log"
)

type Particle struct {
	position     mgl64.Vec3
	velocity     mgl64.Vec3
	acceleration mgl64.Vec3

	outside bool
}

func (p *Particle) calcOutside() {
	neigbors := p.getNeigbors(spacing * 1.5)
	log.Print(len(neigbors))
	if len(neigbors) < 10 {
		p.outside = true
	}
}

func (p *Particle) getNeigbors(neigborDist float64) []*Particle {

	neigbors := make([]*Particle, 0)
	for _, particle := range particles {
		if particle != p && particle.position.Sub(p.position).Len() <= neigborDist {
			neigbors = append(neigbors, particle)
		}
	}
	return neigbors
}
