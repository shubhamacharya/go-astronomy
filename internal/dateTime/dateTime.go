package datetime

import (
	"fmt"
	"math"
)

var daysOfWeek = [...]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

func CalculateDateOfEaster(year int) (int, int) {
	// CalculateDateOfEaster calculates the date of Easter Sunday for a given year.
	// The algorithm used is the "Anonymous Gregorian algorithm" also known as the "Meeus/Jones/Butcher algorithm".
	// It returns the day and month as integers.

	// Calculate the "Golden Number" for the year
	goldenNumber := year % 19

	// Century and related calculations
	century := year / 100
	centuryDiv4 := century / 4
	centuryMod4 := century % 4

	// Corrections for leap years and moon cycle
	correctionLeapYears := (century + 8) / 25
	correctionMoonCycle := (century - correctionLeapYears + 1) / 3

	// Calculate the epact (age of the moon on January 1st)
	epact := (19*goldenNumber + century - centuryDiv4 - correctionMoonCycle + 15) % 30

	// Additional corrections for lunar and solar cycles
	quarterCenturyDiv := year % 100 / 4
	quarterCenturyMod := year % 100 % 4

	// Calculate the "Dominical Number" (corresponding Sunday)
	dominicalNumber := (32 + 2*centuryMod4 + 2*quarterCenturyDiv - epact - quarterCenturyMod) % 7

	// Determine the month and day of Easter
	monthOffset := (goldenNumber + 11*epact + 22*dominicalNumber) / 451
	month := (epact + dominicalNumber - 7*monthOffset + 114) / 31
	day := (epact+dominicalNumber-7*monthOffset+114)%31 + 1

	return day, month
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

func ConvertJulianDateToGreenwichDate(julianDate float64) (float64, float64, float64) {
	julianDate += 0.5
	intPart, fracPart := math.Modf(julianDate)
	A := 0.0
	B := intPart
	if intPart > 2299160 {
		A = math.Trunc((intPart - 1867216.25) / 36524.25)
		B = intPart + A - math.Trunc(A/4) + 1
	}

	C := B + 1524
	D := math.Trunc((C - 122.1) / 365.25)
	E := math.Trunc(365.25 * D)
	G := math.Trunc((C - E) / 30.6001)
	d := C - E + fracPart - math.Trunc(30.6001*G)
	m := 0.0
	y := 0.0
	if G < 13.5 {
		m = G - 1
	} else {
		m = G - 13
	}

	if m > 2.5 {
		y = D - 4716
	} else {
		y = D - 4715
	}

	return d, m, y
}

func GetNameOfTheDayOfMonth(day float64, month int, year int) string {
	julianDate := ConvertGreenwichDateToJulianDate(day, month, year)
	A := ((julianDate + 1.5) / 7)
	_, fractOfA := math.Modf(A)
	return daysOfWeek[int(math.Round(fractOfA*7))]
}

func ConvertHrsMinSecToDecimalHrs(Hrs int, min int, sec float64, is12HrClock bool, AMPM string) float64 {
	// Convert seconds to a fractional part of a minute
	secPart := sec / 60.0

	// Convert minutes and the seconds part to a fractional part of an hour
	minPart := (float64(min) + secPart) / 60.0

	// Calculate total hours
	A := float64(Hrs) + minPart

	// Adjust for 12-hour clock format
	if is12HrClock {
		// Handle the case for 12 AM and 12 PM specifically
		if AMPM == "AM" && Hrs == 12 {
			A = minPart // 12 AM is 0 hours
		} else if AMPM == "PM" && Hrs != 12 {
			A += 12.0 // Add 12 hours for PM times, except for 12 PM itself
		}
	}

	return A
}

func ConvertDecimalHrsToHrsMinSec(decimalHrs float64) (float64, float64, float64) {
	hrs, fractPart := math.Modf(decimalHrs)
	min, minFract := math.Modf(fractPart * 60)
	sec := (math.Round(minFract*60) * 10000) / 10000
	return hrs, min, sec
}

func ConvertLocalTimeToUniversalTime(day int, month int, year int, hrs int, min int, sec float64, daylightsavingHrs int, daylightsavingMin int, zoneOffset float64) (float64, float64, float64, float64, float64, float64) {
	// Adjust for daylight saving time
	hrs -= daylightsavingHrs
	min -= daylightsavingMin

	// Normalize the time if minutes or hours are negative
	if min < 0 {
		hrs -= 1
		min += 60
	}
	if hrs < 0 {
		day -= 1
		hrs += 24
	}

	// Convert to decimal hours
	decimalHrs := ConvertHrsMinSecToDecimalHrs(hrs, min, sec, false, "")

	// Adjust for the time zone offset
	UT := decimalHrs - zoneOffset

	// If UT is negative, it means we are in the previous day
	if UT < 0 {
		day -= 1
		UT += 24
	}
	// If UT exceeds 24 hours, it means we are in the next day
	if UT >= 24 {
		day += 1
		UT -= 24
	}

	// Adjust the Greenwich calendar day based on the UT
	Gday := float64(day) + (UT / 24.0)

	// Calculate Julian Date from Greenwich calendar day
	julianDate := ConvertGreenwichDateToJulianDate(Gday, month, year)

	// Convert Julian Date back to Greenwich calendar date
	UTDay, UTMonth, UTYear := ConvertJulianDateToGreenwichDate(julianDate)

	// Extract time from UT
	decimalTime := (UTDay - math.Trunc(UTDay)) * 24
	UTHrs, UTMin, UTSec := ConvertDecimalHrsToHrsMinSec(decimalTime)
	UTDay = math.Trunc(UTDay) // Truncate UTDay to get the whole day number

	return UTDay, UTMonth, UTYear, UTHrs, UTMin, UTSec
}

func ConvertUniversalTimeToLocalTime(day float64, month int, year int, hrs int, min int, sec float64, daylightsavingHrs int, daylightsavingMin int, zoneOffset float64) (float64, float64, float64, float64, float64, float64) {
	decimalHrs := ConvertHrsMinSecToDecimalHrs(hrs, min, sec, false, "") + zoneOffset + float64(daylightsavingHrs) + float64(daylightsavingMin)
	julianDate := ConvertGreenwichDateToJulianDate(day, month, year) + (decimalHrs / 24)
	calDay, calMonth, calYear := ConvertJulianDateToGreenwichDate(julianDate)

	Gday, GTime := math.Modf(calDay)

	GHrs, GMin, GSec := ConvertDecimalHrsToHrsMinSec(GTime * 24)

	return Gday, calMonth, calYear, GHrs, GMin, GSec
}

func ConvertUniversalTimeToGreenwichSiderealTime(day float64, month int, year int, hrs int, min int, sec float64) (float64, float64, float64) {
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

	utInDecimalHours := ConvertHrsMinSecToDecimalHrs(hrs, min, sec, false, "") * 1.002737909
	gst := gstAtZeroUT + utInDecimalHours

	// Normalize GST to the range [0, 24) hours again after adding UT
	for gst < 0 {
		gst += 24
	}
	for gst >= 24 {
		gst -= 24
	}

	GSTHrs, GSTMin, GSTSec := ConvertDecimalHrsToHrsMinSec(gst)

	return GSTHrs, GSTMin, GSTSec
}

func ConvertGreenwichSiderealTimeToUniversalTime(day float64, month int, year int, hrs int, min int, sec float64) (float64, float64, float64) {
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

	gstInDecimalHours := ConvertHrsMinSecToDecimalHrs(hrs, min, sec, false, "")
	utInDecimalHours := gstInDecimalHours - gstAtZeroUT

	// Normalize UT to the range [0, 24) hours
	for utInDecimalHours < 0 {
		utInDecimalHours += 24
	}
	for utInDecimalHours >= 24 {
		utInDecimalHours -= 24
	}

	utInDecimalHours *= 0.9972695663

	UTHrs, UTMin, UTSec := ConvertDecimalHrsToHrsMinSec(utInDecimalHours)

	return UTHrs, UTMin, UTSec
}

func CalculateLocalSiderealTimeUsingGreenwichSideralTime(hrs int, min int, sec float64, geoLong float64) (float64, float64, float64) {
	// TODO: Add Logngitude directions
	decimalTime := ConvertHrsMinSecToDecimalHrs(hrs, min, sec, false, "")
	if geoLong < 0 {
		decimalTime += (geoLong / 15)

		for decimalTime < 0 {
			decimalTime += 24
		}
		for decimalTime >= 24 {
			decimalTime -= 24
		}

		LSTHrs, LSTMin, LSTSec := ConvertDecimalHrsToHrsMinSec(decimalTime)
		return LSTHrs, LSTMin, LSTSec

	} else {
		// Return error : Geo Longitude must be negative
		fmt.Println("\nGeo Longitude must be negative")
		panic("\nGeo Longitude must be negative\n")
	}
	// return 0.0, 0.0, 0.0
}
