package tests

import (
	"go-astronomy/internal/planets"
	"math"
	"testing"
)

func TestCalculateCoordinatesOfPlanet(t *testing.T) {
	raHrs, raMins, raSecs, decDeg, decMin, decSec := planets.CalculateCoordinatesOfPlanet(22.0, 11, 2003, "Jupiter", 1, 1, 2010)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(raHrs)-11.0) > tolerance || math.Abs(float64(raMins)-11.0) > tolerance || math.Abs(raSecs-13.75) > tolerance &&
		math.Abs(float64(decDeg)-6.0) > tolerance || math.Abs(float64(decMin)-21.0) > tolerance || math.Abs(decSec-23.88) > tolerance {
		t.Fatalf(`Error while Calculating Coordinates Of Planet Jupiter. Required: Rising = %d %d %f  Setting = %d %d %f   Got: Rising = %d %d %f  Setting = %d %d %f`, 11, 11, 13.75, 6, 21, 23.88, raHrs, raMins, raSecs, decDeg, decMin, decSec)
	}

	raHrs, raMins, raSecs, decDeg, decMin, decSec = planets.CalculateCoordinatesOfPlanet(22.0, 11, 2003, "Mercury", 1, 1, 2010)

	if math.Abs(float64(raHrs)-16.0) > tolerance || math.Abs(float64(raMins)-49.0) > tolerance || math.Abs(raSecs-12.26) > tolerance &&
		math.Abs(float64(decDeg)+24.0) > tolerance || math.Abs(float64(decMin)-30.0) > tolerance || math.Abs(decSec-3.14) > tolerance {
		t.Fatalf(`Error while Calculating Coordinates Of Planet Mercury. Required: Rising = %d %d %f  Setting = %d %d %f   Got: Rising = %d %d %f  Setting = %d %d %f`, 16, 49, 12.26, -24, 30, 3.14, raHrs, raMins, raSecs, decDeg, decMin, decSec)
	}
}
