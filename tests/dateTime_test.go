package tests

import (
	datetime "go-astronomy/internal/dateTime"
	"math"
	"testing"
)

const tolerance float64 = 0.01

// TestCalculateDateofEaster tests the calculation of Easter date for multiple years including edge cases.
func TestCalculateDateofEaster(t *testing.T) {
	tests := []struct {
		year          int
		expectedDay   int
		expectedMonth int
	}{
		{2009, 12, 4},
		{2023, 9, 4},  // Recent year
		{1900, 15, 4}, // Year close to the 20th century
		{2100, 28, 3}, // Year close to the 22nd century
	}

	for _, test := range tests {
		day, month := datetime.CalculateDateOfEaster(test.year)
		if day != test.expectedDay || month != test.expectedMonth {
			t.Fatalf(`Error while calculating Date of easter for year %d. Expected: %d-%d    Got: %d-%d`, test.year, test.expectedDay, test.expectedMonth, day, month)
		}
	}
}

// TestConvertToJulianDate tests the conversion of Gregorian date to Julian date with various input cases.
func TestConvertToJulianDate(t *testing.T) {
	tests := []struct {
		day                float64
		month              int
		year               int
		expectedJulianDate float64
	}{
		{19.75, 6, 2009, 2455002.25},
		{1.0, 1, 2000, 2451544.5},   // Start of 21st century
		{31.0, 12, 1999, 2451543.5}, // End of 20th century
		{29.5, 2, 2020, 2458909.0},  // Leap year
	}

	for _, test := range tests {
		julianDate := datetime.ConvertGreenwichDateToJulianDate(test.day, test.month, test.year)
		if math.Abs(julianDate-test.expectedJulianDate) > 0.0001 {
			t.Fatalf(`Error while converting Greenwich Date to Julian. Expected: %f    Got: %f`, test.expectedJulianDate, julianDate)
		}
	}
}

// TestCalculateDayNumberSinceEpoch tests the calculation of day number since epoch for various dates.
func TestCalculateDayNumberSinceEpoch(t *testing.T) {
	tests := []struct {
		day           float64
		month, year   int
		expectedDayNo float64
	}{
		{27, 7, 2003, 207},
		{1, 1, 2000, 0},     // Start of epoch year
		{31, 12, 1999, 364}, // End of year before epoch
	}

	for _, test := range tests {
		dayNo := datetime.CalculateDayNumber(test.day, test.month, test.year)
		if math.Abs(dayNo-test.expectedDayNo) > 0.01 {
			t.Fatalf(`Error while calculating Day Number Since Epoch. Expected: %f    Got: %f`, test.expectedDayNo, dayNo)
		}
	}
}

// TestConvertToGreenwichDate tests the conversion of Julian date to Greenwich date with various cases.
func TestConvertToGreenwichDate(t *testing.T) {
	tests := []struct {
		julianDate    float64
		expectedDay   float64
		expectedMonth int
		expectedYear  int
	}{
		{2455002.25, 19.75, 6, 2009},
		{2451544.5, 1.0, 1, 2000},
		{2451543.5, 31.0, 12, 1999},
		{2458909.0, 29.5, 2, 2020},
	}

	for _, test := range tests {
		day, month, year := datetime.ConvertJulianDateToGreenwichDate(test.julianDate)
		if day != test.expectedDay || month != test.expectedMonth || year != test.expectedYear {
			t.Fatalf(`Error while converting Julian to Greenwich Date. Expected: %f-%d-%d    Got: %f-%d-%d`, test.expectedDay, test.expectedMonth, test.expectedYear, day, month, year)
		}
	}
}

// TestGetNameOfTheDayOfMonth tests the retrieval of the weekday name for various dates.
func TestGetNameOfTheDayOfMonth(t *testing.T) {
	tests := []struct {
		day             float64
		month           int
		year            int
		expectedDayName string
	}{
		{19.0, 6, 2009, "Friday"},
		{1.0, 1, 2000, "Saturday"},
		{31.0, 12, 1999, "Friday"},
		{29.5, 2, 2020, "Saturday"}, // Leap year
	}

	for _, test := range tests {
		dayName := datetime.GetNameOfTheDayOfMonth(test.day, test.month, test.year)
		if dayName != test.expectedDayName {
			t.Fatalf(`Error while getting name of the day in the week. Expected: %s    Got: %s`, test.expectedDayName, dayName)
		}
	}
}

