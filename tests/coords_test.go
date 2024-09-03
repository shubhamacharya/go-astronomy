package tests

import (
	"go-astronomy/internal/coords"
	datetime "go-astronomy/internal/dateTime"
	"go-astronomy/internal/macros"
	"math"
	"testing"
)

func TestConvertDecimalDegToDegMinSec(t *testing.T) {
	Deg, Mins, Sec := macros.ConvertDecimalDegToDegMinSec(182.524167)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(Deg-182) > tolerance || math.Abs(Mins-31) > tolerance || math.Abs(Sec-27) > tolerance {
		t.Fatalf(`Error while converting degrees to Deg Min Sec. Required: %f %f %f    Got: %f %f %f`, 182.0, 31.0, 27.0, Deg, Mins, Sec)
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
	decimalDeg := coords.ConvertDecimalHrsToDecimalDegress(datetime.ConvertHrsMinSecToDecimalHrs(9, 36, 10.2, false, ""))
	deg, min, sec := macros.ConvertDecimalDegToDegMinSec(decimalDeg)

	if math.Abs(deg-144.0) > 0.001 || math.Abs(min-2.0) > 0.001 || math.Abs(sec-33.0) > 0.001 {
		t.Fatalf(`Error while converting Decimal Hrs To Decimal Degress. Required: %f %f %f   Got: %f %f %f`, 144.0, 2.0, 33.0, deg, min, sec)
	}
}

func TestConvertDecimalDegressToDecimalHrs(t *testing.T) {
	decimalHrs := macros.ConvertDecimalDegressToDecimalHrs(macros.ConvertDegMinSecToDecimalDeg(144.0, 2.0, 33.0))
	hrs, min, sec := datetime.ConvertDecimalHrsToHrsMinSec(decimalHrs)
	const tolerance = 0.000001 // Define an acceptable error range

	if math.Abs(hrs-9.0) > tolerance || math.Abs(min-36.0) > tolerance || math.Abs(sec-10.2) > tolerance {
		t.Fatalf(`Error while converting Decimal Deg To Hrs Min Sec. Required: %f %f %f   Got: %f %f %f`, 9.0, 36.0, 10.2, hrs, min, sec)
	}
}

func TestConverRightAscensionToHourAngle(t *testing.T) {
	haHrs, haMin, haSec, _ := coords.ConverRightAscensionToHourAngle(22, 4, 1980, 14, 36, 51.67, 18.0, 32.0, 21.0, -64, 0, 0, -4, true)
	// hrs, min, sec := datetime.ConvertDecimalHrsToHrsMinSec(decimalHrs)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(haHrs-9.0) > tolerance || math.Abs(haMin-52.0) > tolerance || math.Abs(haSec-23.66) > tolerance {
		t.Fatalf(`Error while converting Right Ascension To Hour Angle. Required: %f %f %f   Got: %f %f %f`, 9.0, 52.0, 23.66, haHrs, haMin, haSec)
	}
}

func TestConverHourAngleToRightAscension(t *testing.T) {
	haHrs, haMin, haSec, _ := coords.ConverHourAngleToRightAscension(22, 4, 1980, 14, 36, 51.67, 9.0, 52.0, 23.66, -64, 0, 0, -4)
	// hrs, min, sec := datetime.ConvertDecimalHrsToHrsMinSec(decimalHrs)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(haHrs-18.0) > tolerance || math.Abs(haMin-32.0) > tolerance || math.Abs(haSec-21.0) > tolerance {
		t.Fatalf(`Error while converting Right Ascension To Hour Angle. Required: %f %f %f   Got: %f %f %f`, 18.0, 32.0, 21.0, haHrs, haMin, haSec)
	}
}

