package macros

import (
	"fmt"
	datetime "go-astronomy/internal/dateTime"
	"math"
)

// roundToNDecimals rounds a float64 to n decimal places
func RoundToNDecimals(value float64, n int) float64 {
	factor := math.Pow(10, float64(n))
	return math.Round(value*factor) / factor
}

func IsLeapYear(year int) bool {
	return (year)%4 == 0 && year%100 != 0 || (year)%400 == 0
}

func AdjustAngleRange(angle float64, lowestVal, highestVal int) float64 {
	for angle > float64(highestVal) || angle < float64(lowestVal) {
		if angle > 360 {
			angle -= float64(highestVal)
		} else {
			angle += float64(highestVal)
		}
	}

	return angle
}

func CalculateEgWgAnde(GDay float64, GMonth, GYear int, UTHrs, UTMins int, UTSec float64) (Eg, Wg, e float64) {
	julianDate := datetime.ConvertGreenwichDateToJulianDate(GDay, int(GMonth), int(GYear))
	T := (julianDate - 2415020.0) / 36525
	Eg = AdjustAngleRange(279.6966778+(36000.76892*T)+(0.0003025*math.Pow(T, 2)), 0, 360)
	Wg = AdjustAngleRange(281.2208444+(1.719175*T)+(0.000452778*math.Pow(T, 2)), 0, 360)
	e = AdjustAngleRange(0.01675104-(0.0000418*T)-(0.000000126*math.Pow(T, 2)), 0, 360)

	return Eg, Wg, e
}

func DaysElapsedSinceEpoch(epochYear, targetYear int) float64 {
	days := 0.0
	isNegative := false
	if epochYear > targetYear {
		epochYear, targetYear = targetYear, epochYear
		isNegative = true
	}

	for year := epochYear; year < targetYear; year++ {
		if IsLeapYear(year) {
			days += 366
		} else {
			days += 365
		}
	}
	if isNegative {
		days *= -1
	}
	return days
}

func CalculatePositionOfSunHelper(GDay float64, GMonth, GYear, UTHrs, UTMins int, UTSec float64) float64 {
	daysElapsedSinceStartOfYear := datetime.CalculateDayNumber(GDay, GMonth, GYear)
	daysElapsedSinceEpoch := DaysElapsedSinceEpoch(2010, GYear)

	// Calculating epoch at 0h of Jan 2010
	Eg, Wg, e := CalculateEgWgAnde(31, 12, 2009, UTHrs, UTMins, UTSec)

	N := AdjustAngleRange((360/365.242191)*(daysElapsedSinceStartOfYear+daysElapsedSinceEpoch), 0, 360)
	M := AdjustAngleRange(N+Eg-Wg, 0, 360)
	Ec := (360 / 3.1415927) * e * math.Sin(M*(math.Pi/180)) // Convert M to radiance
	lambda := AdjustAngleRange(N+Ec+Eg, 0, 360)
	fmt.Printf("\ndaysElapsedSinceStartOfYear : %f\ndaysElapsedSinceEpoch : %f\nEg : %f\nWg : %f\ne : %f\nN : %f\nM : %f\nEc : %f\nlambda : %f\n", daysElapsedSinceStartOfYear, daysElapsedSinceStartOfYear+daysElapsedSinceEpoch, Eg, Wg, e, N, M, Ec, lambda)
	return lambda
}

func ConvertDecimalDegToDegMinSec(decimalDeg float64) (float64, float64, float64) {
	Deg, fractPart := math.Modf(math.Abs(decimalDeg))
	Mins, minFracts := math.Modf(fractPart * 60)
	Sec := minFracts * 60

	if Deg < 0 {
		return Deg, math.Abs(Mins), math.Abs(Sec)
	}
	return Deg, Mins, Sec
}

func CalculateEclipticMeanObliquity(Gday float64, GMonth, GYear int) (float64, float64, float64, float64) {
	julianDate := datetime.ConvertGreenwichDateToJulianDate(Gday, GMonth, GYear)
	timeElapsed := (julianDate - 2451545.0) / 36525.0
	meanObliquity := 23.439292 - (((46.815 * timeElapsed) + (0.0006 * math.Pow(timeElapsed, 2)) - (0.00181 * math.Pow(timeElapsed, 3))) / 3600)
	obliquityDeg, obliquityMin, obliquitySec := ConvertDecimalDegToDegMinSec(meanObliquity)

	return obliquityDeg, obliquityMin, obliquitySec, meanObliquity
}

