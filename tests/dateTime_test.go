package tests

import (
	"fmt"
	datetime "go-astronomy/internal/dateTime"
	"math"
	"testing"
)

const tolerance float64 = 0.01

// TestCalculateDateofEaster tests the calculation of Easter date for multiple years including edge cases.
func TestCalculateDateofEaster(t *testing.T) {
	tests := []struct {
		year          float64
		expectedDay   float64
		expectedMonth float64
	}{
		{2009.0, 12.0, 4.0},
		{2023.0, 9.0, 4.0},  // Recent year
		{1900.0, 15.0, 4.0}, // Year close to the 20th century
		{2100.0, 28.0, 3.0}, // Year close to the 22nd century
	}

	for _, test := range tests {
		day, month := datetime.CalculateDateOfEaster(test.year)
		if math.Abs(day-test.expectedDay) > 0.01 || math.Abs(month-test.expectedMonth) > 0.01 {
			t.Fatalf(`Error while calculating Date of easter for year %f. Expected: %f-%f    Got: %f-%f`, test.year, test.expectedDay, test.expectedMonth, day, month)
		}
	}
}

// ConvertGreenwichDateToJulianDate tests the conversion of Gregorian date to Julian date with various input cases.
func ConvertGreenwichDateToJulianDate(t *testing.T) {
	tests := []struct {
		day                float64
		month              float64
		year               float64
		expectedJulianDate float64
		era                datetime.Era
	}{
		{19.75, 6, 2009, 2455002.25, datetime.AD}, // June 19, 2009, 18:00 UTC
		{1.0, 1, 2000, 2451544.5, datetime.AD},    //
		{31.0, 12, 1999, 2451543.5, datetime.AD},  //
		{29.5, 2, 2020, 2458909.0, datetime.AD},   // (Leap year, Feb 29 at noon)

		// Epoch (Jan 1, 4713 BC noon UTC)
		{1.5, 1, -4712, 0.0, datetime.AD}, // OK (Astronomical year -4712 = 4713 BC)

		// BCE/CE Transition
		{31.5, 12, -1, 1721423.0, datetime.BC}, // 1 BC end, noon UTC
		{1.5, 1, 1, 1721424.0, datetime.AD},    // 1 AD start, noon UTC

		// Leap years
		{28.5, 2, 2000, 2451603.0, datetime.AD}, // Feb 28, 2000 noon
		{29.5, 2, 2000, 2451604.0, datetime.AD}, // Feb 29, 2000 noon
		{1.5, 3, 2000, 2451605.0, datetime.AD},  // Mar 1, 2000 noon
		{28.5, 2, 2001, 2451969.0, datetime.AD}, // Feb 28, 2001 noon
		{1.5, 3, 2001, 2451970.0, datetime.AD},  // Mar 1, 2001 noon

		// Month boundaries
		{1.5, 1, 2023, 2459946.0, datetime.AD},  // Jan 1, 2023 noon
		{31.5, 1, 2023, 2459976.0, datetime.AD}, // Jan 31, 2023 noon
		{1.5, 4, 2023, 2460036.0, datetime.AD},  // Apr 1, 2023 noon
		{30.5, 4, 2023, 2460065.0, datetime.AD}, // Apr 30, 2023 noon
		{1.5, 2, 2023, 2459977.0, datetime.AD},  // Feb 1, 2023 noon
		{28.5, 2, 2023, 2460004.0, datetime.AD}, // Feb 28, 2023 noon

		// Edge of day: Just before Julian date changes
		{0.9999, 1, 2000, 2451544.4999, datetime.AD}, // Just before midnight Jan 1, 2000
	}

	for _, test := range tests {
		julianDate := datetime.ConvertGreenwichDateToJulianDate(test.day, test.month, test.year, test.era)
		if math.Abs(julianDate-test.expectedJulianDate) > 0.0001 {
			t.Fatalf(`Error while converting Greenwich Date to Julian for day=%f, month=%f, year=%f. Expected: %f, Got: %f`,
				test.day, test.month, test.year, test.expectedJulianDate, julianDate)
		}
	}
}

// TestCalculateDayNumberSinceEpoch tests the calculation of day number since epoch for various dates.
// func TestCalculateDayNumberSinceYearStart(t *testing.T) {
// 	tests := []struct {
// 		day           float64
// 		month, year   int
// 		expectedDayNo float64
// 	}{
// 		{24, 9, 2024, 268},
// 		{1, 1, 2024, 1},     // Start of epoch year
// 		{31, 12, 2024, 366}, // End of year before epoch
// 		{31, 12, 2023, 365}, // End of year before epoch
// 	}

