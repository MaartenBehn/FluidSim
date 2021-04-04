package Simulation

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Particle struct {
	position mgl64.Vec3
	velocity mgl64.Vec3
	force    mgl64.Vec3
	mass     float64

	springNeigbors []*Particle
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
func (p *Particle) calcSpringNeigbors() {
	for i := len(p.springNeigbors) - 1; i >= 0; i-- {
		neigbor := p.springNeigbors[i]
		distance := neigbor.position.Sub(p.position).Len()

		if distance > springNeigborMaxDist {
			p.springNeigbors = append(p.springNeigbors[:i], p.springNeigbors[i+1:]...)
		}
	}

	newNeigbors, _ := p.getNeigbors(springNeigborMinDist)
	for _, neigbor := range newNeigbors {
		isAlreadyInList := false
		for _, springNeigbor := range p.springNeigbors {
			if neigbor == springNeigbor {
				isAlreadyInList = true
				break
			}
		}

		if !isAlreadyInList {
			p.springNeigbors = append(p.springNeigbors, neigbor)
		}
	}
}

func (p *Particle) calcForce() {
	p.force = mgl64.Vec3{}

	for _, neigbor := range p.springNeigbors {

		distance := neigbor.position.Sub(p.position).Len()
		norm := neigbor.position.Sub(p.position).Normalize()

		totalSpringForce := p.getSpringForce(distance) + p.getSpringDampingForce(neigbor)
		p.force = p.force.Add(norm.Mul(totalSpringForce))
	}
}
func (p *Particle) getSpringForce(dist float64) float64 {
	return springStiffness * (dist - springRestingLenght)
}
func (p *Particle) getSpringDampingForce(particle *Particle) float64 {
	norm := particle.position.Sub(p.position).Normalize()
	velocityDelta := particle.velocity.Sub(p.velocity)
	damingForce := norm.Dot(velocityDelta) * springDampingFactor
	return damingForce
}

func (p *Particle) calcVelocity() {
	p.velocity = p.velocity.Add(p.force.Mul(dt / p.mass))
}

func (p *Particle) calcPosition() {
	p.position = p.position.Add(p.velocity.Mul(dt))
}
