package tests

import (
	sun "go-astronomy/internal/sun"
	"math"
	"testing"
)

func TestCalculatePositionOfSun(t *testing.T) {
	raHrs, raMin, raSec, decDeg, decMin, decSec, _ := sun.CalculatePositionOfSun(27.0, 7.0, 2003.0, 0.0, 0.0, 0.0)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(raHrs-8.0) > tolerance || math.Abs(raMin-23.0) > tolerance || math.Abs(raSec-33.65) > tolerance &&
		math.Abs(decDeg-19.0) > tolerance || math.Abs(decMin-21.0) > tolerance || math.Abs(decSec-10.38) > tolerance {
		t.Fatalf(`Error while Calculating Position Of Sun. Required:  %f %f %f    %f %f %f   Got: %f %f %f    %f %f %f`, 8.0, 23.0, 33.65, 19.0, 21.0, 10.38, raHrs, raMin, raSec, decDeg, decMin, decSec)
	}
}
