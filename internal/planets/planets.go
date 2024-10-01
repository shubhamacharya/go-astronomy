package planets

import (
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
