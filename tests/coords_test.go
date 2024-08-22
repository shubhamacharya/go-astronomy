package tests

import (
	coords "go-astronomy/internal/coords"
	datetime "go-astronomy/internal/dateTime"
	"math"
	"testing"
)

func TestConvertDecimalDegToDegMinSec(t *testing.T) {
	Deg, Mins, Sec := coords.ConvertDecimalDegToDegMinSec(182.524167)
	if Deg != 182 || Mins != 31 || Sec != 27 {
		t.Fatalf(`Error while converting degrees to Deg Min Sec. Required: %f %f %f    Got: %f %f %f`, 182.0, 31.0, 27.0, Deg, Mins, Sec)
	}
}

func TestConvertDegMinSecToDecimalDeg(t *testing.T) {
	decimalDeg := coords.ConvertDegMinSecToDecimalDeg(182, 31, 27)
	var expectedDeg float64 = 182.524167
	const tolerance = 0.000001 // Define an acceptable error range

	if math.Abs(decimalDeg-expectedDeg) > tolerance {
		t.Fatalf(`Error while converting Deg Min Sec to degrees. Required: %f    Got: %f`, expectedDeg, decimalDeg)
	}
}

func TestConvertDecimalHrsToDecimalDegress(t *testing.T) {
	decimalDeg := coords.ConvertDecimalHrsToDecimalDegress(datetime.ConvertHrsMinSecToDecimalHrs(9, 36, 10.2, false, ""))
	deg, min, sec := coords.ConvertDecimalDegToDegMinSec(decimalDeg)

	if math.Abs(deg-144.0) > 0.001 || math.Abs(min-2.0) > 0.001 || math.Abs(sec-33.0) > 0.001 {
		t.Fatalf(`Error while converting Decimal Hrs To Decimal Degress. Required: %f %f %f   Got: %f %f %f`, 144.0, 2.0, 33.0, deg, min, sec)
	}
}

func TestConvertDecimalDegressToDecimalHrs(t *testing.T) {
	decimalHrs := coords.ConvertDecimalDegressToDecimalHrs(coords.ConvertDegMinSecToDecimalDeg(144.0, 2.0, 33.0))
	hrs, min, sec := datetime.ConvertDecimalHrsToHrsMinSec(decimalHrs)
	const tolerance = 0.000001 // Define an acceptable error range

	if math.Abs(hrs-9.0) > tolerance || math.Abs(min-36.0) > tolerance || math.Abs(sec-10.2) > tolerance {
		t.Fatalf(`Error while converting Decimal Deg To Hrs Min Sec. Required: %f %f %f   Got: %f %f %f`, 9.0, 36.0, 10.2, hrs, min, sec)
	}
}

func TestConverRightAscensionToHourAngle(t *testing.T) {
	haHrs, haMin, haSec, _ := coords.ConverRightAscensionToHourAngle(22, 4, 1980, 14, 36, 51.67, 18.0, 32.0, 21.0, -64, 0, 0, -4)
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