// TestConvertHrsMinSecToDecimalHrs tests the conversion of time to decimal hours with various cases.
func TestConvertHrsMinSecToDecimalHrs(t *testing.T) {
	tests := []struct {
		hrs         int
		min         int
		sec         float64
		isPM        bool
		expectedOpt float64
	}{
		{6, 31, 27.0, true, 18.524167},   // 6:31:27 PM -> 18.524167 decimal hours
		{12, 0, 0.0, true, 12.0},         // 12:00:00 PM -> 12.0 decimal hours
		{11, 59, 59.0, false, 11.999722}, // 11:59:59 AM -> 11.999722 decimal hours
		{0, 0, 0.0, false, 0.0},          // 12:00:00 AM -> 0.0 decimal hours
	}

	for _, test := range tests {
		decimalHrs := math.Round(datetime.ConvertHrsMinSecToDecimalHrs(test.hrs, test.min, test.sec, true, test.isPM)*1000000) / 1000000
		if math.Abs(decimalHrs-test.expectedOpt) > 0.0001 {
			t.Fatalf("Error while converting Hrs Min Sec to Decimal. Expected: %f    Got: %f", test.expectedOpt, decimalHrs)
		}
	}
}

// TestConvertDecimalHrsToHrsMinSec tests the conversion of decimal hours to hours, minutes, and seconds.
func TestConvertDecimalHrsToHrsMinSec(t *testing.T) {
	tests := []struct {
		decimalHrs  float64
		expectedHrs int
		expectedMin int
		expectedSec float64
	}{
		{18.524167, 18, 31, 27},
		{12.0, 12, 0, 0},
		{11.999722, 11, 59, 59},
		{0.0, 0, 0, 0},
	}

	for _, test := range tests {
		Hrs, min, sec := datetime.ConvertDecimalHrsToHrsMinSec(test.decimalHrs)
		if math.Abs(float64(Hrs)-float64(test.expectedHrs)) > 0.01 || math.Abs(float64(min)-float64(test.expectedMin)) > 0.01 || math.Abs(sec-float64(test.expectedSec)) > 0.01 {
			t.Fatalf("Error while converting Decimal to Hrs Min Sec. Expected: %d %d %f  Got: %d %d %f", test.expectedHrs, test.expectedMin, test.expectedSec, Hrs, min, sec)
		}
	}
}

// TestConvertLocalTimeToUniversalTime tests the conversion of local time to universal time with various input cases.
func TestConvertLocalTimeToUniversalTime(t *testing.T) {
	tests := []struct {
		day                                                   float64
		month, year, hrs, min                                 int
		sec                                                   float64
		daylightSavingHrs, daylightSavingMin                  int
		timeZoneOffsetHrs, expectedDay                        float64
		expectedMonth, expectedYear, expectedHrs, expectedMin int
		expectedSec                                           float64
	}{
		{1, 7, 2013, 3, 37, 0.0, 1, 0, 4.0, 30, 6, 2013, 22, 37, 00.0},
		{31, 12, 1999, 23, 0, 0.0, 0, 0, 2.0, 31, 12, 1999, 21, 0, 0.0},
		{1, 1, 2000, 1, 0, 0.0, 0, 0, -1.0, 1, 1, 2000, 2, 0, 0.0},
	}

	for _, test := range tests {
		UTDay, UTMon, UTYear, UTHrs, UTMin, UTSec, _ := datetime.ConvertLocalTimeToUniversalTime(test.day, test.month, test.year, test.hrs, test.min, test.sec, test.daylightSavingHrs, test.daylightSavingMin, test.timeZoneOffsetHrs)
		// datetime.ConvertLocalTimeToUniversalTime(test.day, test.month, test.year, test.hrs, test.min, test.sec, test.daylightSavingHrs, test.daylightSavingMin, test.timeZoneOffsetHrs)
		if UTDay != test.expectedDay || UTMon != test.expectedMonth || UTYear != test.expectedYear || UTHrs != test.expectedHrs || UTMin != test.expectedMin || UTSec != test.expectedSec {
			t.Fatalf("Error while converting Local Time to Universal Time. Expected: %f-%d-%d %d:%d:%f    Got: %f-%d-%d %d:%d:%f",
				test.expectedDay, test.expectedMonth, test.expectedYear, test.expectedHrs, test.expectedMin, test.expectedSec,
				UTDay, UTMon, UTYear, UTHrs, UTMin, UTSec)
		}
	}
}