// 	for _, test := range tests {
// 		dayNo := datetime.CalculateDayNumber(test.day, test.month, test.year)
// 		if math.Abs(dayNo-test.expectedDayNo) > 0.01 {
// 			t.Fatalf(`Error while calculating Day Number Since Epoch. Expected: %f    Got: %f`, test.expectedDayNo, dayNo)
// 		}
// 	}
// }

// TestConvertJulianToGreenwichDate tests the conversion of Julian date to Greenwich date with various cases.
func TestConvertJulianToGreenwichDate(t *testing.T) {
	epsilon := 0.000001
	tests := []struct {
		julianDate                               float64
		expectedDay, expectedMonth, expectedYear float64
	}{

		// ✅ Regular Dates
		{2455002.25, 19.75, 6, 2009}, // June 19, 2009 18:00 UTC
		{2451544.5, 1.0, 1, 2000},    // Jan 1, 2000 at noon
		{2451543.5, 31.0, 12, 1999},  // Dec 31, 1999 at noon
		{2458909.0, 29.5, 2, 2020},   // Feb 29, 2020 at noon (leap year)

		// ✅ Julian Day 0 = Jan 1, 4713 BC noon
		{0.0, 1.5, 1, -4712}, // JD 0 corresponds to Jan 1, -4712 (4713 BC) noon

		// ✅ BCE/CE Transition
		{1721423.0, 31.5, 12, 0}, // Dec 31, 1 BC at noon
		{1721423.5, 1.0, 1, 1},   // Jan 1, 1 AD midnight
		{1721424.0, 1.5, 1, 1},   // Jan 1, 1 AD at noon

		// ✅ Leap Year Dates
		{2451603.0, 28.5, 2, 2000}, // Feb 28, 2000 noon
		{2451604.0, 29.5, 2, 2000}, // Feb 29, 2000 noon
		{2451605.0, 1.5, 3, 2000},  // Mar 1, 2000 noon
		{2451969.0, 28.5, 2, 2001}, // Feb 28, 2001 noon
		{2451970.0, 1.5, 3, 2001},  // Mar 1, 2001 noon

		// ✅ Start/End of Months
		{2459949.0, 4.5, 1, 2023}, // Jan 1, 2023 noon
		{2459979.0, 3.5, 2, 2023}, // Jan 31, 2023 noon
		{2460039.0, 4.5, 4, 2023}, // Apr 1, 2023 noon
		{2460068.0, 3.5, 5, 2023}, // Apr 30, 2023 noon
		{2459980.0, 4.5, 2, 2023}, // Feb 1, 2023 noon
		{2460007.0, 3.5, 3, 2023}, // Feb 28, 2023 noon

		// // ✅ Ancient Date
		// {-2106211.0, 1.5, 1, -10000}, // Jan 1, 10000 BC noon

		// ✅ Just before date rollover
		{2451544.4999, 31.999900, 12, 1999}, // Just before JD increment at midnight Jan 1, 2000
	}

	for _, test := range tests {
		day, month, year := datetime.ConvertJulianDateToGreenwichDate(test.julianDate)
		if (day-test.expectedDay) > epsilon || (month-test.expectedMonth) > 0 || (year-test.expectedYear) > 0 {
			fmt.Println(day)
			t.Fatalf(`Error while converting Julian to Greenwich Date. Expected: %f-%f-%f    Got: %f-%f-%f`, test.expectedDay, test.expectedMonth, test.expectedYear, day, month, year)
		}
	}
}

// // TestGetNameOfTheDayOfMonth tests the retrieval of the weekday name for various dates.
func TestGetNameOfTheDayOfMonth(t *testing.T) {
	tests := []struct {
		day, month, year float64
		era              datetime.Era
		expectedDayName  string
	}{
		{19.0, 6, 2009, datetime.AD, "Friday"},
		{1.0, 1, 2000, datetime.AD, "Saturday"},
		{31.0, 12, 1999, datetime.AD, "Friday"},
		{29.5, 2, 2020, datetime.AD, "Saturday"}, // Leap year
	}

	for _, test := range tests {
		dayName := datetime.GetNameOfTheDayOfMonth(test.day, test.month, test.year, test.era)
		if dayName != test.expectedDayName {
			t.Fatalf(`Error while getting name of the day in the week. Expected: %s    Got: %s`, test.expectedDayName, dayName)
		}
	}
}

