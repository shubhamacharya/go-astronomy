package tests

import (
	"go-astronomy/internal/coords"
	datetime "go-astronomy/internal/dateTime"
	"go-astronomy/internal/macros"
	"math"
	"testing"
)

func TestConvertDecimalDegToDegMinSec(t *testing.T) {
	tests := []struct {
		input                           float64
		expectedD, expectedM, expectedS float64
	}{
		{182.524167, 182, 31, 27.0},    // standard case
		{0.0, 0, 0, 0.0},               // zero degrees
		{45.7625, 45, 45, 45.0},        // exact seconds
		{359.999999, 359, 59, 59.9964}, // edge near 360°
		{-73.987654, -73, 59, 15.5544}, // negative angle
		{180.5, 180, 30, 0.0},          // exact half degree
		{123.000278, 123, 0, 1.0008},   // very small seconds
		{90.25, 90, 15, 0.0},           // quarter degree
		{89.9999, 89, 59, 59.64},       // near integer rollover
	}

	const tolerance = 0.01

	for _, tt := range tests {
		deg, min, sec := macros.ConvertDecimalDegToDegMinSec(tt.input)
		if math.Abs(deg-tt.expectedD) > tolerance ||
			math.Abs(min-tt.expectedM) > tolerance ||
			math.Abs(sec-tt.expectedS) > tolerance {
			t.Errorf(`For input %.6f, expected %f° %f' %.4f", got %f° %f' %.4f"`,
				tt.input, tt.expectedD, tt.expectedM, tt.expectedS, deg, min, sec)
		}
	}
}

func TestConvertDegMinSecToDecimalDeg(t *testing.T) {
	decimalDeg := macros.ConvertDegMinSecToDecimalDeg(182, 31, 27)
	var expectedDeg float64 = 182.524167
	const tolerance = 0.000001 // Define an acceptable error range

	if math.Abs(decimalDeg-expectedDeg) > tolerance {
		t.Fatalf(`Error while converting Deg Min Sec to degrees. Required: %f    Got: %f`, expectedDeg, decimalDeg)
	}
}

func TestConvertDecimalHrsToDecimalDegress(t *testing.T) {
	decimalDeg := coords.ConvertDecimalHrsToDecimalDegress(datetime.ConvertHrsMinSecToDecimalHrs(9, 36, 10.2, false, false))
	deg, min, sec := macros.ConvertDecimalDegToDegMinSec(decimalDeg)

	if math.Abs(float64(deg)-144.0) > 0.001 || math.Abs(float64(min)-2.0) > 0.001 || math.Abs(sec-32.982) > 0.001 {
		t.Fatalf(`Error while converting Decimal Hrs To Decimal Degress. Required: %f %f %f   Got: %f %f %f`, 144.0, 2.0, 33.0, deg, min, sec)
	}
}

func TestConvertDecimalDegressToDecimalHrs(t *testing.T) {
	decimalHrs := macros.ConvertDecimalDegressToDecimalHrs(macros.ConvertDegMinSecToDecimalDeg(144.0, 2.0, 33.0))
	hrs, min, sec := datetime.ConvertDecimalHrsToHrsMinSec(decimalHrs)
	const tolerance = 0.000001 // Define an acceptable error range

	if math.Abs(float64(hrs)-9) > tolerance || math.Abs(float64(min)-36) > tolerance || math.Abs(sec-10.2) > tolerance {
		t.Fatalf(`Error while converting Decimal Deg To Hrs Min Sec. Required: %f %f %f   Got: %f %f %f`, 9.0, 36.0, 10.2, hrs, min, sec)
	}
}

func TestConverRightAscensionToHourAngle(t *testing.T) {
	haHrs, haMin, haSec, _ := coords.ConverRightAscensionToHourAngle(22.0, 4.0, 1980.0, datetime.AD, 14.0, 36.0, 51.67, 18.0, 32.0, 21.0, -64.0, 0.0, 0.0, -4.0, true)
	// hrs, min, sec := datetime.ConvertDecimalHrsToHrsMinSec(decimalHrs)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(haHrs)-9.0) > tolerance || math.Abs(float64(haMin)-52.0) > tolerance || math.Abs(haSec-23.66) > tolerance {
		t.Fatalf(`Error while converting Right Ascension To Hour Angle. Required: %f %f %f   Got: %f %f %f`, 9.0, 52.0, 23.66, haHrs, haMin, haSec)
	}
}

