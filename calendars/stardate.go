// STARDATE (Yes, from Star Trek™) CALENDAR
/*

The science fiction universe of Star Trek™ introduced a calendar system which
was simultaneously futuristic (for the time) and arbitrary. Over the decades
since its initial use on screen, especially with the growing popularity of the
franchise, the "stardate" has been analyzed and explored and refined into a
number of different variants, each trying to reconcile the arbitrariness of the
on-screen system into something consistent and usable.

This calendar system implementation is an attempt to combine all these variants
into a single system, using the format parameter to select which variant to use.
It was originally ported in 2018 from code by Aaron Chong (2015 version), under
provisions of the MIT License. My thanks for Aaron's use of the MIT License on
the original code, which allowed me to port it cleanly and legally.

Original source: http://rinsanity.weebly.com/files/theme/stardate_public.js

Supported Input Types:
  - string
  - []byte
  - int
  - float64
  - math/big.Float

Supported Format Strings:
  - main         - One of the older, more widely-accepted variants. Alternately
   								 called the "issue number style" stardate, it's a combined
   								 TOS/TNG variant, and the one used by Google Calendar. It was
   								 originally devised by Anhrew Main in CE 1994, with revisions
   								 made through CE 1997. See
   								 http://starchive.cs.umanitoba.ca/?stardates/ for the full
   								 explanation of this variant.
  - kennedy      - In 2006, Richie Kennedy released another combined variant,
   this one designed to have a single continuous count, more
   like the Julian Day Count than Main's issue number system.
  - pugh90s      - Steve Pugh devised 2 separate variants, one of them in the
  								 1990s, and the other later on. The original version used an
   unadjusted Gregorian year as the basis for the duration of a
   given range of stardates, meaning that 0.05 units refer to a
   larger duration of time during a leap year than it would
   otherwise.
  - pughfixed    - The later of Steve Pugh's systems noted the discrepancy, and
   opted to adjust the year length value to the actual average
   length of a Gregorian year, 365.2425 days. This means 0.05
   units are always the same duration, but does mean that the
   Gregorian year doesn't generally start at the same point in
   consecutive stardate ranges.
  - schmidt      - A joint effort between Andreas Schmidt and Graham Kennedy,
   this variant only covers TNG-era stardates, and while it can
   be used proleptically, it ignores the alternate format used
   prior to TNG.
  - guide-equiv  - One of five variants proposed by TrekGuide.com, this is the
   "out-of-universe equivalent" calculation. It isn't intended
   to be accurate for any use other than personal entertainment.
  - guide-tng    - The second of the five TrekGuide variants, this one is the
   current scale listed for TNG-era stardates, and is
   show-accurate (or at least as close to it as feasible with an
   entirely arbitrary system). Note, however, that it is only
   accurate for TNG-era dates.
  - guide-tos    - The third variant, then, covers the TOS-era stardates. Again,
   it is only accurate to the TOS era.
  - guide-oldtng - The fourth variant is no longer displayed on the TrekGuide
   site, and was actually pulled from a previous version of the
   stardates page. It covers the TNG era only, and uses slightly
   different numbers in its calculations than the current
   approach - specifically, it assumes Earth years cover 1000
   stardates.
  - guide-oldtos - NOTE: The fifth TrekGuide variant actually isn't implemented,
   yet. Representing the very first set of calculations
   available in archives of the TrekGuide site, it assumes that
   1000 stardates are one Earth year in the TOS era, and
   calculates dates based on that assumption. This variant was
   replaced within seven months of that first archival, after it
   was noticed that TOS-era stardates don't fit a 1000-stardate
   model.
  - aldrich      - A proof of concept originally written in C#, this variant
 									 results in dates very close to those produced by Pugh's and
 									 Schmidt's, but uses a more simplified calculation to do it.
 	- red-dragon   - A system devised by/for the Red Dragon Inn roleplaying forum
 									 site, it uses a fixed ratio of roughly two and three
 									 quarters stardates per Earth day. It makes no representations
   about accuracy outside the context of the site itself.
 	- sto-hynes    - John Hynes, creator of the Digital Time site, offers a
 									 calculation for STO stardates which appears to be the most
 									 accurate variant for those interested in generating those.
 									 The system doesn't represent itself as accurate outside the
 									 game, but is intentionally proleptic.
  - sto-academy  - Based on an online calculator provided by the STO Academy
   game help site, it is only accurate for stardates within the
   game, and does not offer to calculate dates for the rest of
   the franchise.
  - sto-tom      - Another variant intended only to calculate STO stardates,
   this one was attributed to Major Tom, and hosted as a Wolfram
   Alpha widget.
  - sto-anthodev - Another STO variant, hosted on GitHub.
*/
package calendars

