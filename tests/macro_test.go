package tests

import (
	"fmt"
	"go-astronomy/internal/macros"
	"strconv"
	"testing"
)

var testJulianDate string = "2460516.5000000"
var testDay int = 25
var testMonth int = 7
var testYear int = 2024
var testDayOfWeek string = "Thursday"
var testHrs int = 14
var testMin int = 14
var testSec int = 14
var testDecimalHrs float64 = 14.237222222222222

func TestConvertGregorianToJulian(t *testing.T) {
	// Test 25 July 2024
	julianDate := fmt.Sprintf("%.7f", macros.ConvertGregorianToJulian(testDay, 7, 2024))
	if julianDate != testJulianDate || julianDate == "" {
		t.Fatalf(`Error while converting UT to Julian. Required: %s    Got: %s`, testJulianDate, julianDate)
	}
}

func TestGetJulianDay(t *testing.T) {
	floatJulian, _ := strconv.ParseFloat(testJulianDate, 64)
	day := macros.GetJulianDay(floatJulian)

	if day == 0 || day != testDay {
		t.Fatalf(`Error while fetching day part from Julian. Required: %d    Got: %d`, testDay, day)
	}
}

func TestGetJulianMonth(t *testing.T) {
	floatJulian, _ := strconv.ParseFloat(testJulianDate, 64)
	month := macros.GetJulianMonth(floatJulian)
	if month == 0 || month != testMonth {
		t.Fatalf(`Error while fetching month part from Julian. Required: %d    Got: %d`, testMonth, month)
	}
}

func TestGetJulianYear(t *testing.T) {
	floatJulian, _ := strconv.ParseFloat(testJulianDate, 64)
	year := macros.GetJulianYear(floatJulian)
	if year == 0 || year != 2024 {
		t.Fatalf(`Error while fetching year part from Julian. Required: %d    Got: %d`, testYear, year)
	}
}

func TestGetDayOfWeekFromJulian(t *testing.T) {
	floatJulian, _ := strconv.ParseFloat(testJulianDate, 64)
	dayOfWeek := macros.GetDayOfWeekFromJulian(floatJulian)

	if dayOfWeek == "" || dayOfWeek != testDayOfWeek {
		t.Fatalf(`Error while fetching day of week from Julian. Required: %s    Got: %s`, testDayOfWeek, dayOfWeek)

	}
}

func TestConvertTimeToDecimal(t *testing.T) {
	decimalHours := macros.ConvertTimeToDecimal(testHrs, testMin, testSec)
	if decimalHours == 0.0 || decimalHours != testDecimalHrs {
		t.Fatalf(`Error while converting time to decimal. Required: %f    Got: %f`, testDecimalHrs, decimalHours)
	}
}

func TestGetHourFromDecimalHour(t *testing.T) {
	decimalHrs := macros.GetHourFromDecimalHour(testDecimalHrs)
	if decimalHrs == 0 || decimalHrs != int(testHrs) {
		t.Fatalf(`Error while fetching Hours from decimal Hrs. Required: %d    Got: %d`, testHrs, decimalHrs)
	}
}

func TestGetMinutesFromDecimalHours(t *testing.T) {
	decimalMins := macros.GetMinutesFromDecimalHours(testDecimalHrs)
	if decimalMins == 0 || decimalMins != int(testMin) {
		t.Fatalf(`Error while fetching Minutes from decimal Hrs. Required: %d    Got: %d`, testMin, decimalMins)
	}
}

func TestGetSecondsFromDecimalHours(t *testing.T) {
	decimalSecs := macros.GetSecondsFromDecimalHours(testDecimalHrs)
	if decimalSecs == 0 || decimalSecs != int(testSec) {
		t.Fatalf(`Error while fetching Seconds from decimal Hrs. Required: %d    Got: %d`, testMin, decimalSecs)
	}
}
