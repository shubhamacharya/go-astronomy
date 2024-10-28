package planets

import (
	"fmt"
	"go-astronomy/internal/coords"
	datetime "go-astronomy/internal/dateTime"
	"go-astronomy/internal/macros"
	"go-astronomy/internal/sun"
	"math"
)

func CalculateCoordinatesOfPlanet(day float64, month, year int, planetName string, epochDay float64, epochMonth, epochYear int) (raHrs, raMins int, raSecs float64, decDeg, decMin int, decSec float64) {
	daysSinceYearStart := datetime.CalculateDayNumber(day, month, year)
	daysSinceEpoch := macros.DaysElapsedSinceEpoch(epochYear, year)
	totalDays := daysSinceYearStart + daysSinceEpoch
	planetValues := GetPlanetData(planetName)
	earthValues := GetPlanetData("Earth")

	_, _, Vp := macros.CalculateN_M_V(planetValues, totalDays)

	// Np := macros.AdjustAngleRange((360/365.242191)*(totalDays/planetValues["Tp"].(float64)), 0, 360)
	// Mp := Np + planetValues["Long"].(float64) - planetValues["Peri"].(float64)
	// Vp := macros.AdjustAngleRange(Mp+((360/math.Pi)*planetValues["Ecc"].(float64)*math.Sin(macros.ConvertDegreesToRadiance(Mp))), 0, 360)

	Lp, r, _ := macros.CalculatePerturbationsInPlanetsOrbitHelper(planetValues, Vp, 0)

	// Lp := macros.AdjustAngleRange(Vp+planetValues["Peri"].(float64), 0, 360)
	// r := (planetValues["Axis"].(float64) * (1 - math.Pow(planetValues["Ecc"].(float64), 2))) / (1 + (planetValues["Ecc"].(float64) * math.Cos(macros.ConvertDegreesToRadiance(Vp))))

	// Calculate the values for Earths
	_, _, Ve := macros.CalculateN_M_V(earthValues, totalDays)
	// Ne := macros.AdjustAngleRange((360/365.242191)*(totalDays/earthValues["Tp"].(float64)), 0, 360)
	// Me := Ne + earthValues["Long"].(float64) - earthValues["Peri"].(float64)
	// Ve := macros.AdjustAngleRange(Me+((360/math.Pi)*earthValues["Ecc"].(float64)*math.Sin(macros.ConvertDegreesToRadiance(Me))), 0, 360)

	// Le := macros.AdjustAngleRange(Ve+earthValues["Peri"].(float64), 0, 360)
	// R := (earthValues["Axis"].(float64) * (1 - math.Pow(earthValues["Ecc"].(float64), 2))) / (1 + (earthValues["Ecc"].(float64) * math.Cos(macros.ConvertDegreesToRadiance(Ve))))

	Le, R, si := macros.CalculatePerturbationsInPlanetsOrbitHelper(earthValues, Ve, 0)

	// fmt.Printf("\ndaysSinceYearStart : %f\ndaysSinceEpoch : %f\ntotalDays : %f\nNp : %f\n", daysSinceYearStart, daysSinceEpoch, totalDays, Np)
	// fmt.Printf("\nMp : %f\nVp : %f\nLp : %f\nr : %f\n", Mp, Vp, Lp, r)
	// fmt.Printf("\nMe : %f\nVe : %f\nLe : %f\nR : %f\n", Me, Ve, Le, R)

	// si := macros.ConvertRadianceToDegree(math.Asin(math.Sin(macros.ConvertDegreesToRadiance(Lp-planetValues["Node"].(float64))) * (math.Sin(macros.ConvertDegreesToRadiance(planetValues["Incl"].(float64))))))
	y := math.Sin(macros.ConvertDegreesToRadiance(Lp-planetValues["Node"].(float64))) * (math.Cos(macros.ConvertDegreesToRadiance(planetValues["Incl"].(float64))))
	x := math.Cos(macros.ConvertDegreesToRadiance(Lp - planetValues["Node"].(float64)))
	tanInv := macros.AdjustAngleInQuadrant(x, y, macros.ConvertRadianceToDegree(math.Atan(y/x)))
	ldash := tanInv + planetValues["Node"].(float64)
	rdash := r * math.Cos(macros.ConvertDegreesToRadiance(si))
	// fmt.Printf("\nsi : %f\ny : %f\nx : %f\ntanInv : %f\nldash : %f\nrdash : %f\n", si, y, x, tanInv, ldash, rdash)

	lambdaDecDeg, betaDecDeg := 0.0, 0.0
	lambdaDeg, lambdaMin, lambdaSec := 0, 0, 0.0
	betaDeg, betaMin, betaSec := 0, 0, 0.0

	if planetName == "Mercury" || planetName == "Venus" {
		lambdaDecDeg = macros.AdjustAngleRange(180+Le+macros.ConvertRadianceToDegree(math.Atan((rdash*math.Sin(macros.ConvertDegreesToRadiance(Le-ldash)))/(R-(rdash*math.Cos(macros.ConvertDegreesToRadiance(Le-ldash)))))), 0, 360)
		lambdaDeg, lambdaMin, lambdaSec = macros.ConvertDecimalDegToDegMinSec(lambdaDecDeg)

	} else {
		lambdaDecDeg = macros.AdjustAngleRange(macros.ConvertRadianceToDegree(math.Atan((R*math.Sin(macros.ConvertDegreesToRadiance(ldash-Le)))/(rdash-(R*math.Cos(macros.ConvertDegreesToRadiance(ldash-Le)))))+macros.ConvertDegreesToRadiance(ldash)), 0, 360)
		lambdaDeg, lambdaMin, lambdaSec = macros.ConvertDecimalDegToDegMinSec(lambdaDecDeg)
	}
	betaDecDeg = macros.ConvertRadianceToDegree(math.Atan((rdash * math.Tan(macros.ConvertDegreesToRadiance(si)) * math.Sin(macros.ConvertDegreesToRadiance(lambdaDecDeg-ldash))) / (R * math.Sin(macros.ConvertDegreesToRadiance(ldash-Le)))))
	betaDeg, betaMin, betaSec = macros.ConvertDecimalDegToDegMinSec(betaDecDeg)

	raHrs, raMins, raSecs, decDeg, decMin, decSec = macros.ConvertEclipticCoordinatesToEquatorial(day, month, year, lambdaDeg, lambdaMin, lambdaSec, betaDeg, betaMin, betaSec, epochDay, epochMonth, epochYear)
	// fmt.Printf("\nlambda : %f\nbeta : %f\n", lambdaDecDeg, betaDecDeg)
	// fmt.Printf("\nraHrs : %d\nraMin : %d\nraSec : %f\ndecDeg : %d\ndecMin : %d\ndecSec : %f\n", raHrs, raMins, raSecs, decDeg, decMin, decSec)
	return raHrs, raMins, raSecs, decDeg, decMin, decSec
}