import (
	"fmt"
	"math/big"
)

var jdcBaseGregorian = big.NewFloat(1721425.5)

var stardateMainCutoffs = []big.Float{
	*big.NewFloat(2510717.5),
	*big.NewFloat(2550185.5),
	*big.NewFloat(2555185.5),
	*big.NewFloat(2569517.5),
}

var stardateMainValues = []big.Float{
	*big.NewFloat(0.0),
	*big.NewFloat(197340.0),
	*big.NewFloat(197840.0),
	*big.NewFloat(2100000.0),
}

var stardateMainRates = []big.Float{
	*big.NewFloat(5.0),
	*big.NewFloat(0.1),
	*big.NewFloat(0.5),
	*big.NewFloat(1000.0 / 365.2425),
}

var stardateKennedyCutoffs = []big.Float{
	*big.NewFloat(2548704.5), // 2266-01-06
	*big.NewFloat(2550229.5), // 2270-03-11
	*big.NewFloat(2555479.5), // 2284-07-25
	*big.NewFloat(2575499.5), // 2339-05-19
}

var stardateKennedyValues = []big.Float{
	*big.NewFloat(0.0),
	*big.NewFloat(7320.0),
	*big.NewFloat(8076.0),
	*big.NewFloat(17685.6),
}

var stardateKennedyRates = []big.Float{
	*big.NewFloat(4.8),
	*big.NewFloat(0.144),
	*big.NewFloat(0.48),
	*big.NewFloat(144.0 / 55.0),
}

var stardatePughEpoch = big.NewFloat(2569517.5)

var stardateFixedRateCutoffs = map[string]*big.Float{
	"guide-equiv":  big.NewFloat(2446991.5),
	"guide-tng":    big.NewFloat(2567877.0),
	"guide-tos":    big.NewFloat(2548448.5),
	"guide-oldtng": big.NewFloat(2569296.5),
	// "guide-oldtos": big.NewFloat(2569296.5),
	"aldrich":      big.NewFloat(2569517.5),
	"red-dragon":   big.NewFloat(2569517.5),
	"sto-academy":  big.NewFloat(2455340.5),
	"sto-tom":      big.NewFloat(2423199.5),
	"sto-anthodev": big.NewFloat(2423199.5),
}

var stardateFixedRateRates = map[string]*big.Float{
	"guide-equiv":  big.NewFloat(1000 / 365.25),
	"guide-tng":    big.NewFloat(86400 / 34367.0564),
	"guide-tos":    big.NewFloat(2635.10833 / 365.2422),
	"guide-oldtng": big.NewFloat(1000 / 365.2422),
	// "guide-oldtos": big.NewFloat(1000 / 365.2422),
	"aldrich":      big.NewFloat(1000 / 365.2422),
	"red-dragon":   big.NewFloat(2.73973),
	"sto-academy":  big.NewFloat(2.736),
	"sto-tom":      big.NewFloat(2.7378508),
	"sto-anthodev": big.NewFloat(1000 / 365.2527),
}

var stardateFixedRateOffsets = map[string]*big.Float{
	"guide-equiv":  big.NewFloat(41000),
	"guide-tng":    big.NewFloat(0),
	"guide-tos":    big.NewFloat(0),
	"guide-oldtng": big.NewFloat(0),
	// "guide-oldtos": big.NewFloat(0),
	"aldrich":      big.NewFloat(0),
	"red-dragon":   big.NewFloat(0),
	"sto-academy":  big.NewFloat(87998.3079),
	"sto-tom":      big.NewFloat(0),
	"sto-anthodev": big.NewFloat(0),
}

func init() {
	RegisterElements(
		// name
		"stardate",
		stardateToInternal,
		stardateFromInternal,
		stardateOffset,
		// defaultFormat
		"main",
	)
}

