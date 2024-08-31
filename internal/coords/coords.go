package coords

import (
	datetime "go-astronomy/internal/dateTime"
	"go-astronomy/internal/macros"
	"math"
)

func ConvertDecimalHrsToDecimalDegress(decimalHrs float64) float64 {
	const degreesPerHour = 15.0
	return decimalHrs * degreesPerHour
}

func ConverRightAscensionToHourAngle(localDay int, localMonth int, localYear int, localHrs int, localMin int, localSec, raHrs, raMin, raSec, geoLong float64, daylightsavingHrs int, daylightsavingMin int, zoneOffset float64, adjustRange bool) (float64, float64, float64, float64) {
	GDay, GMonth, GYear, GHrs, GMin, GSec, _ := datetime.ConvertLocalTimeToUniversalTime(localDay, localMonth, localYear, localHrs, localMin, localSec, daylightsavingHrs, daylightsavingMin, zoneOffset)
	GSTHrs, GSTMin, GSTSec, _ := datetime.ConvertUniversalTimeToGreenwichSiderealTime(GDay, int(GMonth), int(GYear), int(GHrs), int(GMin), GSec)
	_, _, _, LSTDecimalTime := datetime.CalculateLocalSiderealTimeUsingGreenwichSideralTime(int(GSTHrs), int(GSTMin), GSTSec, geoLong)
	decimalRA := datetime.ConvertHrsMinSecToDecimalHrs(int(raHrs), int(raMin), raSec, false, "")
	decimalHourAngle := LSTDecimalTime - decimalRA

	if adjustRange {
		if decimalHourAngle < 0 {
			decimalHourAngle += 24
		}
	}

	haHrs, haMin, haSec := datetime.ConvertDecimalHrsToHrsMinSec(decimalHourAngle)
	return haHrs, haMin, haSec, decimalHourAngle
}

func ConverHourAngleToRightAscension(localDay int, localMonth int, localYear int, localHrs int, localMin int, localSec, haHrs, haMin, haSec, geoLong float64, daylightsavingHrs int, daylightsavingMin int, zoneOffset float64) (float64, float64, float64, float64) {
	GDay, GMonth, GYear, GHrs, GMin, GSec, _ := datetime.ConvertLocalTimeToUniversalTime(localDay, localMonth, localYear, localHrs, localMin, localSec, daylightsavingHrs, daylightsavingMin, zoneOffset)
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
	decimalDeclination := macros.ConvertDegMinSecToDecimalDeg(decDegrees, decMinutes, decSeconds)

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
	_, altitudeMin, altitudeSec = macros.ConvertDecimalDegToDegMinSec(altitudeDeg)
	_, azimuthMin, azimuthSec = macros.ConvertDecimalDegToDegMinSec(azimuthDeg)

	return altitudeDeg, altitudeMin, altitudeSec, azimuthDeg, azimuthMin, azimuthSec
}

