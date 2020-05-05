package gen

import (
	"fmt"
	"math/rand"
)

// CityStateZipGenerator generates a City state zip line in the form: City, ST 99999
func CityStateZipGenerator() string {
	var state = StateNameGenerator()
	return fmt.Sprintf("%s, %s %s", CityNameGenerator(), state, ZipcodeGenerator(state))
}

// CityStateZipGenerator2 generates a City state zip line in the form: City, ST 99999 given a state
func CityStateZipGenerator2(state string) string {
	return fmt.Sprintf("%s, %s %s", CityNameGenerator(), state, ZipcodeGenerator(state))
}

// CityNameGenerator generates a prefix for a female name
func CityNameGenerator() string {
	return cityData[rand.Intn(len(cityData))]
}

var cityData = []string{
	"Abilene", "Akron", "Albuquerque", "Alexandria", "Allen", "Allentown", "Amarillo",
	"Anaheim", "Anchorage", "Ann Arbor", "Antioch", "Arlington", "Arvada", "Athens",
	"Atlanta", "Augusta", "Aurora", "Austin", "Bakersfield", "Baltimore", "Baton Rouge",
	"Beaumont", "Bellevue", "Berkeley", "Billings", "Birmingham", "Boise", "Boston",
	"Boulder", "Bridgeport", "Broken Arrow", "Brownsville", "Buffalo", "Burbank", "Cambridge",
	"Cape Coral", "Carlsbad", "Carrollton", "Cary", "Cedar Rapids", "Centennial", "Chandler",
	"Charleston", "Charlotte", "Chattanooga", "Chesapeake", "Chicago", "Chula Vista", "Cincinnati",
	"Clarksville", "Clearwater", "Cleveland", "Clinton", "Clovis", "College Station", "Colorado Springs",
	"Columbia", "Columbus", "Concord", "Coral Springs", "Corona", "Corpus Christi", "Costa Mesa",
	"Dallas", "Daly City", "Davenport", "Davie", "Dayton", "Denton", "Denver",
	"Des Moines", "Detroit", "Downey", "Durham", "Edison", "El Cajon", "El Monte",
	"El Paso", "Elgin", "Elizabeth", "Elk Grove", "Escondido", "Eugene", "Evansville",
	"Everett", "Fairfield", "Fargo", "Fayetteville", "Fontana", "Fort Collins", "Fort Lauderdale",
	"Fort Wayne", "Fort Worth", "Fremont", "Fresno", "Frisco", "Fullerton", "Gainesville",
	"Garden Grove", "Garland", "Gilbert", "Glendale", "Grand Prairie", "Grand Rapids", "Greeley",
	"Green Bay", "Greensboro", "Gresham", "Hampton", "Hartford", "Hayward", "Henderson",
	"Hialeah", "High Point", "Hillsboro", "Hollywood", "Honolulu", "Houston", "Huntington Beach",
	"Huntsville", "Independence", "Indianapolis", "Inglewood", "Irvine", "Irving", "Jackson",
	"Jacksonville", "Jersey City", "Joliet", "Jurupa Valley", "Kansas City", "Kenosha", "Kent",
	"Killeen", "Knoxville", "Lafayette", "Lakeland", "Lakewood", "Lancaster", "Lansing",
	"Laredo", "Las Cruces", "Las Vegas", "League City", "Lewisville", "Lexington", "Lincoln",
	"Little Rock", "Long Beach", "Los Angeles", "Louisville", "Lowell", "Lubbock", "Macon",
	"Madison", "Manchester", "McAllen", "McKinney", "Memphis", "Meridian", "Mesa",
	"Mesquite", "Miami", "Miami Gardens", "Midland", "Milwaukee", "Minneapolis", "Miramar",
	"Mobile", "Modesto", "Montgomery", "Moreno Valley", "Murfreesboro", "Murrieta", "Naperville",
	"Nashville", "New Haven", "New Orleans", "New York", "Newark", "Newport News", "Norfolk",
	"Norman", "North Charleston", "North Las Vegas", "Norwalk", "Oakland", "Oceanside", "Odessa",
	"Oklahoma City", "Olathe", "Omaha", "Ontario", "Orange", "Orlando", "Overland Park",
	"Oxnard", "Palm Bay", "Palmdale", "Pasadena", "Paterson", "Pearland", "Pembroke Pines",
	"Peoria", "Philadelphia", "Phoenix", "Pittsburgh", "Plano", "Pomona", "Pompano Beach",
	"Port St. Lucie", "Portland", "Providence", "Provo", "Pueblo", "Raleigh", "Rancho Cucamonga",
	"Reno", "Renton", "Rialto", "Richardson", "Richmond", "Riverside", "Rochester",
	"Rockford", "Roseville", "Round Rock", "Sacramento", "Saint Paul", "Salem", "Salinas",
	"Salt Lake City", "San Angelo", "San Antonio", "San Bernardino", "San Diego", "San Francisco", "San Jose",
	"San Mateo", "Sandy Springs", "Santa Ana", "Santa Clara", "Santa Clarita", "Santa Maria", "Santa Rosa",
	"Savannah", "Scottsdale", "Seattle", "Shreveport", "Simi Valley", "Sioux Falls", "South Bend",
	"Sparks", "Spokane", "Springfield", "St. Louis", "St. Petersburg", "Stamford", "Sterling Heights",
	"Stockton", "Sugar Land", "Sunnyvale", "Surprise", "Syracuse", "Tacoma", "Tallahassee",
	"Tampa", "Temecula", "Tempe", "Thornton", "Thousand Oaks", "Toledo", "Topeka",
	"Torrance", "Tucson", "Tulsa", "Tuscaloosa", "Tyler", "Vacaville", "Vallejo",
	"Vancouver", "Ventura", "Victorville", "Virginia Beach", "Visalia", "Vista", "Waco",
	"Warren", "Washington", "Waterbury", "West Covina", "West Jordan", "West Palm Beach", "West Valley City",
	"Westminster", "Wichita", "Wichita Falls", "Wilmington", "Winston–Salem", "Woodbridge", "Worcester",
	"Yonkers",
}

