package Simulation

import (
	"github.com/go-gl/mathgl/mgl32"
	"math"
)

func normlizeVectorWithoutLength(v1 mgl32.Vec3, len float32) mgl32.Vec3 {
	l := 1.0 / len
	return mgl32.Vec3{v1[0] * l, v1[1] * l, v1[2] * l}
}

func Wpoly6(radiusSquared float32) float32 {

	coefficient := 315.0 / (64.0 * math.Pi * math.Pow(smoothingDistance, 9))
	hSquared := smoothingDistance * smoothingDistance

	return float32(coefficient * math.Pow(hSquared-float64(radiusSquared), 3))
}

func Wpoly6Gradient(diffPosition mgl32.Vec3, radiusSquared float32) mgl32.Vec3 {

	coefficient := -945.0 / (32.0 * math.Pi * math.Pow(smoothingDistance, 9))
	hSquared := smoothingDistance * smoothingDistance

	return diffPosition.Mul(float32(coefficient * math.Pow(hSquared-float64(radiusSquared), 2)))
}

func Wpoly6Laplacian(radiusSquared float32) float32 {

	coefficient := -945.0 / (32.0 * math.Pi * math.Pow(smoothingDistance, 9))
	hSquared := smoothingDistance * smoothingDistance

	return float32(coefficient * (hSquared - float64(radiusSquared)) * (3.0*hSquared - 7.0*float64(radiusSquared)))
}

func WspikyGradient(diffPosition mgl32.Vec3, radius float32) mgl32.Vec3 { //

	coefficient := -45.0 / (math.Pi * math.Pow(smoothingDistance, 6))

	return diffPosition.Mul(float32(coefficient * math.Pow(smoothingDistance-float64(radius), 2) / float64(radius)))
}

/*

double PARTICLE_SYSTEM::WviscosityLaplacian(double radiusSquared) {

static double coefficient = 45.0/(M_PI*pow(h,6));

double radius = sqrt(radiusSquared);

return coefficient * (h - radius);
}


*/