func CalculateApproximatePositionOfPlanet(day float64, month, year int, planetName string, epochDay float64, epochMonth, epochYear int) (raHrs, raMins int, raSecs float64, decDeg, decMin int, decSec float64) {
	daysSinceYearStart := datetime.CalculateDayNumber(day, month, year)
	daysSinceEpoch := macros.DaysElapsedSinceEpoch(epochYear, year)
	totalDays := daysSinceYearStart + daysSinceEpoch
	planetValues := GetPlanetData(planetName)
	earthValues := GetPlanetData("Earth")

	l := macros.AdjustAngleRange(((360/365.242191)*(totalDays/planetValues["Tp"].(float64)))+planetValues["Long"].(float64), 0, 360)
	L := macros.AdjustAngleRange(((360/365.242191)*(totalDays/earthValues["Tp"].(float64)))+earthValues["Long"].(float64), 0, 360)

	lambdaDecDeg := macros.AdjustAngleRange(macros.ConvertRadianceToDegree(math.Atan(math.Sin(macros.ConvertDegreesToRadiance(l-L))/(planetValues["Axis"].(float64)-math.Cos(macros.ConvertDegreesToRadiance(l-L)))))+l, 0, 360)
	lambdaDeg, lambdaMin, lambdaSec := macros.ConvertDecimalDegToDegMinSec(lambdaDecDeg)
	raHrs, raMins, raSecs, decDeg, decMin, decSec = macros.ConvertEclipticCoordinatesToEquatorial(day, month, year, lambdaDeg, lambdaMin, lambdaSec, 0, 0, 0.0, epochDay, epochMonth, epochYear)

	// fmt.Printf("\ndaysSinceYearStart : %f\ndaysSinceEpoch : %f\ntotalDays : %f\nlambda : %f\n", daysSinceYearStart, daysSinceEpoch, totalDays, lambdaDecDeg)
	// fmt.Printf("\nraHrs : %d\nraMin : %d\nraSec : %f\ndecDeg : %d\ndecMin : %d\ndecSec : %f\n", raHrs, raMins, raSecs, decDeg, decMin, decSec)

	return raHrs, raMins, raSecs, decDeg, decMin, decSec
}