func ConvertHorizonCoordinatesToEquatorial(GSTHrs, GSTMin, GSec, altitudeDeg, altitudeMin, altitudeSec, azimuthDeg, azimuthMin, azimuthSec, latitude float64) (float64, float64, float64, float64, float64, float64) {
	altitudeDecimalDeg := macros.ConvertDegMinSecToDecimalDeg(altitudeDeg, altitudeMin, altitudeSec)
	azimuthDecimalDeg := macros.ConvertDegMinSecToDecimalDeg(azimuthDeg, azimuthMin, azimuthSec)

	declination := macros.ConvertRadianceToDegree(math.Asin((math.Sin(macros.ConvertDegreesToRadiance(altitudeDecimalDeg)) * math.Sin(macros.ConvertDegreesToRadiance(latitude))) + (math.Cos(macros.ConvertDegreesToRadiance(altitudeDecimalDeg)) * math.Cos(macros.ConvertDegreesToRadiance(latitude)) * math.Cos(macros.ConvertDegreesToRadiance(azimuthDecimalDeg)))))
	cosInvHourAngle := macros.ConvertRadianceToDegree(math.Acos((math.Sin(macros.ConvertDegreesToRadiance(altitudeDecimalDeg)) - (math.Sin(macros.ConvertDegreesToRadiance(latitude)) * math.Sin(macros.ConvertDegreesToRadiance(declination)))) / (math.Cos(macros.ConvertDegreesToRadiance(latitude)) * math.Cos(macros.ConvertDegreesToRadiance(declination)))))

	hourAngleInDecimalDeg := 0.0
	if macros.ConvertRadianceToDegree(math.Sin(macros.ConvertDegreesToRadiance(azimuthDecimalDeg))) < 0 {
		hourAngleInDecimalDeg = cosInvHourAngle
	} else {
		hourAngleInDecimalDeg = 360 - cosInvHourAngle
	}
	hourAngleInDecimalHrs := macros.ConvertDecimalDegressToDecimalHrs(hourAngleInDecimalDeg)

	haHrs, haMin, haSec := datetime.ConvertDecimalHrsToHrsMinSec(hourAngleInDecimalHrs)
	decDeg, decMin, decSec := macros.ConvertDecimalDegToDegMinSec(declination)

	return haHrs, haMin, haSec, decDeg, decMin, decSec
}

func ConvertEquatorialCoordinatesToEcliptic(Gday float64, GMonth, GYear, raHrs, raMin int, raSec, decDeg, decMin, decSec float64) (eclipticLongDeg, eclipticLongMin, eclipticLongSec, eclipticLatDeg, eclipticLatMin, eclipticLatSec float64) {
	raDecimalDeg := ConvertDecimalHrsToDecimalDegress(datetime.ConvertHrsMinSecToDecimalHrs(raHrs, raMin, raSec, false, ""))
	decDecimalDeg := macros.ConvertDegMinSecToDecimalDeg(decDeg, decMin, decSec)
	_, _, _, meanObliquity := macros.CalculateEclipticMeanObliquity(Gday, GMonth, GYear)
	// fmt.Printf("meanObliquity : %f\n", meanObliquity)

	latDecimal := macros.ConvertRadianceToDegree(math.Asin((math.Sin(macros.ConvertDegreesToRadiance(decDecimalDeg)) * math.Cos(macros.ConvertDegreesToRadiance(meanObliquity))) - (math.Cos(macros.ConvertDegreesToRadiance(decDecimalDeg)) * math.Sin(macros.ConvertDegreesToRadiance(meanObliquity)) * math.Sin(macros.ConvertDegreesToRadiance(raDecimalDeg)))))
	// fmt.Printf("sineB : %f\n", macros.ConvertRadianceToDegree(sineB))

	y := (math.Sin(macros.ConvertDegreesToRadiance(raDecimalDeg)) * math.Cos(macros.ConvertDegreesToRadiance(meanObliquity))) + (math.Tan(macros.ConvertDegreesToRadiance(decDecimalDeg)) * math.Sin(macros.ConvertDegreesToRadiance(meanObliquity)))
	x := math.Cos(macros.ConvertDegreesToRadiance(raDecimalDeg))
	longDecimal := macros.ConvertRadianceToDegree(math.Atan2(y, x))

	latDeg, latMin, latSec := macros.ConvertDecimalDegToDegMinSec(latDecimal)
	longDeg, longMin, longSec := macros.ConvertDecimalDegToDegMinSec(longDecimal)

	return latDeg, latMin, latSec, longDeg, longMin, longSec
}

