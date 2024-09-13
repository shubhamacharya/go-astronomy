package sun

import (
	"fmt"
	"go-astronomy/internal/coords"
	datetime "go-astronomy/internal/dateTime"
	"go-astronomy/internal/macros"
	"math"
)

// Values at epoch 2010.0
const semiMajorAxis = 1.495985e8
const angularDiameter = 0.533128

func CalculatePositionOfSun(GDay float64, GMonth, GYear int, UTHrs, UTMins int, UTSec float64) (raHrs, raMin int, raSec, decDeg, decMin, decSec, lambda float64) {

	lambda = macros.CalculatePositionOfSunHelper(GDay, GMonth, GYear, UTHrs, UTMins, UTSec)
	lambdaDeg, lambdaMin, lambdaSec := macros.ConvertDecimalDegToDegMinSec(lambda)
	raHrs, raMin, raSec, decDeg, decMin, decSec = macros.ConvertEclipticCoordinatesToEquatorial(GDay, GMonth, GYear, lambdaDeg, lambdaMin, lambdaSec, 0, 0, 0)
	fmt.Printf("\nra : %d %d %f\ndec : %f %f %f\n", raHrs, raMin, raSec, decDeg, decMin, decSec)
	return raHrs, raMin, raSec, decDeg, decMin, decSec, lambda
}

func CalculatePrecisePositionOfSun(GDay float64, GMonth, GYear int, UTHrs, UTMins int, UTSec float64) (raHrs, raMins int, raSecs, decDeg, decMin, decSec, lambda0 float64) {
	Eg, Wg, e := macros.CalculateEgWgAnde(GDay, GMonth, GYear, UTHrs, UTMins, UTSec)
	MRad := macros.ConvertDegreesToRadiance(macros.AdjustAngleRange(Eg-Wg, 0, 360))
	eccentricAnomaly := macros.ConvertRadianceToDegree(macros.CalculateEccentricAnomaly(MRad, e))
	V := macros.ConvertRadianceToDegree(math.Atan(math.Sqrt((1+e)/(1-e))*math.Tan(macros.ConvertDegreesToRadiance(eccentricAnomaly/2))) * 2)
	if V < 0 {
		V += 360
	}

	lambda0 = V + Wg
	if lambda0 > 360 {
		lambda0 -= 360
	}
	lambda0Deg, lambda0Min, lambda0Sec := macros.ConvertDecimalDegToDegMinSec(lambda0)
	raHrs, raMins, raSecs, decDeg, decMin, decSec = macros.ConvertEclipticCoordinatesToEquatorial(GDay, GMonth, GYear, lambda0Deg, lambda0Min, lambda0Sec, 0, 0, 0)

	return raHrs, raMins, raSecs, decDeg, decMin, decSec, lambda0
}

func CalculateSunsDistanceAndAngularSize(GDay float64, GMonth, GYear int, UTHrs, UTMins int, UTSec float64) (r, thetaDeg, thetaMin, thetaSec, theta float64) {
	_, Wg, e := macros.CalculateEgWgAnde(GDay, GMonth, GYear, UTHrs, UTMins, UTSec)
	_, _, _, _, _, _, lambda0 := CalculatePrecisePositionOfSun(GDay, GMonth, GYear, UTHrs, UTMins, UTSec)
	V := lambda0 - Wg
	if V < 0 {
		V += 360
	}
	F := ((1 + (e * math.Cos(macros.ConvertDegreesToRadiance(V)))) / (1 - math.Pow(e, 2)))
	r = semiMajorAxis / F
	theta = angularDiameter * F
	thetaDeg, thetaMin, thetaSec = macros.ConvertDecimalDegToDegMinSec(theta)
	return r, thetaDeg, thetaMin, thetaSec, theta
}

func CalculateSunsRiseAndSet(GDay float64, GMonth, GYear int, UTHrs, UTMins int, UTSec, geoLongW, geoLatN, refractionInArcMin, daylightsavingHrs, daylightsavingMin, timeZone float64) (riseHrs, riseMin int, riseSec float64, SetHrs, SetMin int, SetSec float64) {
	raHrs, raMins, raSecs, decDeg, decMin, decSec, _ := CalculatePositionOfSun(GDay, GMonth, GYear, UTHrs, UTMins, UTSec)
	geoLongE := 360 - geoLongW
	fmt.Printf("\nra : %d %d %f\ndec : %f %f %f\n", raHrs, raMins, raSecs, decDeg, decMin, decSec)
	UTrHrs, UTrMin, UTrSec, UTsHrs, UTsMin, UTsSec := coords.CalculateRisingAndSettingTime(GDay, GMonth, GYear, int(raHrs), int(raMins), raSecs, decDeg, decMin, decSec, geoLatN, geoLongE, refractionInArcMin)
	_, _, _, riseHrs, riseMin, riseSec = datetime.ConvertUniversalTimeToLocalTime(GDay, int(GMonth), int(GYear), int(UTrHrs), int(UTrMin), UTrSec, int(daylightsavingHrs), int(daylightsavingMin), timeZone)
	_, _, _, SetHrs, SetMin, SetSec = datetime.ConvertUniversalTimeToLocalTime(GDay, int(GMonth), int(GYear), int(UTsHrs), int(UTsMin), UTsSec, int(daylightsavingHrs), int(daylightsavingMin), timeZone)
	fmt.Printf("\nraDecimal : %f\ndecDeg : %f\n", datetime.ConvertHrsMinSecToDecimalHrs(int(raHrs), int(raMins), raSecs, false, false), macros.ConvertDegMinSecToDecimalDeg(decDeg, decMin, decSec))
	return riseHrs, riseMin, riseSec, SetHrs, SetMin, SetSec
}

