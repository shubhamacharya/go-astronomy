package macros

import (
	datetime "go-astronomy/internal/dateTime"
	"math"
)

// roundToNDecimals rounds a float64 to n decimal places
func RoundToNDecimals(value float64, n int) float64 {
	factor := math.Pow(10, float64(n))
	return math.Round(value*factor) / factor
}

func IsLeapYear(year float64) bool {
	return (int(year)%4 == 0 && int(year)%100 != 0) || (int(year)%400 == 0)
}

func CalculatePositionOfSunHelper(GDay, GMonth, GYear float64, UTHrs, UTMins int, UTSec float64) float64 {
	daysElapsedSinceStartOfYear := datetime.CalculateDayNumber(GDay, GMonth, GYear)
	daysElapsedSinceEpoch := 0.0
	Eg := 279.557208
	Wg := 283.112438
	e := 0.016705

	if GYear < 2010.0 {
		for itr := GYear; itr < 2010.0; itr++ {
			if IsLeapYear(itr) {
				daysElapsedSinceEpoch += 366
			} else {
				daysElapsedSinceEpoch += 365
			}
		}
	} else {
		for itr := 2010.0; itr > GYear; itr++ {
			if IsLeapYear(itr) {
				daysElapsedSinceEpoch += 366
			} else {
				daysElapsedSinceEpoch += 365
			}
		}
	}
	N := (360 / 365.242191) * (daysElapsedSinceStartOfYear - daysElapsedSinceEpoch)
	for N < 0 || N > 360 {
		if N < 0 {
			N += 360
		} else {
			N -= 360
		}
	}
	M := N + Eg - Wg
	if M < 0 {
		M += 360
	}
	Ec := (360 / 3.1415927) * e * math.Sin(M*(math.Pi/180)) // Convert M to radiance
	lambda := N + Ec + Eg
	if lambda > 360 {
		lambda -= 360
	}
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

func CalculatePositionOfSun(GDay, GMonth, GYear float64, UTHrs, UTMins int, UTSec float64) (float64, float64, float64, float64, float64, float64, float64) {

	lambda := CalculatePositionOfSunHelper(GDay, GMonth, GYear, UTHrs, UTMins, UTSec)
	lambdaDeg, lambdaMin, lambdaSec := ConvertDecimalDegToDegMinSec(lambda)
	raHrs, raMin, raSec, decDeg, decMin, decSec := ConvertEclipticCoordinatesToEquatorial(GDay, int(GMonth), int(GYear), lambdaDeg, lambdaMin, lambdaSec, 0, 0, 0)
	return raHrs, raMin, raSec, decDeg, decMin, decSec, lambda
}