// StateNameGenerator generates a prefix for a female name
func StateNameGenerator() string {
	return stateData[rand.Intn(len(stateData))]
}

var stateData = []string{
	"AK", "AL", "AR", "AZ", "CA", "CO", "CT",
	"DC", "DE", "FL", "GA", "HI", "IA", "ID",
	"IL", "IN", "KS", "KY", "LA", "MA", "MD",
	"ME", "MI", "MN", "MO", "MS", "MT", "NC",
	"ND", "NE", "NH", "NJ", "NM", "NV", "NY",
	"OH", "OK", "OR", "PA", "RI", "SC", "SD",
	"TN", "TX", "UT", "VA", "VT", "WA", "WI",
	"WV", "WY",
}

// ZipcodeGenerator generates a prefix for a female name
func ZipcodeGenerator(state string) string {
	var zip int
	if rng, found := zipcodeData[state]; found {
		zip = rand.Intn(rng.max-rng.min) + rng.min
	} else {
		zip = rand.Intn(99999)
	}
	return fmt.Sprintf("%5.5d", zip)
}

type zipRange struct {
	min int
	max int
}

var zipcodeData = map[string]zipRange{
	"AK": {min: 99501, max: 99950},
	"AL": {min: 35004, max: 36925},
	"AR": {min: 71601, max: 75502},
	"AZ": {min: 85001, max: 86556},
	"CA": {min: 90001, max: 96162},
	"CO": {min: 80001, max: 81658},
	"CT": {min: 6001, max: 6928},
	"DC": {min: 20001, max: 20799},
	"DE": {min: 19701, max: 19980},
	"FL": {min: 32004, max: 34997},
	"GA": {min: 30001, max: 39901},
	"HI": {min: 96701, max: 96898},
	"IA": {min: 50001, max: 68120},
	"ID": {min: 83201, max: 83876},
	"IL": {min: 60001, max: 62999},
	"IN": {min: 46001, max: 47997},
	"KS": {min: 66002, max: 67954},
	"KY": {min: 40003, max: 42788},
	"LA": {min: 70001, max: 71497},
	"MA": {min: 01001, max: 05544},
	"MD": {min: 20331, max: 21930},
	"ME": {min: 3901, max: 4992},
	"MI": {min: 48001, max: 49971},
	"MN": {min: 55001, max: 56763},
	"MO": {min: 63001, max: 65899},
	"MS": {min: 38601, max: 39776},
	"MT": {min: 59001, max: 59937},
	"NC": {min: 27006, max: 28909},
	"ND": {min: 58001, max: 58856},
	"NE": {min: 68001, max: 69367},
	"NH": {min: 03031, max: 3897},
	"NJ": {min: 07001, max: 8989},
	"NM": {min: 87001, max: 88441},
	"NV": {min: 88901, max: 89883},
	"NY": {min: 10001, max: 14975},
	"OH": {min: 43001, max: 45999},
	"OK": {min: 73001, max: 74966},
	"OR": {min: 97001, max: 97920},
	"PA": {min: 15001, max: 19640},
	"RI": {min: 2801, max: 2940},
	"SC": {min: 29001, max: 29948},
	"SD": {min: 57001, max: 57799},
	"TN": {min: 37010, max: 38589},
	"TX": {min: 73301, max: 79999},
	"UT": {min: 84001, max: 84784},
	"VA": {min: 20040, max: 24658},
	"VT": {min: 5001, max: 5907},
	"WA": {min: 98001, max: 99403},
	"WI": {min: 53001, max: 54990},
	"WV": {min: 24701, max: 26886},
	"WY": {min: 82001, max: 83128},
}