func TestConverHourAngleToRightAscension(t *testing.T) {
	haHrs, haMin, haSec, _ := coords.ConverHourAngleToRightAscension(22, 4, 1980, datetime.AD, 14, 36, 51.67, 9.0, 52.0, 23.66, -64, 0, 0, -4)
	// hrs, min, sec := datetime.ConvertDecimalHrsToHrsMinSec(decimalHrs)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(haHrs)-18.0) > tolerance || math.Abs(float64(haMin)-32.0) > tolerance || math.Abs(haSec-21.0) > tolerance {
		t.Fatalf(`Error while converting Right Ascension To Hour Angle. Required: %f %f %f   Got: %f %f %f`, 18.0, 32.0, 21.0, haHrs, haMin, haSec)
	}
}

func TestConvertEquatorialToHorizonCoordinates(t *testing.T) {
	altitudeDeg, altitudeMin, altitudeSec, azimuthDeg, azimuthMin, azimuthSec := coords.ConvertEquatorialToHorizonCoordinates(5, 51, 44, 23, 13, 10.00, 52)
	// hrs, min, sec := datetime.ConvertDecimalHrsToHrsMinSec(decimalHrs)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(altitudeDeg)-19) > tolerance || math.Abs(float64(altitudeMin)-20) > tolerance || math.Abs(altitudeSec-3.64) > tolerance &&
		math.Abs(float64(azimuthDeg)-283) > tolerance || math.Abs(float64(azimuthMin)-16) > tolerance || math.Abs(azimuthSec-15.69) > tolerance {
		t.Fatalf(`Error while converting Equatorial To Horizon Coordinates. Required: %f %f %f, %f %f %f   Got: %f %f %f, %f %f %f`, 19.0, 20.0, 3.64, 283.0, 16.0, 15.69, altitudeDeg, altitudeMin, altitudeSec, azimuthDeg, azimuthMin, azimuthSec)
	}
}

func TestConvertHorizonCoordinatesToEquatorial(t *testing.T) {
	tests := []struct {
		GSTHrs, GSTMin, GSec                           float64
		altitudeDeg, altitudeMin, altitudeSec          float64
		azimuthDeg, azimuthMin, azimuthSec, latitude   float64
		expectedHaHrs, expectedHaMin, expectedHaSec    float64
		expectedDecDeg, expectedDecMin, expectedDecSec float64
	}{
		{0, 24.0, 05.0, 19.0, 20.0, 03.64, 283.0, 16.0, 15.7, 52.0, 5, 51, 44.0, 23, 13, 9.98},
	}
	for _, test := range tests {
		haHrs, haMin, haSec, decDeg, decMin, decSec := coords.ConvertHorizonCoordinatesToEquatorial(test.GSTHrs, test.GSTMin, test.GSec, test.altitudeDeg, test.altitudeMin, test.altitudeSec, test.azimuthDeg, test.azimuthMin, test.azimuthSec, test.latitude)
		const tolerance = 0.01 // Define an acceptable error range
		if math.Abs(float64(haHrs)-float64(test.expectedHaHrs)) > tolerance || math.Abs(float64(haMin)-float64(test.expectedHaMin)) > tolerance || math.Abs(haSec-test.expectedHaSec) > tolerance &&
			math.Abs(float64(decDeg)-float64(test.expectedDecDeg)) > tolerance || math.Abs(float64(decMin)-float64(test.expectedDecMin)) > tolerance || math.Abs(decSec-test.expectedDecSec) > tolerance {
			t.Fatalf(`Error while converting Horizon To Equatorial Coordinates. Required: %f %f %f, %f %f %f   Got: %f %f %f, %f %f %f`, test.expectedHaHrs, test.expectedHaMin, test.expectedHaSec, test.expectedDecDeg, test.expectedDecMin, test.expectedDecSec, haHrs, haMin, haSec, decDeg, decMin, decSec)
		}
	}
}