func ConvertEquatorialCoordinateToGalactic(raHrs, raMin int, raSec, decDeg, decMin, decSec float64) (float64, float64, float64, float64, float64, float64) {
	raDecimalDeg := ConvertDecimalHrsToDecimalDegress(datetime.ConvertHrsMinSecToDecimalHrs(raHrs, raMin, raSec, false, ""))
	decDecimalDeg := macros.ConvertDegMinSecToDecimalDeg(decDeg, decMin, decSec)

	b := macros.ConvertRadianceToDegree(math.Asin((math.Cos(macros.ConvertDegreesToRadiance(decDecimalDeg)) * math.Cos(macros.ConvertDegreesToRadiance(27.4)) * math.Cos(macros.ConvertDegreesToRadiance(raDecimalDeg)-macros.ConvertDegreesToRadiance(192.25))) + (math.Sin(macros.ConvertDegreesToRadiance(decDecimalDeg)) * math.Sin(macros.ConvertDegreesToRadiance(27.4)))))

	y := math.Sin(macros.ConvertDegreesToRadiance(decDecimalDeg)) - math.Sin(macros.ConvertDegreesToRadiance(b))*math.Sin(macros.ConvertDegreesToRadiance(27.4))
	x := math.Cos(macros.ConvertDegreesToRadiance(decDecimalDeg)) * math.Sin(macros.ConvertDegreesToRadiance(raDecimalDeg)-macros.ConvertDegreesToRadiance(192.25)) * math.Cos(macros.ConvertDegreesToRadiance(27.4))
	l := macros.ConvertRadianceToDegree(math.Atan2(y, x)) + 33.0

	if l < 0 {
		l += 360
	}
	if l >= 360 {
		l -= 360
	}

	lDeg, lMin, lSec := macros.ConvertDecimalDegToDegMinSec(l)
	bDeg, bMin, bSec := macros.ConvertDecimalDegToDegMinSec(b)

	return lDeg, lMin, lSec, bDeg, bMin, bSec
}

func ConvertGalacticCoordinateToEquatorial(lHrs, lMin, lSec, bDeg, bMin, bSec float64) (float64, float64, float64, float64, float64, float64) {
	lDecimalDeg := macros.ConvertDegMinSecToDecimalDeg(lHrs, lMin, lSec)
	bDecimalDeg := macros.ConvertDegMinSecToDecimalDeg(bDeg, bMin, bSec)
	decDecimalDeg := macros.ConvertRadianceToDegree(math.Asin((math.Cos(macros.ConvertDegreesToRadiance(bDecimalDeg)) * math.Cos(macros.ConvertDegreesToRadiance(27.4)) * math.Sin(macros.ConvertDegreesToRadiance(lDecimalDeg)-macros.ConvertDegreesToRadiance(33.0))) + (math.Sin(macros.ConvertDegreesToRadiance(bDecimalDeg)) * math.Sin(macros.ConvertDegreesToRadiance(27.4)))))

	y := math.Cos(macros.ConvertDegreesToRadiance(bDecimalDeg)) * math.Cos(macros.ConvertDegreesToRadiance(lDecimalDeg-33.0))
	x := (math.Sin(macros.ConvertDegreesToRadiance(bDecimalDeg)) * math.Cos(macros.ConvertDegreesToRadiance(27.4))) - (math.Cos(macros.ConvertDegreesToRadiance(bDecimalDeg)) * math.Sin(macros.ConvertDegreesToRadiance(27.4)) * math.Sin(macros.ConvertDegreesToRadiance(lDecimalDeg-33.0)))
	raDecimalDeg := macros.ConvertRadianceToDegree(math.Atan2(y, x)) + 192.25

	if raDecimalDeg < 0 {
		raDecimalDeg += 360
	}
	if raDecimalDeg >= 360 {
		raDecimalDeg -= 360
	}

	raHrs, raMin, raSec := datetime.ConvertDecimalHrsToHrsMinSec(macros.ConvertDecimalDegressToDecimalHrs(raDecimalDeg))
	decDeg, decMin, decSec := macros.ConvertDecimalDegToDegMinSec(decDecimalDeg)

	return raHrs, raMin, raSec, decDeg, decMin, decSec
}