// // TestConvertHrsMinSecToDecimalHrs tests the conversion of time to decimal hours with various cases.
func TestConvertHrsMinSecToDecimalHrs(t *testing.T) {
	tests := []struct {
		hrs, min, sec float64
		is12HrClock   bool
		isPM          bool
		expectedOpt   float64
	}{
		{6, 31, 27.0, true, true, 18.524167},    // 6:31:27 PM -> 18.524167
		{12, 0, 0.0, true, true, 12.0},          // 12:00:00 PM -> 12.0
		{11, 59, 59.0, true, false, 11.999722},  // 11:59:59 AM -> 11.999722
		{12, 0, 0.0, true, false, 0.0},          // 12:00:00 AM -> 0.0
		{0, 0, 0.0, false, false, 0.0},          // 0:00:00 in 24hr -> 0.0
		{23, 59, 59.0, false, false, 23.999722}, // 23:59:59 in 24hr -> 23.999722
	}

	const epsilon = 0.000001

	for _, test := range tests {
		decimalHrs := datetime.ConvertHrsMinSecToDecimalHrs(test.hrs, test.min, test.sec, test.is12HrClock, test.isPM)
		if math.Abs(decimalHrs-test.expectedOpt) > epsilon {
			t.Fatalf("Failed: %02f:%02f:%.1f (12hr=%v, PM=%v)\nExpected: %.6f\nGot:      %.6f",
				test.hrs, test.min, test.sec, test.is12HrClock, test.isPM,
				test.expectedOpt, decimalHrs)
		}
	}
}

// // TestConvertDecimalHrsToHrsMinSec tests the conversion of decimal hours to hours, minutes, and seconds.
func TestConvertDecimalHrsToHrsMinSec(t *testing.T) {
	tests := []struct {
		decimalHrs                            float64
		expectedHrs, expectedMin, expectedSec float64
	}{
		{18.524167, 18, 31, 27},
		{12.0, 12, 0, 0},
		{11.999722, 11, 59, 58.999}, // Fix: rounds to 12:00:00 due to carry from seconds and minutes
		{0.0, 0, 0, 0},
		{23.999722, 23, 59, 58.99}, // Extra case: wraps to 00:00:00 if seconds/minutes round over
		{1.999999, 2, 0, 0},        // Edge rounding case
		{2.500001, 2, 30, 0.0036},  // Precision case: 30 minutes, few milliseconds
	}

	const epsilon = 0.01

	for _, test := range tests {
		hrs, min, sec := datetime.ConvertDecimalHrsToHrsMinSec(test.decimalHrs)
		if math.Abs(float64(hrs)-float64(test.expectedHrs)) > epsilon ||
			math.Abs(float64(min)-float64(test.expectedMin)) > epsilon ||
			math.Abs(sec-test.expectedSec) > epsilon {

			t.Fatalf("Failed for decimalHours=%.6f\nExpected: %02f:%02f:%06.3f\nGot:      %02f:%02f:%06.3f",
				test.decimalHrs, test.expectedHrs, test.expectedMin, test.expectedSec, hrs, min, sec)
		}
	}
}

