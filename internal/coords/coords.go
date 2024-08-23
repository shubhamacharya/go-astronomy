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

func ConvertDegreesToRadiance(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

func ConvertRadianceToDegree(radians float64) float64 {
	return radians * (180 / math.Pi)
}

func ConvertHorizonCoordinatesToEquatorial(GSTHrs, GSTMin, GSec, altitudeDeg, altitudeMin, altitudeSec, azimuthDeg, azimuthMin, azimuthSec, latitude float64) (float64, float64, float64, float64, float64, float64) {
	altitudeDecimalDeg := ConvertDegMinSecToDecimalDeg(altitudeDeg, altitudeMin, altitudeSec)
	azimuthDecimalDeg := ConvertDegMinSecToDecimalDeg(azimuthDeg, azimuthMin, azimuthSec)

	declination := ConvertRadianceToDegree(math.Asin((math.Sin(ConvertDegreesToRadiance(altitudeDecimalDeg)) * math.Sin(ConvertDegreesToRadiance(latitude))) + (math.Cos(ConvertDegreesToRadiance(altitudeDecimalDeg)) * math.Cos(ConvertDegreesToRadiance(latitude)) * math.Cos(ConvertDegreesToRadiance(azimuthDecimalDeg)))))
	cosInvHourAngle := ConvertRadianceToDegree(math.Acos((math.Sin(ConvertDegreesToRadiance(altitudeDecimalDeg)) - (math.Sin(ConvertDegreesToRadiance(latitude)) * math.Sin(ConvertDegreesToRadiance(declination)))) / (math.Cos(ConvertDegreesToRadiance(latitude)) * math.Cos(ConvertDegreesToRadiance(declination)))))

	hourAngleInDecimalDeg := 0.0
	if ConvertRadianceToDegree(math.Sin(ConvertDegreesToRadiance(azimuthDecimalDeg))) < 0 {
		hourAngleInDecimalDeg = cosInvHourAngle
	} else {
		hourAngleInDecimalDeg = 360 - cosInvHourAngle
	}
	hourAngleInDecimalHrs := ConvertDecimalDegressToDecimalHrs(hourAngleInDecimalDeg)

	haHrs, haMin, haSec := datetime.ConvertDecimalHrsToHrsMinSec(hourAngleInDecimalHrs)
	decDeg, decMin, decSec := ConvertDecimalDegToDegMinSec(declination)

	return haHrs, haMin, haSec, decDeg, decMin, decSec
}

func CalculateEclipticMeanObliquity(Gday float64, GMonth, GYear int) (float64, float64, float64, float64) {
	julianDate := datetime.ConvertGreenwichDateToJulianDate(Gday, GMonth, GYear)
	timeElapsed := (julianDate - 2451545.0) / 36525.0
	meanObliquity := 23.439292 - (((46.815 * timeElapsed) + (0.0006 * math.Pow(timeElapsed, 2)) - (0.00181 * math.Pow(timeElapsed, 3))) / 3600)
	obliquityDeg, obliquityMin, obliquitySec := ConvertDecimalDegToDegMinSec(meanObliquity)

	return obliquityDeg, obliquityMin, obliquitySec, meanObliquity

}

func ConvertEclipticCoordinatesToEquatorial(day float64, month, year int, eclipticLongDeg, eclipticLongMin, eclipticLongSec, eclipticLatDeg, eclipticLatMin, eclipticLatSec float64) (float64, float64, float64, float64, float64, float64) {
	_, _, _, meanObliquity := CalculateEclipticMeanObliquity(day, month, year)
	eclipticLongDecimalDeg := ConvertDegMinSecToDecimalDeg(eclipticLongDeg, eclipticLongMin, eclipticLongSec)

	eclipticLatDecimalDeg := ConvertDegMinSecToDecimalDeg(eclipticLatDeg, eclipticLatMin, eclipticLatSec)

	decDecimalDeg := ConvertRadianceToDegree(math.Asin((math.Sin(ConvertDegreesToRadiance(eclipticLatDecimalDeg)) * math.Cos(ConvertDegreesToRadiance(meanObliquity))) + (math.Cos(ConvertDegreesToRadiance(eclipticLatDecimalDeg)) * math.Sin(ConvertDegreesToRadiance(meanObliquity)) * math.Sin(ConvertDegreesToRadiance(eclipticLongDecimalDeg)))))

	y := (math.Sin(ConvertDegreesToRadiance(eclipticLongDecimalDeg)) * math.Cos(ConvertDegreesToRadiance(meanObliquity))) - (math.Tan(ConvertDegreesToRadiance(eclipticLatDecimalDeg)) * math.Sin(ConvertDegreesToRadiance(meanObliquity)))

	x := math.Cos(ConvertDegreesToRadiance(eclipticLongDecimalDeg))

	raDeg := ConvertRadianceToDegree(math.Atan2(ConvertDegreesToRadiance(y), ConvertDegreesToRadiance(x)))

	raDecimalHrs := ConvertDecimalDegressToDecimalHrs(raDeg)

	decDeg, decMin, decSec := ConvertDecimalDegToDegMinSec(decDecimalDeg)
	raHrs, raMins, raSecs := datetime.ConvertDecimalHrsToHrsMinSec(raDecimalHrs)

	return raHrs, raMins, raSecs, decDeg, decMin, decSec
}
