package datetime

import (
	"math"
)

var daysOfWeek = [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

type YearLabel int

const (
	AD  YearLabel = iota // 0
	BC  YearLabel = 1    // 1
	CE  YearLabel = 2    // 0
	BCE YearLabel = 3    // 1
)

func (c YearLabel) String(year float64) float64 {
	switch c {
	case AD:
		return year
	case BC:
		if year > 0 {
			return (year - 1) * -1
		} else if year < 0 {
			return (year + 1) * -1
		}
		return 0
	case CE:
		return year
	case BCE:
		if year > 0 {
			return (year - 1) * -1
		} else if year < 0 {
			return (year + 1) * -1
		}
		return 0
	default:
		return year
	}
}

func CalculateDateOfEaster(year float64) (day, month float64) {
	// Calculate the date of Easter Sunday for a given year.
	// The algorithm used is the "Anonymous Gregorian algorithm" (Meeus/Jones/Butcher algorithm).
	// Returns the day and month as integers.

	// Calculate the "Golden Number" for the year
	goldenNumber := int(year) % 19

	// Century and related calculations
	century := year / 100
	leapYearCorrection := century / 4
	centuryRemainder := int(century) % 4

	// Corrections for leap years and the moon cycle
	leapYearAdjustment := (century + 8) / 25
	moonCycleCorrection := (century - leapYearAdjustment + 1) / 3

	// Calculate the epact (age of the moon on January 1st)
	epact := int(19*float64(goldenNumber)+century-leapYearCorrection-moonCycleCorrection+15) % 30

	// Additional corrections for lunar and solar cycles
	yearRemainder := int(year) % 100
	quarterCentury := yearRemainder / 4
	yearModFour := yearRemainder % 4

	// Calculate the "Dominical Number" (corresponding Sunday)
	dominicalNumber := (32 + 2*centuryRemainder + 2*quarterCentury - epact - yearModFour) % 7

	// Determine the month and day of Easter
	monthOffset := (goldenNumber + 11*epact + 22*dominicalNumber) / 451
	month = float64(epact+dominicalNumber-7*monthOffset+114) / 31
	day = float64(int(epact+dominicalNumber-7*monthOffset+114)%31 + 1)

	return day, month
}

func IsLeapYear(inputYear float64) bool {
	// Check if the given year is a leap year based on the rules of the Gregorian calendar.
	return (int(inputYear)%4 == 0 && int(inputYear)%100 != 0) || (int(inputYear)%400 == 0)
}

func CalculateDayNumber(day, month, year float64) float64 {
	var dayNumber float64
	isLeapYear := IsLeapYear(year)
	if month > 2 {
		month++
		dayNumber = math.Floor(float64(month) * 30.6)
		if isLeapYear {
			dayNumber -= 62
		} else {
			dayNumber -= 63
		}
	} else {
		month--
		dayNumber = math.Floor(float64(month) * 63 * 0.5)
		if isLeapYear {
			dayNumber = math.Floor(float64(month) * 62 * 0.5)
		}
	}

	dayNumber += day
	return dayNumber
}

func ConvertGreenwichDateToJulianDate(day, month, year float64, yearLabel YearLabel) float64 {
	var a, c, y, m float64
	year = yearLabel.String(year)
	if month <= 2 {
		y = year - 1
		m = month + 12
	} else {
		y = year
		m = month
	}

	a = math.Trunc(y / 100)
	b := 0.0
	if year > 1582 || (year == 1582 && month > 10) || (year == 1582 && month == 10 && day >= 15) {
		b = 2 - a + math.Trunc(a/4)
	}

	if year < 0 {
		c = math.Trunc((365.25 * y) - 0.75)
	} else {
		c = math.Trunc(365.25 * y)
	}
	d := math.Trunc(30.6001 * (m + 1))
	jd := b + c + d + day + 1720994.5
	// jd := (365.2500 * (y + 4716)) + (30.6001 * (m + 1)) + day + b - 1524.5000

	return jd
}

func ConvertJulianDateToGreenwichDate(julianDate float64) (day float64, month, year int) {
	// Adjust the Julian date to start the calculation
	julianDate += 0.5
	integerPart, fractionalPart := math.Modf(julianDate)

	// Initialize variables
	var correctionFactor, adjustedJulianDate float64
	if integerPart > 2299160 {
		correctionFactor = math.Trunc((integerPart - 1867216.25) / 36524.25)
		adjustedJulianDate = integerPart + correctionFactor - math.Trunc(correctionFactor/4) + 1
	} else {
		adjustedJulianDate = integerPart
	}

	// Perform the conversion to the Greenwich date
	calendarDate := adjustedJulianDate + 1524
	yearEstimate := math.Trunc((calendarDate - 122.1) / 365.25)
	dayOfYear := math.Trunc(365.25 * yearEstimate)
	monthEstimate := math.Trunc((calendarDate - dayOfYear) / 30.6001)

	// Calculate day, month, and year
	day = calendarDate - dayOfYear + fractionalPart - math.Trunc(30.6001*monthEstimate)

	if monthEstimate < 13.5 {
		month = int(monthEstimate - 1.0)
	} else {
		month = int(monthEstimate - 13)
	}

	if float64(month) > 2.5 {
		year = int(yearEstimate - 4716)
	} else {
		year = int(yearEstimate - 4715)
	}

	return day, month, year
}

func GetNameOfTheDayOfMonth(day, month, year float64, yearLabel YearLabel) string {
	// Convert the Greenwich date to a Julian Date Number
	julianDate := ConvertGreenwichDateToJulianDate(day, month, year, yearLabel)

	// Calculate the day of the week (0 = Monday, 6 = Sunday)
	dayOfWeek := int(math.Mod(julianDate+1.5, 7))

	// Return the corresponding name of the day of the week
	return daysOfWeek[dayOfWeek]
}

func ConvertHrsMinSecToDecimalHrs(Hrs int, min int, sec float64, is12HrClock bool, isPM bool) float64 {
	// Convert seconds to a fractional part of a minute
	secPart := sec / 60.0

	// Convert minutes and the seconds part to a fractional part of an hour
	minPart := (float64(min) + secPart) / 60.0

	// Calculate total hours
	A := float64(Hrs) + minPart

	// Adjust for 12-hour clock format
	if is12HrClock {
		// Handle the case for 12 AM and 12 PM specifically
		if isPM {
			if Hrs != 12 {
				A += 12.0 // Add 12 hours for PM times, except for 12 PM itself
			}
		} else {
			if Hrs == 12 {
				A = minPart // 12 AM is 0 hours
			}
		}
	}
	factor := math.Pow(10, float64(6))
	return math.Round(A*factor) / factor
}

func ConvertDecimalHrsToHrsMinSec(decimalHours float64) (hours, minutes int, seconds float64) {
	// Split decimal hours into the integer part (hours) and fractional part (fractionalHours)
	hoursFloat, fractionalHours := math.Modf(decimalHours)
	hours = int(hoursFloat)

	// Convert fractional hours to minutes
	minutesFloat, fractionalMinutes := math.Modf(fractionalHours * 60)

	// Convert fractional minutes to seconds
	factor := math.Pow(10, float64(6))
	seconds = math.Round((fractionalMinutes*60)*factor) / factor

	if math.Round(seconds) == 60 {
		seconds = 0
		minutesFloat += 1
	}

	if math.Round(minutesFloat) == 60 {
		minutesFloat = 0
		hours += 1
	}

	minutes = int(minutesFloat)
	return hours, minutes, seconds
}

func ConvertLocalTimeToUniversalTime(day, month, year float64, yearLabel YearLabel, hrs int, min int, sec float64, daylightsavingHrs int, daylightsavingMin int, zoneOffset float64) (UTDay float64, UTMonth, UTYear, UTHrs, UTMin int, UTSec, decimalTime float64) {
	// Adjust for daylight saving time
	hrs -= daylightsavingHrs
	min -= daylightsavingMin

	// Convert to decimal hours
	decimalHrs := ConvertHrsMinSecToDecimalHrs(hrs, min, sec, false, false)

	// Correct time zone adjustment by subtracting the zone offset
	UT := decimalHrs - zoneOffset

	// Adjust the Greenwich calendar day based on the UT
	Gday := (UT / 24) + day

	// Calculate Julian Date from Greenwich calendar day
	julianDate := ConvertGreenwichDateToJulianDate(Gday, month, year, yearLabel)

	// Convert Julian Date back to Greenwich calendar date
	UTDay, UTMonth, UTYear = ConvertJulianDateToGreenwichDate(julianDate)

	decimalUTTime := (Gday - math.Trunc(Gday)) * 24

	UTHrs, UTMin, UTSec = ConvertDecimalHrsToHrsMinSec(decimalUTTime)

	// Handle leap second case: if seconds are exactly 60.0, adjust it to 59.999999 without rolling over
	if sec >= 60.0 {
		UTSec = 59.999999
	} else {
		// Handle small floating-point precision issues
		UTSec = math.Round(UTSec*1e6) / 1e6
	}

	// Truncate UTDay to get the whole day number
	UTDay = math.Trunc(UTDay)

	// Return the correct UT values
	return UTDay, UTMonth, UTYear, UTHrs, UTMin, UTSec, decimalTime
}

func ConvertUniversalTimeToLocalTime(day, month, year float64, yearLabel YearLabel, hrs int, min int, sec float64, daylightsavingHrs int, daylightsavingMin int, zoneOffset float64) (Gday float64, calMonth, calYear, GHrs, GMin int, GSec float64) {
	factor := math.Pow(10, float64(6))
	decimalHrs := ConvertHrsMinSecToDecimalHrs(hrs, min, sec, false, false) + zoneOffset + float64(daylightsavingHrs) + float64(daylightsavingMin)
	decimalHrs = math.Round(decimalHrs*factor) / factor

	julianDate := ConvertGreenwichDateToJulianDate(day, month, year, yearLabel) + (decimalHrs / 24)
	julianDate = math.Round(julianDate*factor) / factor

	calDay, calMonth, calYear := ConvertJulianDateToGreenwichDate(julianDate)
	calDay = math.Round(calDay*factor) / factor

	Gday, GTime := math.Modf(calDay)
	Gday = math.Round(Gday*factor) / factor
	GTime = math.Round(GTime*factor) / factor

	GHrs, GMin, GSec = ConvertDecimalHrsToHrsMinSec(GTime * 24)

	return Gday, calMonth, calYear, GHrs, GMin, GSec
}

func ConvertUniversalTimeToGreenwichSiderealTime(day, month, year float64, yearLabel YearLabel, hrs int, min int, sec float64) (GSTHrs, GSTMin int, GSTSec, gst float64) {
	julianDate := ConvertGreenwichDateToJulianDate(day, month, year, yearLabel)
	elapsedDays := julianDate - 2451545.0
	centuriesSinceJ2000 := elapsedDays / 36525.0
	gstAtZeroUT := 6.697374558 + (2400.051336 * centuriesSinceJ2000) + (0.000025862 * math.Pow(centuriesSinceJ2000, 2))

	// Normalize GST to the range [0, 24) hours
	for gstAtZeroUT < 0 {
		gstAtZeroUT += 24
	}
	for gstAtZeroUT >= 24 {
		gstAtZeroUT -= 24
	}

	utInDecimalHours := ConvertHrsMinSecToDecimalHrs(hrs, min, sec, false, false) * 1.002737909
	gst = gstAtZeroUT + utInDecimalHours

	// Normalize GST to the range [0, 24) hours again after adding UT
	for gst < 0 {
		gst += 24
	}
	for gst >= 24 {
		gst -= 24
	}

	GSTHrs, GSTMin, GSTSec = ConvertDecimalHrsToHrsMinSec(gst)

	return GSTHrs, GSTMin, GSTSec, gst
}

func ConvertGreenwichSiderealTimeToUniversalTime(day, month, year float64, yearLabel YearLabel, hrs int, min int, sec float64) (UTHrs, UTMin int, UTSec float64) {
	julianDate := ConvertGreenwichDateToJulianDate(day, month, year, yearLabel)
	centuriesSinceJ2000 := ((julianDate - 2451545.0) / 36525.0)
	factor := math.Pow(10, float64(6))
	centuriesSinceJ2000 = math.Round(centuriesSinceJ2000*factor) / factor
	gstAtZeroUT := 6.697374558 + (2400.051336 * centuriesSinceJ2000) + (0.000025862 * math.Pow(centuriesSinceJ2000, 2))
	gstAtZeroUT = math.Round(gstAtZeroUT*factor) / factor

	// Normalize GST to the range [0, 24) hours
	for gstAtZeroUT < 0 {
		gstAtZeroUT += 24
	}
	for gstAtZeroUT >= 24 {
		gstAtZeroUT -= 24
	}

	gstInDecimalHours := ConvertHrsMinSecToDecimalHrs(hrs, min, sec, false, false)
	utInDecimalHours := gstInDecimalHours - gstAtZeroUT

	// Normalize UT to the range [0, 24) hours
	for utInDecimalHours < 0 {
		utInDecimalHours += 24
	}
	for utInDecimalHours >= 24 {
		utInDecimalHours -= 24
	}

	utInDecimalHours *= 0.9972695663

	UTHrs, UTMin, UTSec = ConvertDecimalHrsToHrsMinSec(utInDecimalHours)

	return UTHrs, UTMin, UTSec
}

func CalculateLocalSiderealTimeUsingGreenwichSiderealTime(hours, minutes int, seconds, geoLongitude float64) (LSTHours, LSTMinutes int, LSTSeconds, decimalLST float64) {
	// Convert Greenwich Sidereal Time to decimal hours
	decimalGST := ConvertHrsMinSecToDecimalHrs(hours, minutes, seconds, false, false)

	// Adjust for geographical longitude (in degrees)
	decimalLST = decimalGST + (geoLongitude / 15)

	// Normalize the Local Sidereal Time to the range [0, 24) hours
	for decimalLST < 0 {
		decimalLST += 24
	}
	for decimalLST >= 24 {
		decimalLST -= 24
	}

	// Convert decimal Local Sidereal Time back to hours, minutes, and seconds
	LSTHours, LSTMinutes, LSTSeconds = ConvertDecimalHrsToHrsMinSec(decimalLST)

	return LSTHours, LSTMinutes, LSTSeconds, decimalLST
}

func CalculateGreenwichSiderealTimeUsingLocalSiderealTime(hours, minutes int, seconds, geoLongitude float64) (GSTHours, GSTMinutes int, GSTSeconds, decimalGST float64) {
	// Convert Local Sidereal Time to decimal hours
	decimalLST := ConvertHrsMinSecToDecimalHrs(hours, minutes, seconds, false, false)

	// Adjust for geographical longitude (in degrees)
	decimalGST = decimalLST - (geoLongitude / 15)

	// Normalize the Greenwich Sidereal Time to the range [0, 24) hours
	for decimalGST < 0 {
		decimalGST += 24
	}
	for decimalGST >= 24 {
		decimalGST -= 24
	}

	// Convert decimal Greenwich Sidereal Time back to hours, minutes, and seconds
	GSTHours, GSTMinutes, GSTSeconds = ConvertDecimalHrsToHrsMinSec(decimalGST)

	return GSTHours, GSTMinutes, GSTSeconds, decimalGST
}
