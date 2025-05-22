Here’s the updated documentation with `Parameters` and `Returns` sections, and without the `Syntax` block:

---

## Function: `CalculateDateOfEaster`

### Description

Calculates the date of Easter Sunday for a given Gregorian calendar year using the **Anonymous Gregorian algorithm**, also known as the **Meeus/Jones/Butcher algorithm**. This method relies on solar and lunar calendar cycles to determine Easter's month and day.

### Parameters

* `year` (`float64`): The Gregorian calendar year for which Easter Sunday is to be calculated.

### Returns

* `day` (`float64`): The day of the month on which Easter Sunday falls.
* `month` (`float64`): The month (as a number, e.g., 4 for April) in which Easter Sunday falls.

---

## Function: `ConvertGreenwichDateToJulianDate`

### Description

Tests the conversion of a Gregorian calendar date to the corresponding Julian date. This function validates the correctness of the conversion logic against a variety of input scenarios.

### Parameters

* (Dependent on implementation): Likely includes year, month, day, and time components such as hours, minutes, and seconds — all as `float64` or `int` types.

### Returns

* `julianDate` (`float64`): The calculated Julian date corresponding to the provided Gregorian date.

---

## Function: `ConvertJulianDateToGreenwichDate`

### Description

Converts a given Julian date to the equivalent Gregorian calendar date. This function is useful for translating astronomical Julian dates into a standard calendar format recognizable for civil or historical date references.

### Parameters

* `julianDate` (`float64`): The Julian date to be converted. It typically represents the number of days since January 1, 4713 BCE (Julian calendar).

### Returns

* `day` (`float64`): The day of the Gregorian calendar corresponding to the input Julian date.
* `month` (`float64`): The month (1–12) of the converted date.
* `year` (`float64`): The year in the Gregorian calendar corresponding to the input Julian date.

---