// PhoneGenerator generates a phone number for a state
func PhoneGenerator(state string) string {
	if areas, found := areacodeData[state]; found {
		area := areas[rand.Intn(len(areas))]
		return fmt.Sprintf("(%s) %3.3d-%4.4d", area, rand.Intn(1000), rand.Intn(10000))
	}
	return fmt.Sprintf("(%3.3d) %3.3d-%4.4d", rand.Intn(1000), rand.Intn(1000), rand.Intn(10000))
}

var areacodeData = map[string][]string{
	"AL": {"205", "251", "256", "334", "938"},
	"AK": {"907"},
	"AZ": {"480", "520", "602", "623", "928"},
	"AR": {"479", "501", "870"},
	"CA": {"209", "213", "310", "323", "408", "415", "424", "442", "510", "530", "559", "562", "619", "626", "628", "650", "657", "661", "669", "707", "714", "747", "760", "805", "818", "831", "858", "909", "916", "925", "949", "951"},
	"CO": {"303", "719", "720", "970"},
	"CT": {"203", "475", "860", "959"},
	"DE": {"302"},
	"DC": {"202"},
	"FL": {"239", "305", "321", "352", "386", "407", "561", "727", "754", "772", "786", "813", "850", "863", "904", "941", "954"},
	"GA": {"229", "404", "470", "478", "678", "706", "762", "770", "912"},
	"HI": {"808"},
	"ID": {"208"},
	"IL": {"217", "224", "309", "312", "331", "618", "630", "708", "773", "779", "815", "847", "872"},
	"IN": {"219", "260", "317", "463", "574", "765", "812", "930"},
	"IA": {"319", "515", "563", "641", "712"},
	"KS": {"316", "620", "785", "913"},
	"KY": {"270", "364", "502", "606", "859"},
	"LA": {"225", "318", "337", "504", "985"},
	"ME": {"207"},
	"MD": {"240", "301", "410", "443", "667"},
	"MA": {"339", "351", "413", "508", "617", "774", "781", "857", "978"},
	"MI": {"231", "248", "269", "313", "517", "586", "616", "679", "734", "810", "906", "947", "989"},
	"MN": {"218", "320", "507", "612", "651", "763", "952"},
	"MS": {"228", "601", "662", "769"},
	"MO": {"314", "417", "573", "636", "660", "816"},
	"MT": {"406"},
	"NE": {"308", "402", "531"},
	"NV": {"702", "725", "775"},
	"NH": {"603"},
	"NJ": {"201", "551", "609", "732", "848", "856", "862", "908", "973"},
	"NM": {"505", "575"},
	"NY": {"212", "315", "332", "347", "516", "518", "585", "607", "631", "646", "680", "716", "718", "838", "845", "914", "917", "929", "934"},
	"NC": {"252", "336", "704", "743", "828", "910", "919", "980", "984"},
	"ND": {"701"},
	"OH": {"216", "220", "234", "330", "380", "419", "440", "513", "567", "614", "740", "937"},
	"OK": {"405", "539", "580", "918"},
	"OR": {"458", "503", "541", "971"},
	"PA": {"215", "223", "267", "272", "412", "484", "570", "610", "717", "724", "814", "878"},
	"RI": {"401"},
	"SC": {"803", "843", "854", "864"},
	"SD": {"605"},
	"TN": {"423", "615", "629", "731", "865", "901", "931"},
	"TX": {"210", "214", "254", "281", "325", "346", "361", "409", "430", "432", "469", "512", "682", "713", "726", "737", "806", "817", "830", "832", "903", "915", "936", "940", "956", "972", "979"},
	"UT": {"385", "435", "801"},
	"VT": {"802"},
	"VA": {"276", "434", "540", "571", "703", "757", "804"},
	"WA": {"206", "253", "360", "425", "509", "564"},
	"WV": {"304", "681"},
	"WI": {"262", "414", "534", "608", "715", "920"},
	"WY": {"307"},
}