func TestConvertEquatorialToHorizonCoordinates(t *testing.T) {
	altitudeDeg, altitudeMin, altitudeSec, azimuthDeg, azimuthMin, azimuthSec := coords.ConvertEquatorialToHorizonCoordinates(5, 51, 44, 23, 13, 10.00, 52)
	// hrs, min, sec := datetime.ConvertDecimalHrsToHrsMinSec(decimalHrs)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(altitudeDeg-19.33) > tolerance || math.Abs(altitudeMin-20.0) > tolerance || math.Abs(altitudeSec-3.64) > tolerance &&
		math.Abs(azimuthDeg-283.27) > tolerance || math.Abs(azimuthMin-16.0) > tolerance || math.Abs(azimuthSec-15.69) > tolerance {
		t.Fatalf(`Error while converting Equatorial To Horizon Coordinates. Required: %f %f %f, %f %f %f   Got: %f %f %f, %f %f %f`, 19.33, 20.0, 3.64, 283.27, 16.0, 15.69, altitudeDeg, altitudeMin, altitudeSec, azimuthDeg, azimuthMin, azimuthSec)
	}
}

func TestConvertHorizonCoordinatesToEquatorial(t *testing.T) {
	haHrs, haMin, haSec, decDeg, decMin, decSec := coords.ConvertHorizonCoordinatesToEquatorial(0, 24.0, 05.0, 19.0, 20.0, 03.64, 283.0, 16.0, 15.7, 52.0)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(haHrs-5.0) > tolerance || math.Abs(haMin-51.0) > tolerance || math.Abs(haSec-44.0) > tolerance &&
		math.Abs(decDeg-23.0) > tolerance || math.Abs(decMin-13.0) > tolerance || math.Abs(decSec-10.0) > tolerance {
		t.Fatalf(`Error while converting Horizon To Equatorial Coordinates. Required: %f %f %f, %f %f %f   Got: %f %f %f, %f %f %f`, 5.0, 51.0, 44.0, 23.0, 13.0, 10.0, haHrs, haMin, haSec, decDeg, decMin, decSec)
	}
}

func TestCalculateEclipticMeanObliquity(t *testing.T) {
	obliquityDeg, obliquityMin, obliquitySec, _ := macros.CalculateEclipticMeanObliquity(6.0, 7, 2009)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(obliquityDeg-23.0) > tolerance || math.Abs(obliquityMin-26.0) > tolerance || math.Abs(obliquitySec-17.0) > tolerance {
		t.Fatalf(`Error while Calculating Ecliptic Mean Obliquity. Required: %f %f %f  Got: %f %f %f`, 23.0, 26.0, 17.0, obliquityDeg, obliquityMin, obliquitySec)
	}
}

func TestConvertEclipticCoordinatesToEquatorial(t *testing.T) {
	raHrs, raMin, raSec, decDeg, decMin, decSec := macros.ConvertEclipticCoordinatesToEquatorial(6.0, 7, 2009, 139.0, 41.0, 10.0, 4.0, 52.0, 31.0)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(raHrs-9.0) > tolerance || math.Abs(raMin-34.0) > tolerance || math.Abs(raSec-53.32) > tolerance &&
		math.Abs(decDeg-19.0) > tolerance || math.Abs(decMin-32.0) > tolerance || math.Abs(decSec-6.01) > tolerance {
		t.Fatalf(`Error while converting Horizon To Equatorial Coordinates. Required: %f %f %f, %f %f %f   Got: %f %f %f, %f %f %f`, 9.0, 34.0, 53.32, 19.0, 32.0, 6.01, raHrs, raMin, raSec, decDeg, decMin, decSec)
	}
}

func TestConvertEquatorialCoordinatesToEcliptic(t *testing.T) {
	latDeg, latMin, latSec, longDeg, longMin, longSec := coords.ConvertEquatorialCoordinatesToEcliptic(6.0, 7, 2009, 9.0, 34.0, 53.32, 19.0, 32.0, 6.01)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(longDeg-139.0) > tolerance || math.Abs(longMin-41.0) > tolerance || math.Abs(longSec-9.98) > tolerance &&
		math.Abs(latDeg-4.0) > tolerance || math.Abs(latMin-52.0) > tolerance || math.Abs(latSec-30.99) > tolerance {
		t.Fatalf(`Error while convert Equatorial Coordinates to Ecliptic. Required: %f %f %f, %f %f %f   Got: %f %f %f, %f %f %f`, 139.0, 41.0, 9.98, 4.0, 52.0, 30.99, longDeg, longMin, longSec, latDeg, latMin, latSec)
	}
}