// TestConvertUniversalTimeToLocalTime tests the conversion of universal time to local time with various input cases.
func TestConvertUniversalTimeToLocalTime(t *testing.T) {
	tests := []struct {
		day                                                   float64
		month, year, hrs, min                                 int
		sec                                                   float64
		daylightSavingHrs, daylightSavingMin                  int
		timeZoneOffsetHrs                                     float64
		expectedDay                                           float64
		expectedMonth, expectedYear, expectedHrs, expectedMin int
		expectedSec                                           float64
	}{
		{30.0, 06, 2013, 22, 37, 0.0, 0, 0, 1, 30.0, 6, 2013, 23, 37, 0.0},
		{31.0, 12, 1999, 21, 0, 0.0, 0, 0, 2, 31.0, 12, 1999, 23, 0, 0.0},
		{1.0, 1, 2000, 2, 0, 0.0, 0, 0, -1, 1, 1.0, 2000, 1, 0, 0.0},
	}

	for _, test := range tests {
		GDay, GMon, GYear, GHrs, GMin, GSec := datetime.ConvertUniversalTimeToLocalTime(test.day, test.month, test.year, test.hrs, test.min, test.sec, test.daylightSavingHrs, test.daylightSavingMin, test.timeZoneOffsetHrs)
		if GDay != test.expectedDay || GMon != test.expectedMonth || GYear != test.expectedYear || GHrs != test.expectedHrs || GMin != test.expectedMin || math.Trunc(GSec) != test.expectedSec {
			t.Fatalf("Error while converting Universal Time to Local Time. Expected: %f-%d-%d %d:%d:%f    Got: %f-%d-%d %d:%d:%f",
				test.expectedDay, test.expectedMonth, test.expectedYear, test.expectedHrs, test.expectedMin, test.expectedSec,
				GDay, GMon, GYear, GHrs, GMin, GSec)
		}
	}
}

// TestConvertUniversalTimeToGreenwichSiderealTime tests the conversion of universal time to Greenwich sidereal time with various input cases.
func TestConvertUniversalTimeToGreenwichSiderealTime(t *testing.T) {
	tests := []struct {
		day                      float64
		month, year, hrs, min    int
		sec                      float64
		expectedHrs, expectedMin int
		expectedSec              float64
	}{
		{22, 04, 1980, 14, 36, 51.67, 4, 40, 5.23},
		{1, 1, 2000, 0, 0, 0.0, 6, 39, 52.27},      // Test case for start of epoch year
		{31, 12, 1999, 23, 59, 59.0, 6, 39, 51.26}, // Test case for end of year before epoch
	}

	for _, test := range tests {
		GHrs, GMin, GSec, _ := datetime.ConvertUniversalTimeToGreenwichSiderealTime(test.day, test.month, test.year, test.hrs, test.min, test.sec)
		if GHrs != test.expectedHrs || GMin != test.expectedMin || math.Abs(GSec-test.expectedSec) > tolerance {
			t.Fatalf("Error while converting Universal Time to Greenwich Sidereal Time. Expected: %d:%d:%f    Got: %d:%d:%f",
				test.expectedHrs, test.expectedMin, test.expectedSec, GHrs, GMin, GSec)
		}
	}
}

