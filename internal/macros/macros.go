package macros

import (
	"math"
	"strings"
)

func ConvertGregorianToJulian(day int, month int, year int) float64 {
	/*
		Convert a Gregorian Date (day, month, year) to Julian Date
		Original macro name: GERG2JULI
	*/
	var adjustedYear, adjustedMonth int
	var centuryAdjustment, leapYearAdjustment, dayOfYear, monthAdjustment float64

	if month < 3 {
		adjustedYear = year - 1
		adjustedMonth = month + 12
	} else {
		adjustedYear = year
		adjustedMonth = month
	}

	if year > 1582 || (year == 1582 && (month > 10 || (month == 10 && day >= 15))) {
		centuryAdjustment = math.Floor(float64(adjustedYear) / 100)
		leapYearAdjustment = 2 - centuryAdjustment + math.Floor(centuryAdjustment/4)
	} else {
		leapYearAdjustment = 0
	}

	if adjustedYear < 0 {
		dayOfYear = math.Floor(365.25*float64(adjustedYear)) - 0.75
	} else {
		dayOfYear = math.Floor(365.25 * float64(adjustedYear))
	}

	monthAdjustment = math.Floor(30.6001 * (float64(adjustedMonth) + 1))

	return leapYearAdjustment + dayOfYear + monthAdjustment + float64(day) + 1720994.5
}

func GetJulianDay(julianDate float64) int {
	/*
		Returns the day part of a Julian Date
	*/
	integerPart := math.Floor(julianDate + 0.5)
	fractionalPart := julianDate + 0.5 - integerPart

	var adjustedInteger float64
	if integerPart > 2299160 {
		centuryAdjustment := math.Floor((integerPart - 1867216.25) / 36524.25)
		adjustedInteger = integerPart + 1 + centuryAdjustment - math.Floor(centuryAdjustment/4)
	} else {
		adjustedInteger = integerPart
	}

	correctedDate := adjustedInteger + 1524
	century := math.Floor((correctedDate - 122.1) / 365.25)
	dayOfYear := math.Floor(365.25 * century)
	monthPart := math.Floor((correctedDate - dayOfYear) / 30.6001)

	return int(correctedDate - dayOfYear + fractionalPart - math.Floor(30.6001*monthPart))
}

func GetJulianMonth(julianDate float64) int {
	/*
		Returns the month part of a Julian Date
	*/
	var adjustedInteger float64

	integerPart := math.Floor(julianDate + 0.5)
	centuryAdjustment := math.Floor((integerPart - 1867216.25) / 36524.25)
	if integerPart > 2299160 {
		adjustedInteger = integerPart + 1 + centuryAdjustment - math.Floor(centuryAdjustment/4)
	} else {
		adjustedInteger = integerPart
	}

	correctedDate := adjustedInteger + 1524
	century := math.Floor((correctedDate - 122.1) / 365.25)
	dayOfYear := math.Floor(365.25 * century)
	monthPart := math.Floor((correctedDate - dayOfYear) / 30.6001)

	if monthPart < 13.5 {
		return int(monthPart - 1)
	} else {
		return int(monthPart - 13)
	}
}

func GetJulianYear(julianDate float64) int {
	/*
		Returns the year part of a Julian Date
	*/
	var monthPart, year float64

	integerPart := math.Floor(julianDate + 0.5)
	centuryAdjustment := math.Floor((integerPart - 1867216.25) / 36524.25)
	if integerPart > 2299160 {
		monthPart = integerPart + 1 + centuryAdjustment - math.Floor(centuryAdjustment/4)
	} else {
		monthPart = integerPart
	}

	correctedDate := monthPart + 1524
	century := math.Floor((correctedDate - 122.1) / 365.25)
	dayOfYear := math.Floor(365.25 * century)
	monthPart = math.Floor((correctedDate - dayOfYear) / 30.6001)

	if monthPart < 13.5 {
		year = monthPart - 1
	} else {
		year = monthPart - 13
	}

	if year > 2.5 {
		return int(century - 4716)
	} else {
		return int(century - 4715)
	}
}

// GetDayOfWeekFromJulian converts a Julian Date to a Day-of-Week (e.g., Sunday).
func GetDayOfWeekFromJulian(julianDate float64) string {
	// Calculate the Julian Day Number (JDN) by flooring the input Julian Date.
	julianDayNumber := math.Floor(julianDate-0.5) + 0.5

	// Calculate the day of the week (0=Sunday, 1=Monday, ..., 6=Saturday).
	dayOfWeekIndex := math.Mod(julianDayNumber+1.5, 7)

	// Return the corresponding day of the week as a string.
	switch dayOfWeekIndex {
	case 0:
		return "Sunday"
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	default:
		return "Unknown"
	}
}

// ConvertTimeToDecimal converts a civil time (hours, minutes, seconds) to decimal hours.
func ConvertTimeToDecimal(hours int, minutes int, seconds int) float64 {
	// Calculate the fractional minutes from seconds.
	fractionalMinutes := math.Abs(float64(seconds)) / 60
	// Calculate the fractional hours from minutes and fractional minutes.
	fractionalHours := (math.Abs(float64(minutes)) + fractionalMinutes) / 60
	// Calculate the total decimal hours.
	decimalHours := math.Abs(float64(hours)) + fractionalHours

	// If any of the time components are negative, return the negative decimal hours.
	if hours < 0 || minutes < 0 || seconds < 0 {
		return -decimalHours
	} else {
		return decimalHours
	}
}

// GetHourFromDecimalHour returns the hour part of decimal hours.
func GetHourFromDecimalHour(decimalHours float64) int {
	// Take the absolute value of decimal hours.
	absoluteDecimalHours := math.Abs(decimalHours)
	// Convert decimal hours to total seconds.
	totalSeconds := absoluteDecimalHours * 3600

	// If totalSeconds modulo 3600 is 60, adjust the values.
	if math.Mod(totalSeconds, 3600) == 60 {
		totalSeconds += 60
	}

	// Calculate the hour part.
	hourPart := int(math.Floor(totalSeconds / 3600))

	// Adjust for negative input.
	if decimalHours < 0 {
		return -hourPart
	}
	return hourPart
}

// GetMinutesFromDecimalHours returns the minutes part of decimal hours.
func GetMinutesFromDecimalHours(decimalHours float64) int {
	// Take the absolute value of decimal hours.
	absoluteDecimalHours := math.Abs(decimalHours)
	// Convert decimal hours to total seconds.
	totalSeconds := absoluteDecimalHours * 3600

	// Calculate the minute part.
	minutesPart := math.Mod(math.Floor(totalSeconds/60), 60)

	return int(minutesPart)
}

// GetSecondsFromDecimalHours returns the seconds part of decimal hours.
func GetSecondsFromDecimalHours(decimalHours float64) int {
	// Take the absolute value of decimal hours.
	absoluteDecimalHours := math.Abs(decimalHours)
	// Convert decimal hours to total seconds.
	totalSeconds := absoluteDecimalHours * 3600
	// Calculate the minute part.
	remainingSeconds := math.Mod(totalSeconds, 60)

	return int(math.Round(remainingSeconds))
}

// ConvertLocalTimeToUTC converts local time to UTC and returns the time as an integer representing the number of hours.
func ConvertLocalTimeToUTC(localHours, localMinutes, localSeconds int, daylightSavingOffset, zoneCorrection float64, localDay, localMonth, localYear int) int {
	// Convert local time to decimal hours
	decimalHours := ConvertTimeToDecimal(localHours, localMinutes, localSeconds)

	// Adjust the decimal hours by subtracting daylight saving and zone correction
	adjustedDecimalHours := decimalHours - daylightSavingOffset - zoneCorrection

	// Calculate the Julian Day Number
	dayOfYear := float64(localDay) + (adjustedDecimalHours / 24)
	julianDay := ConvertGregorianToJulian(int(dayOfYear), localMonth, localYear)
	julianDayNumber := GetJulianDay(julianDay)

	// Calculate the fractional part of the Julian Day Number and convert to hours
	fractionalDay := julianDayNumber - int(math.Floor(julianDay))
	hours := int(24 * fractionalDay)

	return hours
}

// convertUTCToLocalTime converts UTC time to local time and returns the time as an integer representing the number of hours.
func ConvertUTCToLocalTime(utcHours, utcMinutes, utcSeconds int, daylightSavingOffset, zoneCorrection float64, greenwichDay, greenwichMonth, greenwichYear int) int {
	// Convert UTC time to decimal hours
	decimalUTC := ConvertTimeToDecimal(utcHours, utcMinutes, utcSeconds)

	// Adjust the decimal hours by adding zone correction and daylight saving
	adjustedDecimalHours := decimalUTC + zoneCorrection + daylightSavingOffset

	// Calculate the Julian Day Number for the given Greenwich date
	julianDayNumber := ConvertGregorianToJulian(greenwichDay, greenwichMonth, greenwichYear)

	// Adjust Julian Day Number by adding the fractional part from the decimal hours
	adjustedJulianDayNumber := julianDayNumber + (adjustedDecimalHours / 24)

	// Get the fractional part of the Julian Day Number and convert it to hours
	fractionalDay := adjustedJulianDayNumber - math.Floor(adjustedJulianDayNumber)
	localHours := int(24 * fractionalDay)

	return localHours
}

// GetLocalCivilDayForUT calculates the local civil day for a given Universal Time (UT).
func GetLocalCivilDayForUT(utHours, utMinutes, utSeconds int, daylightSavingOffset, zoneCorrection float64, greenwichDay, greenwichMonth, greenwichYear int) int {
	// Convert UTC time to decimal hours
	decimalUT := ConvertTimeToDecimal(utHours, utMinutes, utSeconds)

	// Adjust for zone correction and daylight saving
	adjustedDecimalTime := decimalUT + zoneCorrection + daylightSavingOffset

	// Calculate the Julian Day Number for the given Greenwich date and adjust by the decimal time
	greekJulianDay := ConvertGregorianToJulian(greenwichDay, greenwichMonth, greenwichYear)
	adjustedJulianDay := greekJulianDay + (adjustedDecimalTime / 24)

	// Return the Julian Day Number, rounded down to the nearest whole number
	localCivilDay := int(math.Floor(adjustedJulianDay))

	return localCivilDay
}

// GetLocalCivilMonthForUT calculates the local civil Month for a given Universal Time (UT).
func GetLocalCivilMonthForUT(utHours, utMinutes, utSeconds int, daylightSavingOffset, zoneCorrection float64, greenwichDay, greenwichMonth, greenwichYear int) int {
	// Convert UTC time to decimal hours
	decimalUT := ConvertTimeToDecimal(utHours, utMinutes, utSeconds)

	// Adjust for zone correction and daylight saving
	adjustedDecimalTime := decimalUT + zoneCorrection + daylightSavingOffset

	// Calculate the Julian Day Number for the given Greenwich date and adjust by the decimal time
	greekJulianDay := ConvertGregorianToJulian(greenwichDay, greenwichMonth, greenwichYear)
	adjustedJulianDay := greekJulianDay + (adjustedDecimalTime / 24)

	// Return the Julian Month Number, rounded down to the nearest whole number
	return GetJulianMonth(adjustedJulianDay)
}

