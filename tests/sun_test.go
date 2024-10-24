package tests

import (
	sun "go-astronomy/internal/sun"
	"math"
	"testing"
)

func TestCalculatePositionOfSun(t *testing.T) {

	const tolerance = 0.1 // Define an acceptable error range

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
		{27.0, 7, 2003, 0, 0, 0, 8, 23, 36.03, 19, 20, 59.72, 0.0, 1, 2010},
		// {10.0, 3, 1986, 0, 0, 0.0, 0, 40, 3.2, -4, 18, 41.28, 1.0, 1, 2010},
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

	if math.Abs(float64(raHrs)-8.0) > tolerance || math.Abs(float64(raMin)-26.0) > tolerance || math.Abs(raSec-3.26) > tolerance &&
		math.Abs(float64(decDeg-19.0)) > tolerance || math.Abs(float64(decMin-12.0)) > tolerance || math.Abs(decSec-44.38) > tolerance {
		t.Fatalf(`Error while Calculating Precise Position Of Sun. Required:  %d %d %f    %d %d %f   Got: %d %d %f    %d %d %f`, 8, 26, 3.26, 19, 12, 44.38, raHrs, raMin, raSec, decDeg, decMin, decSec)
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
	riseHrs, riseMin, riseSec, SetHrs, SetMin, SetSec := sun.CalculateSunsRiseAndSet(10, 3, 1986, 12, 0, 0.0, -71.05, 42.37, 34.0, 0, 0, -5.0, 1, 1, 2010)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(riseHrs)-6.0) > tolerance || math.Abs(float64(riseMin)-7.0) > tolerance || math.Abs(riseSec-45.36) > tolerance &&
		math.Abs(float64(SetHrs)-17.0) > tolerance || math.Abs(float64(SetMin)-43.0) > tolerance || math.Abs(SetSec-22.09) > tolerance {
		t.Fatalf(`Error while Calculating Suns Rise And Set. Required:  %d %d %f    %d %d %f   Got: %d %d %f    %d %d %f`, 6, 7, 45.36, 17, 43, 22.09, riseHrs, riseMin, riseSec, SetHrs, SetMin, SetSec)
	}

	riseHrs, riseMin, riseSec, SetHrs, SetMin, SetSec = sun.CalculateSunsRiseAndSet(15.0, 10, 2024, 12, 0, 0.0, 73.8567, 18.5204, 34.0, 0, 0, 5.5, 1, 1, 2010)

	if math.Abs(float64(riseHrs)-6.0) > tolerance || math.Abs(float64(riseMin)-33.0) > tolerance || math.Abs(riseSec-45.22) > tolerance &&
		math.Abs(float64(SetHrs)-18.0) > tolerance || math.Abs(float64(SetMin)-12.0) > tolerance || math.Abs(SetSec-22.87) > tolerance {
		t.Fatalf(`Error while Calculating Suns Rise And Set. Required:  %d %d %f    %d %d %f   Got: %d %d %f    %d %d %f`, 6, 33, 45.22, 18, 12, 22.87, riseHrs, riseMin, riseSec, SetHrs, SetMin, SetSec)
	}
}

func TestCalculateCalculateSunTwilight(t *testing.T) {
	riseTwilightHrs, riseTwilightMin, riseTwilightSec, setTwilightHrs, setTwilightMin, setTwilightSec := sun.CalculateCalculateSunTwilight(7.0, 9, 1979, 0, 0, 0.0, 0, 52.0, 34.0, 0, 0, 0.0, 0, 1, 2010)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(riseTwilightHrs)-3.0) > tolerance || math.Abs(float64(riseTwilightMin)-12.0) > tolerance || math.Abs(riseTwilightSec-30.136076) > tolerance &&
		math.Abs(float64(setTwilightHrs)-20.0) > tolerance || math.Abs(float64(setTwilightMin)-39.0) > tolerance || math.Abs(setTwilightSec-79.362118) > tolerance {
		t.Fatalf(`Error while Calculating Suns Twilight. Required:  %d %d %f    %d %d %f   Got: %d %d %f    %d %d %f`, 3, 12, 30.136076, 20, 39, 79.362118, riseTwilightHrs, riseTwilightMin, riseTwilightSec, setTwilightHrs, setTwilightMin, setTwilightSec)
	}
}

func TestCalculateTheEquationOfTime(t *testing.T) {
	eqHrs, eqMin, eqSec := sun.CalculateTheEquationOfTime(27.5, 7, 2010, 12, 0, 0.0, 0, 52.0, 34.0, 0, 0, 0.0, 0, 1, 2010)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(eqHrs)-0.0) > tolerance || math.Abs(float64(eqMin)-4.0) > tolerance || math.Abs(eqSec-30.541071) > tolerance {
		t.Fatalf(`Error while Calculating Suns Twilight. Required:  %d %d %f  Got: %d %d %f`, 0, 4, 30.541071, eqHrs, eqMin, eqSec)
	}
}