func TestConvertEquatorialCoordinateToGalactic(t *testing.T) {
	lDeg, lMin, lSec, bDeg, bMin, bSec := coords.ConvertEquatorialCoordinateToGalactic(10.0, 21.0, 0.0, 10.0, 3.0, 11.00)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(lDeg-232.0) > tolerance || math.Abs(lMin-14.0) > tolerance || math.Abs(lSec-52.38) > tolerance &&
		math.Abs(bDeg-51.0) > tolerance || math.Abs(bMin-7.0) > tolerance || math.Abs(bSec-20.16) > tolerance {
		t.Fatalf(`Error while convert Equatorial Coordinate To Galactic. Required: %f %f %f, %f %f %f   Got: %f %f %f, %f %f %f`, 232.0, 14.0, 52.38, 51.0, 7.0, 20.16, lDeg, lMin, lSec, bDeg, bMin, bSec)
	}
}

func TestConvertGalacticCoordinateToEquatorial(t *testing.T) {
	lDeg, lMin, lSec, bDeg, bMin, bSec := coords.ConvertGalacticCoordinateToEquatorial(232.0, 14.0, 52.0, 51.0, 7.0, 20.00)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(lDeg-10.0) > tolerance || math.Abs(lMin-20.0) > tolerance || math.Abs(lSec-59.98) > tolerance &&
		math.Abs(bDeg-10.0) > tolerance || math.Abs(bMin-3.0) > tolerance || math.Abs(bSec-11.11) > tolerance {
		t.Fatalf(`Error while convert Galactic Coordinate To Equatorial. Required: %f %f %f, %f %f %f   Got: %f %f %f, %f %f %f`, 10.0, 20.0, 59.98, 10.0, 3.0, 11.11, lDeg, lMin, lSec, bDeg, bMin, bSec)
	}
}

func TestCalculateAngleBetweenTwoCelestialObjects(t *testing.T) {
	Deg, Min, Sec := coords.CalculateAngleBetweenTwoCelestialObjects(5.0, 13.0, 31.7, -8.0, 13.0, 30.0, 6.0, 44.0, 13.4, -16.0, 41.0, 11.0)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(Deg-23.0) > tolerance || math.Abs(Min-40.0) > tolerance || math.Abs(Sec-25.85) > tolerance {
		t.Fatalf(`Error while Calculating Angle Between Two Celestial Objects. Required: %f %f %f   Got: %f %f %f`, 23.0, 40.0, 25.85, Deg, Min, Sec)
	}
}

func TestCalculateRisingAndSettingTime(t *testing.T) {
	UTrHrs, UTrMin, UTrSec, UTsHrs, UTsMin, UTsSec := coords.CalculateRisingAndSettingTime(24.0, 8, 2010, 23.0, 39.0, 20.0, 21.0, 42.0, 0.0, 30.0, 64.0, 34)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(UTrHrs-14.0) > tolerance || math.Abs(UTrMin-16.0) > tolerance || math.Abs(UTrSec-18.02) > tolerance &&
		math.Abs(UTsHrs-4.0) > tolerance || math.Abs(UTsMin-10.0) > tolerance || math.Abs(UTsSec-1.15) > tolerance {
		t.Fatalf(`Error while Calculating Rising And Setting Time. Required: Rising = %f %f %f  Setting = %f %f %f   Got: Rising = %f %f %f  Setting = %f %f %f`, 14.0, 16.0, 18.02, 4.0, 10.0, 1.15, UTrHrs, UTrMin, UTrSec, UTsHrs, UTsMin, UTsSec)
	}
}

