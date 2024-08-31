package sun

import (
	"go-astronomy/internal/macros"
)

func CalculatePositionOfSun(GDay, GMonth, GYear float64, UTHrs, UTMins int, UTSec float64) (float64, float64, float64, float64, float64, float64, float64) {

	lambda := macros.CalculatePositionOfSunHelper(GDay, GMonth, GYear, UTHrs, UTMins, UTSec)
	lambdaDeg, lambdaMin, lambdaSec := macros.ConvertDecimalDegToDegMinSec(lambda)
	raHrs, raMin, raSec, decDeg, decMin, decSec := macros.ConvertEclipticCoordinatesToEquatorial(GDay, int(GMonth), int(GYear), lambdaDeg, lambdaMin, lambdaSec, 0, 0, 0)
	return raHrs, raMin, raSec, decDeg, decMin, decSec, lambda
}
