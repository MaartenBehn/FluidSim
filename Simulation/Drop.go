package main

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

type Drop struct {
	position  mgl64.Vec3
	radius    float64
	particles []*Particle
}

func (d *Drop) calcPosition() {
	d.position = mgl64.Vec3{}
	for _, particle := range d.particles {
		d.position = d.position.Add(particle.position)
	}
	d.position = d.position.Mul(1 / float64(len(d.particles)))
}
func (d *Drop) checkDrop() {
	for _, drop := range drops {
		if drop == d {
			continue
		}

		distance := drop.position.Sub(d.position).Len()
		if distance < (d.radius+drop.radius)*mergeFactor {
			d.mergeDrop(drop)
			break
		}
	}
}

func (d *Drop) mergeDrop(drop *Drop) {
	d.particles = append(d.particles, drop.particles...)
	for _, particle := range drop.particles {
		particle.drop = d
	}

	for i, testDrop := range drops {
		if testDrop == drop {
			drops = append(drops[:i], drops[i+1:]...)
			break
		}
	}
}

func (d *Drop) calcDropRadius() {
	d.radius = math.Sqrt(float64(len(d.particles))) * radiusFactor
}
