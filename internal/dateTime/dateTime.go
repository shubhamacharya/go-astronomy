package datetime

import "math"

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
	A := float64(Hrs) + ((float64(min) + (sec / 60)) / 60)
	if is12HrClock && AMPM == "PM" {
		return A + 12.0
	}
	return A
}