func CalculatePerturbationsInPlanetsOrbit(day float64, month, year int, planetName string, epochDay float64, epochMonth, epochYear int) (raHrs, raMins int, raSecs float64, decDeg, decMin int, decSec float64) {
	julianDate := datetime.ConvertGreenwichDateToJulianDate(day, month, year)
	T := macros.RoundToNDecimals(((julianDate - 2415020) / 36525), 6)
	A := (T / 5) + 0.1
	P := (3034.906100 * T) + 237.47555
	Q := (1222.1139 * T) + 265.91650
	V := (5 * Q) - (2 * P)
	B := Q - P
	deltaL := 0.0
	if planetName == "Jupiter" {
		deltaL = ((0.3314 - (0.0103 * A)) * math.Sin(macros.ConvertDegreesToRadiance(V))) - (0.0644 * A * math.Cos(macros.ConvertDegreesToRadiance(V)))
	} else if planetName == "Saturn" {
		deltaL = (((0.1609 * A) - 0.0105) * math.Cos(macros.ConvertDegreesToRadiance(V))) + (((0.0182 * A) - 0.8142) * math.Sin(macros.ConvertDegreesToRadiance(V))) - (0.1488 * math.Sin(macros.ConvertDegreesToRadiance(B))) - (0.0408 * math.Sin(macros.ConvertDegreesToRadiance(2*B))) + (0.0856 * math.Sin(macros.ConvertDegreesToRadiance(B)) * math.Cos(macros.ConvertDegreesToRadiance(Q))) + (0.0813 * math.Cos(macros.ConvertDegreesToRadiance(B)) * math.Sin(macros.ConvertDegreesToRadiance(Q)))
	}
	daysSinceYearStart := datetime.CalculateDayNumber(day, month, year)
	daysSinceEpoch := macros.DaysElapsedSinceEpoch(epochYear, year)
	totalDays := daysSinceYearStart + daysSinceEpoch

	planetValues := GetPlanetData(planetName)
	earthValues := GetPlanetData("Earth")

	_, _, Vp := macros.CalculateN_M_V(planetValues, totalDays)

	Lp, r, si := macros.CalculatePerturbationsInPlanetsOrbitHelper(planetValues, Vp, deltaL)

	_, _, Ve := macros.CalculateN_M_V(earthValues, totalDays)

	Le, R, _ := macros.CalculatePerturbationsInPlanetsOrbitHelper(earthValues, Ve, deltaL)
	y := math.Sin(macros.ConvertDegreesToRadiance(Lp-planetValues["Node"].(float64))) * (math.Cos(macros.ConvertDegreesToRadiance(planetValues["Incl"].(float64))))
	x := math.Cos(macros.ConvertDegreesToRadiance(Lp - planetValues["Node"].(float64)))
	tanInv := macros.AdjustAngleInQuadrant(x, y, macros.ConvertRadianceToDegree(math.Atan(y/x)))
	ldash := tanInv + planetValues["Node"].(float64)
	rdash := r * math.Cos(macros.ConvertDegreesToRadiance(si))
	// fmt.Printf("\nsi : %f\ny : %f\nx : %f\ntanInv : %f\nldash : %f\nrdash : %f\n", si, y, x, tanInv, ldash, rdash)

	lambdaDecDeg, betaDecDeg := 0.0, 0.0
	lambdaDeg, lambdaMin, lambdaSec := 0, 0, 0.0
	betaDeg, betaMin, betaSec := 0, 0, 0.0

	if planetName == "Mercury" || planetName == "Venus" {
		lambdaDecDeg = macros.AdjustAngleRange(180+Le+macros.ConvertRadianceToDegree(math.Atan((rdash*math.Sin(macros.ConvertDegreesToRadiance(Le-ldash)))/(R-(rdash*math.Cos(macros.ConvertDegreesToRadiance(Le-ldash)))))), 0, 360)
		lambdaDeg, lambdaMin, lambdaSec = macros.ConvertDecimalDegToDegMinSec(lambdaDecDeg)

	} else {
		lambdaDecDeg = macros.AdjustAngleRange(macros.ConvertRadianceToDegree(math.Atan((R*math.Sin(macros.ConvertDegreesToRadiance(ldash-Le)))/(rdash-(R*math.Cos(macros.ConvertDegreesToRadiance(ldash-Le)))))+macros.ConvertDegreesToRadiance(ldash)), 0, 360)
		lambdaDeg, lambdaMin, lambdaSec = macros.ConvertDecimalDegToDegMinSec(lambdaDecDeg)
	}
	betaDecDeg = macros.ConvertRadianceToDegree(math.Atan((rdash * math.Tan(macros.ConvertDegreesToRadiance(si)) * math.Sin(macros.ConvertDegreesToRadiance(lambdaDecDeg-ldash))) / (R * math.Sin(macros.ConvertDegreesToRadiance(ldash-Le)))))
	betaDeg, betaMin, betaSec = macros.ConvertDecimalDegToDegMinSec(betaDecDeg)

	raHrs, raMins, raSecs, decDeg, decMin, decSec = macros.ConvertEclipticCoordinatesToEquatorial(day, month, year, lambdaDeg, lambdaMin, lambdaSec, betaDeg, betaMin, betaSec, epochDay, epochMonth, epochYear)
	// fmt.Printf("\nlambda : %f\nbeta : %f\n", lambdaDecDeg, betaDecDeg)
	// fmt.Printf("\njulianDate : %f\nT : %f\nA : %f\nP : %f\nQ : %f\nV : %f\ndeltaL : %f\nMp : %f\nVp : %f\nLp : %f\n", julianDate, T, A, P, Q, V, deltaL, Mp, Vp, Lp)
	// fmt.Printf("\nraHrs : %d\nraMin : %d\nraSec : %f\ndecDeg : %d\ndecMin : %d\ndecSec : %f\n", raHrs, raMins, raSecs, decDeg, decMin, decSec)
	// fmt.Printf("\nMe : %f\nVe : %f\nLe : %f\n", Me, Ve, Le)
	return raHrs, raMins, raSecs, decDeg, decMin, decSec
}

