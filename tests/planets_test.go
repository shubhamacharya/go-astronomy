package tests

var lctHour float64 = 11
var lctMin float64 = 52
var lctSec float64 = 0
var isDaylightSaving bool = false
var zoneCorrectionHrs float64 = 0
var localDateDay float64 = 8
var localDateMonth float64 = 8
var localDateYear float64 = 2024
var planetName string = "Mars"

// func TestGetApproximatePositionOfPlanet(t *testing.T) {
// 	// Test 08 August 2024
// 	planetRAHour, planetRAMin, planetRASec, planetDecDeg, planetDecMin, planetDecSec := planets.GetApproximatePositionOfPlanet(lctHour, lctMin, lctSec,
// 		isDaylightSaving, zoneCorrectionHrs,
// 		localDateDay, localDateMonth, localDateYear, planetName)

// 	fmt.Printf("RA : %d %d %.2f\t\t Dec : %d %d %.2f\n", planetRAHour, planetRAMin, planetRASec, planetDecDeg, planetDecMin, planetDecSec)

// 	if planetRAHour != 4 || planetRAMin != 43 || planetRASec != 11.56 ||
// 		planetDecDeg != 21 || planetDecMin != 45 || planetDecSec != 47.00 {
// 		t.Fatalf(`Wrong Planetory  Positions.`)
// 	}
// }

// func TestGetPrecisePositionOfPlanet(t *testing.T) {
// 	// Test 08 August 2024
// 	planetRAHour, planetRAMin, planetRASec, planetDecDeg, planetDecMin, planetDecSec := planets.GetPrecisePositionOfPlanet(lctHour, lctMin, lctSec,
// 		isDaylightSaving, zoneCorrectionHrs,
// 		localDateDay, localDateMonth, localDateYear, planetName)

// 	fmt.Printf("RA : %d %d %.2f\t\t Dec : %d %d %.2f\n", planetRAHour, planetRAMin, planetRASec, planetDecDeg, planetDecMin, planetDecSec)

// 	// if planetRAHour != 4 || planetRAMin != 43 || planetRASec != 11.56 ||
// 	// 	planetDecDeg != 21 || planetDecMin != 45 || planetDecSec != 47.00 {
// 	// 	t.Fatalf(`Wrong Planetory  Positions.`)
// 	// }
// }
