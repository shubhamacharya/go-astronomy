package macros

import (
	datetime "go-astronomy/internal/dateTime"
	"math"
)

// Days in each month (non-leap year)
var daysInMonth = map[int]int{
	1: 31, 2: 28, 3: 31, 4: 30,
	5: 31, 6: 30, 7: 31, 8: 31,
	9: 30, 10: 31, 11: 30, 12: 31,
}

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
		if angle > float64(highestVal) {
			angle -= float64(highestVal)
		} else {
			angle += float64(highestVal)
		}
	}
	return angle
}

func CalculateEgWgAnde(GDay, GMonth, GYear float64, yearLabel datetime.YearLabel, UTHrs, UTMins int, UTSec float64, epochDay, epochMonth, epochYear float64) (Eg, Wg, e, r0, theta0 float64) {
	julianDate := datetime.ConvertGreenwichDateToJulianDate(GDay, GMonth, GYear, yearLabel)
	julianDateEpoch := RoundToNDecimals(datetime.ConvertGreenwichDateToJulianDate(epochDay, epochMonth, epochYear, yearLabel), 6)

	T := RoundToNDecimals(((julianDate - 2415020.0) / 36525), 6)
	Tepoch := RoundToNDecimals((julianDateEpoch-2415020.0)/36525, 6)

	r0 = 1.495985e8 // 1 AU in kilometers
	d := 1392000.0  // Diameter of the Sun in km
	Eg = RoundToNDecimals(AdjustAngleRange(279.6966778+(36000.76892*Tepoch)+(0.0003025*math.Pow(Tepoch, 2)), 0, 360), 6)
	Wg = RoundToNDecimals(AdjustAngleRange(281.2208444+(1.719175*T)+(0.000452778*math.Pow(T, 2)), 0, 360), 6)
	e = RoundToNDecimals(0.01675104-(0.0000418*Tepoch)-(0.000000126*math.Pow(Tepoch, 2)), 6)
	theta0 = RoundToNDecimals((d/r0)*(180/math.Pi), 6)

	return Eg, Wg, e, r0, theta0
}
func CalculateL_E_V(Eg, Wg, e float64) (eccentricAnomaly, V, lambda0 float64) {
	MRad := ConvertDegreesToRadiance(AdjustAngleRange(Eg-Wg, 0, 360))
	eccentricAnomaly = ConvertRadianceToDegree(CalculateEccentricAnomaly(MRad, e))
	V = ConvertRadianceToDegree(math.Atan(math.Sqrt((1+e)/(1-e))*math.Tan(ConvertDegreesToRadiance(eccentricAnomaly/2))) * 2)
	if V < 0 {
		V += 360
	}

	lambda0 = V + Wg
	if lambda0 > 360 {
		lambda0 -= 360
	}
	return eccentricAnomaly, V, lambda0
}