func TestCalculatePrecession(t *testing.T) {
	alpha1Hrs, alpha1Min, alpha1Sec, delta1Deg, delta1Min, delta1Sec := coords.CalculatePrecession(1979.5, 1950.0, 9.0, 10.0, 43.0, 14.0, 23.0, 25.0)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(alpha1Hrs-9.0) > tolerance || math.Abs(alpha1Min-12.0) > tolerance || math.Abs(alpha1Sec-20.47) > tolerance &&
		math.Abs(delta1Deg-14.0) > tolerance || math.Abs(delta1Min-16.0) > tolerance || math.Abs(delta1Sec-7.83) > tolerance {
		t.Fatalf(`Error while Calculating Precession. Required:  %f %f %f   %f %f %f   Got: Rising = %f %f %f  Setting = %f %f %f`, 9.0, 12.0, 20.47, 14.0, 16.0, 7.83, alpha1Hrs, alpha1Min, alpha1Sec, delta1Deg, delta1Min, delta1Sec)
	}
}

func TestCalculateNutation(t *testing.T) {
	nutationInLong, nutationInObliquity := coords.CalculateNutation(1.0, 9, 1988)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(nutationInLong-5.49) > tolerance || math.Abs(nutationInObliquity-9.24) > tolerance {
		t.Fatalf(`Error while Calculating Nutation. Required: %f  %f   Got: %f  %f`, 5.49, 9.24, nutationInLong, nutationInObliquity)
	}
}

func TestCalculateAberration(t *testing.T) {
	correctedLambdaDeg, correctedLambdaMin, correctedLambdaSec, correctedBetaDeg, correctedBetaMin, correctedBetaSec := coords.CalculateAberration(8.0, 9, 1988, 352.0, 37.0, 10.1, -1, 32, 56.4, 165.0, 33.0, 44.1)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(correctedLambdaDeg-352.0) > tolerance || math.Abs(correctedLambdaMin-37.0) > tolerance || math.Abs(correctedLambdaSec-30.45) > tolerance &&
		math.Abs(correctedBetaDeg-(-1.0)) > tolerance || math.Abs(correctedBetaMin-32.0) > tolerance || math.Abs(correctedBetaSec-56.33) > tolerance {
		t.Fatalf(`Error while Calculating Aberration. Required:  %f %f %f    %f %f %f   Got: %f %f %f    %f %f %f`, 352.0, 37.0, 30.45, -1.0, 32.0, 56.33, correctedLambdaDeg, correctedLambdaMin, correctedLambdaSec, correctedBetaDeg, correctedBetaMin, correctedBetaSec)
	}
}