func TestCalculateEclipticMeanObliquity(t *testing.T) {
	obliquityDeg, obliquityMin, obliquitySec, _ := macros.CalculateEclipticMeanObliquity(6.0, 7, 2009, datetime.AD)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs((obliquityDeg)-23.0) > tolerance || math.Abs((obliquityMin)-26.0) > tolerance || math.Abs(obliquitySec-17.0) > tolerance {
		t.Fatalf(`Error while Calculating Ecliptic Mean Obliquity. Required: %f %f %f  Got: %f %f %f`, 23.0, 26.0, 17.0, obliquityDeg, obliquityMin, obliquitySec)
	}
}

func TestConvertEclipticCoordinatesToEquatorial(t *testing.T) {
	const tolerance = 0.1 // Define an acceptable error range
	tests := []struct {
		day, month, year                                                   float64
		YearLabel                                                          datetime.YearLabel
		eclipticLongDeg, eclipticLongMin, eclipticLongSec                  float64
		eclipticLatDeg, eclipticLatMin, eclipticLatSec, epochDay           float64
		epochMonth, epochYear, expectedRaHrs, expectedRaMin, expectedRaSec float64
		expectedDecDeg, expectedDecMin, expectedDecSec                     float64
	}{
		{6.0, 7, 2009, datetime.AD, 139.0, 41.0, 10.0, 4.0, 52.0, 31.0, 1, 1, 2010, 9, 34, 53.32, 19, 32, 5.89},
		{25.0, 9, 2024, datetime.AD, 182.0, 2.0, 27.2688, 0.0, 0.0, 0.0, 1, 1, 2010, 12, 7, 29.43, 0, 48, 41.88},
	}
	for _, test := range tests {
		raHrs, raMin, raSec, decDeg, decMin, decSec := macros.ConvertEclipticCoordinatesToEquatorial(test.day, test.month, test.year, test.YearLabel, test.eclipticLongDeg, test.eclipticLongMin, test.eclipticLongSec, test.eclipticLatDeg, test.eclipticLatMin, test.eclipticLatSec) //, test.epochDay, test.epochMonth, test.epochYear)
		if math.Abs(float64(raHrs)-float64(test.expectedRaHrs)) > tolerance || math.Abs(float64(raMin)-float64(test.expectedRaMin)) > tolerance || math.Abs(raSec-test.expectedRaSec) > tolerance &&
			math.Abs(float64(decDeg)-float64(test.expectedDecDeg)) > tolerance || math.Abs(float64(decMin)-float64(test.expectedDecMin)) > tolerance || math.Abs(decSec-test.expectedDecSec) > tolerance {
			t.Fatalf(`Error while converting Horizon To Equatorial Coordinates. Required: %f %f %f, %f %f %f   Got: %f %f %f, %f %f %f`, test.expectedRaHrs, test.expectedRaMin, test.expectedRaSec, test.expectedDecDeg, test.expectedDecMin, test.expectedDecSec, raHrs, raMin, raSec, decDeg, decMin, decSec)
		}
	}
}

func TestConvertEquatorialCoordinatesToEcliptic(t *testing.T) {
	latDeg, latMin, latSec, longDeg, longMin, longSec := coords.ConvertEquatorialCoordinatesToEcliptic(6.0, 7, 2009, datetime.AD, 9.0, 34.0, 53.32, 19.0, 32.0, 6.01, 1, 1, 2010)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(longDeg)-139) > tolerance || math.Abs(float64(longMin)-41) > tolerance || math.Abs(longSec-9.98) > tolerance &&
		math.Abs(float64(latDeg)-4) > tolerance || math.Abs(float64(latMin)-52) > tolerance || math.Abs(latSec-30.99) > tolerance {
		t.Fatalf(`Error while convert Equatorial Coordinates to Ecliptic. Required: %f %f %f, %f %f %f   Got: %f %f %f, %f %f %f`, 139.0, 41.0, 9.98, 4.0, 52.0, 30.99, longDeg, longMin, longSec, latDeg, latMin, latSec)
	}
}