func stardateToInternal(date interface{}, format string) (stamp TAI64NAXURTime, err error) {
	var jdc big.Float
	var in string

	switch date.(type) {
	// TODO - other types
	case big.Float:
		tmp := date.(big.Float)
		in = tmp.String()
	case *big.Float:
		in = date.(*big.Float).String()
	case float64:
		in = fmt.Sprintf("%f", date.(float64))
	case int:
		in = fmt.Sprintf("%d", date.(int))
	case []byte:
		in = string(date.([]byte))
	case string:
		in = date.(string)
	default:
		err = ErrUnsupportedInput
		return
	}

	switch format {
	case "main":
		jdc = stardateMainToJDC(in)
	case "kennedy":
		jdc = stardateKennedyToJDC(in)
	case "pugh90s":
		jdc = stardatePughToJDC(in, false)
	case "pughfixed":
		jdc = stardatePughToJDC(in, true)
	case "schmidt":
		jdc = stardateSchmidtToJDC(in)
	case "guide-equiv":
		jdc = stardateFixedRateToJDC(in, format)
	case "guide-tng":
		jdc = stardateFixedRateToJDC(in, format)
	case "guide-tos":
		jdc = stardateFixedRateToJDC(in, format)
	case "guide-oldtng":
		jdc = stardateFixedRateToJDC(in, format)
	case "guide-oldtos":
		err = ErrInvalidFormat
	case "aldrich":
		jdc = stardateFixedRateToJDC(in, format)
	case "red-dragon":
		jdc = stardateFixedRateToJDC(in, format)
	case "sto-hynes":
		jdc = stardateSTOHynesToJDC(in)
	case "sto-academy":
		jdc = stardateFixedRateToJDC(in, format)
	case "sto-tom":
		jdc = stardateFixedRateToJDC(in, format)
	case "sto-anthodev":
		jdc = stardateFixedRateToJDC(in, format)
	default:
		err = ErrInvalidFormat
	}
	if err != nil {
		return
	}

	stamp, err = jdcToInternal(jdc, "full")

	return
}

func stardateFromInternal(stamp TAI64NAXURTime, format string) (date string, err error) {
	var jdcFloatP *big.Float
	var jdcFloat big.Float
	var jdcString string
	jdcString, err = jdcFromInternal(stamp, "full")
	if err != nil {
		return
	}
	jdcFloatP, _, _ = big.ParseFloat(jdcString, 10, 188, big.ToNearestAway)
	jdcFloat = *jdcFloatP

	switch format {
	case "main":
		date = stardateJDCToMain(jdcFloat)
	case "kennedy":
		date = stardateJDCToKennedy(jdcFloat)
	case "pugh90s":
		date = stardateJDCToPugh(jdcFloat, false)
	case "pughfixed":
		date = stardateJDCToPugh(jdcFloat, true)
	case "schmidt":
		date = stardateJDCToSchmidt(jdcFloat)
	case "guide-equiv":
		date = stardateJDCToFixedRate(jdcFloat, format)
	case "guide-tng":
		date = stardateJDCToFixedRate(jdcFloat, format)
	case "guide-tos":
		date = stardateJDCToFixedRate(jdcFloat, format)
	case "guide-oldtng":
		date = stardateJDCToFixedRate(jdcFloat, format)
	case "guide-oldtos":
		err = ErrInvalidFormat
	case "aldrich":
		date = stardateJDCToFixedRate(jdcFloat, format)
	case "red-dragon":
		date = stardateJDCToFixedRate(jdcFloat, format)
	case "sto-hynes":
		date = stardateJDCToSTOHynes(jdcFloat)
	case "sto-academy":
		date = stardateJDCToFixedRate(jdcFloat, format)
	case "sto-tom":
		date = stardateJDCToFixedRate(jdcFloat, format)
	case "sto-anthodev":
		date = stardateJDCToFixedRate(jdcFloat, format)
	default:
		err = ErrInvalidFormat
	}

	return
}

func stardateOffset(in TAI64NAXURTime, offset interface{}) (out TAI64NAXURTime, err error) {
	var date, format, mod string
	var adjust TAI64NAXURTime

	switch offset.(type) {
	case []byte:
		mod = string(offset.([]byte))
	case string:
		mod = offset.(string)
	default:
		err = ErrUnsupportedInput
		return
	}

	_, err = fmt.Sscanf(mod, "%s %s", &date, &format)
	if err != nil && err.Error() != "EOF" {
		return
	}
	err = nil

	adjust, err = stardateToInternal(date, format)
	if err != nil {
		adjust, err = stardateToInternal(format, date)
		if err != nil {
			return
		}
	}

	out = in.Add(adjust)

	return
}

// Internal Logic

func stardateMainToJDC(stardate string) big.Float {
	var region, issue int
	var sdFloat64 float64
	var sdBigFloat = big.NewFloat(0.0)
	var tmp = big.NewFloat(0.0).SetMode(big.ToNegativeInf)

	fmt.Sscanf(stardate, "[%d]%f", &issue, &sdFloat64)
	sdBigFloat = big.NewFloat(sdFloat64)

	if issue >= 21 {
		region = 3
		sdBigFloat.Add(sdBigFloat, big.NewFloat(float64(100000*issue)))
	} else {
		sdBigFloat.Add(sdBigFloat, big.NewFloat(float64(10000*issue)))
		region = 0
		for region < 2 && sdBigFloat.Cmp(&stardateMainValues[region+1]) >= 0 {
			region++
		}
	}

	return *sdBigFloat.Add(&stardateMainCutoffs[region], tmp.Quo(tmp.Sub(sdBigFloat, &stardateMainValues[region]), &stardateMainRates[region]))
}