// TestConvertGreenwichSiderealTimeToUniversalTime tests the conversion of Greenwich sidereal time to universal time with various input cases.
func TestConvertGreenwichSiderealTimeToUniversalTime(t *testing.T) {
	tests := []struct {
		day                      float64
		month, year, hrs, min    int
		sec                      float64
		expectedHrs, expectedMin int
		expectedSec              float64
	}{
		{22, 04, 1980, 4, 40, 5.23, 14, 36, 51.67},
		{1, 1, 2000, 6, 39, 43.8, 23, 55, 55.64}, // Test case for start of epoch year
		{31, 12, 1999, 6, 39, 42.8, 0, 3, 46.46}, // Test case for end of year before epoch
	}

	for _, test := range tests {
		GHrs, GMin, GSec := datetime.ConvertGreenwichSiderealTimeToUniversalTime(test.day, test.month, test.year, test.hrs, test.min, test.sec)
		if GHrs != test.expectedHrs || GMin != test.expectedMin || (GSec-test.expectedSec) > tolerance {
			t.Fatalf("Error while converting Greenwich Sidereal to Universal Time. Expected: %d:%d:%f    Got: %d:%d:%f",
				test.expectedHrs, test.expectedMin, test.expectedSec, GHrs, GMin, GSec)
		}
	}
}

// TestCalculateLocalSiderealTimeUsingGreenwichSiderealTime tests the calculation of local sidereal time using Greenwich sidereal time.
func TestCalculateLocalSiderealTimeUsingGreenwichSiderealTime(t *testing.T) {
	tests := []struct {
		GHrs, GMin                     int
		GSec, longitude                float64
		expectedLSTHrs, expectedLSTMin int
		expectedLSTSec                 float64
	}{
		{4, 40, 5.23, -64, 0, 24, 5.23},
		{6, 39, 43.8, 0, 6, 39, 43.8},    // Greenwich meridian (0° longitude)
		{6, 39, 43.8, 180, 18, 39, 43.8}, // Opposite side of the Earth
	}

	for _, test := range tests {
		LSTHrs, LSTMin, LSTSec, _ := datetime.CalculateLocalSiderealTimeUsingGreenwichSiderealTime(test.GHrs, test.GMin, test.GSec, test.longitude)
		if LSTHrs != test.expectedLSTHrs || LSTMin != test.expectedLSTMin || (LSTSec-test.expectedLSTSec) > tolerance {
			t.Fatalf("Error while converting Local Sidereal Time Using Greenwich Sidereal Time. Expected: %d:%d:%f    Got: %d:%d:%f",
				test.expectedLSTHrs, test.expectedLSTMin, test.expectedLSTSec, LSTHrs, LSTMin, LSTSec)
		}
	}
}

// TestCalculateGreenwichSiderealTimeUsingLocalSiderealTime tests the calculation of Greenwich sidereal time using local sidereal time.
func TestCalculateGreenwichSiderealTimeUsingLocalSiderealTime(t *testing.T) {
	tests := []struct {
		LSTHrs, LSTMin             int
		LSTSec, longitude          float64
		expectedGHrs, expectedGMin int
		expectedGSec               float64
	}{
		{0, 24, 5.23, -64, 4, 40, 5.23},
		{6, 39, 43.8, 0, 6, 39, 43.8},    // Greenwich meridian (0° longitude)
		{18, 39, 43.8, 180, 6, 39, 43.8}, // Opposite side of the Earth
	}

	for _, test := range tests {
		GHrs, GMin, GSec, _ := datetime.CalculateGreenwichSiderealTimeUsingLocalSiderealTime(test.LSTHrs, test.LSTMin, test.LSTSec, test.longitude)
		if GHrs != test.expectedGHrs || GMin != test.expectedGMin || (GSec-test.expectedGSec) > tolerance {
			t.Fatalf("Error while converting Greenwich Sidereal Time Using Local Sidereal Time. Expected: %d:%d:%f    Got: %d:%d:%f",
				test.expectedGHrs, test.expectedGMin, test.expectedGSec, GHrs, GMin, GSec)
		}
	}
}