func TestConvertEquatorialCoordinateToGalactic(t *testing.T) {
	const tolerance = 0.01 // Define an acceptable error range
	tests := []struct {
		raHrs, raMin, raSec                      float64
		decDeg, decMin, decSec                   float64
		expectedLDeg, expectedLMin, expectedLSec float64
		expectedBDeg, expectedBMin, expectedBSec float64
	}{
		{10.0, 21.0, 0.0, 10.0, 3.0, 11.00, 232, 14, 52.47, 51, 7, 20.32},
	}
	for _, test := range tests {
		lDeg, lMin, lSec, bDeg, bMin, bSec := coords.ConvertEquatorialCoordinateToGalactic(10.0, 21.0, 0.0, 10.0, 3.0, 11.00)
		if math.Abs((lDeg)-(test.expectedLDeg)) > tolerance || math.Abs((lMin)-(test.expectedLMin)) > tolerance || math.Abs(lSec-test.expectedLSec) > tolerance &&
			math.Abs((bDeg)-(test.expectedBDeg)) > tolerance || math.Abs((bMin)-(test.expectedBMin)) > tolerance || math.Abs(bSec-test.expectedBSec) > tolerance {
			t.Fatalf(`Error while convert Equatorial Coordinate To Galactic. Required: %f %f %f, %f %f %f   Got: %f %f %f, %f %f %f`, test.expectedLDeg, test.expectedLMin, test.expectedLSec, test.expectedBDeg, test.expectedBMin, test.expectedBSec, lDeg, lMin, lSec, bDeg, bMin, bSec)
		}
	}
}

func TestConvertGalacticCoordinateToEquatorial(t *testing.T) {
	lDeg, lMin, lSec, bDeg, bMin, bSec := coords.ConvertGalacticCoordinateToEquatorial(232.0, 14.0, 52.0, 51.0, 7.0, 20.00)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(lDeg)-10) > tolerance || math.Abs(float64(lMin)-21) > tolerance || math.Abs(lSec-0.0) > tolerance &&
		math.Abs(float64(bDeg)-10) > tolerance || math.Abs(float64(bMin)-3) > tolerance || math.Abs(bSec-11.11) > tolerance {
		t.Fatalf(`Error while convert Galactic Coordinate To Equatorial. Required: %f %f %f, %f %f %f   Got: %f %f %f, %f %f %f`, 10.0, 21.0, 0.0, 10.0, 3.0, 11.11, lDeg, lMin, lSec, bDeg, bMin, bSec)
	}
}

func TestCalculateAngleBetweenTwoCelestialObjects(t *testing.T) {
	Deg, Min, Sec := coords.CalculateAngleBetweenTwoCelestialObjects(5.0, 13.0, 31.7, -8.0, 13.0, 30.0, 6.0, 44.0, 13.4, -16.0, 41.0, 11.0)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs((Deg)-23) > tolerance || math.Abs((Min)-40) > tolerance || math.Abs(Sec-25.89) > tolerance {
		t.Fatalf(`Error while Calculating Angle Between Two Celestial Objects. Required: %f %f %f   Got: %f %f %f`, 23.0, 40.0, 25.89, Deg, Min, Sec)
	}
}

func TestCalculateRisingAndSettingTime(t *testing.T) {
	UTrHrs, UTrMin, UTrSec, UTsHrs, UTsMin, UTsSec, _, _ := coords.CalculateRisingAndSettingTime(24.0, 8, 2010, datetime.AD, 23.0, 39.0, 20.0, 21.0, 42.0, 0.0, 30.0, 64.00, 34)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs((UTrHrs)-14.0) > tolerance || math.Abs((UTrMin)-16.0) > tolerance || math.Abs(UTrSec-18.02) > tolerance &&
		math.Abs((UTsHrs)-4.0) > tolerance || math.Abs((UTsMin)-10.0) > tolerance || math.Abs(UTsSec-1.15) > tolerance {
		t.Fatalf(`Error while Calculating Rising And Setting Time. Required: Rising = %f %f %f  Setting = %f %f %f   Got: Rising = %f %f %f  Setting = %f %f %f`, 14.0, 16.0, 18.02, 4.0, 10.0, 1.15, UTrHrs, UTrMin, UTrSec, UTsHrs, UTsMin, UTsSec)
	}
}