func CalculatePlantesDistanceLightTravelAndAngularSize(day float64, month, year int, planetName string, epochDay float64, epochMonth, epochYear int) (distance float64, timeMin int, timeSec, diameter float64) {
	julianDate := datetime.ConvertGreenwichDateToJulianDate(day, month, year)
	T := macros.RoundToNDecimals(((julianDate - 2415020) / 36525), 6)
	A := (T / 5) + 0.1
	P := (3034.906100 * T) + 237.47555
	Q := (1222.1139 * T) + 265.91650
	V := (5 * Q) - (2 * P)
	B := Q - P
	deltaL := 0.0
	if planetName == "Jupiter" {
		deltaL = ((0.3314 - (0.0103 * A)) * math.Sin(macros.ConvertDegreesToRadiance(V))) - (0.0644 * A * math.Cos(macros.ConvertDegreesToRadiance(V)))
	} else if planetName == "Saturn" {
		deltaL = (((0.1609 * A) - 0.0105) * math.Cos(macros.ConvertDegreesToRadiance(V))) + (((0.0182 * A) - 0.8142) * math.Sin(macros.ConvertDegreesToRadiance(V))) - (0.1488 * math.Sin(macros.ConvertDegreesToRadiance(B))) - (0.0408 * math.Sin(macros.ConvertDegreesToRadiance(2*B))) + (0.0856 * math.Sin(macros.ConvertDegreesToRadiance(B)) * math.Cos(macros.ConvertDegreesToRadiance(Q))) + (0.0813 * math.Cos(macros.ConvertDegreesToRadiance(B)) * math.Sin(macros.ConvertDegreesToRadiance(Q)))
	}
	daysSinceYearStart := datetime.CalculateDayNumber(day, month, year)
	daysSinceEpoch := macros.DaysElapsedSinceEpoch(epochYear, year)
	totalDays := daysSinceYearStart + daysSinceEpoch
	planetValues := GetPlanetData(planetName)
	earthValues := GetPlanetData("Earth")

	_, _, Vp := macros.CalculateN_M_V(planetValues, totalDays)

	Lp, r, si := macros.CalculatePerturbationsInPlanetsOrbitHelper(planetValues, Vp, deltaL)

	_, _, Ve := macros.CalculateN_M_V(earthValues, totalDays)

	Le, R, _ := macros.CalculatePerturbationsInPlanetsOrbitHelper(earthValues, Ve, deltaL)
	fmt.Printf("\nLp : %v\nr : %v\nsi : %v\nLe : %f\nR : %v", Lp, r, si, Le, R)

	distance = math.Sqrt(math.Pow(R, 2) + math.Pow(r, 2) - (2 * R * r * (math.Cos((Lp - Le)) * math.Cos((si)))))
	_, timeMin, timeSec = datetime.ConvertDecimalHrsToHrsMinSec(0.1386 * distance)
	diameter = planetValues["Theta0"].(float64) / distance
	fmt.Printf("\ndistance : %v\ntime : %v min %v sec\ndiameter : %v", distance, timeMin, timeSec, diameter)
	return diameter, timeMin, timeSec, diameter
}

