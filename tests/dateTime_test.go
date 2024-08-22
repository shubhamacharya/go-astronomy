package tests

import (
	datetime "go-astronomy/internal/dateTime"
	"math"
	"testing"
)

func TestCalculateDateofEaster(t *testing.T) {
	day, month := datetime.CalculateDateOfEaster(2009)
	if day != 12 || month != 4 {
		t.Fatalf(`Error while calculating Date of easter for year 2009. Required: %d-%d    Got: %d-%d`, 12, 4, day, month)
	}
}

func TestConvertToJulianDate(t *testing.T) {
	julianDate := datetime.ConvertGreenwichDateToJulianDate(19.75, 6, 2009)
	if julianDate != 2455002.25 {
		t.Fatalf(`Error while converting Greenwich Date to julian. Required: %f    Got: %f`, 2455002.25, julianDate)
	}
}

func TestConvertToGreenwichDate(t *testing.T) {
	day, month, year := datetime.ConvertJulianDateToGreenwichDate(2455002.25)
	if day != 19.75 || month != 6 || year != 2009 {
		t.Fatalf(`Error while converting julian to Greenwich Date. Required: %f-%f-%f    Got: %f-%f-%f`, 19.75, 6.0, 2009.0, day, month, year)
	}
}

func TestGetNameOfTheDayOfMonth(t *testing.T) {
	day := datetime.GetNameOfTheDayOfMonth(19.0, 6, 2009)
	if day != "Friday" {
		t.Fatalf(`Error while getting name of the day in the week. Required: %s    Got: %s`, "Friday", day)
	}
}

func TestConvertHrsMinSecToDecimalHrs(t *testing.T) {
	decimalHrs := math.Round(datetime.ConvertHrsMinSecToDecimalHrs(6, 31, 27.0, true, "PM")*1000000) / 1000000
	expectedOpt := float64(18.524167)
	if math.Abs(decimalHrs-expectedOpt) != 0 {
		t.Fatalf("Error while converting Hrs Min Sec to Decimal. Required: %f    Got: %f", expectedOpt, decimalHrs)
	}
}

func TestConvertDecimalHrsToHrsMinSec(t *testing.T) {
	Hrs, min, sec := datetime.ConvertDecimalHrsToHrsMinSec(18.524167)
	if math.Abs(Hrs-18.0) > 0.01 || math.Abs(min-31) > 0.01 || math.Abs(sec-27) > 0.01 {
		t.Fatalf("Error while converting Decimal to Hrs Min Sec\n. Required: %f %f %f  Got: %f %f %f", 18.0, 31.0, 27.0, Hrs, min, sec)
	}
}

func TestConvertLocalTimeToUniversalTime(t *testing.T) {
	UTDay, UTMon, UTYear, UTHrs, UTMin, UTSec, _ := datetime.ConvertLocalTimeToUniversalTime(1, 7, 2013, 3, 37, 0.0, 1, 0, 4)
	if UTDay != 30 && UTMon != 6 && UTYear != 2013 && UTHrs != 22 && UTMin != 36 && UTSec != 60 {
		t.Fatalf("Error while converting Local Time to Universal Time\n. Required: %f %f %f %f %f %f    Got: %f %f %f %f %f %f",
			30.0, 6.0, 2013.0, 22.0, 36.0, 60.0, UTDay, UTMon, UTYear, UTHrs, UTMin, UTSec)
	}
}

func TestConvertUniversalTimeToLocalTime(t *testing.T) {
	GDay, GMon, GYear, GHrs, GMin, GSec := datetime.ConvertUniversalTimeToLocalTime(30, 06, 2013, 22, 37, 0.0, 1, 0, 4)
	if GDay != 30 && GMon != 6 && GYear != 2013 && GHrs != 22 && GMin != 36 && GSec != 60 {
		t.Fatalf("Error while converting Universal Time to Local Time\n. Required: %f %f %f %f %f %f    Got: %f %f %f %f %f %f",
			1.0, 7.0, 2013.0, 3.0, 36.0, 60.0, GDay, GMon, GYear, GHrs, GMin, GSec)
	}
}

func TestConvertUniversalTimeToGreenwichSiderealTime(t *testing.T) {
	GHrs, GMin, GSec, _ := datetime.ConvertUniversalTimeToGreenwichSiderealTime(22, 04, 1980, 14, 36, 51.67)
	if GHrs != 4 && GMin != 40 && GSec != 5.23 {
		t.Fatalf("Error while converting Universal Time to Greenwich Sidereal Time\n. Required: %f %f %f    Got: %f %f %f",
			4.0, 40.0, 5.23, GHrs, GMin, GSec)
	}
}

func TestConvertGreenwichSiderealTimeToUniversalTime(t *testing.T) {
	GHrs, GMin, GSec := datetime.ConvertGreenwichSiderealTimeToUniversalTime(22, 04, 1980, 4, 40, 5.23)
	if GHrs != 14 && GMin != 36 && GSec != 51.67 {
		t.Fatalf("Error while converting Greenwich Sidereal to Universal Time\n. Required: %f %f %f    Got: %f %f %f",
			14.0, 36.0, 51.67, GHrs, GMin, GSec)
	}
}

func TestCalculateLocalSiderealTimeUsingGreenwichSideralTime(t *testing.T) {
	GHrs, GMin, GSec, _ := datetime.CalculateLocalSiderealTimeUsingGreenwichSideralTime(4, 40, 5.23, -64)
	if GHrs != 0 && GMin != 24 && GSec != 5.23 {
		t.Fatalf("Error while converting Greenwich Sidereal to Universal Time\n. Required: %f %f %f    Got: %f %f %f",
			0.0, 24.0, 5.23, GHrs, GMin, GSec)
	}
}