// GetLocalCivilYearForUT calculates the local civil Year for a given Universal Time (UT).
func GetLocalCivilYearForUT(utHours, utMinutes, utSeconds int, daylightSavingOffset, zoneCorrection float64, greenwichDay, greenwichMonth, greenwichYear int) int {
	// Convert UTC time to decimal hours
	decimalUT := ConvertTimeToDecimal(utHours, utMinutes, utSeconds)

	// Adjust for zone correction and daylight saving
	adjustedDecimalTime := decimalUT + zoneCorrection + daylightSavingOffset

	// Calculate the Julian Day Number for the given Greenwich date and adjust by the decimal time
	greekJulianDay := ConvertGregorianToJulian(greenwichDay, greenwichMonth, greenwichYear)
	adjustedJulianDay := greekJulianDay + (adjustedDecimalTime / 24)

	// Return the Julian Month Number, rounded down to the nearest whole number
	return GetJulianYear(adjustedJulianDay)
}

// ComputeGreenwichDayForLT calculates the Greenwich Day for a given Local Time.
func ComputeGreenwichDayForLT(localHours, localMinutes, localSeconds int, daylightSavingOffset, zoneCorrection float64, localDay, localMonth, localYear int) int {
	// Convert local time to decimal hours
	decimalLocalTime := ConvertTimeToDecimal(localHours, localMinutes, localSeconds)

	// Adjust for zone correction and daylight saving to get UTC time in decimal hours
	decimalUT := decimalLocalTime - zoneCorrection - daylightSavingOffset

	// Compute the day of the year in UTC time
	utcDayOfYear := localDay + int(math.Floor(decimalUT/24))

	// Convert the UTC day of the year to Julian date
	greekJulianDay := ConvertGregorianToJulian(utcDayOfYear, localMonth, localYear)

	// Compute the Julian Day Number (already returned by ConvertGregorianToJulian)
	greekJulianDayNumber := GetJulianDay(greekJulianDay)

	return greekJulianDayNumber
}

// ComputeGreenwichMonthForLT calculates the Greenwich Month for a given Local Time.
func ComputeGreenwichMonthForLT(localHours, localMinutes, localSeconds int, daylightSavingOffset, zoneCorrection float64, localDay, localMonth, localYear int) int {
	// Convert local time to decimal hours
	decimalLocalTime := ConvertTimeToDecimal(localHours, localMinutes, localSeconds)

	// Adjust for zone correction and daylight saving to get UTC time in decimal hours
	decimalUT := decimalLocalTime - zoneCorrection - daylightSavingOffset

	// Compute the day of the year in UTC time
	utcDayOfYear := localDay + int(math.Floor(decimalUT/24))

	// Convert the UTC day of the year to Julian date
	greekJulianDay := ConvertGregorianToJulian(utcDayOfYear, localMonth, localYear)

	// Compute the Julian Month Number (already returned by ConvertGregorianToJulian)
	greekJulianMonthNumber := GetJulianMonth(greekJulianDay)

	return greekJulianMonthNumber
}

// ComputeGreenwichYearForLT calculates the Greenwich Year for a given Local Time.
func ComputeGreenwichYearForLT(localHours, localMinutes, localSeconds int, daylightSavingOffset, zoneCorrection float64, localDay, localMonth, localYear int) int {
	// Convert local time to decimal hours
	decimalLocalTime := ConvertTimeToDecimal(localHours, localMinutes, localSeconds)

	// Adjust for zone correction and daylight saving to get UTC time in decimal hours
	decimalUT := decimalLocalTime - zoneCorrection - daylightSavingOffset

	// Compute the day of the year in UTC time
	utcDayOfYear := localDay + int(math.Floor(decimalUT/24))

	// Convert the UTC day of the year to Julian date
	greekJulianDay := ConvertGregorianToJulian(utcDayOfYear, localMonth, localYear)

	// Compute the Julian Year Number (already returned by ConvertGregorianToJulian)
	greekJulianYearNumber := GetJulianYear(greekJulianDay)

	return greekJulianYearNumber
}

// ConvertUTToGST converts Universal Time (UT) to Greenwich Sidereal Time (GST).
func ConvertUTToGST(utHours, utMinutes, utSeconds, greenwichDay, greenwichMonth, greenwichYear int) float64 {
	// Convert the given Gregorian date to Julian Date
	julianDate := ConvertGregorianToJulian(greenwichDay, greenwichMonth, greenwichYear)

	// Compute Julian Centuries from the J2000 epoch
	julianCenturies := (julianDate - 2451545.0) / 36525.0

	// Calculate the Greenwich Sidereal Time in hours at 0h UT
	gstAtZeroUT := 6.697374558 + (2400.051336 * julianCenturies) + (0.000025862 * julianCenturies * julianCenturies)

	// Normalize GST to the range [0, 24) hours
	gstAtZeroUT = gstAtZeroUT - (24 * math.Floor(gstAtZeroUT/24.0))

	// Convert UT time to decimal hours
	decimalUT := ConvertTimeToDecimal(utHours, utMinutes, utSeconds)

	// Convert UT to GST by applying the sidereal time rate correction
	siderealTimeCorrection := decimalUT * 1.002737909

	// Compute the final Greenwich Sidereal Time
	gst := gstAtZeroUT + siderealTimeCorrection

	// Normalize GST to the range [0, 24) hours
	gst = gst - (24 * math.Floor(gst/24.0))

	return gst
}

// ConvertGSTToLST converts Greenwich Sidereal Time (GST) to Local Sidereal Time (LST).
func ConvertGSTToLST(greenwichHours, greenwichMinutes, greenwichSeconds int, geographicalLongitude float64) float64 {
	// Convert Greenwich Sidereal Time to decimal hours
	decimalGST := ConvertTimeToDecimal(greenwichHours, greenwichMinutes, greenwichSeconds)

	// Convert geographical longitude to sidereal hours
	longitudeInHours := geographicalLongitude / 15.0

	// Calculate Local Sidereal Time by adding longitude correction
	localSiderealTime := decimalGST + longitudeInHours

	// Normalize Local Sidereal Time to the range [0, 24) hours
	localSiderealTime = localSiderealTime - (24 * math.Floor(localSiderealTime/24.0))

	return localSiderealTime
}

// ConvertLSTToGST converts Local Sidereal Time (LST) to Greenwich Sidereal Time (GST).
func ConvertLSTToGST(localHours, localMinutes, localSeconds int, longitude float64) float64 {
	// Convert Local Sidereal Time to decimal hours
	decimalLST := ConvertTimeToDecimal(localHours, localMinutes, localSeconds)

	// Convert geographical longitude to sidereal hours
	longitudeInHours := longitude / 15.0

	// Calculate Greenwich Sidereal Time by subtracting longitude correction
	greenwichSiderealTime := decimalLST - longitudeInHours

	// Normalize Greenwich Sidereal Time to the range [0, 24) hours
	greenwichSiderealTime = greenwichSiderealTime - (24 * math.Floor(greenwichSiderealTime/24.0))

	return greenwichSiderealTime
}

// ConvertGSTToUT converts Greenwich Sidereal Time (GST) to Universal Time (UT).
func ConvertGSTToUT(gstHours, gstMinutes, gstSeconds, greenwichDay, greenwichMonth, greenwichYear int) (float64, string) {
	// Convert the given Gregorian date to Julian Date
	julianDate := ConvertGregorianToJulian(greenwichDay, greenwichMonth, greenwichYear)

	// Compute Julian Centuries from the J2000 epoch
	julianCenturies := (julianDate - 2451545.0) / 36525.0

	// Calculate the Greenwich Sidereal Time in hours at 0h UT
	gstAtZeroUT := 6.697374558 + (2400.051336 * julianCenturies) + (0.000025862 * julianCenturies * julianCenturies)

	// Normalize GST to the range [0, 24) hours
	gstAtZeroUT = gstAtZeroUT - (24 * math.Floor(gstAtZeroUT/24.0))

	// Convert GST time to decimal hours
	decimalGST := ConvertTimeToDecimal(gstHours, gstMinutes, gstSeconds)

	// Calculate the difference between the GST and GST at 0h UT
	decimalGSTDifference := decimalGST - gstAtZeroUT

	// Normalize the difference to the range [0, 24) hours
	normalizedGSTDifference := decimalGSTDifference - (24 * math.Floor(decimalGSTDifference/24.0))

	// Convert the normalized GST difference to UT by applying the sidereal time rate correction
	utTime := normalizedGSTDifference * 0.9972695663

	conversionStatus := GetStatusOfGSTToUTConversion(utTime)

	return utTime, conversionStatus
}

// GetStatusOfGSTToUTConversion check if conversion of Greenwich Sidereal Time to Universal Time is OK or having some error.
func GetStatusOfGSTToUTConversion(utTime float64) string {
	if utTime < (4 / float64(60)) {
		return "Warning"
	} else {
		return "OK"
	}
}

// RAtoHourAngle converts Right Ascension (RA) to Hour Angle (HA).
func RAtoHourAngle(raHours, raMinutes, raSeconds, ltHours, ltMinutes, ltSeconds int, daylightSavings, zoneCorrection, localDay, localMonth, localYear, geographicalLongitude float64) float64 {
	// Convert Local Time to UTC
	utcTime := ConvertLocalTimeToUTC(ltHours, ltMinutes, ltSeconds, daylightSavings, zoneCorrection, int(localDay), int(localMonth), int(localYear))

	// Determine Local Civil Date for UTC
	utcDay := GetLocalCivilDayForUT(ltHours, ltMinutes, ltSeconds, daylightSavings, zoneCorrection, int(localDay), int(localMonth), int(localYear))
	utcMonth := GetLocalCivilMonthForUT(ltHours, ltMinutes, ltSeconds, daylightSavings, zoneCorrection, int(localDay), int(localMonth), int(localYear))
	utcYear := GetLocalCivilYearForUT(ltHours, ltMinutes, ltSeconds, daylightSavings, zoneCorrection, int(localDay), int(localMonth), int(localYear))

	// Convert UTC to Greenwich Sidereal Time (GST)
	gst := ConvertUTToGST(utcTime, 0, 0, utcDay, utcMonth, utcYear)

	// Convert GST to Local Sidereal Time (LST)
	lst := ConvertGSTToLST(int(gst), int((gst-float64(int(gst)))*60), int((((gst-float64(int(gst)))*60)-float64(int((gst-float64(int(gst)))*60)))*60), geographicalLongitude)

	// Convert Right Ascension (RA) to decimal hours
	raDecimalHours := ConvertTimeToDecimal(raHours, raMinutes, raSeconds)

	// Calculate Hour Angle (HA)
	hourAngle := lst - raDecimalHours

	// Normalize Hour Angle to the range [0, 24) hours
	if hourAngle < 0 {
		hourAngle += 24
	}

	return hourAngle
}