func CalculatePhasesOfPlanets(day float64, month, year int, planetName string, epochDay float64, epochMonth, epochYear int) (d, F float64) {
	julianDate := datetime.ConvertGreenwichDateToJulianDate(day, month, year)
	T := macros.RoundToNDecimals(((julianDate - 2415020) / 36525), 6)
	A := (T / 5) + 0.1
	P := (3034.906100 * T) + 237.47555
	Q := (1222.1139 * T) + 265.91650
	V := (5 * Q) - (2 * P)
	B := Q - P
	deltaL := 0.0
	if planetName == "Jupiter" {
		deltaL = ((0.3314 - (0.0103 * A)) * math.Sin(macros.ConvertDegreesToRadiance(V))) - (0.0644 * A * math.Cos(macros.ConvertDegreesToRadiance(V)))
	} else if planetName == "Saturn" {
		deltaL = (((0.1609 * A) - 0.0105) * math.Cos(macros.ConvertDegreesToRadiance(V))) + (((0.0182 * A) - 0.8142) * math.Sin(macros.ConvertDegreesToRadiance(V))) - (0.1488 * math.Sin(macros.ConvertDegreesToRadiance(B))) - (0.0408 * math.Sin(macros.ConvertDegreesToRadiance(2*B))) + (0.0856 * math.Sin(macros.ConvertDegreesToRadiance(B)) * math.Cos(macros.ConvertDegreesToRadiance(Q))) + (0.0813 * math.Cos(macros.ConvertDegreesToRadiance(B)) * math.Sin(macros.ConvertDegreesToRadiance(Q)))
	}
	daysSinceYearStart := datetime.CalculateDayNumber(day, month, year)
	daysSinceEpoch := macros.DaysElapsedSinceEpoch(epochYear, year)
	totalDays := daysSinceYearStart + daysSinceEpoch
	planetValues := GetPlanetData(planetName)

	raHrs, raMins, raSecs, decDeg, decMin, decSec := CalculateCoordinatesOfPlanet(day, month, year, planetName, epochDay, epochMonth, epochYear)
	eclipticLongDeg, eclipticLongMin, eclipticLongSec, _, _, _ := coords.ConvertEquatorialCoordinatesToEcliptic(day, month, year, raHrs, raMins, raSecs, decDeg, decMin, decSec, epochDay, epochMonth, epochYear)
	lambda := macros.ConvertDegMinSecToDecimalDeg(eclipticLongDeg, eclipticLongMin, eclipticLongSec)

	_, _, Vp := macros.CalculateN_M_V(planetValues, totalDays)

	Lp, _, _ := macros.CalculatePerturbationsInPlanetsOrbitHelper(planetValues, Vp, deltaL)
	d = lambda - Lp
	F = 0.5 * (1 + math.Cos(d))

	return d, F
}