func TestCalculatePrecession(t *testing.T) {
	alpha1Hrs, alpha1Min, alpha1Sec, delta1Deg, delta1Min, delta1Sec := coords.CalculatePrecession(1979.5, 1950.0, 9.0, 10.0, 43.0, 14.0, 23.0, 25.0)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs((alpha1Hrs)-9) > tolerance || math.Abs((alpha1Min)-12) > tolerance || math.Abs(alpha1Sec-20.47) > tolerance &&
		math.Abs((delta1Deg)-14) > tolerance || math.Abs((delta1Min)-16) > tolerance || math.Abs(delta1Sec-7.83) > tolerance {
		t.Fatalf(`Error while Calculating Precession. Required:  %f %f %f   %f %f %f   Got: %f %f %f   %f %f %f`, 9.0, 12.0, 20.47, 14.0, 16.0, 7.83, alpha1Hrs, alpha1Min, alpha1Sec, delta1Deg, delta1Min, delta1Sec)
	}
}

func TestCalculateNutation(t *testing.T) {
	nutationInLong, nutationInObliquity := coords.CalculateNutation(1.0, 9, 1988, datetime.AD)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(nutationInLong-5.49) > tolerance || math.Abs(nutationInObliquity-9.24) > tolerance {
		t.Fatalf(`Error while Calculating Nutation. Required: %f  %f   Got: %f  %f`, 5.49, 9.24, nutationInLong, nutationInObliquity)
	}
}

func TestCalculateAberration(t *testing.T) {
	correctedLambdaDeg, correctedLambdaMin, correctedLambdaSec, correctedBetaDeg, correctedBetaMin, correctedBetaSec := coords.CalculateAberration(8.0, 9, 1988, 352.0, 37.0, 10.1, -1, 32, 56.4, 165.0, 33.0, 44.1)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(correctedLambdaDeg)-352) > tolerance || math.Abs(float64(correctedLambdaMin)-37) > tolerance || math.Abs(correctedLambdaSec-30.45) > tolerance &&
		math.Abs(float64(correctedBetaDeg)-(-1)) > tolerance || math.Abs(float64(correctedBetaMin)-32) > tolerance || math.Abs(correctedBetaSec-56.33) > tolerance {
		t.Fatalf(`Error while Calculating Aberration. Required:  %f %f %f    %f %f %f   Got: %f %f %f    %f %f %f`, 352.0, 37.0, 30.45, -1.0, 32.0, 56.33, correctedLambdaDeg, correctedLambdaMin, correctedLambdaSec, correctedBetaDeg, correctedBetaMin, correctedBetaSec)
	}
}

func TestCalculateRefraction(t *testing.T) {
	HaHrs, HaMin, HaSec, DecDeg, DecMin, DecSec := coords.CalculateRefraction(5.0, 51.0, 44.0, 23.0, 13.0, 10.0, 52.0, 13.0, 1008.0)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(float64(HaHrs)-5) > tolerance || math.Abs(float64(HaMin)-51) > tolerance || math.Abs(HaSec-36.26) > tolerance &&
		math.Abs(float64(DecDeg)-23) > tolerance || math.Abs(float64(DecMin)-15) > tolerance || math.Abs(DecSec-13.91) > tolerance {
		t.Fatalf(`Error while Calculating Refraction. Required:  %f %f %f    %f %f %f   Got: %f %f %f    %f %f %f`, 5.0, 51.0, 36.26, -23.0, 15.0, 13.91, HaHrs, HaMin, HaSec, DecDeg, DecMin, DecSec)
	}
}

func TestCalculateGeocentricParallax(t *testing.T) {
	pSin, pCos := coords.CalculateGeocentricParallax(60.0, 100.0, 50.0)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(pSin-0.762422) > tolerance || math.Abs(pCos-0.644060) > tolerance {
		t.Fatalf(`Error while Calculating GeocentricParallax. Required:  %f %f Got: %f %f`, 0.762422, 0.644060, pSin, pCos)
	}
}

