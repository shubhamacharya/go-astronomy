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
	decimalDeg := math.Abs(math.Round(deg)) + (min / 60) + (sec / 3600)

	if deg < 0 {
		return -decimalDeg
	} else {
		return decimalDeg
	}
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

func ConvertEquatorialToHorizonCoordinates(raHours float64, raMinutes float64, raSeconds float64, decDegrees float64, decMinutes float64, decSeconds float64, latitude float64) (altitudeDeg float64, altitudeMin float64, altitudeSec float64, azimuthDeg float64, azimuthMin float64, azimuthSec float64) {
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

func ConvertEquatorialCoordinatesToEcliptic(Gday float64, GMonth, GYear, raHrs, raMin int, raSec, decDeg, decMin, decSec float64) (eclipticLongDeg, eclipticLongMin, eclipticLongSec, eclipticLatDeg, eclipticLatMin, eclipticLatSec float64) {
	raDecimalDeg := ConvertDecimalHrsToDecimalDegress(datetime.ConvertHrsMinSecToDecimalHrs(raHrs, raMin, raSec, false, ""))
	decDecimalDeg := ConvertDegMinSecToDecimalDeg(decDeg, decMin, decSec)
	_, _, _, meanObliquity := CalculateEclipticMeanObliquity(Gday, GMonth, GYear)
	// fmt.Printf("meanObliquity : %f\n", meanObliquity)

	latDecimal := ConvertRadianceToDegree(math.Asin((math.Sin(ConvertDegreesToRadiance(decDecimalDeg)) * math.Cos(ConvertDegreesToRadiance(meanObliquity))) - (math.Cos(ConvertDegreesToRadiance(decDecimalDeg)) * math.Sin(ConvertDegreesToRadiance(meanObliquity)) * math.Sin(ConvertDegreesToRadiance(raDecimalDeg)))))
	// fmt.Printf("sineB : %f\n", ConvertRadianceToDegree(sineB))

	y := (math.Sin(ConvertDegreesToRadiance(raDecimalDeg)) * math.Cos(ConvertDegreesToRadiance(meanObliquity))) + (math.Tan(ConvertDegreesToRadiance(decDecimalDeg)) * math.Sin(ConvertDegreesToRadiance(meanObliquity)))
	x := math.Cos(ConvertDegreesToRadiance(raDecimalDeg))
	longDecimal := ConvertRadianceToDegree(math.Atan2(y, x))

	latDeg, latMin, latSec := ConvertDecimalDegToDegMinSec(latDecimal)
	longDeg, longMin, longSec := ConvertDecimalDegToDegMinSec(longDecimal)

	return latDeg, latMin, latSec, longDeg, longMin, longSec
}

func ConvertEquatorialCoordinateToGalactic(raHrs, raMin int, raSec, decDeg, decMin, decSec float64) (float64, float64, float64, float64, float64, float64) {
	raDecimalDeg := ConvertDecimalHrsToDecimalDegress(datetime.ConvertHrsMinSecToDecimalHrs(raHrs, raMin, raSec, false, ""))
	decDecimalDeg := ConvertDegMinSecToDecimalDeg(decDeg, decMin, decSec)

	b := ConvertRadianceToDegree(math.Asin((math.Cos(ConvertDegreesToRadiance(decDecimalDeg)) * math.Cos(ConvertDegreesToRadiance(27.4)) * math.Cos(ConvertDegreesToRadiance(raDecimalDeg)-ConvertDegreesToRadiance(192.25))) + (math.Sin(ConvertDegreesToRadiance(decDecimalDeg)) * math.Sin(ConvertDegreesToRadiance(27.4)))))

	y := math.Sin(ConvertDegreesToRadiance(decDecimalDeg)) - math.Sin(ConvertDegreesToRadiance(b))*math.Sin(ConvertDegreesToRadiance(27.4))
	x := math.Cos(ConvertDegreesToRadiance(decDecimalDeg)) * math.Sin(ConvertDegreesToRadiance(raDecimalDeg)-ConvertDegreesToRadiance(192.25)) * math.Cos(ConvertDegreesToRadiance(27.4))
	l := ConvertRadianceToDegree(math.Atan2(y, x)) + 33.0

	if l < 0 {
		l += 360
	}
	if l >= 360 {
		l -= 360
	}

	lDeg, lMin, lSec := ConvertDecimalDegToDegMinSec(l)
	bDeg, bMin, bSec := ConvertDecimalDegToDegMinSec(b)

	return lDeg, lMin, lSec, bDeg, bMin, bSec
}

func ConvertGalacticCoordinateToEquatorial(lHrs, lMin, lSec, bDeg, bMin, bSec float64) (float64, float64, float64, float64, float64, float64) {
	lDecimalDeg := ConvertDegMinSecToDecimalDeg(lHrs, lMin, lSec)
	bDecimalDeg := ConvertDegMinSecToDecimalDeg(bDeg, bMin, bSec)
	decDecimalDeg := ConvertRadianceToDegree(math.Asin((math.Cos(ConvertDegreesToRadiance(bDecimalDeg)) * math.Cos(ConvertDegreesToRadiance(27.4)) * math.Sin(ConvertDegreesToRadiance(lDecimalDeg)-ConvertDegreesToRadiance(33.0))) + (math.Sin(ConvertDegreesToRadiance(bDecimalDeg)) * math.Sin(ConvertDegreesToRadiance(27.4)))))

	y := math.Cos(ConvertDegreesToRadiance(bDecimalDeg)) * math.Cos(ConvertDegreesToRadiance(lDecimalDeg-33.0))
	x := (math.Sin(ConvertDegreesToRadiance(bDecimalDeg)) * math.Cos(ConvertDegreesToRadiance(27.4))) - (math.Cos(ConvertDegreesToRadiance(bDecimalDeg)) * math.Sin(ConvertDegreesToRadiance(27.4)) * math.Sin(ConvertDegreesToRadiance(lDecimalDeg-33.0)))
	raDecimalDeg := ConvertRadianceToDegree(math.Atan2(y, x)) + 192.25

	if raDecimalDeg < 0 {
		raDecimalDeg += 360
	}
	if raDecimalDeg >= 360 {
		raDecimalDeg -= 360
	}

	raHrs, raMin, raSec := datetime.ConvertDecimalHrsToHrsMinSec(ConvertDecimalDegressToDecimalHrs(raDecimalDeg))
	decDeg, decMin, decSec := ConvertDecimalDegToDegMinSec(decDecimalDeg)

	return raHrs, raMin, raSec, decDeg, decMin, decSec
}

func CalculateAngleBetweenTwoCelestialObjects(p1RAHrs, p1RAMin int, p1RASec, p1DecDeg, p1DecMin, p1DecSec float64, p2RAHrs, p2RAMin int, p2RASec, p2DecDeg, p2DecMin, p2DecSec float64) (float64, float64, float64) {
	p1RADecimalHrs := datetime.ConvertHrsMinSecToDecimalHrs(p1RAHrs, p1RAMin, p1RASec, false, "")
	p1DecDecimalDeg := ConvertDegMinSecToDecimalDeg(p1DecDeg, p1DecMin, p1DecSec)
	p2RADecimalHrs := datetime.ConvertHrsMinSecToDecimalHrs(p2RAHrs, p2RAMin, p2RASec, false, "")
	p2DecDecimalDeg := ConvertDegMinSecToDecimalDeg(p2DecDeg, p2DecMin, p2DecSec)

	RADiffInDegress := ConvertDecimalHrsToDecimalDegress(p1RADecimalHrs - p2RADecimalHrs)

	angle := ConvertRadianceToDegree(math.Acos((math.Sin(ConvertDegreesToRadiance(p1DecDecimalDeg)) * math.Sin(ConvertDegreesToRadiance(p2DecDecimalDeg))) + (math.Cos(ConvertDegreesToRadiance(p1DecDecimalDeg)) * math.Cos(ConvertDegreesToRadiance(p2DecDecimalDeg)) * math.Cos(ConvertDegreesToRadiance(RADiffInDegress)))))
	angleDeg, angleMin, angleSec := ConvertDecimalDegToDegMinSec(angle)
	return angleDeg, angleMin, angleSec
}

func CalculateRisingAndSettingTime(Gday float64, Gmonth, Gyear float64, raHrs, raMin int, raSec, decDeg, decMin, decSec, geoLatN, geoLongE, refractionInArcMin float64) (float64, float64, float64, float64, float64, float64) {
	decimalRAHrs := datetime.ConvertHrsMinSecToDecimalHrs(int(raHrs), int(raMin), raSec, false, "")
	decimalDECDeg := ConvertDegMinSecToDecimalDeg(decDeg, decMin, decSec)
	refractionDeg := refractionInArcMin / 60 // Converted refraction from arcmin to degress

	cosH := -((math.Sin(ConvertDegreesToRadiance(refractionDeg)) + (math.Sin(ConvertDegreesToRadiance(geoLatN)) * math.Sin(ConvertDegreesToRadiance(decimalDECDeg)))) / (math.Cos(ConvertDegreesToRadiance(geoLatN)) * math.Cos(ConvertDegreesToRadiance(decimalDECDeg))))
	H := 0.0
	if cosH > -1 && cosH < +1 {
		H = ConvertDecimalDegressToDecimalHrs(ConvertRadianceToDegree(math.Acos(cosH)))
	}

	LSTr := decimalRAHrs - H
	LSTs := decimalRAHrs + H

	for LSTr < 0 || LSTr > 24 {
		if LSTr < 0 {
			LSTr += 24
		}

		if LSTr > 24 {
			LSTr -= 24
		}
	}

	for LSTs < 0 || LSTs > 24 {
		if LSTs < 0 {
			LSTs += 24
		}

		if LSTs > 24 {
			LSTs -= 24
		}
	}

	Ar := ConvertRadianceToDegree(math.Acos((math.Sin(ConvertDegreesToRadiance(decimalDECDeg)) + (math.Sin(ConvertDegreesToRadiance(refractionDeg)) * math.Sin(ConvertDegreesToRadiance(geoLatN)))) / (math.Cos(ConvertDegreesToRadiance(refractionDeg)) * math.Cos(ConvertDegreesToRadiance(geoLatN)))))

	for Ar < 0 || Ar > 360 {
		if Ar < 0 {
			Ar += 360
		}

		if Ar > 360 {
			Ar -= 360
		}
	}
	// As := 360 - Ar
	// Convert E longitude to west
	geoLongW := -(360 - geoLongE)
	rHrs, rMin, rSec := datetime.ConvertDecimalHrsToHrsMinSec(LSTr)
	sHrs, sMin, sSec := datetime.ConvertDecimalHrsToHrsMinSec(LSTs)
	GSTrHrs, GSTrMin, GSTrSec, _ := datetime.CalculateGreenwichSideralTimeUsingLocalSiderealTime(int(rHrs), int(rMin), rSec, geoLongW)
	GSTsHrs, GSTsMin, GSTsSec, _ := datetime.CalculateGreenwichSideralTimeUsingLocalSiderealTime(int(sHrs), int(sMin), sSec, geoLongW)
	UTrHrs, UTrMin, UTrSec := datetime.ConvertGreenwichSiderealTimeToUniversalTime(Gday, int(Gmonth), int(Gyear), int(GSTrHrs), int(GSTrMin), GSTrSec)
	UTsHrs, UTsMin, UTsSec := datetime.ConvertGreenwichSiderealTimeToUniversalTime(Gday, int(Gmonth), int(Gyear), int(GSTsHrs), int(GSTsMin), GSTsSec)

	return UTrHrs, UTrMin, UTrSec, UTsHrs, UTsMin, UTsSec
}

func CalculatePrecession(n1, n2 float64, alphaHrs, alphaMin int, alphaSec, deltaDeg, deltaMin, deltaSec float64) (float64, float64, float64, float64, float64, float64) {
	decimalHrs := datetime.ConvertHrsMinSecToDecimalHrs(alphaHrs, alphaMin, alphaSec, false, "")
	decimalHrsTodeg := ConvertDecimalHrsToDecimalDegress(decimalHrs)
	decimalDeg := ConvertDegMinSecToDecimalDeg(deltaDeg, deltaMin, deltaSec)

	S1Hrs := ((3.07327 + (1.33617 * math.Sin(ConvertDegreesToRadiance(decimalHrsTodeg)) * math.Tan(ConvertDegreesToRadiance(decimalDeg)))) * (n1 - n2)) / 3600 // Convert to Hrs
	S2Deg := ((20.0426 * math.Cos(ConvertDegreesToRadiance(decimalHrsTodeg))) * (n1 - n2)) / 3600

	alpha1Hrs, alpha1Min, alpha1Sec := datetime.ConvertDecimalHrsToHrsMinSec(decimalHrs + S1Hrs)
	delta1Deg, delta1Min, delta1Sec := ConvertDecimalDegToDegMinSec(decimalDeg + S2Deg)

	return alpha1Hrs, alpha1Min, alpha1Sec, delta1Deg, delta1Min, delta1Sec
}

func CalculateNutation(day float64, month, year int) (float64, float64) {
	julianDate := datetime.ConvertGreenwichDateToJulianDate(day, month, year)
	T := (julianDate - 2415020.0) / 36525.0
	A := 100.002136 * T
	L := 279.6967 + 360.0*(A-math.Trunc(A)) //Sun Mean Longitude

	for L < 0 || L > 360 {
		if L < 0 {
			L += 360
		}

		if L > 360 {
			L -= 360
		}
	}

	B := 5.372617 * T

	moonNode := 259.1833 - 360.0*(B-math.Trunc(B))

	for moonNode < 0 || moonNode > 360 {
		if moonNode < 0 {
			moonNode += 360
		}

		if moonNode > 360 {
			moonNode -= 360
		}
	}

	nutationInLong := -(17.2 * math.Sin(ConvertDegreesToRadiance(moonNode))) - (1.3 * math.Sin(ConvertDegreesToRadiance(2*L)))
	nutationInObliquity := (9.2 * math.Cos(ConvertDegreesToRadiance(moonNode))) + (0.5 * math.Cos(ConvertDegreesToRadiance(2*L)))

	return nutationInLong, nutationInObliquity
}

func CalculateAberration(day float64, month, year int, trueLambdaDeg, trueLambdaMin, trueLambdaSec, trueBetaDeg, trueBetaMin, trueBetaSec, longDeg, longMin, longSec float64) (float64, float64, float64, float64, float64, float64) {
	trueLambdaDecimalDeg := ConvertDegMinSecToDecimalDeg(trueLambdaDeg, trueLambdaMin, trueLambdaSec)
	trueBetaDecimalDeg := ConvertDegMinSecToDecimalDeg(trueBetaDeg, trueBetaMin, trueBetaSec)
	longDecimalDeg := ConvertDegMinSecToDecimalDeg(longDeg, longMin, longSec)

	trueLambdaDecimalDeg += (-20.5 * math.Cos(ConvertDegreesToRadiance(longDecimalDeg-trueLambdaDecimalDeg)) / math.Cos(ConvertDegreesToRadiance(trueBetaDecimalDeg))) / 3600
	trueBetaDecimalDeg += (-20.5 * math.Sin(ConvertDegreesToRadiance(longDecimalDeg-trueLambdaDecimalDeg)) * math.Sin(ConvertDegreesToRadiance(trueBetaDecimalDeg))) / 3600

	correctedLambdaDeg, correctedLambdaMin, correctedLambdaSec := ConvertDecimalDegToDegMinSec(trueLambdaDecimalDeg)
	correctedBetaDeg, correctedBetaMin, correctedBetaSec := ConvertDecimalDegToDegMinSec(trueBetaDecimalDeg)

	return correctedLambdaDeg, correctedLambdaMin, correctedLambdaSec, correctedBetaDeg, math.Abs(correctedBetaMin), math.Abs(correctedBetaSec)
}

func CalculateRefraction(trueHAHr, trueHAMin, trueHASec, trueDecDeg, trueDecMin, trueDecSec, geoLat, temp, pressure float64) (float64, float64, float64, float64, float64, float64) {
	altitudeDeg, altitudeMin, altitudeSec, azimuthDeg, azimuthMin, azimuthSec := ConvertEquatorialToHorizonCoordinates(trueHAHr, trueHAMin, trueHASec, trueDecDeg, trueDecMin, trueDecSec, geoLat)
	altitudeDecimalDeg := ConvertDegMinSecToDecimalDeg(math.Round(altitudeDeg), altitudeMin, altitudeSec)
	R := 0.0
	z := 90 - altitudeDecimalDeg
	if altitudeDecimalDeg > 15.0 {
		R = (0.00452 * pressure * math.Tan(ConvertDegreesToRadiance(z))) / (273 + temp)
	}

	apperentAltDeg, apperentAltMin, apperentAltSec := ConvertDecimalDegToDegMinSec(R + altitudeDecimalDeg)

	HaHrs, HaMin, HaSec, DecDeg, DecMin, DecSec := ConvertHorizonCoordinatesToEquatorial(0.0, 0, 0, apperentAltDeg, apperentAltMin, apperentAltSec, azimuthDeg, azimuthMin, azimuthSec, geoLat)

	return HaHrs, HaMin, HaSec, DecDeg, DecMin, DecSec
}