func CalculatePositionAngleOfBrightLimb(day float64, month, year int, planetName string, epochDay float64, epochMonth, epochYear int) (angleOfBrightLimb float64) {
	raHrs, raMins, raSecs, decDeg, decMin, decSec := CalculateCoordinatesOfPlanet(day, month, year, planetName, epochDay, epochMonth, epochYear)
	planetRA := datetime.ConvertHrsMinSecToDecimalHrs(raHrs, raMins, raSecs, false, false)
	planetDec := macros.ConvertDegMinSecToDecimalDeg(decDeg, decMin, decSec)
	fmt.Printf("\nra : %v %v %v dec : %v %v %v\n", raHrs, raMins, raSecs, decDeg, decMin, decSec)

	raHrs, raMins, raSecs, decDeg, decMin, decSec, _ = sun.CalculatePositionOfSun(day, month, year, 0, 0, 0, epochDay, epochMonth, epochYear)
	fmt.Printf("\nSun ra : %v %v %v Sun dec : %v %v %v\n", raHrs, raMins, raSecs, decDeg, decMin, decSec)
	sunRA := datetime.ConvertHrsMinSecToDecimalHrs(raHrs, raMins, raSecs, false, false)
	sunDec := macros.ConvertDegMinSecToDecimalDeg(decDeg, decMin, decSec)

	deltaRA := coords.ConvertDecimalHrsToDecimalDegress(sunRA - planetRA)
	fmt.Printf("\ndeltaRA : %v\n", deltaRA)
	y := math.Cos(macros.ConvertDegreesToRadiance(sunDec)) * math.Sin(macros.ConvertDegreesToRadiance(deltaRA))
	x := (math.Cos(macros.ConvertDegreesToRadiance(planetDec)) * math.Sin(macros.ConvertDegreesToRadiance(sunDec))) - (math.Sin(macros.ConvertDegreesToRadiance(planetDec)) * math.Cos(macros.ConvertDegreesToRadiance(sunDec)) * math.Cos(macros.ConvertDegreesToRadiance(deltaRA)))
	angleOfBrightLimb = math.Atan(y / x)
	angleOfBrightLimb = macros.AdjustAngleInQuadrant(x, y, angleOfBrightLimb)
	return angleOfBrightLimb
}

func CalculateApparentBrightnessOfPlanet(day float64, month, year int, planetName string, epochDay float64, epochMonth, epochYear int) float64 {
	daysSinceYearStart := datetime.CalculateDayNumber(day, month, year)
	daysSinceEpoch := macros.DaysElapsedSinceEpoch(epochYear, year)
	totalDays := daysSinceYearStart + daysSinceEpoch
	planetValues := GetPlanetData(planetName)

	_, _, Vp := macros.CalculateN_M_V(planetValues, totalDays)

	_, r, _ := macros.CalculatePerturbationsInPlanetsOrbitHelper(planetValues, Vp, 0)

	distance, _, _, _ := CalculatePlantesDistanceLightTravelAndAngularSize(day, month, year, planetName, epochDay, epochMonth, epochYear)
	_, F := CalculatePhasesOfPlanets(day, month, year, planetName, epochDay, epochMonth, epochYear)

	m := (5 * math.Log10((r*distance)/(math.Sqrt(F)))) + planetValues["V0"].(float64)
	return m
}
