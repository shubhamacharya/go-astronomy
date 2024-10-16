package sun

import (
	"go-astronomy/internal/coords"
	datetime "go-astronomy/internal/dateTime"
	"go-astronomy/internal/macros"
	"math"
)

// Values at epoch 2010.0
const semiMajorAxis = 1.495985e8
const angularDiameter = 0.533128

func CalculatePositionOfSun(GDay float64, GMonth, GYear int, UTHrs, UTMins int, UTSec float64, epochDay float64, epochMonth, epochYear int) (raHrs, raMin int, raSec float64, decDeg, decMin int, decSec, lambda float64) {
	lambda = macros.CalculatePositionOfSunHelper(GDay, GMonth, GYear, UTHrs, UTMins, UTSec, epochDay, epochMonth, epochYear)
	lambdaDeg, lambdaMin, lambdaSec := macros.ConvertDecimalDegToDegMinSec(lambda)
	raHrs, raMin, raSec, decDeg, decMin, decSec = macros.ConvertEclipticCoordinatesToEquatorial(GDay, GMonth, GYear, lambdaDeg, lambdaMin, lambdaSec, 0, 0, 0, epochDay, epochMonth, epochYear)
	// fmt.Printf("\nlambda : %d %d %f\nra : %d %d %f\ndec : %d %d %f\n", lambdaDeg, lambdaMin, lambdaSec, raHrs, raMin, raSec, decDeg, decMin, decSec)
	return raHrs, raMin, raSec, decDeg, decMin, decSec, lambda
}

func CalculatePrecisePositionOfSun(GDay float64, GMonth, GYear int, UTHrs, UTMins int, UTSec float64, epochDay float64, epochMonth, epochYear int) (raHrs, raMins int, raSecs float64, decDeg, decMin int, decSec, lambda0 float64) {
	Eg, Wg, e, _, _ := macros.CalculateEgWgAnde(GDay, GMonth, GYear, UTHrs, UTMins, UTSec, epochDay, epochMonth, epochYear)
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
	raHrs, raMins, raSecs, decDeg, decMin, decSec = macros.ConvertEclipticCoordinatesToEquatorial(GDay, GMonth, GYear, lambda0Deg, lambda0Min, lambda0Sec, 0, 0, 0, epochDay, epochMonth, epochYear)
	return raHrs, raMins, raSecs, decDeg, decMin, decSec, lambda0
}