func TestCalculateRefraction(t *testing.T) {
	HaHrs, HaMin, HaSec, DecDeg, DecMin, DecSec := coords.CalculateRefraction(5.0, 51.0, 44.0, 23.0, 13.0, 10.0, 52.0, 13.0, 1008.0)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(HaHrs-5.0) > tolerance || math.Abs(HaMin-51.0) > tolerance || math.Abs(HaSec-36.26) > tolerance &&
		math.Abs(DecDeg-23.0) > tolerance || math.Abs(DecMin-15.0) > tolerance || math.Abs(DecSec-13.91) > tolerance {
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
	raMoonHrs, raMoonMin, raMoonSec, decMoonDeg, decMoonMin, decMoonSec := coords.CalculateParallaxCorrections(26.0, 2, 1979, 16.0, 45.0, 0.0, 60.0, 100.0, 50.0, 22.0, 35.0, 19.0, -7.0, 41.0, 13.0, 1.0, 1.0, 9.0, 0.0)
	// Test data for sun and other planets
	raHrs, raMin, raSec, decDeg, decMin, decSec := coords.CalculateParallaxCorrections(26.0, 2, 1979, 16.0, 45.0, 0.0, 60.0, 100.0, 50.0, 22.0, 36.0, 44.0, -8.0, 44.0, 24.0, 0.0, 0.0, 0.0, 0.9901)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(raMoonHrs-22.0) > tolerance || math.Abs(raMoonMin-36.0) > tolerance || math.Abs(raMoonSec-43.21) > tolerance &&
		math.Abs(decMoonDeg-(-8.0)) > tolerance || math.Abs(decMoonMin-32.0) > tolerance || math.Abs(decMoonSec-17.39) > tolerance {
		t.Fatalf(`Error while Calculating Parallax Corrections for Moon. Required:  %f %f %f    %f %f %f   Got: %f %f %f    %f %f %f`, 22.0, 36.0, 43.21, -8.0, 32.0, 17.39, raMoonHrs, raMoonMin, raMoonSec, decMoonDeg, decMoonMin, decMoonSec)
	}

	if math.Abs(raHrs-22.0) > tolerance || math.Abs(raMin-36.0) > tolerance || math.Abs(raSec-44.00) > tolerance &&
		math.Abs(decDeg-(-8.0)) > tolerance || math.Abs(decMin-44.0) > tolerance || math.Abs(decSec-31.43) > tolerance {
		t.Fatalf(`Error while Calculating Parallax Corrections for sun and other planets. Required:  %f %f %f    %f %f %f   Got: %f %f %f    %f %f %f`, 22.0, 36.0, 44.00, -8.0, 44.0, 31.43, raHrs, raMin, raSec, decDeg, decMin, decSec)
	}
}

func TestCalculateHeliographicCoordinates(t *testing.T) {
	longitude, latitude := coords.CalculateHeliographicCoordinates(1.0, 5.0, 1988.0, 0, 0, 0, 40.0, 50.0, 37.0, 220.0, 10.5, 0, 15.0, 52.0)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(longitude-(-19.94)) > tolerance || math.Abs(latitude-143.27) > tolerance {
		t.Fatalf(`Error while Calculating Heliographic Coordinates. Required:  %f %f Got: %f %f`, -19.94, 143.27, longitude, latitude)
	}
}

func TestCalculateCarringtonRotationNumbers(t *testing.T) {
	CRN := coords.CalculateCarringtonRotationNumbers(27.0, 1, 1975)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(CRN-1624.0) > tolerance {
		t.Fatalf(`Error while Calculating Carrington Rotation Numbers. Required:  %f Got: %f`, 1624.0, CRN)
	}
}

func TestCalculateSelenographicCoordinatesOfMoon(t *testing.T) {
	le, be, C := coords.CalculateSelenographicCoordinatesOfMoon(1.0, 5, 1988, 209.12, -3.08, 23.4433)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(le-(-4.88)) > tolerance || math.Abs(be-4.04) > tolerance || math.Abs(C-19.78) > tolerance {
		t.Fatalf(`Error while Calculating Selenographic Coordinates of Moon. Required: le : %f\tbe : %f\tC : %f Got: le : %f\tbe : %f\tC : %f`, -4.88, 4.04, 19.78, le, be, C)
	}
}

func TestCalculateSelenographicCoordinatesOfSun(t *testing.T) {
	ls, bs, colongitude := coords.CalculateSelenographicCoordinatesOfSun(1.0, 5, 1988, 0, 0, 0, 209.12, -3.08, 23.4433, 55.952, 1.0076, 40.8437)
	const tolerance = 0.01 // Define an acceptable error range

	if math.Abs(ls-(6.81)) > tolerance || math.Abs(bs-1.18) > tolerance || math.Abs(colongitude-83.18) > tolerance {
		t.Fatalf(`Error while Calculating Selenographic Coordinates of Sun. Required: ls : %f\tbs : %f\tcolongitude : %f Got: ls : %f\tbs : %f\tcolongitude : %f`, 6.81, 1.18, 19.78, ls, bs, colongitude)
	}
}