func CalculateAngleBetweenTwoCelestialObjects(p1RAHrs, p1RAMin int, p1RASec, p1DecDeg, p1DecMin, p1DecSec float64, p2RAHrs, p2RAMin int, p2RASec, p2DecDeg, p2DecMin, p2DecSec float64) (float64, float64, float64) {
	p1RADecimalHrs := datetime.ConvertHrsMinSecToDecimalHrs(p1RAHrs, p1RAMin, p1RASec, false, "")
	p1DecDecimalDeg := macros.ConvertDegMinSecToDecimalDeg(p1DecDeg, p1DecMin, p1DecSec)
	p2RADecimalHrs := datetime.ConvertHrsMinSecToDecimalHrs(p2RAHrs, p2RAMin, p2RASec, false, "")
	p2DecDecimalDeg := macros.ConvertDegMinSecToDecimalDeg(p2DecDeg, p2DecMin, p2DecSec)

	RADiffInDegress := ConvertDecimalHrsToDecimalDegress(p1RADecimalHrs - p2RADecimalHrs)

	angle := macros.ConvertRadianceToDegree(math.Acos((math.Sin(macros.ConvertDegreesToRadiance(p1DecDecimalDeg)) * math.Sin(macros.ConvertDegreesToRadiance(p2DecDecimalDeg))) + (math.Cos(macros.ConvertDegreesToRadiance(p1DecDecimalDeg)) * math.Cos(macros.ConvertDegreesToRadiance(p2DecDecimalDeg)) * math.Cos(macros.ConvertDegreesToRadiance(RADiffInDegress)))))
	angleDeg, angleMin, angleSec := macros.ConvertDecimalDegToDegMinSec(angle)
	return angleDeg, angleMin, angleSec
}