// HourAngleToRA converts Hour Angle (HA) to Right Ascension (RA).
func HourAngleToRA(hourAngleHours, hourAngleMinutes, hourAngleSeconds, ltHours, ltMinutes, ltSeconds int, daylightSaving, zoneCorrection float64, localDay, localMonth, localYear int, geographicalLongitude float64) float64 {
	// Convert Local Time to UTC
	utcTime := ConvertLocalTimeToUTC(ltHours, ltMinutes, ltSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)

	// Determine Local Civil Date for UTC
	utcDay := GetLocalCivilDayForUT(ltHours, ltMinutes, ltSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	utcMonth := GetLocalCivilMonthForUT(ltHours, ltMinutes, ltSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)
	utcYear := GetLocalCivilYearForUT(ltHours, ltMinutes, ltSeconds, daylightSaving, zoneCorrection, localDay, localMonth, localYear)

	// Convert UTC to Greenwich Sidereal Time (GST)
	gst := ConvertUTToGST(utcTime, 0, 0, utcDay, utcMonth, utcYear)

	// Convert GST to Local Sidereal Time (LST)
	lst := ConvertGSTToLST(int(gst), int((gst-float64(int(gst)))*60), int((((gst-float64(int(gst)))*60)-float64(int((gst-float64(int(gst)))*60)))*60), geographicalLongitude)

	// Convert Hour Angle (HA) to decimal hours
	hourAngleDecimal := ConvertTimeToDecimal(hourAngleHours, hourAngleMinutes, hourAngleSeconds)

	// Calculate Right Ascension (RA)
	rightAscension := lst - hourAngleDecimal

	// Normalize Right Ascension to the range [0, 24) hours
	if rightAscension < 0 {
		rightAscension += 24
	}

	return rightAscension
}

// ConvertDegMinSecToDecimalDeg converts degrees, minutes, and seconds to decimal degrees.
func ConvertDegMinSecToDecimalDeg(degrees, minutes, seconds float64) float64 {
	// Convert seconds to decimal minutes
	secondsToMinutes := math.Abs(seconds) / 60.0

	// Convert minutes to decimal degrees, including the converted seconds
	minutesToDegrees := (math.Abs(minutes) + secondsToMinutes) / 60.0

	// Combine degrees with converted minutes
	decimalDegrees := math.Abs(degrees) + minutesToDegrees

	// Apply the sign of the original degrees, minutes, or seconds
	if degrees < 0 || minutes < 0 || seconds < 0 {
		return -decimalDegrees
	}
	return decimalDegrees
}

// GetDegreeOfDecimalDeg extracts the degree component from a decimal degree value.
func GetDegreeOfDecimalDeg(decimalDeg float64) int {
	// Take the absolute value of the input to handle negative degrees
	absoluteDecimalDeg := math.Abs(decimalDeg)

	// Calculate the degree part by flooring the absolute value
	degrees := math.Floor(absoluteDecimalDeg)

	// Determine the sign based on the original input
	if decimalDeg < 0 {
		return -int(degrees)
	} else {
		return int(degrees)
	}
}

// GetMinOfDecimalDeg extracts the minutes component from a decimal degree value.
func GetMinOfDecimalDeg(decimalDeg float64) int {
	// Take the absolute value of the input to handle negative degrees
	absoluteDecimalDeg := math.Abs(decimalDeg)

	// Calculate the total number of seconds
	totalSeconds := absoluteDecimalDeg * 3600

	// Calculate the number of minutes by taking the floor of total seconds divided by 60 and then taking mod 60
	minutes := math.Floor(totalSeconds / 60) // Total minutes
	minutesPart := math.Mod(minutes, 60)     // Minutes part within the hour

	return int(minutesPart)
}

// GetSecOfDecimalDeg extracts the seconds component from a decimal degree value.
func GetSecOfDecimalDeg(decimalDeg float64) float64 {
	// Take the absolute value of the input to handle negative degrees
	absoluteDecimalDeg := math.Abs(decimalDeg)

	// Calculate the total number of seconds
	totalSeconds := absoluteDecimalDeg * 3600

	// Calculate the seconds part within the minute
	secondsPart := math.Mod(totalSeconds, 60)

	// Round the seconds part to handle any precision issues
	roundedSeconds := math.Round(secondsPart)

	// Handle the special case where rounded seconds equals 60
	if roundedSeconds == 60 {
		return 0
	}
	return roundedSeconds
}

// ConvertDecimalDegToHours converts decimal degrees to hours.
func ConvertDecimalDegToHours(decimalDegrees float64) float64 {
	return decimalDegrees / 15.0
}

// ConvertHoursToDecimalDeg converts degree hours to decimal degrees.
func ConvertHoursToDecimalDeg(degHours float64) float64 {
	return degHours * 15.0
}

// ConvertRadiansToDegrees converts radians to degrees.
func ConvertRadiansToDegrees(radians float64) float64 {
	const radiansToDegrees = 180 / math.Pi
	return radians * radiansToDegrees
}

// ConvertRadiansToDegrees converts radians to degrees.
func ConvertDegreesToRadians(degrees float64) float64 {
	const degreesToRadians = math.Pi / 180
	return degrees * degreesToRadians
}

func Atan2(y, x float64) float64 {
	const epsilon = 1e-20
	const pi = math.Pi

	var angle float64

	if math.Abs(x) < epsilon {
		if y < 0 {
			angle = -pi / 2
		} else {
			angle = pi / 2
		}
	} else {
		angle = math.Atan2(y, x)
	}

	if x < 0 {
		angle = pi + angle
	}

	if angle < 0 {
		angle = angle + 2*pi
	}

	return angle
}

// Convert equatorial coordinates to azimuth (in decimal degrees)
func ConvertEquCoordinatesToAzimuth(hourAngleHrs, hourAngleMins, hourAngleSecs int, declinationDeg float64, declinationMins, declinationSecs, geographicalLatitude float64) float64 {
	// Convert hour angle to decimal hours
	hourAngleDecimal := ConvertTimeToDecimal(hourAngleHrs, hourAngleMins, hourAngleSecs)
	// Convert hour angle to degrees (1 hour = 15 degrees)
	hourAngleDeg := hourAngleDecimal * 15.0
	// Convert hour angle to radians
	hourAngleRad := ConvertDegreesToRadians(hourAngleDeg)

	// Convert declination to decimal degrees
	declinationDecimal := ConvertDegMinSecToDecimalDeg(declinationDeg, declinationMins, declinationSecs)
	// Convert declination to radians
	declinationRad := ConvertDegreesToRadians(declinationDecimal)

	// Convert geographical latitude to radians
	latitudeRad := ConvertDegreesToRadians(geographicalLatitude)

	// Calculate intermediate values for azimuth computation
	SinDeclination := math.Sin(declinationRad)
	SinLatitude := math.Sin(latitudeRad)
	CosDeclination := math.Cos(declinationRad)
	CosLatitude := math.Cos(latitudeRad)
	CosHourAngle := math.Cos(hourAngleRad)
	SinHourAngle := math.Sin(hourAngleRad)

	// Calculate the azimuth
	term1 := CosDeclination * CosLatitude * SinHourAngle
	term2 := SinDeclination - (SinLatitude * (SinDeclination*SinLatitude + CosDeclination*CosLatitude*CosHourAngle))
	azimuthRad := math.Atan2(-term1, term2)
	azimuthDeg := ConvertRadiansToDegrees(azimuthRad)

	// Normalize azimuth to [0, 360) degrees
	normalizedAzimuth := azimuthDeg - 360.0*math.Floor(azimuthDeg/360.0)

	return normalizedAzimuth
}

// Convert equatorial coordinates to altitude (in decimal degrees)
func ConvertEquCoordinatesToAltitude(hourAngleHrs, hourAngleMins, hourAngleSecs int, declinationDeg float64, declinationMins, declinationSecs, geographicalLatitude float64) float64 {
	// Convert hour angle to decimal hours
	hourAngleDecimal := ConvertTimeToDecimal(hourAngleHrs, hourAngleMins, hourAngleSecs)
	// Convert hour angle to degrees (1 hour = 15 degrees)
	hourAngleDeg := hourAngleDecimal * 15.0
	// Convert hour angle to radians
	hourAngleRad := ConvertDegreesToRadians(hourAngleDeg)

	// Convert declination to decimal degrees
	declinationDecimal := ConvertDegMinSecToDecimalDeg(declinationDeg, declinationMins, declinationSecs)
	// Convert declination to radians
	declinationRad := ConvertDegreesToRadians(declinationDecimal)

	// Convert geographical latitude to radians
	latitudeRad := ConvertDegreesToRadians(geographicalLatitude)

	// Calculate the Sine and CoSine values
	SinDeclination := math.Sin(declinationRad)
	SinLatitude := math.Sin(latitudeRad)
	CosDeclination := math.Cos(declinationRad)
	CosLatitude := math.Cos(latitudeRad)
	CosHourAngle := math.Cos(hourAngleRad)

	// Calculate the altitude
	altitudeRad := math.ASin(SinDeclination*SinLatitude + CosDeclination*CosLatitude*CosHourAngle)
	altitudeDeg := ConvertRadiansToDegrees(altitudeRad)

	return altitudeDeg
}

// Convert horizon coordinates to declination (in decimal degrees)
func HorizonCoordinatesToDec(azimuthDeg, azimuthMin, azimuthSec, altitudeDeg, altitudeMin, altitudeSec, geographicalLatitude float64) float64 {
	// Convert azimuth to decimal degrees
	azimuthDecimal := ConvertDegMinSecToDecimalDeg(azimuthDeg, azimuthMin, azimuthSec)
	// Convert altitude to decimal degrees
	altitudeDecimal := ConvertDegMinSecToDecimalDeg(altitudeDeg, altitudeMin, altitudeSec)
	// Convert azimuth and altitude to radians
	azimuthRad := ConvertDegreesToRadians(azimuthDecimal)
	altitudeRad := ConvertDegreesToRadians(altitudeDecimal)
	// Convert geographical latitude to radians
	latitudeRad := ConvertDegreesToRadians(geographicalLatitude)

	// Calculate the Sine and CoSine values
	SinAltitude := math.Sin(altitudeRad)
	SinLatitude := math.Sin(latitudeRad)
	CosAltitude := math.Cos(altitudeRad)
	CosLatitude := math.Cos(latitudeRad)
	CosAzimuth := math.Cos(azimuthRad)

	// Calculate the declination
	declinationRad := math.ASin(SinAltitude*SinLatitude + CosAltitude*CosLatitude*CosAzimuth)
	declinationDeg := ConvertRadiansToDegrees(declinationRad)

	return declinationDeg
}

// Convert horizon coordinates to hour angle (in decimal hours)
func HorizonCoordinatesToHourAngle(azimuthDeg, azimuthMin, azimuthSec, altitudeDeg, altitudeMin, altitudeSec, geographicalLatitude float64) float64 {
	// Convert azimuth and altitude to decimal degrees
	azimuthDecimal := ConvertDegMinSecToDecimalDeg(azimuthDeg, azimuthMin, azimuthSec)
	altitudeDecimal := ConvertDegMinSecToDecimalDeg(altitudeDeg, altitudeMin, altitudeSec)

	// Convert decimal degrees to radians
	azimuthRad := ConvertDegreesToRadians(azimuthDecimal)
	altitudeRad := ConvertDegreesToRadians(altitudeDecimal)
	latitudeRad := ConvertDegreesToRadians(geographicalLatitude)

	// Calculate intermediate values for hour angle computation
	SinAltitude := math.Sin(altitudeRad)
	SinLatitude := math.Sin(latitudeRad)
	CosAltitude := math.Cos(altitudeRad)
	CosLatitude := math.Cos(latitudeRad)
	CosAzimuth := math.Cos(azimuthRad)
	SinAzimuth := math.Sin(azimuthRad)

	// Calculate F, G, and H
	F := SinAltitude*SinLatitude + CosAltitude*CosLatitude*CosAzimuth
	G := -CosAltitude * CosLatitude * SinAzimuth
	H := SinAltitude - SinLatitude*F

	// Calculate the hour angle in decimal hours
	hourAngleRad := math.Atan2(G, H)
	hourAngleDeg := ConvertRadiansToDegrees(hourAngleRad)
	hourAngleHours := ConvertDecimalDegToHours(hourAngleDeg)

	// Normalize hour angle to [0, 24) hours
	normalizedHourAngle := hourAngleHours - 24.0*math.Floor(hourAngleHours/24.0)

	return normalizedHourAngle
}

// Calculate nutation in obliquity (in degrees)
func CalculateNutationOfObliquity(greenwichDay, greenwichMonth, greenwichYear int) float64 {
	// Calculate Julian Day and Time
	julianDay := ConvertGregorianToJulian(greenwichDay, greenwichMonth, greenwichYear) - 2415020
	centuries := julianDay / 36525.0
	centuriesSquared := centuries * centuries

	// Calculate arguments in radians

	// Argument L1 (Longitude of the Moon's Mean Node)
	argL1 := 279.6967 + 0.000303*centuriesSquared
	argL1 += 360.0 * (100.0021358*centuries - math.Floor(100.0021358*centuries))
	L1Rad := ConvertDegreesToRadians(2 * argL1)

	// Argument D1 (Mean Distance of the Moon from the Sun)
	argD1 := 270.4342 - 0.001133*centuriesSquared
	argD1 += 360.0 * (1336.855231*centuries - math.Floor(1336.855231*centuries))
	D1Rad := ConvertDegreesToRadians(2 * argD1)

	// Argument M1 (Mean Longitude of the Moon)
	argM1 := 358.4758 - 0.00015*centuriesSquared
	argM1 += 360.0 * (99.99736056*centuries - math.Floor(99.99736056*centuries))
	M1Rad := ConvertDegreesToRadians(argM1)

	// Argument M2 (Mean Longitude of the Sun)
	argM2 := 296.1046 + 0.009192*centuriesSquared
	argM2 += 360.0 * (1325.552359*centuries - math.Floor(1325.552359*centuries))
	M2Rad := ConvertDegreesToRadians(argM2)

	// Argument N1 (Longitude of the Ascending Node)
	argN1 := 259.1833 + 0.002078*centuriesSquared
	argN1 -= 360.0 * (5.372616667*centuries - math.Floor(5.372616667*centuries))
	N1Rad := ConvertDegreesToRadians(argN1)
	N2Rad := 2 * N1Rad

	// Calculate nutation in obliquity
	nutationObliquity := (9.21 + 0.00091*centuries) * math.Cos(N1Rad)
	nutationObliquity += (0.5522 - 0.00029*centuries) * math.Cos(L1Rad)
	nutationObliquity -= 0.0904 * math.Cos(N2Rad)
	nutationObliquity += 0.0884 * math.Cos(D1Rad)
	nutationObliquity += 0.0216 * math.Cos(L1Rad+M1Rad)
	nutationObliquity += 0.0183 * math.Cos(D1Rad-N1Rad)
	nutationObliquity += 0.0113 * math.Cos(D1Rad+M2Rad)
	nutationObliquity -= 0.0093 * math.Cos(L1Rad-M1Rad)
	nutationObliquity -= 0.0066 * math.Cos(L1Rad-N1Rad)

	// Convert nutation from arcseconds to degrees
	return nutationObliquity / 3600.0
}

// Calculate the obliquity of the ecliptic
func CalculateObliquityOfEcliptic(day, month, year int) float64 {
	// Convert Gregorian date to Julian Day
	julianDay := ConvertGregorianToJulian(day, month, year)
	// Calculate Julian Century
	julianCentury := (julianDay - 2415020) / 36525.0
	// Calculate the correction to obliquity of the ecliptic
	obliquityCorrection := julianCentury * (46.815 + julianCentury*(0.0006-julianCentury*0.00181))
	// Convert correction from arcseconds to degrees
	obliquityCorrectionDegrees := obliquityCorrection / 3600.0

	// Calculate obliquity of ecliptic
	obliquityOfEcliptic := 23.43929167 - obliquityCorrectionDegrees + CalculateNutationOfObliquity(day, month, year)

	return obliquityOfEcliptic
}

// CalculateTrueAnomaly solves Kepler's equation and returns the true anomaly in radians
func CalculateTrueAnomaly(meanAnomaly, eccentricity float64) float64 {
	const twoPi = 2 * math.Pi
	// Normalize the mean anomaly to the range 0 to 2*pi
	meanAnomaly = meanAnomaly - twoPi*math.Floor(meanAnomaly/twoPi)

	// Initial guess for Eccentric Anomaly
	eccentricAnomaly := meanAnomaly

	// Solve Kepler's equation uSing Newton's method
	for {
		delta := eccentricAnomaly - eccentricity*math.Sin(eccentricAnomaly) - meanAnomaly
		if math.Abs(delta) < 1e-6 {
			break
		}
		delta /= 1 - eccentricity*math.Cos(eccentricAnomaly)
		eccentricAnomaly -= delta
	}

	// Compute the true anomaly
	trueAnomaly := 2 * math.Atan(math.Sqrt((1+eccentricity)/(1-eccentricity))*math.Tan(eccentricAnomaly/2))
	return trueAnomaly
}

// CalculateEccentricAnomaly solves Kepler's equation and returns the eccentric anomaly in radians
func CalculateEccentricAnomaly(meanAnomaly, eccentricity float64) float64 {
	const twoPi = 2 * math.Pi
	// Normalize the mean anomaly to the range 0 to 2*pi
	meanAnomaly = meanAnomaly - twoPi*math.Floor(meanAnomaly/twoPi)

	// Initial guess for Eccentric Anomaly
	eccentricAnomaly := meanAnomaly

	// Solve Kepler's equation uSing Newton's method
	for {
		delta := eccentricAnomaly - eccentricity*math.Sin(eccentricAnomaly) - meanAnomaly
		if math.Abs(delta) < 1e-6 {
			break
		}
		delta /= 1 - eccentricity*math.Cos(eccentricAnomaly)
		eccentricAnomaly -= delta
	}

	return eccentricAnomaly
}

// CalculateSunEclipticLong calculates the Sun's ecliptic longitude
func CalculateSunEclipticLong(localHour, localMinute, localSecond int, daylightSavings, timeZoneCorrection float64, localDay, localMonth, localYear int) float64 {
	// Calculate the Greenwich Day, Month, and Year for Local Time
	greenwichDay := ComputeGreenwichDayForLT(localHour, localMinute, localSecond, daylightSavings, timeZoneCorrection, localDay, localMonth, localYear)
	greenwichMonth := ComputeGreenwichMonthForLT(localHour, localMinute, localSecond, daylightSavings, timeZoneCorrection, localDay, localMonth, localYear)
	greenwichYear := ComputeGreenwichYearForLT(localHour, localMinute, localSecond, daylightSavings, timeZoneCorrection, localDay, localMonth, localYear)

	// Convert Local Time to UTC
	utc := ConvertLocalTimeToUTC(localHour, localMinute, localSecond, daylightSavings, timeZoneCorrection, localDay, localMonth, localYear)

	// Calculate Julian Date
	julianDate := ConvertTimeToDecimal(greenwichDay, greenwichMonth, greenwichYear) - 2415020

	// Calculate time in Julian centuries Since 1900.0
	timeInJulianCenturies := (julianDate / 36525) + float64(utc/876600)
	timeInJulianCenturiesSquared := timeInJulianCenturies * timeInJulianCenturies

	// Mean Longitude of the Sun
	meanLongitudeSun := 100.00212359 * timeInJulianCenturies
	B := 360 * (meanLongitudeSun - math.Floor(meanLongitudeSun))
	meanLongitude := 279.69668 + 0.0003025*timeInJulianCenturiesSquared + B

	// Mean Anomaly of the Sun
	meanAnomalySun := 99.99736042 * timeInJulianCenturies
	B = 360 * (meanAnomalySun - math.Floor(meanAnomalySun))
	meanAnomaly := 358.47583 - (0.00015+0.0000033*timeInJulianCenturies)*timeInJulianCenturiesSquared + B

	// Eccentricity of the Earth's Orbit
	orbitalEccentricity := 0.01675104 - 0.0000418*timeInJulianCenturies - 0.000000126*timeInJulianCenturiesSquared

	// Calculate the true anomaly and eccentric anomaly
	meanAnomalyRad := ConvertDegreesToRadians(meanAnomaly)
	trueAnomaly := CalculateTrueAnomaly(meanAnomalyRad, orbitalEccentricity)
	// eccentricAnomaly := CalculateEccentricAnomaly(meanAnomalyRad, orbitalEccentricity)

	// Calculate perturbations
	A := 62.55209472 * timeInJulianCenturies
	B = 360 * (A - math.Floor(A))
	perturbationA1 := ConvertDegreesToRadians(153.23 + B)

	A = 125.1041894 * timeInJulianCenturies
	B = 360 * (A - math.Floor(A))
	perturbationB1 := ConvertDegreesToRadians(216.57 + B)

	A = 91.56766028 * timeInJulianCenturies
	B = 360 * (A - math.Floor(A))
	perturbationC1 := ConvertDegreesToRadians(312.69 + B)

	A = 1236.853095 * timeInJulianCenturies
	B = 360 * (A - math.Floor(A))
	perturbationD1 := ConvertDegreesToRadians(350.74 - 0.00144*timeInJulianCenturiesSquared + B)

	perturbationE1 := ConvertDegreesToRadians(231.19 + 20.2*timeInJulianCenturies)

	A = 183.1353208 * timeInJulianCenturies
	B = 360 * (A - math.Floor(A))
	perturbationH1 := ConvertDegreesToRadians(353.4 + B)

	perturbationSum := 0.00134*math.Cos(perturbationA1) + 0.00154*math.Cos(perturbationB1) + 0.002*math.Cos(perturbationC1)
	perturbationSum += 0.00179*math.Sin(perturbationD1) + 0.00178*math.Sin(perturbationE1)

	perturbationSum2 := 0.00000543*math.Sin(perturbationA1) + 0.00001575*math.Sin(perturbationB1)
	perturbationSum2 += 0.00001627*math.Sin(perturbationC1) + 0.00003076*math.Cos(perturbationD1)
	perturbationSum2 += 0.00000927 * math.Sin(perturbationH1)

	// Calculate the Sun's true anomaly corrected for perturbations
	trueAnomalyCorrected := trueAnomaly + ConvertDegreesToRadians(meanLongitude-meanAnomaly+perturbationSum)
	const twoPi = 2 * math.Pi

	// Normalize true anomaly to the range 0 to 2*pi
	trueAnomalyCorrected = trueAnomalyCorrected - twoPi*math.Floor(trueAnomalyCorrected/twoPi)

	// Convert the result to degrees
	return ConvertRadiansToDegrees(trueAnomalyCorrected)
}

// CalculateSunDistanceFromEarthInAU Calculate Sun's distance from the Earth in astronomical units
func CalculateSunDistanceFromEarthInAU(localHour, localMinute, localSecond int, daylightSavings, timeZoneCorrection float64, localDay, localMonth, localYear int) float64 {
	// Calculate the Greenwich Day, Month, and Year for Local Time
	greenwichDay := ComputeGreenwichDayForLT(localHour, localMinute, localSecond, daylightSavings, timeZoneCorrection, localDay, localMonth, localYear)
	greenwichMonth := ComputeGreenwichMonthForLT(localHour, localMinute, localSecond, daylightSavings, timeZoneCorrection, localDay, localMonth, localYear)
	greenwichYear := ComputeGreenwichYearForLT(localHour, localMinute, localSecond, daylightSavings, timeZoneCorrection, localDay, localMonth, localYear)

	// Convert Local Time to UTC
	utc := ConvertLocalTimeToUTC(localHour, localMinute, localSecond, daylightSavings, timeZoneCorrection, localDay, localMonth, localYear)

	// Calculate Julian Date
	julianDate := ConvertTimeToDecimal(greenwichDay, greenwichMonth, greenwichYear) - 2415020

	// Calculate time in Julian centuries Since 1900.0
	timeInJulianCenturies := (julianDate / 36525) + float64(utc/876600)
	timeInJulianCenturiesSquared := timeInJulianCenturies * timeInJulianCenturies

	// Mean Longitude of the Sun
	meanLongitudeSun := 100.00212359 * timeInJulianCenturies
	var B float64 = 360 * (meanLongitudeSun - math.Floor(meanLongitudeSun))
	// meanLongitude := 279.69668 + 0.0003025*timeInJulianCenturiesSquared + B

	// Mean Anomaly of the Sun
	meanAnomalySun := 99.99736042 * timeInJulianCenturies
	B = 360 * (meanAnomalySun - math.Floor(meanAnomalySun))
	meanAnomaly := 358.47583 - (0.00015+0.0000033*timeInJulianCenturies)*timeInJulianCenturiesSquared + B

	// Eccentricity of the Earth's Orbit
	orbitalEccentricity := 0.01675104 - 0.0000418*timeInJulianCenturies - 0.000000126*timeInJulianCenturiesSquared

	// Calculate the true anomaly and eccentric anomaly
	meanAnomalyRad := ConvertDegreesToRadians(meanAnomaly)
	// trueAnomaly := CalculateTrueAnomaly(meanAnomalyRad, orbitalEccentricity)
	eccentricAnomaly := CalculateEccentricAnomaly(meanAnomalyRad, orbitalEccentricity)

	// Calculate perturbations
	A := 62.55209472 * timeInJulianCenturies
	B = 360 * (A - math.Floor(A))
	perturbationA1 := ConvertDegreesToRadians(153.23 + B)

	A = 125.1041894 * timeInJulianCenturies
	B = 360 * (A - math.Floor(A))
	perturbationB1 := ConvertDegreesToRadians(216.57 + B)

	A = 91.56766028 * timeInJulianCenturies
	B = 360 * (A - math.Floor(A))
	perturbationC1 := ConvertDegreesToRadians(312.69 + B)

	A = 1236.853095 * timeInJulianCenturies
	B = 360 * (A - math.Floor(A))
	perturbationD1 := ConvertDegreesToRadians(350.74 - 0.00144*timeInJulianCenturiesSquared + B)

	perturbationE1 := ConvertDegreesToRadians(231.19 + 20.2*timeInJulianCenturies)

	A = 183.1353208 * timeInJulianCenturies
	B = 360 * (A - math.Floor(A))
	perturbationH1 := ConvertDegreesToRadians(353.4 + B)

	perturbationSum := 0.00134*math.Cos(perturbationA1) + 0.00154*math.Cos(perturbationB1) + 0.002*math.Cos(perturbationC1)
	perturbationSum += 0.00179*math.Sin(perturbationD1) + 0.00178*math.Sin(perturbationE1)

	perturbationSum2 := 0.00000543*math.Sin(perturbationA1) + 0.00001575*math.Sin(perturbationB1)
	perturbationSum2 += 0.00001627*math.Sin(perturbationC1) + 0.00003076*math.Cos(perturbationD1)
	perturbationSum2 += 0.00000927 * math.Sin(perturbationH1)

	return 1.0000002*(1-orbitalEccentricity*math.Cos(eccentricAnomaly)) + perturbationSum2
}

// CalculateAngularDiameterOfSunInDeg Calculate Sun's angular diameter in decimal degrees
func CalculateAngularDiameterOfSunInDeg(localHour, localMinute, localSecond int, daylightSavings, timeZoneCorrection float64, localDay, localMonth, localYear int) float64 {
	return  0.533128 / CalculateSunDistanceFromEarthInAU(localHour, localMinute, localSecond int, daylightSavings, timeZoneCorrection float64, localDay, localMonth, localYear int)
}

// EffectOfRefractionHelper calculates the effect of refraction given pressure, temperature, angle, and distance.
func EffectOfRefractionHelper(pressure, temperature, angle, distance float64) float64 {
	const lowerAngleThreshold = -0.087
	const upperAngleThreshold = 0.2617994
	const temperatureOffset = 273.0

	if angle < upperAngleThreshold {
		if angle < lowerAngleThreshold {
			return 0
		}

		angleDegrees := ConvertRadiansToDegrees(angle)
		a := ((0.00002*angleDegrees + 0.0196) * angleDegrees + 0.1594) * pressure
		b := (temperatureOffset + temperature) * ((0.0845*angleDegrees + 0.505) * angleDegrees + 1.0)

		return -ConvertDegreesToRadians(a / b) * distance
	}

	return -distance * 0.00007888888 * pressure / ((temperatureOffset + temperature) * math.Tan(angle))
}

// CalculateEffectOfRefraction calculates the total effect of refraction given an angle, switches, pressure, and temperature.
func CalculateEffectOfRefraction(angleDegrees float64, switches []string, pressure, temperature float64) float64 {
	angleRadians := ConvertDegreesToRadians(angleDegrees)
	direction := 1

	if len(switches) > 0 && strings.ToLower(switches[0]) == "t" {
		direction = -1
	}

	if direction == -1 {
		initialAngle := angleRadians
		currentAngle := angleRadians
		refractionCorrection := 0.0

		for {
			angleRadians = currentAngle + refractionCorrection
			if angleRadians < -0.087 {
				return 0
			}

			newRefractionCorrection := EffectOfRefractionHelper(pressure, temperature, angleRadians, float64(direction))
			if newRefractionCorrection == 0 || math.Abs(newRefractionCorrection-refractionCorrection) < 0.000001 {
				return ConvertRadiansToDegrees(initialAngle + newRefractionCorrection)
			}

			refractionCorrection = newRefractionCorrection
		}
	}

	refractionEffect := EffectOfRefractionHelper(pressure, temperature, angleRadians, float64(direction))
	if angleRadians < -0.087 {
		return 0
	}

	return ConvertRadiansToDegrees(angleRadians + refractionEffect)
}

// CalculateParallaxHourAngleHelper calculates the parallax hour angle correction given various parameters.
func CalculateParallaxHourAngleHelper(hourAngle, declination, rc, rp, rs, tp float64) (float64, float64) {
	CosHourAngle := math.Cos(hourAngle)
	SinDeclination := math.Sin(declination)
	CosDeclination := math.Cos(declination)

	parallaxCorrection := (rc*math.Sin(hourAngle)/(rp*CosDeclination) - rc*CosHourAngle)
	deltaX := math.Atan(parallaxCorrection)

	parallaxHourAngle := hourAngle + deltaX
	CosParallaxHourAngle := math.Cos(parallaxHourAngle)

	// Normalize the parallax hour angle to be within the range of [0, tp)
	parallaxHourAngle = parallaxHourAngle - tp*math.Floor(parallaxHourAngle/tp)
	
	q := math.Atan(CosParallaxHourAngle * (rp*SinDeclination - rs) / (rp*CosDeclination*CosHourAngle - rc))

	return parallaxHourAngle, q
}

// CalculateParallaxHourAngle calculates the corrected hour angle in decimal hours.
func CalculateParallaxHourAngle(hour, minute, second, declinationDeg, declinationMin, declinationSec float64, switches []string, geographicLatitude, height, horizontalParallax float64) float64 {
	latitudeRad := ConvertDegreesToRadians(geographicLatitude)
	CosLatitude := math.Cos(latitudeRad)
	SinLatitude := math.Sin(latitudeRad)

	u := math.Atan(0.996647 * SinLatitude / CosLatitude)
	CosU := math.Cos(u)
	SinU := math.Sin(u)
	heightFactor := height / 6378160.0

	rs := 0.996647*SinU + heightFactor*SinLatitude
	rc := CosU + heightFactor*CosLatitude
	tp := 2 * math.Pi
	rp := 1 / math.Sin(ConvertDegreesToRadians(horizontalParallax))

	hourAngleRad := ConvertDegreesToRadians(ConvertHoursToDecimalDegrees(ConvertTimeToDecimal(hour, minute, second)))
	initialHourAngle := hourAngleRad
	declinationRad := ConvertDegreesToRadians(ConvertDegMinSecToDecimalDeg(declinationDeg, declinationMin, declinationSec))
	initialDeclination := declinationRad

	direction := -1
	if len(switches) > 0 && strings.ToLower(switches[0]) == "t" {
		direction = 1
	}

	if direction == 1 {
		parallaxHourAngle, _ := CalculateParallaxHourAngleHelper(hourAngleRad, declinationRad, rc, rp, rs, tp)
		return ConvertDecimalDegToHours(ConvertRadiansToDegrees(parallaxHourAngle))
	}

	var prevP, prevQ float64
	for {
		parallaxHourAngle, parallaxDeclination := CalculateParallaxHourAngleHelper(hourAngleRad, declinationRad, rc, rp, rs, tp)
		deltaP := parallaxHourAngle - hourAngleRad
		deltaQ := parallaxDeclination - declinationRad

		if math.Abs(deltaP-prevP) < 0.000001 && math.Abs(deltaQ-prevQ) < 0.000001 {
			finalHourAngle := initialHourAngle - deltaP
			return ConvertDecimalDegToHours(ConvertRadiansToDegrees(finalHourAngle))
		}

		hourAngleRad = initialHourAngle - deltaP
		declinationRad = initialDeclination - deltaQ
		prevP = deltaP
		prevQ = deltaQ
	}
}

// CalculateParallaxDecHelper calculates the parallax declination correction given various parameters.
func CalculateParallaxDecHelper(hour, minute, second, declinationDeg, declinationMin, declinationSec float64, switches []string, geographicLatitude, height, horizontalParallax float64) float64 {
	latitudeRad := ConvertDegreesToRadians(geographicLatitude)
	CosLatitude := math.Cos(latitudeRad)
	SinLatitude := math.Sin(latitudeRad)

	u := math.Atan(0.996647 * SinLatitude / CosLatitude)
	CosU := math.Cos(u)
	SinU := math.Sin(u)
	heightFactor := height / 6378160.0

	rs := 0.996647*SinU + heightFactor*SinLatitude
	rc := CosU + heightFactor*CosLatitude
	tp := 2 * math.Pi
	rp := 1 / math.Sin(ConvertDegreesToRadians(horizontalParallax))

	hourAngleRad := ConvertDegreesToRadians(ConvertHoursToDecimalDegrees(ConvertTimeToDecimal(hour, minute, second)))
	declinationRad := ConvertDegreesToRadians(ConvertDegMinSecToDecimalDeg(declinationDeg, declinationMin, declinationSec))

	// Calculate the parallax hour angle correction
	CosHourAngle := math.Cos(hourAngleRad)
	SinDeclination := math.Sin(declinationRad)
	CosDeclination := math.Cos(declinationRad)

	parallaxCorrection := (rc * math.Sin(hourAngleRad) / (rp * CosDeclination) - rc * CosHourAngle)
	deltaX := math.Atan(parallaxCorrection)

	parallaxHourAngle := hourAngleRad + deltaX
	CosParallaxHourAngle := math.Cos(parallaxHourAngle)

	// Normalize the parallax hour angle to be within the range of [0, tp)
	parallaxHourAngle = parallaxHourAngle - tp*math.Floor(parallaxHourAngle/tp)
	
	parallaxDeclination := math.Atan(CosParallaxHourAngle * (rp*SinDeclination - rs) / (rp*CosDeclination*CosHourAngle - rc))

	return parallaxDeclination
}

// CalculateParallaxDec calculates the corrected declination in decimal degrees.
func CalculateParallaxDec(hour, minute, second, declinationDeg, declinationMin, declinationSec float64, switches []string, geographicLatitude, height, horizontalParallax float64) float64 {
	latitudeRad := ConvertDegreesToRadians(geographicLatitude)
	CosLatitude := math.Cos(latitudeRad)
	SinLatitude := math.Sin(latitudeRad)

	u := math.Atan(0.996647 * SinLatitude / CosLatitude)
	CosU := math.Cos(u)
	SinU := math.Sin(u)
	heightFactor := height / 6378160.0

	rs := 0.996647*SinU + heightFactor*SinLatitude
	rc := CosU + heightFactor*CosLatitude
	tp := 2 * math.Pi
	rp := 1 / math.Sin(ConvertDegreesToRadians(horizontalParallax))

	hourAngleRad := ConvertDegreesToRadians(ConvertHoursToDecimalDeg(ConvertTimeToDecimal(hour, minute, second)))
	declinationRad := ConvertDegreesToRadians(ConvertDegMinSecToDecimalDeg(declinationDeg, declinationMin, declinationSec))

	// Determine the direction based on the switch value
	direction := -1
	if len(switches) > 0 && strings.ToLower(switches[0]) == "t" {
		direction = 1
	}

	if direction == 1 {
		_, parallaxDeclination := CalculateParallaxHourAngleHelper(hourAngleRad, declinationRad, rc, rp, rs, tp)
		return ConvertRadiansToDegrees(parallaxDeclination)
	}

	var prevP, prevQ float64
	for {
		parallaxHourAngle, parallaxDeclination := CalculateParallaxHourAngleHelper(hourAngleRad, declinationRad, rc, rp, rs, tp)
		deltaP := parallaxHourAngle - hourAngleRad
		deltaQ := parallaxDeclination - declinationRad

		if math.Abs(deltaP-prevP) < 0.000001 && math.Abs(deltaQ-prevQ) < 0.000001 {
			finalDeclination := declinationRad - deltaQ
			return ConvertRadiansToDegrees(finalDeclination)
		}

		hourAngleRad = hourAngleRad - deltaP
		declinationRad = declinationRad - deltaQ
		prevP = deltaP
		prevQ = deltaQ
	}
}

// UnwindRadians converts an angle in radians to its equivalent angle in the range [0, 2π).
func UnwindRadians(radians float64) float64 {
	return radians - (2 * math.Pi * math.Floor(radians / (2 * math.Pi)))
}

// UnwindDegrees converts an angle in degrees to its equivalent angle in the range [0, 360).
func UnwindDegrees(degrees float64) float64 {
	return degrees - (360 * math.Floor(degrees / 360))
}

// CalculateMoonEclipticLongitude calculates the geocentric ecliptic longitude for the Moon.
func CalculateMoonEclipticLongitude(LH, LM, LS, DS, ZC, DY, MN, YR float64) float64 {
	UT := ConvertLocalTimeToUTC(LH, LM, LS, DS, ZC, DY, MN, YR)
	GD := GetLocalCivilDayForUT(LH, LM, LS, DS, ZC, DY, MN, YR)
	GM := GetLocalCivilMonthForUT(LH, LM, LS, DS, ZC, DY, MN, YR)
	GY := GetLocalCivilYearForUT(LH, LM, LS, DS, ZC, DY, MN, YR)

	JulianDate := ConvertGregorianToJulian(GD, GM, GY)
	T := ((JulianDate - 2415020) / 36525) + (UT / 876600)
	T2 := T * T

	Q := JulianDate - 2415020 + (UT / 24)

	meanLongitude := 360 * (Q / 27.32158213 - math.Floor(Q/27.32158213))
	meanAnomalySun := 360 * (Q / 365.2596407 - math.Floor(Q/365.2596407))
	meanAnomalyMoon := 360 * (Q / 27.55455094 - math.Floor(Q/27.55455094))
	elMeanAnomaly := 360 * (Q / 29.53058868 - math.Floor(Q/29.53058868))
	meanDistance := 360 * (Q / 27.21222039 - math.Floor(Q/27.21222039))
	node := 360 * (Q / 6798.363307 - math.Floor(Q/6798.363307))

	meanLongitude += 270.434164 - (0.001133 - 0.0000019*T)*T2
	meanAnomalySun += 358.475833 - (0.00015 + 0.0000033*T)*T2
	meanAnomalyMoon += 296.104608 + (0.009192 + 0.0000144*T)*T2
	elMeanAnomaly += 350.737486 - (0.001436 - 0.0000019*T)*T2
	meanDistance += 11.250889 - (0.003211 + 0.0000003*T)*T2
	node += 259.183275 + (0.002078 + 0.0000022*T)*T2

	A := ConvertDegreesToRadians(51.2 + 20.2*T)
	S1 := math.Sin(A)
	S2 := math.Sin(ConvertDegreesToRadians(node))
	B := 346.56 + (132.87 - 0.0091731*T)*T
	S3 := 0.003964 * math.Sin(ConvertDegreesToRadians(B))
	C := ConvertDegreesToRadians((node + 275.50) - 2.3*T)
	S4 := math.Sin(C)

	meanLongitude += 0.000233*S1 + S3 + 0.001964*S2
	meanAnomalySun -= 0.001778 * S1
	meanAnomalyMoon += 0.000817*S1 + S3 + 0.002541*S2
	meanDistance += S3 - 0.024691*S2 - 0.004328*S4
	elMeanAnomaly += 0.002011*S1 + S3 + 0.001964*S2
	E := 1 - (0.002495 + 0.00000752*T)*T
	E2 := E * E

	MLRad := ConvertDegreesToRadians(meanLongitude)
	MSRad := ConvertDegreesToRadians(meanAnomalySun)
	NARad := ConvertDegreesToRadians(node)
	MERad := ConvertDegreesToRadians(elMeanAnomaly)
	MFRad := ConvertDegreesToRadians(meanDistance)
	MDRad := ConvertDegreesToRadians(meanAnomalyMoon)

	L := 6.28875*math.Sin(MDRad) + 1.274018*math.Sin(2*MERad-MDRad)
	L += 0.658309*math.Sin(2*MERad) + 0.213616*math.Sin(2*MDRad)
	L -= E * 0.185596 * math.Sin(MSRad)
	L -= 0.114336 * math.Sin(2*MFRad)
	L += 0.058793 * math.Sin(2*(MERad-MDRad))
	L += 0.057212 * E * math.Sin(2*MERad-MSRad-MDRad)
	L += 0.05332 * math.Sin(2*MERad+MDRad)
	L += 0.045874 * E * math.Sin(2*MERad-MSRad)
	L += 0.041024 * E * math.Sin(MDRad-MSRad)
	L -= 0.034718 * math.Sin(MERad)
	L -= E * 0.030465 * math.Sin(MSRad+MDRad)
	L += 0.015326 * math.Sin(2*(MERad-MFRad))
	L -= 0.012528 * math.Sin(2*MFRad+MDRad)
	L -= 0.01098 * math.Sin(2*MFRad-MDRad)
	L += 0.010674 * math.Sin(4*MERad-MDRad)
	L += 0.010034 * math.Sin(3*MDRad)
	L += 0.008548 * math.Sin(4*MERad-2*MDRad)
	L -= E * 0.00791 * math.Sin(MSRad-MDRad+2*MERad)
	L -= E * 0.006783 * math.Sin(2*MERad+MSRad)
	L += 0.005162 * math.Sin(MDRad-MERad)
	L += E * 0.005 * math.Sin(MSRad+MERad)
	L += 0.003862 * math.Sin(4*MERad)
	L += E * 0.004049 * math.Sin(MDRad-MSRad+2*MERad)
	L += 0.003996 * math.Sin(2*(MDRad+MERad))
	L += 0.003665 * math.Sin(2*MERad-3*MDRad)
	L += E * 0.002695 * math.Sin(2*MDRad-MSRad)
	L += 0.002602 * math.Sin(MDRad-2*(MFRad+MERad))
	L += E * 0.002396 * math.Sin(2*(MERad-MDRad)-MSRad)
	L -= 0.002349 * math.Sin(MDRad+MERad)
	L += E2 * 0.002249 * math.Sin(2*(MERad-MSRad))
	L -= E * 0.002125 * math.Sin(2*MDRad+MSRad)
	L -= E2 * 0.002079 * math.Sin(2*MSRad)
	L += E2 * 0.002059 * math.Sin(2*(MERad-MSRad)-MDRad)
	L -= 0.001773 * math.Sin(MDRad+2*(MERad-MFRad))
	L -= 0.001595 * math.Sin(2*(MFRad+MERad))
	L += E * 0.00122 * math.Sin(4*MERad-MSRad-MDRad)
	L -= 0.00111 * math.Sin(2*(MDRad+MFRad))
	L += 0.000892 * math.Sin(MDRad-3*MERad)
	L -= E * 0.000811 * math.Sin(MSRad+MDRad+2*MERad)
	L += E * 0.000761 * math.Sin(4*MERad-MSRad-2*MDRad)
	L += E2 * 0.000704 * math.Sin(MDRad-2*(MSRad+MERad))
	L += E * 0.000693 * math.Sin(MSRad-2*(MDRad-MERad))
	L += E * 0.000598 * math.Sin(2*(MERad-MFRad)-MSRad)
	L += 0.00055 * math.Sin(MDRad+4*MERad)
	L += 0.000538 * math.Sin(4*MDRad)
	L += E * 0.000521 * math.Sin(4*MERad-MSRad)
	L += 0.000486 * math.Sin(2*MDRad-MERad)
	L += E2 * 0.000717 * math.Sin(MDRad-2*MSRad)

	geocentricLongitude := UnwindRadians(MLRad + ConvertDegreesToRadians(L))
	return ConvertRadiansToDegrees(geocentricLongitude)
}

func CalculateMoonEclipticLatitude(LH, LM, LS, DS, ZC, DY, MN, YR float64) float64 {
	// Convert local time to UTC and get the corresponding date components
	UT := ConvertLocalTimeToUTC(LH, LM, LS, DS, ZC, DY, MN, YR)
	day := GetLocalCivilDayForUT(LH, LM, LS, DS, ZC, DY, MN, YR)
	month := GetLocalCivilMonthForUT(LH, LM, LS, DS, ZC, DY, MN, YR)
	year := GetLocalCivilYearForUT(LH, LM, LS, DS, ZC, DY, MN, YR)

	// Calculate Julian date
	JulianDate := ConvertGregorianToJulian(day, month, year)
	T := ((JulianDate - 2415020) / 36525) + (UT / 876600)
	T2 := T * T

	// Define mean motions of the Moon
	periods := []float64{27.32158213, 365.2596407, 27.55455094, 29.53058868, 27.21222039, 6798.363307}
	Q := ConvertGregorianToJulian(day, month, year) - 2415020 + (UT / 24)

	// Calculate mean longitudes
	meanLongitudes := make([]float64, 6)
	for i, period := range periods {
		meanLongitudes[i] = 360 * (Q/period - math.Floor(Q/period))
	}

	// Calculate lunar elements
	ML := 270.434164 + meanLongitudes[0] - (0.001133-0.0000019*T)*T2
	MS := 358.475833 + meanLongitudes[1] - (0.00015+0.0000033*T)*T2
	MD := 296.104608 + meanLongitudes[2] + (0.009192+0.0000144*T)*T2
	ME := 350.737486 + meanLongitudes[3] - (0.001436-0.0000019*T)*T2
	MF := 11.250889 + meanLongitudes[4] - (0.003211+0.0000003*T)*T2
	NA := 259.183275 - meanLongitudes[5] + (0.002078+0.0000022*T)*T2

	// Calculate perturbations
	A := ConvertDegreesToRadians(51.2 + 20.2*T)
	S1 := math.Sin(A)
	S2 := math.Sin(ConvertDegreesToRadians(NA))
	B := 346.56 + (132.87-0.0091731*T)*T
	S3 := 0.003964 * math.Sin(ConvertDegreesToRadians(B))
	C := ConvertDegreesToRadians(NA + 275.05 - 2.3*T)
	S4 := math.Sin(C)

	ML += 0.000233*S1 + S3 + 0.001964*S2
	MS -= 0.001778*S1
	MD += 0.000817*S1 + S3 + 0.002541*S2
	MF += S3 - 0.024691*S2 - 0.004328*S4
	ME += 0.002011*S1 + S3 + 0.001964*S2

	// Calculate eccentricity
	E := 1 - (0.002495+0.00000752*T)*T
	E2 := E * E

	// Convert elements to radians
	ML = ConvertDegreesToRadians(ML)
	MS = ConvertDegreesToRadians(MS)
	NA = ConvertDegreesToRadians(NA)
	ME = ConvertDegreesToRadians(ME)
	MF = ConvertDegreesToRadians(MF)
	MD = ConvertDegreesToRadians(MD)

	// Calculate ecliptic latitude
	G := 5.128189*math.Sin(MF) + 0.280606*math.Sin(MD+MF)
	G += 0.277693*math.Sin(MD-MF) + 0.173238*math.Sin(2*ME-MF)
	G += 0.055413*math.Sin(2*ME+MF-MD) + 0.046272*math.Sin(2*ME-MF-MD)
	G += 0.032573*math.Sin(2*ME+MF) + 0.017198*math.Sin(2*MD+MF)
	G += 0.009267*math.Sin(2*ME+MD-MF) + 0.008823*math.Sin(2*MD-MF)
	G += E * 0.008247 * math.Sin(2*ME-MS-MF)
	G += 0.004323*math.Sin(2*(ME-MD)-MF)
	G += 0.0042*math.Sin(2*ME+MF+MD)
	G += E * 0.003372 * math.Sin(MF-MS-2*ME)
	G += E * 0.002472 * math.Sin(2*ME+MF-MS-MD)
	G += E * 0.002222 * math.Sin(2*ME+MF-MS)
	G += E * 0.002072 * math.Sin(2*ME-MF-MS-MD)
	G += E * 0.001877 * math.Sin(MF-MS+MD)
	G += 0.001828 * math.Sin(4*ME-MF-MD)
	G -= E * 0.001803 * math.Sin(MF+MS)
	G -= 0.00175 * math.Sin(3*MF)
	G += E * 0.00157 * math.Sin(MD-MS-MF)
	G -= 0.001487 * math.Sin(MF+ME)
	G -= E * 0.001481 * math.Sin(MF+MS+MD)
	G += E * 0.001417 * math.Sin(MF-MS-MD)
	G += E * 0.00135 * math.Sin(MF-MS)
	G += 0.00133 * math.Sin(MF-ME)
	G += 0.001106 * math.Sin(MF+3*MD)
	G += 0.00102 * math.Sin(4*ME-MF)
	G += 0.000833 * math.Sin(MF+4*ME-MD)
	G += 0.000781 * math.Sin(MD-3*MF)
	G += 0.00067 * math.Sin(MF+4*ME-2*MD)
	G += 0.000606 * math.Sin(2*ME-3*MF)
	G += 0.000597 * math.Sin(2*(ME+MD)-MF)
	G += E * 0.000492 * math.Sin(2*ME+MD-MS-MF)
	G += 0.00045 * math.Sin(2*(MD-ME)-MF)
	G += 0.000439 * math.Sin(3*MD-MF)
	G += 0.000423 * math.Sin(MF+2*(ME+MD))
	G += 0.000422 * math.Sin(2*ME-MF-3*MD)
	G -= E * 0.000367 * math.Sin(MS+MF+2*ME-MD)
	G -= E * 0.000353 * math.Sin(MS+MF+2*ME)
	G += 0.000331 * math.Sin(MF+4*ME)
	G += E * 0.000317 * math.Sin(2*ME+MF-MS+MD)
	G += E2 * 0.000306 * math.Sin(2*(ME-MS)-MF)
	G -= 0.000283 * math.Sin(MD+3*MF)

	// Calculate nutation terms
	W1 := 0.0004664 * math.Cos(NA)
	W2 := 0.0000754 * math.Cos(C)

	// Calculate and return the Moon's ecliptic latitude
	latitude := ConvertDegreesToRadians(G) * (1 - W1 - W2)
	return ConvertRadiansToDegrees(latitude)
}

func CalculateMoonHorizontalParallax(LH, LM, LS, DS, ZC, DY, MN, YR float64) float64 {
	// Convert local time to UTC and get the corresponding date components
	UT := ConvertLocalTimeToUTC(LH, LM, LS, DS, ZC, DY, MN, YR)
	day := GetLocalCivilDayForUT(LH, LM, LS, DS, ZC, DY, MN, YR)
	month := GetLocalCivilMonthForUT(LH, LM, LS, DS, ZC, DY, MN, YR)
	year := GetLocalCivilYearForUT(LH, LM, LS, DS, ZC, DY, MN, YR)

	// Calculate Julian date
	JulianDate := ConvertGregorianToJulian(day, month, year)
	T := ((JulianDate - 2415020) / 36525) + (UT / 876600)
	T2 := T * T

	// Define mean motions of the Moon
	periods := []float64{27.32158213, 365.2596407, 27.55455094, 29.53058868, 27.21222039, 6798.363307}
	Q := ConvertGregorianToJulian(day, month, year) - 2415020 + (UT / 24)

	// Calculate mean longitudes
	meanLongitudes := make([]float64, 6)
	for i, period := range periods {
		meanLongitudes[i] = 360 * (Q/period - math.Floor(Q/period))
	}

	// Calculate lunar elements
	ML := 270.434164 + meanLongitudes[0] - (0.001133-0.0000019*T)*T2
	MS := 358.475833 + meanLongitudes[1] - (0.00015+0.0000033*T)*T2
	MD := 296.104608 + meanLongitudes[2] + (0.009192+0.0000144*T)*T2
	ME := 350.737486 + meanLongitudes[3] - (0.001436-0.0000019*T)*T2
	MF := 11.250889 + meanLongitudes[4] - (0.003211+0.0000003*T)*T2
	NA := 259.183275 - meanLongitudes[5] + (0.002078+0.0000022*T)*T2

	// Calculate perturbations
	A := ConvertDegreesToRadians(51.2 + 20.2*T)
	S1 := math.Sin(A)
	S2 := math.Sin(ConvertDegreesToRadians(NA))
	B := 346.56 + (132.87-0.0091731*T)*T
	S3 := 0.003964 * math.Sin(ConvertDegreesToRadians(B))
	C := ConvertDegreesToRadians(NA + 275.05 - 2.3*T)
	S4 := math.Sin(C)

	ML += 0.000233*S1 + S3 + 0.001964*S2
	MS -= 0.001778*S1
	MD += 0.000817*S1 + S3 + 0.002541*S2
	MF += S3 - 0.024691*S2 - 0.004328*S4
	ME += 0.002011*S1 + S3 + 0.001964*S2

	// Calculate eccentricity
	E := 1 - (0.002495+0.00000752*T)*T
	E2 := E * E

	// Convert elements to radians
	ML = ConvertDegreesToRadians(ML)
	MS = ConvertDegreesToRadians(MS)
	NA = ConvertDegreesToRadians(NA)
	ME = ConvertDegreesToRadians(ME)
	MF = ConvertDegreesToRadians(MF)
	MD = ConvertDegreesToRadians(MD)

	// Calculate horizontal parallax
	parallax := 0.950724 + 0.051818*math.Cos(MD) + 0.009531*math.Cos(2*ME-MD)
	parallax += 0.007843*math.Cos(2*ME) + 0.002824*math.Cos(2*MD)
	parallax += 0.000857*math.Cos(2*ME+MD) + E*0.000533*math.Cos(2*ME-MS)
	parallax += E*0.000401*math.Cos(2*ME-MD-MS) + E*0.00032*math.Cos(MD-MS)
	parallax -= 0.000271*math.Cos(ME) - E*0.000264*math.Cos(MS+MD)
	parallax -= 0.000198*math.Cos(2*MF-MD) + 0.000173*math.Cos(3*MD)
	parallax += 0.000167*math.Cos(4*ME-MD) - E*0.000111*math.Cos(MS)
	parallax += 0.000103*math.Cos(4*ME-2*MD) - 0.000084*math.Cos(2*MD-2*ME)
	parallax -= E*0.000083*math.Cos(2*ME+MS) + 0.000079*math.Cos(2*ME+2*MD)
	parallax += 0.000072*math.Cos(4*ME) + E*0.000064*math.Cos(2*ME-MS+MD)
	parallax -= E*0.000063*math.Cos(2*ME+MS-MD) + E*0.000041*math.Cos(MS+ME)
	parallax += E*0.000035*math.Cos(2*MD-MS) - 0.000033*math.Cos(3*MD-2*ME)
	parallax -= 0.00003*math.Cos(MD+ME) - 0.000029*math.Cos(2*(MF-ME))
	parallax -= E*0.000029*math.Cos(2*MD+MS) + E2*0.000026*math.Cos(2*(ME-MS))
	parallax -= 0.000023*math.Cos(2*(MF-ME)+MD) + E*0.000019*math.Cos(4*ME-MS-MD)

	return parallax
}

func CalculateEarthToMoonDistance(LH, LM, LS, DS, ZC, DY, MN, YR float64) float64 {
	// Calculate the horizontal parallax of the Moon in radians
	horizontalParallax := ConvertDegreesToRadians(CalculateMoonHorizontalParallax(LH, LM, LS, DS, ZC, DY, MN, YR))
	
	// Calculate the distance from the Earth to the Moon using the parallax angle
	distance := 6378.14 / math.Sin(horizontalParallax)

	return distance
}

func CalculateMoonAngularDiameter(LH, LM, LS, DS, ZC, DY, MN, YR float64) float64 {
	// Calculate the distance from the Earth to the Moon in kilometers
	distance := CalculateEarthToMoonDistance(LH, LM, LS, DS, ZC, DY, MN, YR)

	// Calculate the Moon's angular diameter
	angularDiameter := 384401 * 0.5181 / distance

	return angularDiameter
}

func CalculateSunMeanEclipticLongitude(GD, GM, GY int) float64 {
	// Calculate Julian centuries from the given date
	T := (ConvertGregorianToJulian(GD, GM, GY) - 2415020) / 36525.0
	T2 := T * T
	
	// Calculate the mean ecliptic longitude of the Sun
	meanLongitude := 279.6966778 + 36000.76892 * T + 0.0003025 * T2
	
	// Normalize to the range [0, 360)
	normalizedLongitude := meanLongitude - 360 * math.Floor(meanLongitude / 360)

	return normalizedLongitude
}

func CalculateLongitudeOfSunAtPerigee(GD, GM, GY int) float64 {
	// Calculate Julian centuries from the given date
	T := (ConvertGregorianToJulian(GD, GM, GY) - 2415020) / 36525.0
	T2 := T * T
	
	// Calculate the longitude of the Sun at perigee
	longitudeAtPerigee := 281.2208444 + 1.719175 * T + 0.000452778 * T2
	
	// Normalize to the range [0, 360)
	normalizedLongitude := longitudeAtPerigee - 360 * math.Floor(longitudeAtPerigee / 360)

	return normalizedLongitude
}

func CalculateEccentricityOfSunEarthOrbit(GD, GM, GY int) float64 {
	// Calculate Julian centuries from the given date
	T := (ConvertGregorianToJulian(GD, GM, GY) - 2415020) / 36525.0
	T2 := T * T
	return 0.01675104 - 0.0000418 * T - 0.000000126 * T2
}

func CalculateEclipticDeclination(ELD, ELM, ELS, BD, BM, BS float64, GD, GM, GY int) float64 {
	// Convert ecliptic longitude and latitude from degrees, minutes, and seconds to decimal degrees
	eclipticLongitude := ConvertDegMinSecToDecimalDeg(ELD, ELM, ELS)
	eclipticLatitude := ConvertDegMinSecToDecimalDeg(BD, BM, BS)

	// Convert the angles to radians
	A := ConvertDegreesToRadians(eclipticLongitude)
	B := ConvertDegreesToRadians(eclipticLatitude)

	// Calculate the obliquity of the ecliptic
	obliquity := ConvertDegreesToRadians(CalculateObliquityOfEcliptic(GD, GM, GY))

	// Calculate the declination
	D := math.Sin(B) * math.Cos(obliquity) + math.Cos(B) * math.Sin(obliquity) * math.Sin(A)
	
	// Convert the declination from radians to degrees and return
	return ConvertRadiansToDegrees(math.Asin(D))
}

func CalculateEclipticRightAscension(ELD, ELM, ELS, BD, BM, BS float64, GD, GM, GY int) float64 {
	// Convert ecliptic longitude and latitude from degrees, minutes, and seconds to decimal degrees
	eclipticLongitude := ConvertDegMinSecToDecimalDeg(ELD, ELM, ELS)
	eclipticLatitude := ConvertDegMinSecToDecimalDeg(BD, BM, BS)

	// Convert the angles to radians
	A := ConvertDegreesToRadians(eclipticLongitude)
	B := ConvertDegreesToRadians(eclipticLatitude)

	// Calculate the obliquity of the ecliptic
	obliquity := ConvertDegreesToRadians(CalculateObliquityOfEcliptic(GD, GM, GY))

	// Calculate the right ascension
	D := math.Sin(A) * math.cos(C) - math.Tan(B) * math.Sin(C)
	E := math.Cos(A)

	F := ConvertRadiansToDegrees(math.Asin(D))
	return F - 360 * math.Floor(F / 360)
}

func CalculateSunTrueAnomaly(LCH, LCM, LCS, DS, ZC, LD, LM, LY float64) float64 {
	// Calculate Sun's true anomaly, i.e., how much its orbit deviates from a true circle to an ellipse

	// Compute Greenwich date and time
	GD := GetDComputeGreenwichDayForLT(LCH, LCM, LCS, DS, ZC, LD, LM, LY)
	GM := GetDComputeGreenwichMonthForLT(LCH, LCM, LCS, DS, ZC, LD, LM, LY)
	GY := GetDComputeGreenwichYearForLT(LCH, LCM, LCS, DS, ZC, LD, LM, LY)
	UT := ConvertLocalTimeToUTC(LCH, LCM, LCS, DS, ZC, LD, LM, LY)

	// Calculate Julian date
	JulianDate := ConvertGregorianToJulian(GD, GM, GY) - 2415020

	// Calculate Julian centuries
	T := (JulianDate / 36525) + (UT / 876600)
	T2 := T * T

	// Mean anomaly of the Sun
	meanAnomaly := 358.47583 - (0.00015 + 0.0000033 * T) * T2 + 360 * (100.0021359 * T - math.Floor(100.0021359 * T))

	// Eccentricity of Earth's orbit
	eccentricity := 0.01675104 - 0.0000418 * T - 0.000000126 * T2

	// Convert mean anomaly to radians
	meanAnomalyRad := ConvertDegreesToRadians(meanAnomaly)

	// Calculate true anomaly
	trueAnomaly := ConvertRadiansToDegrees(CalculateTrueAnomaly(meanAnomalyRad, eccentricity))

	return trueAnomaly
}

func CalculateSunMeanAnomaly(LCH, LCM, LCS, DS, ZC, LD, LM, LY float64) float64 {
	// Compute Greenwich date and time
	GD := GetDComputeGreenwichDayForLT(LCH, LCM, LCS, DS, ZC, LD, LM, LY)
	GM := GetDComputeGreenwichMonthForLT(LCH, LCM, LCS, DS, ZC, LD, LM, LY)
	GY := GetDComputeGreenwichYearForLT(LCH, LCM, LCS, DS, ZC, LD, LM, LY)
	UT := ConvertLocalTimeToUTC(LCH, LCM, LCS, DS, ZC, LD, LM, LY)

	// Calculate Julian date
	JulianDate := ConvertGregorianToJulian(GD, GM, GY) - 2415020

	// Calculate Julian centuries
	T := (JulianDate / 36525) + (UT / 876600)
	T2 := T * T
	A := 100.0021359 * T
	B := 360 * (A - math.Floor(A))
	M1 := 358.47583 - (0.00015 + 0.0000033 * T) * T2 + B
	AM := UnwindRadians(ConvertDegreesToRadians(M1))

	return AM
}

// next py fun :sun_true_anomaly
