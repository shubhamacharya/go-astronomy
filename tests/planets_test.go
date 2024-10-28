package tests

import (
	"go-astronomy/internal/planets"
	"math"
	"testing"
)

func TestCalculateCoordinatesOfPlanet(t *testing.T) {
	raHrs, raMins, raSecs, decDeg, decMin, decSec := planets.CalculateCoordinatesOfPlanet(22.0, 11, 2003, "Jupiter", 0, 1, 2010)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(raHrs)-11.0) > tolerance || math.Abs(float64(raMins)-11.0) > tolerance || math.Abs(raSecs-13.72) > tolerance &&
		math.Abs(float64(decDeg)-6.0) > tolerance || math.Abs(float64(decMin)-21.0) > tolerance || math.Abs(decSec-23.21) > tolerance {
		t.Fatalf(`Error while Calculating Coordinates Of Planet Jupiter. Required: Rising = %d %d %f  Setting = %d %d %f   Got: Rising = %d %d %f  Setting = %d %d %f`, 11, 11, 13.72, 6, 21, 23.21, raHrs, raMins, raSecs, decDeg, decMin, decSec)
	}

	raHrs, raMins, raSecs, decDeg, decMin, decSec = planets.CalculateCoordinatesOfPlanet(22.0, 11, 2003, "Mercury", 0, 1, 2010)

	if math.Abs(float64(raHrs)-16.0) > tolerance || math.Abs(float64(raMins)-49.0) > tolerance || math.Abs(raSecs-12.30) > tolerance &&
		math.Abs(float64(decDeg)+24.0) > tolerance || math.Abs(float64(decMin)-30.0) > tolerance || math.Abs(decSec-0.29) > tolerance {
		t.Fatalf(`Error while Calculating Coordinates Of Planet Mercury. Required: Rising = %d %d %f  Setting = %d %d %f   Got: Rising = %d %d %f  Setting = %d %d %f`, 16, 49, 12.30, -24, 30, 0.29, raHrs, raMins, raSecs, decDeg, decMin, decSec)
	}
}

func TestCalculateApproximatePositionOfPlanet(t *testing.T) {
	raHrs, raMins, raSecs, decDeg, decMin, decSec := planets.CalculateApproximatePositionOfPlanet(22.0, 11, 2003, "Jupiter", 0, 1, 2010)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(raHrs)-10.0) > tolerance || math.Abs(float64(raMins)-58.0) > tolerance || math.Abs(raSecs-21.83) > tolerance &&
		math.Abs(float64(decDeg)-6.0) > tolerance || math.Abs(float64(decMin)-34.0) > tolerance || math.Abs(decSec-15.64) > tolerance {
		t.Fatalf(`Error while Calculating Approximate Position Of Planet Jupiter. Required: Rising = %d %d %f  Setting = %d %d %f   Got: Rising = %d %d %f  Setting = %d %d %f`, 10, 58, 21.83, 6, 34, 15.64, raHrs, raMins, raSecs, decDeg, decMin, decSec)
	}
}

func TestCalculatePerturbationsInPlanetsOrbit(t *testing.T) {
	raHrs, raMins, raSecs, decDeg, decMin, decSec := planets.CalculatePerturbationsInPlanetsOrbit(22.0, 11, 2003, "Jupiter", 0, 1, 2010)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(raHrs)-11.0) > tolerance || math.Abs(float64(raMins)-10.0) > tolerance || math.Abs(raSecs-47.00) > tolerance &&
		math.Abs(float64(decDeg)-6.0) > tolerance || math.Abs(float64(decMin)-24.0) > tolerance || math.Abs(decSec-12.00) > tolerance {
		t.Fatalf(`Error while Calculatinge Perturbations In Planets Orbit for Jupiter. Required: Rising = %d %d %f  Setting = %d %d %f   Got: Rising = %d %d %f  Setting = %d %d %f`, 11, 10, 47.00, 6, 24, 12.00, raHrs, raMins, raSecs, decDeg, decMin, decSec)
	}
}

func TestCalculatePlantesDistanceLightTravelAndAngularSize(t *testing.T) {
	distance, timeMin, timeSec, diameter := planets.CalculatePlantesDistanceLightTravelAndAngularSize(22.0, 11, 2003, "Jupiter", 0, 1, 2010)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(distance)-11.0) > tolerance || math.Abs(float64(timeMin)-10.0) > tolerance || math.Abs(timeSec-47.00) > tolerance || math.Abs(diameter-12.00) > tolerance {
		t.Fatalf(`Error while Calculating Plantes Distance Light Travel And Angular Size Of Planet Jupiter. Required:  Distance = %v\t Light Travel Time : %v Min %v Sec\t Diameter = %v\n Got: Distance = %v\t Light Travel Time : %v Min %v Sec\t Diameter = %v\n`, 11, 10, 47.00, 6, distance, timeMin, timeSec, diameter)
	}
}

func TestCalculatePhasesOfPlanets(t *testing.T) {
	const tolerance = 0.01 // Define an acceptable error range
	d, F := planets.CalculatePhasesOfPlanets(22.0, 11, 2003, "Jupiter", 0, 1, 2010)

	if math.Abs(float64(d)-11.0) > tolerance || math.Abs(float64(F)-10.0) > tolerance {
		t.Fatalf(`Error while Calculating Phases Of Planets. Required:  d = %v\t F : %v\n Got: d = %v\t F : %v\n`, 11, 10, d, F)
	}
}

func TestCalculatePositionAngleOfBrightLimb(t *testing.T) {
	const tolerance = 0.01 // Define an acceptable error range
	d := planets.CalculatePositionAngleOfBrightLimb(22.0, 11, 2003, "Mercury", 0, 1, 2010)

	if math.Abs(float64(d)-10.0) > tolerance {
		t.Fatalf(`Error while Calculating Position Angle Of Bright Limb. Required: F : %v\n Got: F : %v\n`, 11, d)
	}
}
