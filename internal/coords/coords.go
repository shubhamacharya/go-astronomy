package coords

import (
	datetime "go-astronomy/internal/dateTime"
	"math"
)

func ConvertDecimalDegToDegMinSec(decimalDeg float64) (float64, float64, float64) {
	Deg, fractPart := math.Modf(decimalDeg)
	Mins, minFracts := math.Modf(fractPart * 60)
	Sec := minFracts * 60
	return Deg, Mins, Sec
}

func ConvertDegMinSecToDecimalDeg(deg, min, sec float64) float64 {
	return float64(deg + ((min + (sec / 60)) / 60))
}

func ConvertDecimalHrsToDecimalDegress(decimalHrs float64) float64 {
	const degreesPerHour = 15.0
	return decimalHrs * degreesPerHour
}

func ConvertDecimalDegressToDecimalHrs(decimalDeg float64) float64 {
	const degreesPerHour = 15.0
	return decimalDeg / degreesPerHour
}

func ConverRightAscensionToHourAngle(localDay int, localMonth int, localYear int, localHrs int, localMin int, localSec, raHrs, raMin, raSec, geoLong float64, daylightsavingHrs int, daylightsavingMin int, zoneOffset float64) (float64, float64, float64, float64) {
	GDay, GMonth, GYear, GHrs, GMin, GSec, _ := datetime.ConvertLocalTimeToUniversalTime(localDay, localMonth, localYear, localHrs, localMin, localSec, 0, 0, -4)
	GSTHrs, GSTMin, GSTSec, _ := datetime.ConvertUniversalTimeToGreenwichSiderealTime(GDay, int(GMonth), int(GYear), int(GHrs), int(GMin), GSec)
	_, _, _, LSTDecimalTime := datetime.CalculateLocalSiderealTimeUsingGreenwichSideralTime(int(GSTHrs), int(GSTMin), GSTSec, geoLong)
	decimalRA := datetime.ConvertHrsMinSecToDecimalHrs(int(raHrs), int(raMin), raSec, false, "")
	decimalHourAngle := LSTDecimalTime - decimalRA
	if decimalHourAngle < 0 {
		decimalHourAngle += 24
	}
	haHrs, haMin, haSec := datetime.ConvertDecimalHrsToHrsMinSec(decimalHourAngle)
	return haHrs, haMin, haSec, decimalHourAngle
}

func ConverHourAngleToRightAscension(localDay int, localMonth int, localYear int, localHrs int, localMin int, localSec, haHrs, haMin, haSec, geoLong float64, daylightsavingHrs int, daylightsavingMin int, zoneOffset float64) (float64, float64, float64, float64) {
	GDay, GMonth, GYear, GHrs, GMin, GSec, _ := datetime.ConvertLocalTimeToUniversalTime(localDay, localMonth, localYear, localHrs, localMin, localSec, 0, 0, -4)
	GSTHrs, GSTMin, GSTSec, _ := datetime.ConvertUniversalTimeToGreenwichSiderealTime(GDay, int(GMonth), int(GYear), int(GHrs), int(GMin), GSec)
	_, _, _, LSTDecimalTime := datetime.CalculateLocalSiderealTimeUsingGreenwichSideralTime(int(GSTHrs), int(GSTMin), GSTSec, geoLong)
	decimalRA := datetime.ConvertHrsMinSecToDecimalHrs(int(haHrs), int(haMin), haSec, false, "")
	decimalRightAscension := LSTDecimalTime - decimalRA
	if decimalRightAscension < 0 {
		decimalRightAscension += 24
	}
	raHrs, raMin, raSec := datetime.ConvertDecimalHrsToHrsMinSec(decimalRightAscension)
	return raHrs, raMin, raSec, decimalRightAscension
}

func ConvertEquatorialToHorizonCoordinates(
	raHours float64, raMinutes float64, raSeconds float64, // Right Ascension in hours, minutes, seconds
	decDegrees float64, decMinutes float64, decSeconds float64, // Declination in degrees, minutes, seconds
	latitude float64, // Observer's latitude in degrees
) (
	altitudeDeg float64, altitudeMin float64, altitudeSec float64, // Altitude in degrees, minutes, seconds
	azimuthDeg float64, azimuthMin float64, azimuthSec float64, // Azimuth in degrees, minutes, seconds
) {
	// Convert Right Ascension (RA) to decimal hours
	decimalRAHours := datetime.ConvertHrsMinSecToDecimalHrs(int(raHours), int(raMinutes), raSeconds, false, "")

	// Convert decimal Right Ascension hours to degrees
	decimalRADegrees := ConvertDecimalHrsToDecimalDegress(decimalRAHours)

	// Convert Declination to decimal degrees
	decimalDeclination := ConvertDegMinSecToDecimalDeg(decDegrees, decMinutes, decSeconds)

	// Convert observer's latitude to radians
	latitudeRad := latitude * math.Pi / 180.0

	// Calculate the altitude (altitude angle) in radians
	sineAlt := math.Asin(
		math.Sin(decimalDeclination*math.Pi/180.0)*math.Sin(latitudeRad) +
			math.Cos(decimalDeclination*math.Pi/180.0)*math.Cos(latitudeRad)*math.Cos(decimalRADegrees*math.Pi/180.0),
	)

	// Calculate the azimuth in radians
	cosAz := (math.Sin(decimalDeclination*math.Pi/180.0) - math.Sin(latitudeRad)*math.Sin(sineAlt)) /
		(math.Cos(latitudeRad) * math.Cos(sineAlt))
	sineAz := math.Sin(decimalRADegrees * math.Pi / 180.0)
	azimuthRad := math.Acos(cosAz)

	// Adjust azimuth for the correct quadrant
	if sineAz > 0 {
		azimuthRad = 2*math.Pi - azimuthRad
	}

	// Convert radians to degrees
	altitudeDeg = sineAlt * 180.0 / math.Pi
	azimuthDeg = azimuthRad * 180.0 / math.Pi

	// Convert decimal degrees to degrees, minutes, seconds
	_, altitudeMin, altitudeSec = ConvertDecimalDegToDegMinSec(altitudeDeg)
	_, azimuthMin, azimuthSec = ConvertDecimalDegToDegMinSec(azimuthDeg)

	return altitudeDeg, altitudeMin, altitudeSec, azimuthDeg, azimuthMin, azimuthSec
}
