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