func CalculateSunsDistanceAndAngularSize(GDay float64, GMonth, GYear int, UTHrs, UTMins int, UTSec float64, epochDay float64, epochMonth, epochYear int) (r float64, thetaDeg int, thetaMin int, thetaSec, theta float64) {
	_, Wg, e, _, _ := macros.CalculateEgWgAnde(GDay, GMonth, GYear, UTHrs, UTMins, UTSec, epochDay, epochMonth, epochYear)
	_, _, _, _, _, _, lambda0 := CalculatePrecisePositionOfSun(GDay, GMonth, GYear, UTHrs, UTMins, UTSec, epochDay, epochMonth, epochYear)
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

func CalculateSunsRiseAndSet(GDay float64, GMonth, GYear int, UTHrs, UTMins int, UTSec, geoLongW, geoLatN, refractionInArcMin, daylightsavingHrs, daylightsavingMin, timeZone float64, epochDay float64, epochMonth, epochYear int) (riseHrs, riseMin int, riseSec float64, SetHrs, SetMin int, SetSec float64) {
	const verticalShiftOfUpperLimb float64 = 0.008333
	raHrs, raMins, raSecs, decDeg, decMin, decSec, _ := CalculatePositionOfSun(GDay, GMonth, GYear, UTHrs, UTMins, UTSec, epochDay, epochMonth, epochYear)
	// fmt.Printf("\nra : %f\ndec :  %f\n", datetime.ConvertHrsMinSecToDecimalHrs(raHrs, raMins, raSecs, false, false), macros.ConvertDegMinSecToDecimalDeg(decDeg, decMin, decSec))
	verticalShiftHrs, verticalShiftMin, verticalShiftSec := datetime.ConvertDecimalHrsToHrsMinSec(macros.ConvertDecimalDegressToDecimalHrs(verticalShiftOfUpperLimb))
	UTrHrs, UTrMin, UTrSec, UTsHrs, UTsMin, UTsSec, _, _ := coords.CalculateRisingAndSettingTime(GDay, GMonth, GYear, raHrs, raMins, raSecs, decDeg, decMin, decSec, geoLatN, geoLongW, refractionInArcMin)
	// fmt.Printf("\nUTriseHrs : %d\nUTriseMin : %d\nUTriseSec : %f\nUTSetHrs : %d\nUTsetMin : %d\nUTsetSec : %f\n", UTrHrs, UTrMin, UTrSec, UTsHrs, UTsMin, UTsSec)

	_, _, _, riseHrs, riseMin, riseSec = datetime.ConvertUniversalTimeToLocalTime(GDay, GMonth, GYear, UTrHrs, UTrMin, UTrSec, int(daylightsavingHrs), int(daylightsavingMin), timeZone)
	_, _, _, SetHrs, SetMin, SetSec = datetime.ConvertUniversalTimeToLocalTime(GDay, GMonth, GYear, UTsHrs, UTsMin, UTsSec, int(daylightsavingHrs), int(daylightsavingMin), timeZone)
	// fmt.Printf("\nriseHrs : %d\nriseMin : %d\nriseSec : %f\nSetHrs : %d\nsetMin : %d\nsetSec : %f\n", riseHrs, riseMin, riseSec, SetHrs, SetMin, SetSec)
	return riseHrs + verticalShiftHrs, riseMin + verticalShiftMin, riseSec + verticalShiftSec, SetHrs + verticalShiftHrs, SetMin + verticalShiftMin, SetSec + verticalShiftSec
}

func CalculateCalculateSunTwilight(GDay float64, GMonth, GYear int, UTHrs, UTMins int, UTSec, geoLongW, geoLatN, refractionInArcMin, daylightsavingHrs, daylightsavingMin, timeZone, epochDay float64, epochMonth, epochYear int) (int, int, float64, int, int, float64) {
	_, _, _, decDeg, decMin, decSec, _ := CalculatePositionOfSun(GDay, GMonth, GYear, UTHrs, UTMins, UTSec, epochDay, epochMonth, epochYear)
	lambda := macros.ConvertDegMinSecToDecimalDeg(decDeg, decMin, decSec)
	riseHrs, riseMin, riseSec, SetHrs, SetMin, SetSec := CalculateSunsRiseAndSet(GDay, GMonth, GYear, UTHrs, UTMins, UTSec, geoLongW, geoLatN, refractionInArcMin, daylightsavingHrs, daylightsavingMin, timeZone, epochDay, epochMonth, epochYear)
	hourAngle := math.Acos(-math.Tan(macros.ConvertDegreesToRadiance(geoLatN)) * math.Tan(macros.ConvertDegreesToRadiance(lambda)))
	hourAngleInv := math.Acos((math.Cos(macros.ConvertDegreesToRadiance(108)) - (math.Sin(macros.ConvertDegreesToRadiance(geoLatN)) * math.Sin(macros.ConvertDegreesToRadiance(lambda)))) / (math.Cos(macros.ConvertDegreesToRadiance(geoLatN)) * math.Cos(macros.ConvertDegreesToRadiance(lambda))))
	tUTDecimalHrs := (macros.ConvertRadianceToDegree(hourAngleInv-hourAngle) / 15) * 0.9973
	tUTHrs, tUTMin, tUTSec := datetime.ConvertDecimalHrsToHrsMinSec(tUTDecimalHrs)
	// fmt.Printf("\nlambda : %f\nhourAngle : %f\nhourAngleInv : %f\ntUTDecimalHrs : %f\n", lambda, hourAngle, hourAngleInv, tUTDecimalHrs)
	return riseHrs - tUTHrs, riseMin - tUTMin, riseSec - tUTSec, SetHrs + tUTHrs, SetMin + tUTMin, SetSec + tUTSec
}

func CalculateTheEquationOfTime(GDay float64, GMonth, GYear int, UTHrs, UTMins int, UTSec, geoLongW, geoLatN, refractionInArcMin, daylightsavingHrs, daylightsavingMin, timeZone, epochDay float64, epochMonth, epochYear int) (eqHrs, eqMin int, eqSec float64) {
	raHrs, raMin, raSec, _, _, _, _ := CalculatePositionOfSun(GDay, GMonth, GYear, UTHrs, UTMins, UTSec, epochDay, epochMonth, epochYear)
	raUTHrs, raUTMin, raUTSec := datetime.ConvertGreenwichSiderealTimeToUniversalTime(GDay, GMonth, GYear, raHrs, raMin, raSec)
	eqHrs, eqMin, eqSec = datetime.ConvertDecimalHrsToHrsMinSec(datetime.ConvertHrsMinSecToDecimalHrs(raUTHrs, raUTMin, raUTSec, false, false) - 12)
	return eqHrs, eqMin, eqSec
}
