package macros

import (
	"math"
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
	sinDeclination := math.Sin(declinationRad)
	sinLatitude := math.Sin(latitudeRad)
	cosDeclination := math.Cos(declinationRad)
	cosLatitude := math.Cos(latitudeRad)
	cosHourAngle := math.Cos(hourAngleRad)
	sinHourAngle := math.Sin(hourAngleRad)

	// Calculate the azimuth
	term1 := cosDeclination * cosLatitude * sinHourAngle
	term2 := sinDeclination - (sinLatitude * (sinDeclination*sinLatitude + cosDeclination*cosLatitude*cosHourAngle))
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

	// Calculate the sine and cosine values
	sinDeclination := math.Sin(declinationRad)
	sinLatitude := math.Sin(latitudeRad)
	cosDeclination := math.Cos(declinationRad)
	cosLatitude := math.Cos(latitudeRad)
	cosHourAngle := math.Cos(hourAngleRad)

	// Calculate the altitude
	altitudeRad := math.Asin(sinDeclination*sinLatitude + cosDeclination*cosLatitude*cosHourAngle)
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

	// Calculate the sine and cosine values
	sinAltitude := math.Sin(altitudeRad)
	sinLatitude := math.Sin(latitudeRad)
	cosAltitude := math.Cos(altitudeRad)
	cosLatitude := math.Cos(latitudeRad)
	cosAzimuth := math.Cos(azimuthRad)

	// Calculate the declination
	declinationRad := math.Asin(sinAltitude*sinLatitude + cosAltitude*cosLatitude*cosAzimuth)
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
	sinAltitude := math.Sin(altitudeRad)
	sinLatitude := math.Sin(latitudeRad)
	cosAltitude := math.Cos(altitudeRad)
	cosLatitude := math.Cos(latitudeRad)
	cosAzimuth := math.Cos(azimuthRad)
	sinAzimuth := math.Sin(azimuthRad)

	// Calculate F, G, and H
	F := sinAltitude*sinLatitude + cosAltitude*cosLatitude*cosAzimuth
	G := -cosAltitude * cosLatitude * sinAzimuth
	H := sinAltitude - sinLatitude*F

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

func CalculateSunEclipticLong(LCH, LCM, LCS int, DS, ZC float64, LD, LM, LY int) float64 {
	AA := ComputeGreenwichDayForLT(LCH, LCM, LCS, DS, ZC, LD, LM, LY)
	BB := ComputeGreenwichMonthForLT(LCH, LCM, LCS, DS, ZC, LD, LM, LY)
	CC := ComputeGreenwichYearForLT(LCH, LCM, LCS, DS, ZC, LD, LM, LY)
	UT := ConvertLSTToGST()
}

// next py fun :dms_dd
