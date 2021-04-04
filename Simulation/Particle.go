package main

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Particle struct {
	position mgl64.Vec3
	velocity mgl64.Vec3
	force    mgl64.Vec3
	mass     float64

	drop *Drop
}

func (p *Particle) calcForce() {
	p.force = mgl64.Vec3{}

	for _, particle := range p.drop.particles {
		if p == particle {
			continue
		}
		delta := particle.position.Sub(p.position)
		distance := delta.Len()
		norm := delta.Mul(1 / distance)

		springForce := springStiffness * (distance - p.drop.radius)

		velocityDelta := particle.velocity.Sub(p.velocity)
		damingForce := norm.Dot(velocityDelta) * springDampingFactor

		totalSpringForce := springForce + damingForce
		p.force = p.force.Add(norm.Mul(totalSpringForce))
	}
}

func (p *Particle) calcVelocity() {
	p.velocity = p.velocity.Add(p.force.Mul(dt / p.mass))
}

func (p *Particle) calcPosition() {
	p.position = p.position.Add(p.velocity.Mul(dt))
}

func (p *Particle) getNeigbors(neigborDist float64) ([]*Particle, []float64) {

	neigbors := make([]*Particle, 0)
	distances := make([]float64, 0)
	for _, particle := range particles {
		if p == particle {
			continue
		}

		distance := particle.position.Sub(p.position).Len()
		if distance <= neigborDist {
			neigbors = append(neigbors, particle)
			distances = append(distances, distance)
		}
	}
	return neigbors, distances
}