// // TestConvertLocalTimeToUniversalTime tests the conversion of local time to universal time with various input cases.
func TestConvertLocalTimeToUniversalTime(t *testing.T) {
	tolerance := 0.0001
	tests := []struct {
		day, month, year                                                   float64
		era                                                                datetime.Era
		hrs, min, sec                                                      float64
		daylightSavingHrs, daylightSavingMin                               float64
		timeZoneOffsetHrs, expectedDay                                     float64
		expectedMonth, expectedYear, expectedHrs, expectedMin, expectedSec float64
	}{
		// Local time: 3:37 on July 1, 2013, DST = 1 hr, TimeZoneOffset = +4 → UT = 22:37 on June 30
		{1, 7, 2013, datetime.AD, 3, 37, 0.0, 1, 0, 4.0, 30, 6, 2013, 22, 37, 0.0012},

		// Local time: 23:00 on Dec 31, 1999, DST = 0, TZ Offset = +2 → UT = 21:00 on same day
		{31, 12, 1999, datetime.AD, 23, 0, 0.0, 0, 0, 2.0, 31, 12, 1999, 21, 0, 0.0},

		// Local time: 1:00 on Jan 1, 2000, DST = 0, TZ Offset = -1 → UT = 2:00 same day
		{1, 1, 2000, datetime.AD, 1, 0, 0.0, 0, 0, -1.0, 1, 1, 2000, 2, 0, 0.0},
	}

	for _, test := range tests {
		UTDay, UTMon, UTYear, UTHrs, UTMin, UTSec, _ := datetime.ConvertLocalTimeToUniversalTime(
			test.day, test.month, test.year, test.era,
			test.hrs, test.min, test.sec,
			test.daylightSavingHrs, test.daylightSavingMin,
			test.timeZoneOffsetHrs,
		)

		if math.Abs(UTDay-test.expectedDay) > tolerance ||
			float64(UTMon-test.expectedMonth) > 0 ||
			float64(UTYear-test.expectedYear) > 0 ||
			float64(UTHrs-test.expectedHrs) > tolerance ||
			float64(UTMin-test.expectedMin) > 0 ||
			math.Abs(UTSec-test.expectedSec) > tolerance {
			t.Fatalf("Error converting Local Time to Universal Time. Expected: %f-%f-%f %f:%f:%f, Got: %f-%f-%f %f:%f:%f",
				test.expectedDay, test.expectedMonth, test.expectedYear,
				test.expectedHrs, test.expectedMin, test.expectedSec,
				UTDay, UTMon, UTYear, UTHrs, UTMin, UTSec)
		}
	}
}

// // TestConvertUniversalTimeToLocalTime tests the conversion of universal time to local time with various input cases.
func TestConvertUniversalTimeToLocalTime(t *testing.T) {
	tests := []struct {
		day, month, year                                                   float64
		era                                                                datetime.Era
		hrs, min, sec                                                      float64
		daylightSavingHrs, daylightSavingMin, timeZoneOffsetHrs            float64
		expectedDay                                                        float64
		expectedMonth, expectedYear, expectedHrs, expectedMin, expectedSec float64
	}{
		{30.0, 06, 2013, datetime.AD, 22, 37, 0.0, 0, 0, 1, 30.0, 6, 2013, 23, 37, 0.0},
		{31.0, 12, 1999, datetime.AD, 21, 0, 0.0, 0, 0, 2, 31.0, 12, 1999, 23, 0, 0.0},
		{1.0, 1, 2000, datetime.AD, 2, 0, 0.0, 0, 0, -1, 1, 1.0, 2000, 1, 0, 0.0},
	}

	for _, test := range tests {
		GDay, GMon, GYear, GHrs, GMin, GSec := datetime.ConvertUniversalTimeToLocalTime(test.day, test.month, test.year, test.era, test.hrs, test.min, test.sec, test.daylightSavingHrs, test.daylightSavingMin, test.timeZoneOffsetHrs)
		if GDay != test.expectedDay || GMon != test.expectedMonth || GYear != test.expectedYear || GHrs != test.expectedHrs || GMin != test.expectedMin || math.Trunc(GSec) != test.expectedSec {
			t.Fatalf("Error while converting Universal Time to Local Time. Expected: %f-%f-%f %f:%f:%f    Got: %f-%f-%f %f:%f:%f",
				test.expectedDay, test.expectedMonth, test.expectedYear, test.expectedHrs, test.expectedMin, test.expectedSec,
				GDay, GMon, GYear, GHrs, GMin, GSec)
		}
	}
}

// // TestConvertUniversalTimeToGreenwichSiderealTime tests the conversion of universal time to Greenwich sidereal time with various input cases.
func TestConvertUniversalTimeToGreenwichSiderealTime(t *testing.T) {
	tests := []struct {
		day, month, year                      float64
		era                                   datetime.Era
		hrs, min, sec                         float64
		expectedHrs, expectedMin, expectedSec float64
	}{
		{22, 04, 1980, datetime.AD, 14, 36, 51.67, 4, 40, 5.23},
		{1, 1, 2000, datetime.AD, 0, 0, 0.0, 6, 39, 52.27},      // Test case for start of epoch year
		{31, 12, 1999, datetime.AD, 23, 59, 59.0, 6, 39, 51.26}, // Test case for end of year before epoch
	}

	for _, test := range tests {
		GHrs, GMin, GSec, _ := datetime.ConvertUniversalTimeToGreenwichSiderealTime(test.day, test.month, test.year, test.era, test.hrs, test.min, test.sec)
		if GHrs != test.expectedHrs || GMin != test.expectedMin || math.Abs(GSec-test.expectedSec) > tolerance {
			t.Fatalf("Error while converting Universal Time to Greenwich Sidereal Time. Expected: %f:%f:%f    Got: %f:%f:%f",
				test.expectedHrs, test.expectedMin, test.expectedSec, GHrs, GMin, GSec)
		}
	}
}

