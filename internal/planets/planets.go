package planets

import (
	"fmt"
	datetime "go-astronomy/internal/dateTime"
	"go-astronomy/internal/macros"
	"math"
)

func CalculateCoordinatesOfPlanet(day float64, month, year int, planetName string, epochDay float64, epochMonth, epochYear int) (raHrs, raMins int, raSecs float64, decDeg, decMin int, decSec float64) {
	daysSinceYearStart := datetime.CalculateDayNumber(day, month, year)
	daysSinceEpoch := macros.DaysElapsedSinceEpoch(epochYear, year)
	totalDays := daysSinceYearStart + daysSinceEpoch
	planetValues := GetPlanetData(planetName)
	earthValues := GetPlanetData("Earth")

	Np := macros.AdjustAngleRange((360/365.242191)*(totalDays/planetValues["Tp"].(float64)), 0, 360)
	Mp := Np + planetValues["Long"].(float64) - planetValues["Peri"].(float64)
	Vp := macros.AdjustAngleRange(Mp+((360/math.Pi)*planetValues["Ecc"].(float64)*math.Sin(macros.ConvertDegreesToRadiance(Mp))), 0, 360)
	Lp := macros.AdjustAngleRange(Vp+planetValues["Peri"].(float64), 0, 360)
	r := (planetValues["Axis"].(float64) * (1 - math.Pow(planetValues["Ecc"].(float64), 2))) / (1 + (planetValues["Ecc"].(float64) * math.Cos(macros.ConvertDegreesToRadiance(Vp))))

	// Calculate the values for Earths
	Ne := macros.AdjustAngleRange((360/365.242191)*(totalDays/earthValues["Tp"].(float64)), 0, 360)
	Me := Ne + earthValues["Long"].(float64) - earthValues["Peri"].(float64)
	Ve := macros.AdjustAngleRange(Me+((360/math.Pi)*earthValues["Ecc"].(float64)*math.Sin(macros.ConvertDegreesToRadiance(Me))), 0, 360)
	Le := macros.AdjustAngleRange(Ve+earthValues["Peri"].(float64), 0, 360)

	R := (earthValues["Axis"].(float64) * (1 - math.Pow(earthValues["Ecc"].(float64), 2))) / (1 + (earthValues["Ecc"].(float64) * math.Cos(macros.ConvertDegreesToRadiance(Ve))))

	// fmt.Printf("\ndaysSinceYearStart : %f\ndaysSinceEpoch : %f\ntotalDays : %f\nNp : %f\n", daysSinceYearStart, daysSinceEpoch, totalDays, Np)
	// fmt.Printf("\nMp : %f\nVp : %f\nLp : %f\nr : %f\n", Mp, Vp, Lp, r)
	// fmt.Printf("\nMe : %f\nVe : %f\nLe : %f\nR : %f\n", Me, Ve, Le, R)

	si := macros.ConvertRadianceToDegree(math.Asin(math.Sin(macros.ConvertDegreesToRadiance(Lp-planetValues["Node"].(float64))) * (math.Sin(macros.ConvertDegreesToRadiance(planetValues["Incl"].(float64))))))
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
	// earthValues := GetPlanetData("Earth")

	Np := macros.AdjustAngleRange((360/365.242191)*(totalDays/planetValues["Tp"].(float64)), 0, 360)
	Mp := Np + planetValues["Long"].(float64) - planetValues["Peri"].(float64)

	Ep, Wp, e := macros.CalculateEgWgAnde(day, month, year, 0, 0, 0.0)
	// Eg, Wg, e := macros.AdjustAngleRange(planetValues["Long"].(float64), 0, 360), macros.AdjustAngleRange(planetValues["Peri"].(float64), 0, 360), macros.AdjustAngleRange(planetValues["Ecc"].(float64), 0, 360)
	// fmt.Printf("\n%f = %f\n%f = %f\n%f = %f\n", Eg, planetValues["Long"].(float64), Wg, planetValues["Peri"].(float64), e, planetValues["Ecc"].(float64))
	fmt.Printf("\nEg : %f\nWg : %f\ne : %f\n", Ep, Wp, e)
	MRad := macros.ConvertDegreesToRadiance(macros.AdjustAngleRange(Ep-Wp, 0, 360))
	eccentricAnomaly := macros.ConvertRadianceToDegree(macros.CalculateEccentricAnomaly(MRad, e))
	Vp := macros.ConvertRadianceToDegree(math.Atan(math.Sqrt((1+e)/(1-e))*math.Tan(macros.ConvertDegreesToRadiance(eccentricAnomaly/2))) * 2)

	if Vp < 0 {
		Vp += 360
	}
	Lp := macros.AdjustAngleRange(Vp+planetValues["Peri"].(float64), 0, 360)

	fmt.Printf("\njulianDate : %f\nT : %f\nA : %f\nP : %f\nQ : %f\nV : %f\ndeltaL : %f\nMp : %f\nVp : %f\nLp : %f\n", julianDate, T, A, P, Q, V, deltaL, Mp, Vp, Lp)
	return raHrs, raMins, raSecs, decDeg, decMin, decSec
}