func DaysElapsedSinceEpoch(epochYear, targetYear float64) float64 {
	days := 0.0
	isNegative := false
	if epochYear > targetYear {
		epochYear, targetYear = targetYear, epochYear
		isNegative = true
	}

	for year := epochYear; year < targetYear; year++ {
		if IsLeapYear(int(year)) {
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

func CalculatePositionOfSunHelper(GDay, GMonth, GYear float64, yearLabel datetime.YearLabel, UTHrs, UTMins, UTSec float64, epochDay, epochMonth, epochYear float64) float64 {
	daysElapsedSinceStartOfYear := datetime.CalculateDayNumber(GDay, GMonth, GYear)
	daysElapsedSinceEpoch := DaysElapsedSinceEpoch(epochYear, GYear)

	// Calculating epoch at 0h of Jan 2010
	// _, _, adjustedYear := AdjustDate(epochDay, epochMonth, epochYear)
	Eg, Wg, e, _, _ := CalculateEgWgAnde(GDay, GMonth, GYear, yearLabel, 0, 0, 0, epochDay, epochMonth, epochYear)

	N := RoundToNDecimals((360/365.242191)*(daysElapsedSinceStartOfYear+daysElapsedSinceEpoch), 6)
	M := RoundToNDecimals((N + Eg - Wg), 6)
	Ec := RoundToNDecimals((360 * e * math.Sin(ConvertDegreesToRadiance(M)) / RoundToNDecimals(math.Pi, 7)), 6) // Convert M to radiance
	lambda := RoundToNDecimals((N + Ec + Eg), 6)
	// fmt.Printf("\ndaysElapsedSinceStartOfYear : %v\ndaysElapsedSinceEpoch : %v\ndaysElapsedSinceEpochDifference : %v\nEg : %v\nWg : %v\ne : %v\nr0 : %v\ntheta : %v\nN : %v\nM : %v\nEc : %v\nlambda : %v\n", daysElapsedSinceStartOfYear, daysElapsedSinceEpoch, daysElapsedSinceStartOfYear+daysElapsedSinceEpoch, Eg, Wg, e, r0, theta, N, M, Ec, lambda)
	return lambda
}

func CalculateN_M_V(planetValues map[string]interface{}, totalDays float64) (N, M, V float64) {
	N = RoundToNDecimals(AdjustAngleRange((360/365.242191)*(totalDays/planetValues["Tp"].(float64)), 0, 360), 6)
	M = RoundToNDecimals(N+planetValues["Long"].(float64)-planetValues["Peri"].(float64), 6)
	V = AdjustAngleRange(RoundToNDecimals(M+((360/math.Pi)*planetValues["Ecc"].(float64)*math.Sin(ConvertDegreesToRadiance(M))), 6), 0, 360)

	if V < 0 {
		V += 360
	}
	return N, M, V
}

func CalculatePerturbationsInPlanetsOrbitHelper(planetValues map[string]interface{}, V, deltaL float64) (L, r, si float64) {
	L = RoundToNDecimals(AdjustAngleRange(V+planetValues["Peri"].(float64), 0, 360)+deltaL, 6)
	r = RoundToNDecimals((planetValues["Axis"].(float64)*(1-math.Pow(planetValues["Ecc"].(float64), 2)))/(1+(planetValues["Ecc"].(float64)*math.Cos(ConvertDegreesToRadiance(V)))), 6)
	si = RoundToNDecimals(ConvertRadianceToDegree(math.Asin(math.Sin(ConvertDegreesToRadiance(L-planetValues["Node"].(float64)))*(math.Sin(ConvertDegreesToRadiance(planetValues["Incl"].(float64)))))), 6)
	return L, r, si
}

func AdjustDate(day float64, month, year int) (float64, int, int) {
	if day > 1 {
		day--
	} else {
		month--
		if month < 1 {
			month = 12
			year--
		}
		if month == 2 && IsLeapYear(year) {
			day = 29
		} else {
			day = float64(daysInMonth[month])
		}
	}
	return day, month, year
}

func ConvertDecimalDegToDegMinSec(decimalDeg float64) (deg, min, sec float64) {
	// Handle negative values by storing the absolute value and adjusting the sign later
	decimalDegAbs := math.Abs(decimalDeg)
	Deg, fractPart := math.Modf(decimalDegAbs)
	Mins, minFracts := math.Modf(fractPart * 60)
	sec = minFracts * 60

	// Convert fractional part to int
	deg = math.Round(Deg)
	min = math.Round(Mins)

	// Apply sign to degrees
	if decimalDeg < 0 {
		deg = -deg
	}

	return
}

func CalculateEclipticMeanObliquity(Gday, GMonth, GYear float64, yearLabel datetime.YearLabel) (obliquityDeg, obliquityMin, obliquitySec, meanObliquity float64) {
	julianDate := datetime.ConvertGreenwichDateToJulianDate(Gday, GMonth, GYear, yearLabel)
	timeElapsed := RoundToNDecimals(((julianDate - 2451545.0) / 36525.0), 6)
	meanObliquity = RoundToNDecimals(23.439292-(((46.815*timeElapsed)-(0.0006*math.Pow(timeElapsed, 2))+(0.00181*math.Pow(timeElapsed, 3)))/3600), 6)
	obliquityDeg, obliquityMin, obliquitySec = ConvertDecimalDegToDegMinSec(meanObliquity)

	return obliquityDeg, obliquityMin, obliquitySec, meanObliquity
}

func ConvertDegMinSecToDecimalDeg(deg, min, sec float64) float64 {
	decimalDeg := RoundToNDecimals(math.Abs(float64(deg))+float64(min)/60+(sec/3600), 6)
	if deg < 0 {
		return -decimalDeg
	}
	return decimalDeg
}

func ConvertRadianceToDegree(radians float64) float64 {
	return radians * (180 / math.Pi)
}

func ConvertDegreesToRadiance(degrees float64) float64 {
	return RoundToNDecimals(degrees*(math.Pi/180), 6)
}

func ConvertDecimalDegressToDecimalHrs(decimalDeg float64) float64 {
	const degreesPerHour = 15.0
	return decimalDeg / degreesPerHour
}

func ConvertEclipticCoordinatesToEquatorial(day, month, year float64, yearLabel datetime.YearLabel, eclipticLongDeg, eclipticLongMin, eclipticLongSec, eclipticLatDeg, eclipticLatMin, eclipticLatSec float64) (raHrs, raMins, raSecs, decDeg, decMin, decSec float64) {
	_, _, _, meanObliquity := CalculateEclipticMeanObliquity(day, month, year, yearLabel)
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

// func ConvertEclipticCoordinatesToEquatorial(day float64, month, year, eclipticLongDeg, eclipticLongMin int, eclipticLongSec float64, eclipticLatDeg, eclipticLatMin int, eclipticLatSec float64, epochDay float64, epochMonth, epochYear int) (raHrs int, raMins int, raSecs float64, decDeg int, decMin int, decSec float64) {
// 	_, _, _, meanObliquity := CalculateEclipticMeanObliquity(epochDay, epochMonth, epochYear)
// 	eclipticLongDecimalDeg := ConvertDegMinSecToDecimalDeg(eclipticLongDeg, eclipticLongMin, eclipticLongSec)
// 	eclipticLatDecimalDeg := ConvertDegMinSecToDecimalDeg(eclipticLatDeg, eclipticLatMin, eclipticLatSec)

// 	// Convert meanObliquity to radians
// 	meanObliquityRad := ConvertDegreesToRadiance(meanObliquity)

// 	// Convert ecliptic longitude and latitude to radians
// 	eclipticLongRad := ConvertDegreesToRadiance(eclipticLongDecimalDeg)
// 	eclipticLatRad := ConvertDegreesToRadiance(eclipticLatDecimalDeg)

// 	// Calculate Declination
// 	decRad := math.Asin((math.Sin(eclipticLatRad) * math.Cos(meanObliquityRad)) + (math.Cos(eclipticLatRad) * math.Sin(meanObliquityRad) * math.Sin(eclipticLongRad)))
// 	decDecimalDeg := ConvertRadianceToDegree(decRad)
// 	decDeg, decMin, decSec = ConvertDecimalDegToDegMinSec(decDecimalDeg)

// 	// Calculate Right Ascension
// 	y := (math.Sin(eclipticLongRad) * math.Cos(meanObliquityRad)) - (math.Tan(eclipticLatRad) * math.Sin(meanObliquityRad))
// 	x := math.Cos(eclipticLongRad)
// 	raRad := math.Atan(y / x)
// 	raDeg := AdjustAngleInQuadrant(x, y, ConvertRadianceToDegree(raRad))
// 	raDecimalHrs := raDeg / 15.0                                       // Convert degrees to hours
// 	raHrs, raMins, raSecs = ConvertDecimalDegToDegMinSec(raDecimalHrs) // Convert back to hours, minutes, and seconds
// 	// fmt.Printf("\nmeanObliquity : %f\neclipticLongDecimalDeg : %f\neclipticLatDecimalDeg : %f\ndecRad : %f\ndecDecimalDeg : %f\nx : %f\ny : %f\nraDeg : %f\nraDecimalHrs : %f\n", meanObliquity, eclipticLongDecimalDeg, eclipticLatDecimalDeg, decRad, decDecimalDeg, x, y, raDeg, raDecimalHrs)

// 	return raHrs, raMins, raSecs, decDeg, decMin, decSec
// }

// func CalculatePositionOfSun(GDay float64, GMonth, GYear, UTHrs, UTMins int, UTSec float64, epochDay float64, epochMonth, epochYear int) (raHrs, raMin int, raSec float64, decDeg, decMin int, decSec, lambda float64) {
// 	lambda = CalculatePositionOfSunHelper(GDay, GMonth, GYear, UTHrs, UTMins, UTSec, epochDay, epochMonth, epochYear)
// 	lambdaDeg, lambdaMin, lambdaSec := ConvertDecimalDegToDegMinSec(lambda)
// 	raHrs, raMin, raSec, decDeg, decMin, decSec = ConvertEclipticCoordinatesToEquatorial(GDay, int(GMonth), int(GYear), lambdaDeg, lambdaMin, lambdaSec, 0, 0, 0, epochDay, epochMonth, epochYear)
// 	return raHrs, raMin, raSec, decDeg, decMin, decSec, lambda
// }

func AdjustAngleInQuadrant(x, y, A float64) float64 {
	// Check the signs of x and y to determine the quadrant
	if x < 0 && y > 0 {
		// Second quadrant
		if A < 0 {
			A = 180 + A
		} else {
			A = 180 - A
		}
	} else if x < 0 && y < 0 {
		// Third quadrant
		if A < 0 {
			A = 180 + A
		} else {
			A = 180 + A
		}
	} else if x > 0 && y < 0 {
		// Fourth quadrant
		if A < 0 {
			A = 360 + A
		} else {
			A = 360 - A
		}
	} else {
		// First quadrant
		if A < 0 {
			A = A
		} else {
			A = A
		}
	}

	// Normalize the angle to be within [0, 360)
	if A < 0 {
		A += 360
	} else if A >= 360 {
		A -= 360
	}

	return A
}

func CalculateEccentricAnomaly(M, e float64) (Erad float64) {
	Erad = M // Initial guess for E is M

	for {
		delta := Erad - (e * math.Sin(Erad)) - M

		deltaE := delta / (1 - (e * math.Cos(Erad)))

		Erad -= deltaE

		if math.Abs(deltaE) < 1e-6 {
			break
		}
	}

	return Erad
}
