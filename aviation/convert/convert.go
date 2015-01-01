// Package convert provides basic aviation conversion utilities.
package convert

// A statute mile is 5,280 feet, as defined in 14 CFR Part 298.2 and mentioned in the
// Pilot's Handbook of Aeronautical Knowledge (http://www.faa.gov/regulations_policies/handbooks_manuals/aviation/pilot_handbook/media/PHAK%20-%20Chapter%2015.pdf).
const StatuteMileInFeet = 5280

// A nautical mile is mentioned as 6076.1 in the PHAK (http://www.faa.gov/regulations_policies/handbooks_manuals/aviation/pilot_handbook/media/PHAK%20-%20Chapter%2015.pdf).
// The NGA calls it 6,076.11549 (http://msi.nga.mil/MSISiteContent/StaticFiles/NAV_PUBS/DBP/endtables.pdf).
// Wikipedia calls it 6,076.12 (http://en.wikipedia.org/wiki/Nautical_mile#Conversions_to_other_units).
const NauticalMileInFeet = 6076.1

const statuteMileConversionFactor = StatuteMileInFeet / NauticalMileInFeet

const nauticalMileConversionFactor = NauticalMileInFeet / StatuteMileInFeet

// Convert statute miles to nautical miles.
func StatuteToNauticalMiles(statute float64) float64 {
	return statute * nauticalMileConversionFactor
}

// Convert nautical miles to statute miles.
func NauticalToStatuteMiles(nautical float64) float64 {
	return nautical * statuteMileConversionFactor
}
