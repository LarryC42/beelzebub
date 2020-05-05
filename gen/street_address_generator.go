package gen

import (
	"fmt"
	"math/rand"
	"strings"
)

// StreetAddressGenerator generates a prefix for a female name
func StreetAddressGenerator() string {
	var num = strings.TrimPrefix(StreetNumberGenerator()+" ", " ")
	var dir = strings.TrimPrefix(StreetDirectionGenerator()+" ", " ")
	var name = StreetNameGenerator()
	var suffix = strings.TrimSuffix(" "+StreetSuffixGenerator(), " ")
	var unit = strings.TrimSuffix(" "+UnitGenerator(), " ")

	return fmt.Sprintf("%s%s%s%s%s", num, dir, name, suffix, unit)
}

// StreetDirectionGenerator generates a prefix for a female name
func StreetDirectionGenerator() string {
	return streetDirectionData[rand.Intn(len(streetDirectionData))]
}

var streetDirectionData = []string{
	"N", "S", "E", "W", "NE", "NW", "SE", "SW",
	"", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "",
}

// StreetNameGenerator generates a prefix for a female name
func StreetNameGenerator() string {
	return streetNameData[rand.Intn(len(streetNameData))]
}

var streetNameData = []string{
	"10th", "11th", "12th", "13th", "1st", "2nd", "3rd",
	"4th", "5th", "6th", "7th", "8th", "9th", "Academy",
	"Adams", "Beech", "Belmont", "Berkshire", "Bridge", "Broad", "Broadway",
	"Brookside", "Buckingham", "Cambridge", "Cedar", "Center", "Central", "Charles",
	"Cherry", "Chestnut", "Church", "Clark", "Clinton", "College", "Colonial",
	"Court", "Creek", "Delaware", "Division", "Dogwood", "Durham", "Elizabeth",
	"Elm", "Essex", "Fairview", "Fairway", "Franklin", "Front", "George",
	"Grant", "Green", "Grove", "Hamilton", "Harrison", "Heather", "Hickory",
	"High", "Highland", "Hill", "Hillcrest", "Hillside", "Hilltop", "Holly",
	"Jackson", "James", "Jefferson", "King", "Lafayette", "Lake", "Lakeview",
	"Laurel", "Liberty", "Lincoln", "Linden", "Locust", "Madison", "Magnolia",
	"Main", "Maple", "Market", "Meadow", "Mill", "Monroe", "Mulberry",
	"Myrtle", "New", "Oak", "Park", "Pearl", "Penn", "Pennsylvania",
	"Pine", "Pleasant", "Poplar", "Primrose", "Prospect", "Railroad", "Ridge",
	"River", "Riverside", "School", "Shady", "Sherwood", "Spring", "Spruce",
	"State", "Summit", "Sunset", "Surrey", "Union", "Valley", "View",
	"Vine", "Virginia", "Walnut", "Warren", "Washington", "Water", "Williams",
	"Willow", "Winding", "Wood", "Woodland",
}

// StreetNumberGenerator generates a prefix for a female name
func StreetNumberGenerator() string {
	hasLetter := rand.Intn(10)
	number := fmt.Sprintf("%d", rand.Intn(1000))
	if hasLetter == 1 {
		number = fmt.Sprintf("%s-%s", number, (string)(rand.Intn(26)+65))
	}
	return number
}

// StreetSuffixGenerator generates a prefix for a female name
func StreetSuffixGenerator() string {
	return streetSuffixData[rand.Intn(len(streetSuffixData))]
}

var streetSuffixData = []string{
	"Alley", "Avenue", "Bayou", "Beach", "Bend", "Bluff",
	"Bottom", "Boulevard", "Branch", "Bridge", "Brook", "Burg",
	"Camp", "Canyon", "Cape", "Causeway", "Center",
	"Circle", "Cliff", "Club", "Common", "Corner", "Course", "Court",
	"Cove", "Creek", "Crescent", "Crest", "Crossing", "Crossroad", "Curve",
	"Dale", "Dam", "Divide", "Drive", "Estate",
	"Expressway", "Extension", "Fall", "Ferry", "Field", "Flat", "Ford",
	"Forest", "Forge", "Fork", "Fort", "Freeway", "Garden", "Gateway",
	"Glen", "Green", "Grove", "Harbor", "Haven", "Heights", "Highway",
	"Hill", "Hollow", "Inlet", "Island", "Isle", "Junction", "Key",
	"Knoll", "Lake", "Land", "Landing", "Lane", "Light", "Loaf",
	"Lock", "Lodge", "Loop", "Mall", "Manor", "Meadow",
	"Mill", "Mission", "Motorway", "Mount", "Mountain", "Neck", "Orchard",
	"Oval", "Overpass", "Park", "Parkway", "Pass", "Passage", "Path",
	"Pike", "Pine", "Place", "Plain", "Plaza", "Point", "Port",
	"Prairie", "Radial", "Ramp", "Ranch", "Rapid", "Rest", "Ridge",
	"River", "Road", "Route", "Row", "Rue", "Run", "Shoal",
	"Shore", "Skyway", "Spring", "Spur", "Square", "Station", "Stream",
	"Street", "Summit", "Terrace", "Throughway", "Trace", "Track", "Trafficway",
	"Trail", "Trailer", "Tunnel", "Turnpike", "Union", "Valley", "Viaduct",
	"View", "Village", "Ville", "Vista", "Walk", "Wall", "Way",
	"Well",
}

// UnitGenerator generates unit designations for addresses
func UnitGenerator() string {
	unit := unitData[rand.Intn(len(unitData))]
	if unit != "" {
		unit = fmt.Sprintf("%s %d", unit, rand.Intn(10))
		if rand.Intn(10) == 1 {
			unit = fmt.Sprintf("%s%s", unit, (string)(rand.Intn(26)+65))
		}
	}
	return unit
}

var unitData = []string{
	"Apt", "Bldg", "Dept", "Fl", "Rm", "Ste", "Unit",
	"", "", "", "", "", "", "",
	"", "", "", "", "", "", "",
	"", "", "", "", "", "", "",
	"", "", "", "", "", "", "",
	"", "", "", "", "", "", "",
	"", "", "", "", "", "", "",
	"", "", "", "", "", "", "",
	"", "", "", "", "", "", "",
	"", "", "", "", "", "", "",
	"", "", "", "", "", "", "",
}