func stardateJDCToMain(jdc big.Float) string {
	var region = 0
	var issue int64
	var stardate = big.NewFloat(0.0)
	var tmp = big.NewFloat(0.0).SetMode(big.ToNegativeInf)
	var format string

	for region < 3 && jdc.Cmp(&stardateMainCutoffs[region+1]) >= 0 {
		region++
	}

	stardate.Add(&stardateMainValues[region], tmp.Mul(tmp.Sub(&jdc, &stardateMainCutoffs[region]), &stardateMainRates[region]))

	if region == 3 {
		issue, _ = tmp.Quo(stardate, big.NewFloat(100000.0)).Int64()
		stardate.Sub(stardate, big.NewFloat(float64(100000*issue)))
		format = "[%d]%05.6g"
	} else {
		issue, _ = tmp.Quo(stardate, big.NewFloat(10000.0)).Int64()
		stardate.Sub(stardate, big.NewFloat(float64(10000*issue)))
		if stardate.Cmp(big.NewFloat(0)) < 0 {
			issue--
			stardate.Add(stardate, big.NewFloat(10000))
		}
		format = "[%d]%04.6g"
	}

	return fmt.Sprintf(format, issue, stardate)
}

func stardateKennedyToJDC(stardate string) big.Float {
	var region = 0
	var sdBigFloat = big.NewFloat(0.0)
	var tmp = big.NewFloat(0.0).SetMode(big.ToNegativeInf)

	sdBigFloat, _, _ = big.ParseFloat(stardate, 10, 188, big.ToNearestAway)

	for region < 3 && sdBigFloat.Cmp(&stardateKennedyValues[region+1]) >= 0 {
		region++
	}

	return *sdBigFloat.Add(&stardateKennedyCutoffs[region], tmp.Quo(tmp.Sub(sdBigFloat, &stardateKennedyValues[region]), &stardateKennedyRates[region]))
}

func stardateJDCToKennedy(jdc big.Float) string {
	var region = 0
	var stardate = big.NewFloat(0.0)
	var tmp = big.NewFloat(0.0).SetMode(big.ToNegativeInf)

	for region < 3 && jdc.Cmp(&stardateKennedyCutoffs[region+1]) >= 0 {
		region++
	}

	stardate.Add(&stardateKennedyValues[region], tmp.Mul(tmp.Sub(&jdc, &stardateKennedyCutoffs[region]), &stardateKennedyRates[region]))

	return stardate.String()
}

func stardateSchmidtToJDC(stardate string) big.Float {
	date, _, _ := big.ParseFloat(stardate, 10, 188, big.ToNearestAway)
	return yearfToJDC(*new(big.Float).Add(new(big.Float).Quo(date, big.NewFloat(1000.0)), big.NewFloat(2323.0)))
}

func stardateJDCToSchmidt(jdc big.Float) string {
	var yearf = jdcToYearf(jdc)

	return new(big.Float).Mul(new(big.Float).Sub(&yearf, big.NewFloat(2323.0)), big.NewFloat(1000.0)).String()
}

func stardateSTOHynesToJDC(stardate string) big.Float {
	date, _, _ := big.ParseFloat(stardate, 10, 188, big.ToNearestAway)
	return yearfToJDC(*new(big.Float).Add(new(big.Float).Sub(new(big.Float).Quo(date, big.NewFloat(1000.0)), new(big.Float).Quo(big.NewFloat(221.0), big.NewFloat(365.0))), big.NewFloat(1923.0)))
}

func stardateJDCToSTOHynes(jdc big.Float) string {
	var yearf = jdcToYearf(jdc)

	return new(big.Float).Mul(new(big.Float).Add(new(big.Float).Sub(&yearf, big.NewFloat(1923.0)), new(big.Float).Quo(big.NewFloat(221.0), big.NewFloat(365.0))), big.NewFloat(1000.0)).String()
}

func stardateFixedRateToJDC(stardate, mode string) big.Float {
	date, _, _ := big.ParseFloat(stardate, 10, 188, big.ToNearestAway)
	return *date.Add(date.Quo(date.Sub(date, stardateFixedRateOffsets[mode]), stardateFixedRateRates[mode]), stardateFixedRateCutoffs[mode])
}