func TestCalculateParallaxCorrections(t *testing.T) {
	// Test data for moon
	raMoonHrs, raMoonMin, raMoonSec, decMoonDeg, decMoonMin, decMoonSec := coords.CalculateParallaxCorrections(26.0, 2, 1979, datetime.AD, 16.0, 45.0, 0.0, 60.0, 100.0, 50.0, 22.0, 35.0, 19.0, -7.0, 41.0, 13.0, 1.0, 1.0, 9.0, 0.0)
	// Test data for sun and other planets
	raHrs, raMin, raSec, decDeg, decMin, decSec := coords.CalculateParallaxCorrections(26.0, 2, 1979, datetime.AD, 16.0, 45.0, 0.0, 60.0, 100.0, 50.0, 22.0, 36.0, 44.0, -8.0, 44.0, 24.0, 0.0, 0.0, 0.0, 0.9901)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs((raMoonHrs)-22.0) > tolerance || math.Abs((raMoonMin)-36.0) > tolerance || math.Abs(raMoonSec-43.21) > tolerance &&
		math.Abs((decMoonDeg)-(-8.0)) > tolerance || math.Abs((decMoonMin)-32.0) > tolerance || math.Abs(decMoonSec-17.39) > tolerance {
		t.Fatalf(`Error while Calculating Parallax Corrections for Moon. Required:  %f %f %f    %f %f %f   Got: %f %f %f    %f %f %f`, 22.0, 36.0, 43.21, -8.0, 32.0, 17.39, raMoonHrs, raMoonMin, raMoonSec, decMoonDeg, decMoonMin, decMoonSec)
	}

	if math.Abs((raHrs)-22.0) > tolerance || math.Abs((raMin)-36.0) > tolerance || math.Abs(raSec-44.00) > tolerance &&
		math.Abs((decDeg)-(-8.0)) > tolerance || math.Abs((decMin)-44.0) > tolerance || math.Abs(decSec-31.43) > tolerance {
		t.Fatalf(`Error while Calculating Parallax Corrections for sun and other planets. Required:  %f %f %f    %f %f %f   Got: %f %f %f    %f %f %f`, 22.0, 36.0, 44.00, -8.0, 44.0, 31.43, raHrs, raMin, raSec, decDeg, decMin, decSec)
	}
}

func TestCalculateHeliographicCoordinates(t *testing.T) {
	longitude, latitude := coords.CalculateHeliographicCoordinates(1.0, 5, 1988, datetime.AD, 0, 0, 0, 40.0, 50.0, 37.0, 220.0, 10.5, 0, 15.0, 52.0, 0.5, 1, 1900)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(longitude-(-19.95)) > tolerance || math.Abs(latitude-143.57) > tolerance {
		t.Fatalf(`Error while Calculating Heliographic Coordinates. Required:  %f %f Got: %f %f`, -19.95, 143.57, longitude, latitude)
	}
}

func TestCalculateCarringtonRotationNumbers(t *testing.T) {
	CRN := coords.CalculateCarringtonRotationNumbers(27.0, 1, 1975, datetime.AD)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(CRN-1624.0) > tolerance {
		t.Fatalf(`Error while Calculating Carrington Rotation Numbers. Required:  %f Got: %f`, 1624.0, CRN)
	}
}

func TestCalculateSelenographicCoordinatesOfMoon(t *testing.T) {
	le, be, C := coords.CalculateSelenographicCoordinatesOfMoon(1.0, 5, 1988, datetime.AD, 209.12, -3.08, 23.4433)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(le-(-4.88)) > tolerance || math.Abs(be-4.04) > tolerance || math.Abs(C-19.78) > tolerance {
		t.Fatalf(`Error while Calculating Selenographic Coordinates of Moon. Required: le : %f\tbe : %f\tC : %f Got: le : %f\tbe : %f\tC : %f`, -4.88, 4.04, 19.78, le, be, C)
	}
}

func TestCalculateSelenographicCoordinatesOfSun(t *testing.T) {
	ls, bs, colongitude := coords.CalculateSelenographicCoordinatesOfSun(1.0, 5, 1988, datetime.AD, 0, 0, 0, 209.12, -3.08, 23.4433, 55.952, 1.0076, 40.8437)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(ls-(6.81)) > tolerance || math.Abs(bs-1.18) > tolerance || math.Abs(colongitude-83.18) > tolerance {
		t.Fatalf(`Error while Calculating Selenographic Coordinates of Sun. Required: ls : %f\tbs : %f\tcolongitude : %f Got: ls : %f\tbs : %f\tcolongitude : %f`, 6.81, 1.18, 19.78, ls, bs, colongitude)
	}
}
