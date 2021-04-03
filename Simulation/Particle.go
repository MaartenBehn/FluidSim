package Simulation

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

type Particle struct {
	position mgl64.Vec3
	velocity mgl64.Vec3

	outside bool
}

func (p *Particle) calcOutside() {
	neigbors := p.getNeigbors(spacing * 1.1)
	if len(neigbors) <= 5 {
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

func (p *Particle) calcAcceleration() {
	displace := mgl64.Vec3{}
	for _, particle := range particles {
		if particle == p {
			continue
		}

		delta := p.position.Sub(particle.position)
		len := delta.Len()
		norm := delta.Mul(1 / len)
		fact := getGravitionalRejection(len)
		displace = displace.Add(norm.Mul(fact))

		if p.outside && particle.outside {
			fact = getElasticAttraction(len)
			displace = displace.Sub(norm.Mul(fact))
		}
	}
	p.velocity = displace
}

func (p *Particle) calcPosition() {
	len := p.velocity.Len()
	norm := p.velocity.Mul(1 / len)
	p.position = p.position.Add(norm.Mul(math.Min(len, maxVelocity)))
}

func getElasticAttraction(x float64) float64 {
	return math.Pow(x, 2) / Ke
}
func getGravitionalRejection(x float64) float64 {
	return math.Pow(Kg, 2) / x
}