func stardateJDCToFixedRate(jdc big.Float, mode string) string {
	return jdc.Add(stardateFixedRateOffsets[mode], jdc.Mul(jdc.Sub(&jdc, stardateFixedRateCutoffs[mode]), stardateFixedRateRates[mode])).String()
}

func stardatePughToJDC(stardate string, fixed bool) big.Float {
	var prefix int64
	var suffix = big.NewFloat(0.0)
	var sdBigFloat *big.Float

	sdBigFloat, _, _ = big.ParseFloat(stardate, 10, 188, big.ToNearestAway)

	prefix, _ = new(big.Float).Quo(sdBigFloat, big.NewFloat(1000.0)).Int64()
	suffix = new(big.Float).Abs(new(big.Float).Sub(sdBigFloat, new(big.Float).SetInt64(prefix*1000)))

	if fixed {
		return *new(big.Float).Add(stardatePughEpoch, new(big.Float).Mul(new(big.Float).Add(new(big.Float).SetInt64(prefix), new(big.Float).Quo(suffix, big.NewFloat(1000.0))), big.NewFloat(365.2425)))
	}

	return yearfToJDC(*new(big.Float).Add(new(big.Float).Add(new(big.Float).SetInt64(prefix), big.NewFloat(2323.0)), new(big.Float).Quo(suffix, big.NewFloat(1000.0))))
}

func stardateJDCToPugh(jdc big.Float, fixed bool) string {
	var prefix int64
	var suffix = big.NewFloat(0.0)
	var working = big.NewFloat(0.0)

	if fixed {
		working.Quo(new(big.Float).Sub(&jdc, stardatePughEpoch), big.NewFloat(365.2425))
	} else {
		var yearf = jdcToYearf(jdc)
		working.Sub(&yearf, big.NewFloat(2323.0))
	}

	prefix, _ = working.SetMode(big.ToNegativeInf).Int64()
	suffix = new(big.Float).Mul(new(big.Float).Sub(working, new(big.Float).SetInt64(prefix)), big.NewFloat(1000.0))
	return new(big.Float).Add(new(big.Float).SetInt64((prefix-2)*1000), suffix.Abs(suffix)).String()
}

// Even further-internal logic

func gregorianLeap(year int64) int64 {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return 1
	}

	return 0
}

func yearfToJDC(yearf big.Float) big.Float {
	var year, _ = yearf.SetMode(big.ToNegativeInf).Int64()
	var yearStart = new(big.Float).Add(jdcBaseGregorian, new(big.Float).SetInt64((365*(year-1))+((year-1)/4)-((year-1)/100)+((year-1)/400)))

	return *new(big.Float).Add(yearStart, new(big.Float).Mul(new(big.Float).Sub(&yearf, new(big.Float).SetInt64(year)), new(big.Float).SetInt64(365+gregorianLeap(year))))
}

func jdcToYearf(jdc big.Float) big.Float {
	var yearf = big.NewFloat(0.0)
	var dd = big.NewFloat(0.0)
	var dayf = big.NewFloat(0.0)
	var tmpI int64
	var cent int64
	var year int64

	tmpI, _ = new(big.Float).Sub(&jdc, big.NewFloat(0.5)).SetMode(big.ToNegativeInf).Int64()
	dd.Add(new(big.Float).SetInt64(tmpI), big.NewFloat(0.5))
	dayf.Sub(&jdc, dd)
	dd.Sub(dd, jdcBaseGregorian)
	cent, _ = new(big.Float).Quo(new(big.Float).Add(new(big.Float).Mul(dd, big.NewFloat(4.0)), big.NewFloat(3.0)), big.NewFloat(146097.0)).SetMode(big.ToNegativeInf).Int64()
	dd.Sub(dd, new(big.Float).SetInt64((cent*146097)/4))
	year, _ = new(big.Float).Quo(new(big.Float).Add(new(big.Float).Mul(dd, big.NewFloat(4.0)), big.NewFloat(3.0)), big.NewFloat(1461.0)).SetMode(big.ToNegativeInf).Int64()
	dayf.Add(dayf, new(big.Float).Sub(dd, new(big.Float).SetInt64(((year*1461)/4)+1)))
	year += ((100 * cent) + 1)
	return *yearf.Add(new(big.Float).SetInt64(year), new(big.Float).Quo(new(big.Float).Add(dayf, big.NewFloat(1.0)), new(big.Float).SetInt64(365+gregorianLeap(year))))
}
