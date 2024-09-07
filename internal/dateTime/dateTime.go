package datetime

import (
	"math"
)

var daysOfWeek = [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

func CalculateDateOfEaster(year int) (day, month int) {
	// Calculate the date of Easter Sunday for a given year.
	// The algorithm used is the "Anonymous Gregorian algorithm" (Meeus/Jones/Butcher algorithm).
	// Returns the day and month as integers.

	// Calculate the "Golden Number" for the year
	goldenNumber := year % 19

	// Century and related calculations
	century := year / 100
	leapYearCorrection := century / 4
	centuryRemainder := century % 4

	// Corrections for leap years and the moon cycle
	leapYearAdjustment := (century + 8) / 25
	moonCycleCorrection := (century - leapYearAdjustment + 1) / 3

	// Calculate the epact (age of the moon on January 1st)
	epact := (19*goldenNumber + century - leapYearCorrection - moonCycleCorrection + 15) % 30

	// Additional corrections for lunar and solar cycles
	yearRemainder := year % 100
	quarterCentury := yearRemainder / 4
	yearModFour := yearRemainder % 4

	// Calculate the "Dominical Number" (corresponding Sunday)
	dominicalNumber := (32 + 2*centuryRemainder + 2*quarterCentury - epact - yearModFour) % 7

	// Determine the month and day of Easter
	monthOffset := (goldenNumber + 11*epact + 22*dominicalNumber) / 451
	month = (epact + dominicalNumber - 7*monthOffset + 114) / 31
	day = (epact+dominicalNumber-7*monthOffset+114)%31 + 1

	return day, month
}

func IsLeapYear(inputYear float64) bool {
	// Check if the given year is a leap year based on the rules of the Gregorian calendar.
	return (int(inputYear)%4 == 0 && int(inputYear)%100 != 0) || (int(inputYear)%400 == 0)
}

func CalculateDayNumber(day, month, year float64) float64 {
	var dayNumber float64
	isLeapYear := IsLeapYear(year)

	// Days in each month for a non-leap year
	daysInMonth := []float64{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// If it's a leap year, adjust February's days
	if isLeapYear {
		daysInMonth[1] = 29
	}

	if month > 2 {
		// For months after February, calculate day number based on the month and day
		dayNumber = day
		for i := 0; i < int(month)-1; i++ {
			dayNumber += daysInMonth[i]
		}
	} else {
		// For January and February, directly calculate the day number
		dayNumber = day
		for i := 0; i < int(month)-1; i++ {
			dayNumber += daysInMonth[i]
		}
	}

	return dayNumber - 1 // Subtract 1 to make the first day of the year as day 0
}

func ConvertGreenwichDateToJulianDate(day float64, month, year int) float64 {
	var correctionFactor, daysInYear, daysInMonth float64

	// Adjust for January and February as the 13th and 14th months of the previous year
	if month == 1 || month == 2 {
		year -= 1
		month += 12
	}

	// Calculate century correction for dates after October 15, 1582
	if year > 1582 || (year == 1582 && (month > 10 || (month == 10 && day >= 15))) {
		century := float64(year) / 100.0
		correctionFactor = 2 - math.Floor(century) + math.Floor(century/4.0)
	} else {
		correctionFactor = 0
	}

	// Calculate the days from the years
	daysInYear = math.Floor(365.25 * float64(year))

	// Calculate the days from the months
	daysInMonth = math.Floor(30.6001 * float64(month+1))

	// Calculate the Julian date
	julianDate := correctionFactor + daysInYear + daysInMonth + float64(day) + 1720994.5

	return julianDate
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

func GetNameOfTheDayOfMonth(day float64, month int, year int) string {
	// Convert the Greenwich date to a Julian Date Number
	julianDate := ConvertGreenwichDateToJulianDate(day, month, year)

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

	return A
}

func ConvertDecimalHrsToHrsMinSec(decimalHours float64) (hours, minutes int, seconds float64) {
	// Split decimal hours into the integer part (hours) and fractional part (fractionalHours)
	hoursFloat, fractionalHours := math.Modf(decimalHours)
	hours = int(hoursFloat)

	// Convert fractional hours to minutes
	minutesFloat, fractionalMinutes := math.Modf(fractionalHours * 60)

	// Convert fractional minutes to seconds
	seconds = fractionalMinutes * 60

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

func ConvertLocalTimeToUniversalTime(day float64, month int, year int, hrs int, min int, sec float64, daylightsavingHrs int, daylightsavingMin int, zoneOffset float64) (UTDay float64, UTMonth, UTYear, UTHrs, UTMin int, UTSec, decimalTime float64) {
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
	julianDate := ConvertGreenwichDateToJulianDate(Gday, month, year)

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

func ConvertUniversalTimeToLocalTime(day float64, month int, year int, hrs int, min int, sec float64, daylightsavingHrs int, daylightsavingMin int, zoneOffset float64) (Gday float64, calMonth, calYear, GHrs, GMin int, GSec float64) {
	decimalHrs := ConvertHrsMinSecToDecimalHrs(hrs, min, sec, false, false) + zoneOffset + float64(daylightsavingHrs) + float64(daylightsavingMin)
	julianDate := ConvertGreenwichDateToJulianDate(day, month, year) + (decimalHrs / 24)
	calDay, calMonth, calYear := ConvertJulianDateToGreenwichDate(julianDate)

	Gday, GTime := math.Modf(calDay)

	GHrs, GMin, GSec = ConvertDecimalHrsToHrsMinSec(GTime * 24.0)

	return Gday, calMonth, calYear, GHrs, GMin, GSec
}

func ConvertUniversalTimeToGreenwichSiderealTime(day float64, month int, year int, hrs int, min int, sec float64) (GSTHrs, GSTMin int, GSTSec, gst float64) {
	julianDate := ConvertGreenwichDateToJulianDate(day, month, year)
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

func ConvertGreenwichSiderealTimeToUniversalTime(day float64, month int, year int, hrs int, min int, sec float64) (UTHrs, UTMin int, UTSec float64) {
	julianDate := ConvertGreenwichDateToJulianDate(day, month, year)
	centuriesSinceJ2000 := (julianDate - 2451545.0) / 36525.0
	gstAtZeroUT := 6.697374558 + (2400.051336 * centuriesSinceJ2000) + (0.000025862 * math.Pow(centuriesSinceJ2000, 2))

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
