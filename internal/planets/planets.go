package planets

import (
	"go-astronomy/internal/macros"
	"math"
)

func GetApproximatePositionOfPlanet(lctHour, lctMin, lctSec float64, isDaylightSaving bool, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear float64, planetName string) (planetRAHour, planetRAMin int, planetRASec float64, planetDecDeg, planetDecMin int, planetDecSec float64) {
	/*
		Calculate approximate position of a planet.

		Arguments:
			lct_hour -- Local civil time, in hours.
			lct_min -- Local civil time, in minutes.
			lct_sec -- Local civil time, in seconds.
			is_daylight_saving -- Is daylight savings in effect?
			zone_correction_hours -- Time zone correction, in hours.
			local_date_day -- Local date, day part.
			local_date_month -- Local date, month part.
			local_date_year -- Local date, year part.
			planet_name -- Name of planet, e.g., "Jupiter"

		Returns:
			planet_ra_hour -- Right ascension of planet (hour part)
			planet_ra_min -- Right ascension of planet (minutes part)
			planet_ra_sec -- Right ascension of planet (seconds part)
			planet_dec_deg -- Declination of planet (degrees part)
			planet_dec_min -- Declination of planet (minutes part)
			planet_dec_sec -- Declination of planet (seconds part)
	*/
	var daylightSaving float64
	if isDaylightSaving {
		daylightSaving = 1
	}

	planetData := GetPlanetData(planetName)
	planetTp := planetData["Tp"].(float64)
	planetLong := planetData["Long"].(float64)
	planetPeri := planetData["Peri"].(float64)
	planetEcc := planetData["Ecc"].(float64)
	planetAxis := planetData["Axis"].(float64)
	planetIncl := planetData["Incl"].(float64)
	planetNode := planetData["Node"].(float64)

	gdateDay := macros.ComputeGreenwichDayForLT(lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	gdateMonth := macros.ComputeGreenwichMonthForLT(lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	gdateYear := macros.ComputeGreenwichYearForLT(lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)

	utHours := macros.ConvertLocalTimeToUTC(lctHour, lctMin, lctSec, daylightSaving, zoneCorrectionHours, localDateDay, localDateMonth, localDateYear)
	dDays := macros.ConvertGregorianToJulian(gdateDay+(utHours/24), gdateMonth, gdateYear) - macros.ConvertGregorianToJulian(0, 1, 2010)
	npDeg1 := 360 * dDays / (365.242191 * planetTp)
	npDeg2 := npDeg1 - 360*math.Floor(npDeg1/360)
	mpDeg := npDeg2 + planetLong - planetPeri
	lpDeg1 := npDeg2 + (360 * planetEcc * math.Sin(macros.ConvertDegreesToRadians(mpDeg)) / math.Pi) + planetLong
	lpDeg2 := lpDeg1 - 360*math.Floor(lpDeg1/360)
	planetTrueAnomalyDeg := lpDeg2 - planetPeri
	rAU := planetAxis * (1 - (planetEcc * planetEcc)) / (1 + planetEcc*math.Cos(macros.ConvertDegreesToRadians(planetTrueAnomalyDeg)))

	earthData := GetPlanetData("Earth")
	earthTp := earthData["Tp"].(float64)
	earthLong := earthData["Long"].(float64)
	earthPeri := earthData["Peri"].(float64)
	earthEcc := earthData["Ecc"].(float64)
	earthAxis := earthData["Axis"].(float64)

	neDeg1 := 360 * dDays / (365.242191 * earthTp)
	neDeg2 := neDeg1 - 360*math.Floor(neDeg1/360)
	meDeg := neDeg2 + earthLong - earthPeri
	leDeg1 := neDeg2 + earthLong + 360*earthEcc*math.Sin(macros.ConvertDegreesToRadians(meDeg))/math.Pi
	leDeg2 := leDeg1 - 360*math.Floor(leDeg1/360)
	earthTrueAnomalyDeg := leDeg2 - earthPeri
	rAU2 := earthAxis * (1 - (earthEcc * earthEcc)) / (1 + earthEcc*math.Cos(macros.ConvertDegreesToRadians(earthTrueAnomalyDeg)))

	lpNodeRad := macros.ConvertDegreesToRadians(lpDeg2 - planetNode)
	psiRad := math.Asin(math.Sin(lpNodeRad) * math.Sin(macros.ConvertDegreesToRadians(planetIncl)))
	y := math.Sin(lpNodeRad) * math.Cos(macros.ConvertDegreesToRadians(planetIncl))
	x := math.Cos(lpNodeRad)
	ldDeg := macros.ConvertRadiansToDegrees(math.Atan2(y, x)) + planetNode
	rdAU := rAU * math.Cos(psiRad)
	leLdRad := macros.ConvertDegreesToRadians(leDeg2 - ldDeg)

	atan2Type1 := math.Atan2(rdAU*math.Sin(leLdRad), rAU2-rdAU*math.Cos(leLdRad))
	atan2Type2 := math.Atan2(rAU2*math.Sin(-leLdRad), rdAU-rAU2*math.Cos(leLdRad))
	var aRad float64
	if rdAU < 1 {
		aRad = atan2Type1
	} else {
		aRad = atan2Type2
	}
	var lamdaDeg1 float64
	if rdAU < 1 {
		lamdaDeg1 = 180 + leDeg2 + macros.ConvertRadiansToDegrees(aRad)
	} else {
		lamdaDeg1 = macros.ConvertRadiansToDegrees(aRad) + ldDeg
	}
	lamdaDeg2 := lamdaDeg1 - 360*math.Floor(lamdaDeg1/360)
	betaDeg := macros.ConvertRadiansToDegrees(math.Atan(rdAU * math.Tan(psiRad) * math.Sin(macros.ConvertDegreesToRadians(lamdaDeg2-ldDeg)) / (rAU2 * math.Sin(-leLdRad))))

	raHours := macros.ConvertDecimalDegToHours(macros.CalculateEclipticRightAscension(lamdaDeg2, 0, 0, betaDeg, 0, 0, gdateDay, gdateMonth, gdateYear))
	decDeg := macros.CalculateEclipticDeclination(lamdaDeg2, 0, 0, betaDeg, 0, 0, gdateDay, gdateMonth, gdateYear)

	planetRAHour = macros.GetHourFromDecimalHour(raHours)
	planetRAMin = macros.GetMinutesFromDecimalHours(raHours)
	planetRASec = macros.GetSecondsFromDecimalHours(raHours)
	planetDecDeg = macros.GetDegreeOfDecimalDeg(decDeg)
	planetDecMin = macros.GetMinOfDecimalDeg(decDeg)
	planetDecSec = macros.GetSecOfDecimalDeg(decDeg)

	return planetRAHour, planetRAMin, planetRASec, planetDecDeg, planetDecMin, planetDecSec
}

func GetPrecisePositionOfPlanet(lct_hour, lct_min, lct_sec, is_daylight_saving, zone_correction_hours, local_date_day, local_date_month, local_date_year, planet_name float64) {

}

func GetVisualAspectsOfPlanet(lct_hour, lct_min, lct_sec, is_daylight_saving, zone_correction_hours, local_date_day, local_date_month, local_date_year, planet_name float64) {

}