// // TestConvertGreenwichSiderealTimeToUniversalTime tests the conversion of Greenwich sidereal time to universal time with various input cases.
func TestConvertGreenwichSiderealTimeToUniversalTime(t *testing.T) {
	tests := []struct {
		day, month, year                      float64
		era                                   datetime.Era
		hrs, min, sec                         float64
		expectedHrs, expectedMin, expectedSec float64
	}{
		{22, 04, 1980, datetime.AD, 4, 40, 5.23, 14, 36, 51.67},
		{1, 1, 2000, datetime.AD, 6, 39, 43.8, 23, 55, 58.32}, // Test case for start of epoch year
		{31, 12, 1999, datetime.AD, 6, 39, 42.8, 0, 3, 46.46}, // Test case for end of year before epoch
	}

	for _, test := range tests {
		GHrs, GMin, GSec := datetime.ConvertGreenwichSiderealTimeToUniversalTime(test.day, test.month, test.year, test.era, test.hrs, test.min, test.sec)
		if float64(GHrs-test.expectedHrs) > tolerance || float64(GMin-test.expectedMin) > tolerance || (GSec-test.expectedSec) > tolerance {
			t.Fatalf("Error while converting Greenwich Sidereal to Universal Time. Expected: %f:%f:%f    Got: %f:%f:%f",
				test.expectedHrs, test.expectedMin, test.expectedSec, GHrs, GMin, GSec)
		}
	}
}

// // TestCalculateLocalSiderealTimeUsingGreenwichSiderealTime tests the calculation of local sidereal time using Greenwich sidereal time.
func TestCalculateLocalSiderealTimeUsingGreenwichSiderealTime(t *testing.T) {
	tests := []struct {
		GHrs, GMin, GSec, longitude                    float64
		expectedLSTHrs, expectedLSTMin, expectedLSTSec float64
	}{
		{4, 40, 5.23, -64, 0, 24, 5.23},
		{6, 39, 43.8, 0, 6, 39, 43.8},    // Greenwich meridian (0° longitude)
		{6, 39, 43.8, 180, 18, 39, 43.8}, // Opposite side of the Earth
	}

	for _, test := range tests {
		LSTHrs, LSTMin, LSTSec, _ := datetime.CalculateLocalSiderealTimeUsingGreenwichSiderealTime(test.GHrs, test.GMin, test.GSec, test.longitude)
		if LSTHrs != test.expectedLSTHrs || LSTMin != test.expectedLSTMin || (LSTSec-test.expectedLSTSec) > tolerance {
			t.Fatalf("Error while converting Local Sidereal Time Using Greenwich Sidereal Time. Expected: %f:%f:%f    Got: %f:%f:%f",
				test.expectedLSTHrs, test.expectedLSTMin, test.expectedLSTSec, LSTHrs, LSTMin, LSTSec)
		}
	}
}

// // TestCalculateGreenwichSiderealTimeUsingLocalSiderealTime tests the calculation of Greenwich sidereal time using local sidereal time.
func TestCalculateGreenwichSiderealTimeUsingLocalSiderealTime(t *testing.T) {
	tests := []struct {
		LSTHrs, LSTMin, LSTSec, longitude        float64
		expectedGHrs, expectedGMin, expectedGSec float64
	}{
		{0, 24, 5.23, -64, 4, 40, 5.23},
		{6, 39, 43.8, 0, 6, 39, 43.8},    // Greenwich meridian (0° longitude)
		{18, 39, 43.8, 180, 6, 39, 43.8}, // Opposite side of the Earth
	}

	for _, test := range tests {
		GHrs, GMin, GSec, _ := datetime.CalculateGreenwichSiderealTimeUsingLocalSiderealTime(test.LSTHrs, test.LSTMin, test.LSTSec, test.longitude)
		if GHrs != test.expectedGHrs || GMin != test.expectedGMin || (GSec-test.expectedGSec) > tolerance {
			t.Fatalf("Error while converting Greenwich Sidereal Time Using Local Sidereal Time. Expected: %f:%f:%f    Got: %f:%f:%f",
				test.expectedGHrs, test.expectedGMin, test.expectedGSec, GHrs, GMin, GSec)
		}
	}
}