func CalculateRisingAndSettingTime(Gday float64, Gmonth, Gyear float64, raHrs, raMin int, raSec, decDeg, decMin, decSec, geoLatN, geoLongE, refractionInArcMin float64) (float64, float64, float64, float64, float64, float64) {
	decimalRAHrs := datetime.ConvertHrsMinSecToDecimalHrs(int(raHrs), int(raMin), raSec, false, "")
	decimalDECDeg := macros.ConvertDegMinSecToDecimalDeg(decDeg, decMin, decSec)
	refractionDeg := refractionInArcMin / 60 // Converted refraction from arcmin to degress

	cosH := -((math.Sin(macros.ConvertDegreesToRadiance(refractionDeg)) + (math.Sin(macros.ConvertDegreesToRadiance(geoLatN)) * math.Sin(macros.ConvertDegreesToRadiance(decimalDECDeg)))) / (math.Cos(macros.ConvertDegreesToRadiance(geoLatN)) * math.Cos(macros.ConvertDegreesToRadiance(decimalDECDeg))))
	H := 0.0
	if cosH > -1 && cosH < +1 {
		H = macros.ConvertDecimalDegressToDecimalHrs(macros.ConvertRadianceToDegree(math.Acos(cosH)))
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

	Ar := macros.ConvertRadianceToDegree(math.Acos((math.Sin(macros.ConvertDegreesToRadiance(decimalDECDeg)) + (math.Sin(macros.ConvertDegreesToRadiance(refractionDeg)) * math.Sin(macros.ConvertDegreesToRadiance(geoLatN)))) / (math.Cos(macros.ConvertDegreesToRadiance(refractionDeg)) * math.Cos(macros.ConvertDegreesToRadiance(geoLatN)))))

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
	decimalDeg := macros.ConvertDegMinSecToDecimalDeg(deltaDeg, deltaMin, deltaSec)

	S1Hrs := ((3.07327 + (1.33617 * math.Sin(macros.ConvertDegreesToRadiance(decimalHrsTodeg)) * math.Tan(macros.ConvertDegreesToRadiance(decimalDeg)))) * (n1 - n2)) / 3600 // Convert to Hrs
	S2Deg := ((20.0426 * math.Cos(macros.ConvertDegreesToRadiance(decimalHrsTodeg))) * (n1 - n2)) / 3600

	alpha1Hrs, alpha1Min, alpha1Sec := datetime.ConvertDecimalHrsToHrsMinSec(decimalHrs + S1Hrs)
	delta1Deg, delta1Min, delta1Sec := macros.ConvertDecimalDegToDegMinSec(decimalDeg + S2Deg)

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

	nutationInLong := -(17.2 * math.Sin(macros.ConvertDegreesToRadiance(moonNode))) - (1.3 * math.Sin(macros.ConvertDegreesToRadiance(2*L)))
	nutationInObliquity := (9.2 * math.Cos(macros.ConvertDegreesToRadiance(moonNode))) + (0.5 * math.Cos(macros.ConvertDegreesToRadiance(2*L)))

	return nutationInLong, nutationInObliquity
}

func CalculateAberration(day float64, month, year int, trueLambdaDeg, trueLambdaMin, trueLambdaSec, trueBetaDeg, trueBetaMin, trueBetaSec, longDeg, longMin, longSec float64) (float64, float64, float64, float64, float64, float64) {
	trueLambdaDecimalDeg := macros.ConvertDegMinSecToDecimalDeg(trueLambdaDeg, trueLambdaMin, trueLambdaSec)
	trueBetaDecimalDeg := macros.ConvertDegMinSecToDecimalDeg(trueBetaDeg, trueBetaMin, trueBetaSec)
	longDecimalDeg := macros.ConvertDegMinSecToDecimalDeg(longDeg, longMin, longSec)

	trueLambdaDecimalDeg += (-20.5 * math.Cos(macros.ConvertDegreesToRadiance(longDecimalDeg-trueLambdaDecimalDeg)) / math.Cos(macros.ConvertDegreesToRadiance(trueBetaDecimalDeg))) / 3600
	trueBetaDecimalDeg += (-20.5 * math.Sin(macros.ConvertDegreesToRadiance(longDecimalDeg-trueLambdaDecimalDeg)) * math.Sin(macros.ConvertDegreesToRadiance(trueBetaDecimalDeg))) / 3600

	correctedLambdaDeg, correctedLambdaMin, correctedLambdaSec := macros.ConvertDecimalDegToDegMinSec(trueLambdaDecimalDeg)
	correctedBetaDeg, correctedBetaMin, correctedBetaSec := macros.ConvertDecimalDegToDegMinSec(trueBetaDecimalDeg)

	return correctedLambdaDeg, correctedLambdaMin, correctedLambdaSec, correctedBetaDeg, math.Abs(correctedBetaMin), math.Abs(correctedBetaSec)
}

func CalculateRefraction(trueHAHr, trueHAMin, trueHASec, trueDecDeg, trueDecMin, trueDecSec, geoLat, temp, pressure float64) (float64, float64, float64, float64, float64, float64) {
	altitudeDeg, altitudeMin, altitudeSec, azimuthDeg, azimuthMin, azimuthSec := ConvertEquatorialToHorizonCoordinates(trueHAHr, trueHAMin, trueHASec, trueDecDeg, trueDecMin, trueDecSec, geoLat)
	altitudeDecimalDeg := macros.ConvertDegMinSecToDecimalDeg(math.Round(altitudeDeg), altitudeMin, altitudeSec)
	R := 0.0
	z := 90 - altitudeDecimalDeg
	if altitudeDecimalDeg > 15.0 {
		R = (0.00452 * pressure * math.Tan(macros.ConvertDegreesToRadiance(z))) / (273 + temp)
	}

	apperentAltDeg, apperentAltMin, apperentAltSec := macros.ConvertDecimalDegToDegMinSec(R + altitudeDecimalDeg)

	HaHrs, HaMin, HaSec, DecDeg, DecMin, DecSec := ConvertHorizonCoordinatesToEquatorial(0.0, 0, 0, apperentAltDeg, apperentAltMin, apperentAltSec, azimuthDeg, azimuthMin, azimuthSec, geoLat)

	return HaHrs, HaMin, HaSec, DecDeg, DecMin, DecSec
}

func CalculateGeocentricParallax(heightFromSeaLevel, longW, latN float64) (float64, float64) {
	u := macros.ConvertRadianceToDegree(math.Atan(0.996647 * math.Tan(macros.ConvertDegreesToRadiance(latN))))
	hInv := heightFromSeaLevel / 6378140
	pSin := (0.996647 * math.Sin(macros.ConvertDegreesToRadiance(u))) + (hInv * math.Sin(macros.ConvertDegreesToRadiance(latN)))
	pCos := math.Cos(macros.ConvertDegreesToRadiance(u)) + (hInv * math.Cos(macros.ConvertDegreesToRadiance(latN)))
	return pSin, pCos
}

func CalculateParallaxCorrections(day float64, month, year, UTHrs, UTMin int, UTSec, heighSeaLevel, longW, latN float64,
	geoRAHrs, geoRAMin int, geoRASec, geoDecDeg, geoDecMin, geoDecSec, parallaxDeg, parallaxMin, parallaxSec, distanceAU float64) (float64, float64, float64, float64, float64, float64) {

	localDay, localMonth, localYear, localHrs, localMin, localSec := datetime.ConvertUniversalTimeToLocalTime(day, month, year, UTHrs, UTMin, UTSec, 0, 0, 0.0)

	raDecimalHrs := datetime.ConvertHrsMinSecToDecimalHrs(geoRAHrs, geoRAMin, geoRASec, false, "")
	decDecimalDeg := macros.ConvertDegMinSecToDecimalDeg(geoDecDeg, geoDecMin, geoDecSec)

	_, _, _, hourAngle := ConverRightAscensionToHourAngle(int(localDay), int(localMonth), int(localYear), int(localHrs), int(localMin), localSec, float64(geoRAHrs), float64(geoRAMin), geoRASec, -longW, 0, 0, 0.0, false)
	hourAngleDeg := ConvertDecimalHrsToDecimalDegress(hourAngle)

	pSin, pCos := CalculateGeocentricParallax(heighSeaLevel, longW, latN)
	var raInvDecimalHrs, decDegInv float64
	if distanceAU != 0.0 {
		// Calculations for all planets and comets except moon
		piValHrs := macros.ConvertDecimalDegressToDecimalHrs((8.794 / distanceAU) / 3600)
		delta1 := (piValHrs * math.Sin(macros.ConvertDegreesToRadiance(hourAngleDeg)) * macros.ConvertDegreesToRadiance(pCos)) / math.Cos(macros.ConvertDegreesToRadiance(decDecimalDeg))
		raInvDecimalHrs = raDecimalHrs - delta1
		delta2 := ConvertDecimalHrsToDecimalDegress(piValHrs) * macros.ConvertRadianceToDegree((macros.ConvertDegreesToRadiance(pSin)*math.Cos(macros.ConvertDegreesToRadiance(decDecimalDeg)))-(macros.ConvertDegreesToRadiance(pCos)*math.Cos(macros.ConvertDegreesToRadiance(hourAngleDeg))*math.Sin(macros.ConvertDegreesToRadiance(decDecimalDeg))))
		decDegInv = decDecimalDeg - delta2

	} else {
		parallaxDecimalDeg := macros.ConvertDegMinSecToDecimalDeg(parallaxDeg, parallaxMin, parallaxSec)
		r := (1 / math.Sin(macros.ConvertDegreesToRadiance(parallaxDecimalDeg)))
		delta := macros.ConvertRadianceToDegree(math.Atan(macros.ConvertDegreesToRadiance(pCos)*math.Sin(macros.ConvertDegreesToRadiance(hourAngleDeg))) / (macros.ConvertDegreesToRadiance(r)*math.Cos(macros.ConvertDegreesToRadiance(decDecimalDeg)) - (macros.ConvertDegreesToRadiance(pCos) * math.Cos(macros.ConvertDegreesToRadiance(hourAngleDeg)))))

		hourAngleInvDeg := hourAngleDeg + delta
		deltaHrs := macros.ConvertDecimalDegressToDecimalHrs(delta)

		raInvDecimalHrs = raDecimalHrs - deltaHrs
		a := (macros.ConvertDegreesToRadiance(r) * math.Sin(macros.ConvertDegreesToRadiance(decDecimalDeg))) - macros.ConvertDegreesToRadiance(pSin)
		b := (macros.ConvertDegreesToRadiance(r) * math.Cos(macros.ConvertDegreesToRadiance(decDecimalDeg)) * math.Cos(macros.ConvertDegreesToRadiance(hourAngleDeg))) - macros.ConvertDegreesToRadiance(pCos)
		decDegInv = macros.ConvertRadianceToDegree(math.Atan(math.Cos(macros.ConvertDegreesToRadiance(hourAngleInvDeg)) * (a / b)))
	}

	appreantRAHrs, appreantRAMin, appreantRASec := datetime.ConvertDecimalHrsToHrsMinSec(raInvDecimalHrs)
	appreantDecDeg, appreantDecMin, appreantDecSec := macros.ConvertDecimalDegToDegMinSec(decDegInv)

	return appreantRAHrs, appreantRAMin, appreantRASec, appreantDecDeg, appreantDecMin, appreantDecSec
}

func CalculateHeliographicCoordinates(day float64, month, year, UTHrs, UTMin int, UTSec, geoLongDeg, geoLongMin, geoLongSec, positionAngleTheta, displacementP1, angularRadiusSDeg, angularRadiusSMin, angularRadiusSSec float64) (float64, float64) {
	julianDate := datetime.ConvertGreenwichDateToJulianDate(day, month, year)
	Ideg := 7.25
	T := (julianDate - 2415020.0) / 36525.0
	deltaDeg := (84 * T) / 60
	gamma := macros.ConvertDegMinSecToDecimalDeg(74, 22, 0) + deltaDeg
	lambda := macros.CalculatePositionOfSunHelper(day, float64(month), float64(year), UTHrs, UTMin, UTSec)
	y := math.Sin(macros.ConvertDegreesToRadiance(gamma-lambda)) * math.Cos(macros.ConvertDegreesToRadiance(Ideg))
	x := -math.Cos(math.Sin(macros.ConvertDegreesToRadiance(gamma - lambda)))
	AInv := macros.ConvertRadianceToDegree(math.Atan2(y, x))
	MInv := (360 / 25.38) * (julianDate - 2398220)

	for MInv > 360 {
		MInv -= 360
	}
	M := 360 - MInv
	L0 := M + AInv
	for L0 > 360 {
		L0 -= 360
	}
	B0 := macros.ConvertRadianceToDegree(math.Asin(math.Sin(macros.ConvertDegreesToRadiance(lambda-gamma)) * math.Sin(macros.ConvertDegreesToRadiance(Ideg))))

	theta1 := macros.ConvertRadianceToDegree(math.Atan(-math.Cos(macros.ConvertDegreesToRadiance(lambda)) * math.Tan(macros.ConvertDegreesToRadiance(23.442))))
	theta2 := macros.ConvertRadianceToDegree(math.Atan(-math.Cos(macros.ConvertDegreesToRadiance(gamma-lambda)) * math.Tan(macros.ConvertDegreesToRadiance(Ideg))))
	P := theta1 + theta2

	// Convert Arcmin to decimal deg
	SdecimalArcmin := macros.ConvertDegMinSecToDecimalDeg(angularRadiusSDeg, angularRadiusSMin, angularRadiusSSec) * 60
	sinInv := macros.ConvertRadianceToDegree(math.Asin(displacementP1 / SdecimalArcmin))
	P1 := sinInv - (displacementP1 / 60)
	B := macros.ConvertRadianceToDegree(math.Asin((math.Sin(macros.ConvertDegreesToRadiance(B0)) * math.Cos(macros.ConvertDegreesToRadiance(P1))) + (math.Cos(macros.ConvertDegreesToRadiance(B0)) * math.Sin(macros.ConvertDegreesToRadiance(P1)) * math.Cos(macros.ConvertDegreesToRadiance(P-positionAngleTheta)))))
	A := macros.ConvertRadianceToDegree(math.Asin((math.Sin(macros.ConvertDegreesToRadiance(P1)) * math.Sin(macros.ConvertDegreesToRadiance(P-positionAngleTheta))) / math.Cos(macros.ConvertDegreesToRadiance((B)))))
	L := A + L0
	for L > 360 {
		L -= 360
	}
	return B, L
}
