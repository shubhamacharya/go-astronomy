package tests

import (
	sun "go-astronomy/internal/sun"
	"math"
	"testing"
)

func TestCalculatePositionOfSun(t *testing.T) {

	const tolerance = 0.01 // Define an acceptable error range

	tests := []struct {
		day                        float64
		month, year, hrs, min      int
		sec                        float64
		expectRAHrs, expectRAMin   int
		expectRASec                float64
		expectDecDeg, expectDecMin int
		expectDecSec               float64
		epochDay                   float64
		epochMonth, epochYear      int
	}{
		{27.0, 7, 2003, 0, 0, 0, 8, 23, 33.65, 19, 21, 10.38, 1.0, 1, 2010},
		{10.0, 3, 1986, 0, 0, 0.0, 0, 40, 3.2, -4, 18, 41.28, 1.0, 1, 2010},
	}

	for _, test := range tests {
		// fmt.Printf("\nday : %f\tmonth : %d\tyear : %d\n", test.day, test.month, test.year)
		raHrs, raMin, raSec, decDeg, decMin, decSec, _ := sun.CalculatePositionOfSun(test.day, test.month, test.year, test.hrs, test.min, test.sec, test.epochDay, test.epochMonth, test.epochYear)
		if math.Abs(float64(raHrs)-float64(test.expectRAHrs)) > tolerance || math.Abs(float64(raMin)-float64(test.expectRAMin)) > tolerance || math.Abs(raSec-test.expectRASec) > tolerance &&
			math.Abs(float64(decDeg-test.expectDecDeg)) > tolerance || math.Abs(float64(decMin-test.expectDecMin)) > tolerance || math.Abs(decSec-test.expectDecSec) > tolerance {
			t.Fatalf(`Error while Calculating Position Of Sun. Required:  %d %d %f    %d %d %f   Got: %d %d %f    %d %d %f`, test.expectRAHrs, test.expectRAMin, test.expectRASec, test.expectDecDeg, test.expectDecMin, test.expectDecSec, raHrs, raMin, raSec, decDeg, decMin, decSec)
		}
	}

}

func TestCalculatePrecisePositionOfSun(t *testing.T) {
	raHrs, raMin, raSec, decDeg, decMin, decSec, _ := sun.CalculatePrecisePositionOfSun(27.0, 7.0, 1988.0, 0.0, 0.0, 0.0, 1, 1, 2000)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(raHrs)-8.0) > tolerance || math.Abs(float64(raMin)-26.0) > tolerance || math.Abs(raSec-3.61) > tolerance &&
		math.Abs(float64(decDeg-19.0)) > tolerance || math.Abs(float64(decMin-12.0)) > tolerance || math.Abs(decSec-43.18) > tolerance {
		t.Fatalf(`Error while Calculating Precise Position Of Sun. Required:  %d %d %f    %d %d %f   Got: %d %d %f    %d %d %f`, 8, 26, 3.61, 19, 12, 43.18, raHrs, raMin, raSec, decDeg, decMin, decSec)
	}
}

func TestCalculateSunsDistanceAndAngularSize(t *testing.T) {
	r, thetaDeg, thetaMin, thetaSec, _ := sun.CalculateSunsDistanceAndAngularSize(27.0, 7.0, 1988.0, 0.0, 0.0, 0.0, 1.5, 1, 2000)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(r-151920130.15) > tolerance || math.Abs(float64(thetaDeg-0)) > tolerance || math.Abs(float64(thetaMin-31.0)) > tolerance &&
		math.Abs(thetaSec-29.93) > tolerance {
		t.Fatalf(`Error while Calculating Suns Distance And Angular Size. Required: Distance(r) : %f Angular Size(theta) : %d %d %f   Got: Distance(r) : %f Angular Size(theta) : %d %d %f`, 151920130.15, 0, 31, 29.93, r, thetaDeg, thetaMin, thetaSec)
	}
}

func TestCalculateSunsRiseAndSet(t *testing.T) {
	riseHrs, riseMin, riseSec, SetHrs, SetMin, SetSec := sun.CalculateSunsRiseAndSet(10.0, 3, 1986, 0, 0, 0.0, -71.05, 42.37, 34.0, 0, 0, -5.0, 1.5, 1, 2010)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(riseHrs)-6.0) > tolerance || math.Abs(float64(riseMin)-5.0) > tolerance || math.Abs(riseSec-30.468938) > tolerance &&
		math.Abs(float64(SetHrs)-17.0) > tolerance || math.Abs(float64(SetMin)-38.0) > tolerance || math.Abs(SetSec-13.791434) > tolerance {
		t.Fatalf(`Error while Calculating Suns Rise And Set. Required:  %d %d %f    %d %d %f   Got: %d %d %f    %d %d %f`, 6, 5, 30.468938, 17, 38, 13.791434, riseHrs, riseMin, riseSec, SetHrs, SetMin, SetSec)
	}

	// riseHrs, riseMin, riseSec, SetHrs, SetMin, SetSec := sun.CalculateSunsRiseAndSet(25.0, 9, 2024, 12, 0, 0.0, 17.4065, 78.4772, 34.0, 0, 0, 5.5, 1.5, 1, 2010)

	// if math.Abs(float64(riseHrs)-8.0) > tolerance || math.Abs(float64(riseMin)-26.0) > tolerance || math.Abs(riseSec-3.61) > tolerance &&
	// 	math.Abs(float64(SetHrs)-19.0) > tolerance || math.Abs(float64(SetMin)-12.0) > tolerance || math.Abs(SetSec-43.18) > tolerance {
	// 	t.Fatalf(`Error while Calculating Suns Rise And Set. Required:  %d %d %f    %d %d %f   Got: %d %d %f    %d %d %f`, 8, 26, 3.61, 19, 12, 43.18, riseHrs, riseMin, riseSec, SetHrs, SetMin, SetSec)
	// }
}