func ConvertDegMinSecToDecimalDeg(deg, min, sec float64) float64 {
	decimalDeg := math.Abs(math.Round(deg)) + (min / 60) + (sec / 3600)

	if deg < 0 {
		return -decimalDeg
	} else {
		return decimalDeg
	}
}

func ConvertRadianceToDegree(radians float64) float64 {
	return radians * (180 / math.Pi)
}

func ConvertDegreesToRadiance(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

func ConvertDecimalDegressToDecimalHrs(decimalDeg float64) float64 {
	const degreesPerHour = 15.0
	return decimalDeg / degreesPerHour
}

func ConvertEclipticCoordinatesToEquatorial(day float64, month, year int, eclipticLongDeg, eclipticLongMin, eclipticLongSec, eclipticLatDeg, eclipticLatMin, eclipticLatSec float64) (raHrs, raMins int, raSecs float64, decDeg float64, decMin float64, decSec float64) {
	_, _, _, meanObliquity := CalculateEclipticMeanObliquity(day, month, year)
	eclipticLongDecimalDeg := ConvertDegMinSecToDecimalDeg(eclipticLongDeg, eclipticLongMin, eclipticLongSec)

	eclipticLatDecimalDeg := ConvertDegMinSecToDecimalDeg(eclipticLatDeg, eclipticLatMin, eclipticLatSec)

	decDecimalDeg := ConvertRadianceToDegree(math.Asin((math.Sin(ConvertDegreesToRadiance(eclipticLatDecimalDeg)) * math.Cos(ConvertDegreesToRadiance(meanObliquity))) + (math.Cos(ConvertDegreesToRadiance(eclipticLatDecimalDeg)) * math.Sin(ConvertDegreesToRadiance(meanObliquity)) * math.Sin(ConvertDegreesToRadiance(eclipticLongDecimalDeg)))))

	y := (math.Sin(ConvertDegreesToRadiance(eclipticLongDecimalDeg)) * math.Cos(ConvertDegreesToRadiance(meanObliquity))) - (math.Tan(ConvertDegreesToRadiance(eclipticLatDecimalDeg)) * math.Sin(ConvertDegreesToRadiance(meanObliquity)))

	x := math.Cos(ConvertDegreesToRadiance(eclipticLongDecimalDeg))

	raDeg := ConvertRadianceToDegree(math.Atan2(ConvertDegreesToRadiance(y), ConvertDegreesToRadiance(x)))

	raDecimalHrs := ConvertDecimalDegressToDecimalHrs(raDeg)

	decDeg, decMin, decSec = ConvertDecimalDegToDegMinSec(decDecimalDeg)
	raHrs, raMins, raSecs = datetime.ConvertDecimalHrsToHrsMinSec(raDecimalHrs)

	return raHrs, raMins, raSecs, decDeg, decMin, decSec
}

func CalculatePositionOfSun(GDay float64, GMonth, GYear, UTHrs, UTMins int, UTSec float64) (raHrs, raMin int, raSec, decDeg, decMin, decSec, lambda float64) {

	lambda = CalculatePositionOfSunHelper(GDay, GMonth, GYear, UTHrs, UTMins, UTSec)
	lambdaDeg, lambdaMin, lambdaSec := ConvertDecimalDegToDegMinSec(lambda)
	raHrs, raMin, raSec, decDeg, decMin, decSec = ConvertEclipticCoordinatesToEquatorial(GDay, int(GMonth), int(GYear), lambdaDeg, lambdaMin, lambdaSec, 0, 0, 0)
	return raHrs, raMin, raSec, decDeg, decMin, decSec, lambda
}

func AdjustAngleInQuadrant(x, y, A float64) float64 {
	// Check the signs of x and y to determine the quadrant
	if x < 0 && y > 0 {
		// Second quadrant
		A = 180 - A
	} else if x < 0 && y < 0 {
		// Third quadrant
		A = 180 + A
	} else if x > 0 && y < 0 {
		// Fourth quadrant
		A = 360 - A
	}
	// Normalize the angle to be within [0, 360) if needed
	if A < 0 {
		A += 360
	} else if A >= 360 {
		A -= 360
	}

	return A
}

func CalculateEccentricAnomaly(M, e float64) float64 {
	E := M // Initial guess for E is M

	for {
		delta := E - (e * math.Sin(E)) - M

		deltaE := delta / (1 - (e * math.Cos(E)))

		E -= deltaE

		if math.Abs(deltaE) < 1e-6 {
			break
		}
	}

	return E
}
