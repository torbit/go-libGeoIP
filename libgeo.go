/**
 * libgeo.go
 *
 * Copyright (c) 2010, Nikola Ranchev, Torbit (www.torbit.com)
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 * 	- Redistributions of source code must retain the above copyright
 * 	  notice, this list of conditions and the following disclaimer.
 * 	- Redistributions in binary form must reproduce the above copyright
 * 	  notice, this list of conditions and the following disclaimer in the
 * 	  documentation and/or other materials provided with the distribution.
 * 	- Neither the name of the <organization> nor the
 * 	  names of its contributors may be used to endorse or promote products
 * 	  derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package libgeo

// Dependencies
import (
	"errors"
	"os"
)

// Globals (const arrays that will be initialized inside init())
var (
	countryCode2 = []string{
		"--", "AP", "EU", "AD", "AE", "AF", "AG", "AI", "AL", "AM", "CW", "AO", "AQ", "AR",
		"AS", "AT", "AU", "AW", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ",
		"BM", "BN", "BO", "BR", "BS", "BT", "BV", "BW", "BY", "BZ", "CA", "CC", "CD", "CF",
		"CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU", "CV", "CX", "CY", "CZ",
		"DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE", "EG", "EH", "ER", "ES", "ET", "FI",
		"FJ", "FK", "FM", "FO", "FR", "SX", "GA", "GB", "GD", "GE", "GF", "GH", "GI", "GL",
		"GM", "GN", "GP", "GQ", "GR", "GS", "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR",
		"HT", "HU", "ID", "IE", "IL", "IN", "IO", "IQ", "IR", "IS", "IT", "JM", "JO", "JP",
		"KE", "KG", "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC",
		"LI", "LK", "LR", "LS", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "MG", "MH", "MK",
		"ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW", "MX", "MY",
		"MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP", "NR", "NU", "NZ", "OM",
		"PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN", "PR", "PS", "PT", "PW", "PY",
		"QA", "RE", "RO", "RU", "RW", "SA", "SB", "SC", "SD", "SE", "SG", "SH", "SI", "SJ",
		"SK", "SL", "SM", "SN", "SO", "SR", "ST", "SV", "SY", "SZ", "TC", "TD", "TF", "TG",
		"TH", "TJ", "TK", "TM", "TN", "TO", "TL", "TR", "TT", "TV", "TW", "TZ", "UA", "UG",
		"UM", "US", "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU", "WF", "WS", "YE",
		"YT", "RS", "ZA", "ZM", "ME", "ZW", "A1", "A2", "O1", "AX", "GG", "IM", "JE", "BL",
		"MF", "BQ", "SS"}

	countryCode3 = []string{
		"--", "AP", "EU", "AND", "ARE", "AFG", "ATG", "AIA", "ALB", "ARM", "CUW",
		"AGO", "ATA", "ARG", "ASM", "AUT", "AUS", "ABW", "AZE", "BIH", "BRB",
		"BGD", "BEL", "BFA", "BGR", "BHR", "BDI", "BEN", "BMU", "BRN", "BOL",
		"BRA", "BHS", "BTN", "BVT", "BWA", "BLR", "BLZ", "CAN", "CCK", "COD",
		"CAF", "COG", "CHE", "CIV", "COK", "CHL", "CMR", "CHN", "COL", "CRI",
		"CUB", "CPV", "CXR", "CYP", "CZE", "DEU", "DJI", "DNK", "DMA", "DOM",
		"DZA", "ECU", "EST", "EGY", "ESH", "ERI", "ESP", "ETH", "FIN", "FJI",
		"FLK", "FSM", "FRO", "FRA", "SXM", "GAB", "GBR", "GRD", "GEO", "GUF",
		"GHA", "GIB", "GRL", "GMB", "GIN", "GLP", "GNQ", "GRC", "SGS", "GTM",
		"GUM", "GNB", "GUY", "HKG", "HMD", "HND", "HRV", "HTI", "HUN", "IDN",
		"IRL", "ISR", "IND", "IOT", "IRQ", "IRN", "ISL", "ITA", "JAM", "JOR",
		"JPN", "KEN", "KGZ", "KHM", "KIR", "COM", "KNA", "PRK", "KOR", "KWT",
		"CYM", "KAZ", "LAO", "LBN", "LCA", "LIE", "LKA", "LBR", "LSO", "LTU",
		"LUX", "LVA", "LBY", "MAR", "MCO", "MDA", "MDG", "MHL", "MKD", "MLI",
		"MMR", "MNG", "MAC", "MNP", "MTQ", "MRT", "MSR", "MLT", "MUS", "MDV",
		"MWI", "MEX", "MYS", "MOZ", "NAM", "NCL", "NER", "NFK", "NGA", "NIC",
		"NLD", "NOR", "NPL", "NRU", "NIU", "NZL", "OMN", "PAN", "PER", "PYF",
		"PNG", "PHL", "PAK", "POL", "SPM", "PCN", "PRI", "PSE", "PRT", "PLW",
		"PRY", "QAT", "REU", "ROU", "RUS", "RWA", "SAU", "SLB", "SYC", "SDN",
		"SWE", "SGP", "SHN", "SVN", "SJM", "SVK", "SLE", "SMR", "SEN", "SOM",
		"SUR", "STP", "SLV", "SYR", "SWZ", "TCA", "TCD", "ATF", "TGO", "THA",
		"TJK", "TKL", "TKM", "TUN", "TON", "TLS", "TUR", "TTO", "TUV", "TWN",
		"TZA", "UKR", "UGA", "UMI", "USA", "URY", "UZB", "VAT", "VCT", "VEN",
		"VGB", "VIR", "VNM", "VUT", "WLF", "WSM", "YEM", "MYT", "SRB", "ZAF",
		"ZMB", "MNE", "ZWE", "A1", "A2", "O1", "ALA", "GGY", "IMN", "JEY",
		"BLM", "MAF", "BES", "SSD",
	}

	countryName = []string{
		"N/A", "Asia/Pacific Region", "Europe", "Andorra", "United Arab Emirates",
		"Afghanistan", "Antigua and Barbuda", "Anguilla", "Albania", "Armenia",
		"Curaçao", "Angola", "Antarctica", "Argentina", "American Samoa",
		"Austria", "Australia", "Aruba", "Azerbaijan", "Bosnia and Herzegovina",
		"Barbados", "Bangladesh", "Belgium", "Burkina Faso", "Bulgaria", "Bahrain",
		"Burundi", "Benin", "Bermuda", "Brunei Darussalam", "Bolivia", "Brazil", "Bahamas",
		"Bhutan", "Bouvet Island", "Botswana", "Belarus", "Belize", "Canada",
		"Cocos (Keeling) Islands", "Congo, The Democratic Republic of the",
		"Central African Republic", "Congo", "Switzerland", "Cote D'Ivoire",
		"Cook Islands", "Chile", "Cameroon", "China", "Colombia", "Costa Rica", "Cuba",
		"Cape Verde", "Christmas Island", "Cyprus", "Czech Republic", "Germany",
		"Djibouti", "Denmark", "Dominica", "Dominican Republic", "Algeria", "Ecuador",
		"Estonia", "Egypt", "Western Sahara", "Eritrea", "Spain", "Ethiopia", "Finland",
		"Fiji", "Falkland Islands (Malvinas)", "Micronesia, Federated States of",
		"Faroe Islands", "France", "Sint Maarten (Dutch part)", "Gabon", "United Kingdom",
		"Grenada", "Georgia", "French Guiana", "Ghana", "Gibraltar", "Greenland", "Gambia",
		"Guinea", "Guadeloupe", "Equatorial Guinea", "Greece",
		"South Georgia and the South Sandwich Islands", "Guatemala", "Guam",
		"Guinea-Bissau", "Guyana", "Hong Kong", "Heard Island and McDonald Islands",
		"Honduras", "Croatia", "Haiti", "Hungary", "Indonesia", "Ireland", "Israel", "India",
		"British Indian Ocean Territory", "Iraq", "Iran, Islamic Republic of",
		"Iceland", "Italy", "Jamaica", "Jordan", "Japan", "Kenya", "Kyrgyzstan", "Cambodia",
		"Kiribati", "Comoros", "Saint Kitts and Nevis",
		"Korea, Democratic People's Republic of", "Korea, Republic of", "Kuwait",
		"Cayman Islands", "Kazakhstan", "Lao People's Democratic Republic", "Lebanon",
		"Saint Lucia", "Liechtenstein", "Sri Lanka", "Liberia", "Lesotho", "Lithuania",
		"Luxembourg", "Latvia", "Libya", "Morocco", "Monaco",
		"Moldova, Republic of", "Madagascar", "Marshall Islands",
		"Macedonia", "Mali", "Myanmar", "Mongolia",
		"Macau", "Northern Mariana Islands", "Martinique", "Mauritania", "Montserrat",
		"Malta", "Mauritius", "Maldives", "Malawi", "Mexico", "Malaysia", "Mozambique",
		"Namibia", "New Caledonia", "Niger", "Norfolk Island", "Nigeria", "Nicaragua",
		"Netherlands", "Norway", "Nepal", "Nauru", "Niue", "New Zealand", "Oman", "Panama",
		"Peru", "French Polynesia", "Papua New Guinea", "Philippines", "Pakistan",
		"Poland", "Saint Pierre and Miquelon", "Pitcairn Islands", "Puerto Rico",
		"Palestinian Territory", "Portugal", "Palau", "Paraguay", "Qatar",
		"Reunion", "Romania", "Russian Federation", "Rwanda", "Saudi Arabia",
		"Solomon Islands", "Seychelles", "Sudan", "Sweden", "Singapore", "Saint Helena",
		"Slovenia", "Svalbard and Jan Mayen", "Slovakia", "Sierra Leone", "San Marino",
		"Senegal", "Somalia", "Suriname", "Sao Tome and Principe", "El Salvador",
		"Syrian Arab Republic", "Swaziland", "Turks and Caicos Islands", "Chad",
		"French Southern Territories", "Togo", "Thailand", "Tajikistan", "Tokelau",
		"Turkmenistan", "Tunisia", "Tonga", "Timor-Leste", "Turkey", "Trinidad and Tobago",
		"Tuvalu", "Taiwan", "Tanzania, United Republic of", "Ukraine", "Uganda",
		"United States Minor Outlying Islands", "United States", "Uruguay", "Uzbekistan",
		"Holy See (Vatican City State)", "Saint Vincent and the Grenadines",
		"Venezuela", "Virgin Islands, British", "Virgin Islands, U.S.", "Vietnam",
		"Vanuatu", "Wallis and Futuna", "Samoa", "Yemen", "Mayotte", "Serbia",
		"South Africa", "Zambia", "Montenegro", "Zimbabwe", "Anonymous Proxy",
		"Satellite Provider", "Other", "Aland Islands", "Guernsey", "Isle of Man", "Jersey",
		"Saint Barthelemy", "Saint Martin", "Bonaire, Saint Eustatius and Saba", "South Sudan",
	}

	countryContinentCode = []string{
		"--", "AS", "EU", "EU", "AS", "AS", "NA", "NA", "EU", "AS", "NA",
		"AF", "AN", "SA", "OC", "EU", "OC", "NA", "AS", "EU", "NA",
		"AS", "EU", "AF", "EU", "AS", "AF", "AF", "NA", "AS", "SA",
		"SA", "NA", "AS", "AN", "AF", "EU", "NA", "NA", "AS", "AF",
		"AF", "AF", "EU", "AF", "OC", "SA", "AF", "AS", "SA", "NA",
		"NA", "AF", "AS", "AS", "EU", "EU", "AF", "EU", "NA", "NA",
		"AF", "SA", "EU", "AF", "AF", "AF", "EU", "AF", "EU", "OC",
		"SA", "OC", "EU", "EU", "NA", "AF", "EU", "NA", "AS", "SA",
		"AF", "EU", "NA", "AF", "AF", "NA", "AF", "EU", "AN", "NA",
		"OC", "AF", "SA", "AS", "AN", "NA", "EU", "NA", "EU", "AS",
		"EU", "AS", "AS", "AS", "AS", "AS", "EU", "EU", "NA", "AS",
		"AS", "AF", "AS", "AS", "OC", "AF", "NA", "AS", "AS", "AS",
		"NA", "AS", "AS", "AS", "NA", "EU", "AS", "AF", "AF", "EU",
		"EU", "EU", "AF", "AF", "EU", "EU", "AF", "OC", "EU", "AF",
		"AS", "AS", "AS", "OC", "NA", "AF", "NA", "EU", "AF", "AS",
		"AF", "NA", "AS", "AF", "AF", "OC", "AF", "OC", "AF", "NA",
		"EU", "EU", "AS", "OC", "OC", "OC", "AS", "NA", "SA", "OC",
		"OC", "AS", "AS", "EU", "NA", "OC", "NA", "AS", "EU", "OC",
		"SA", "AS", "AF", "EU", "EU", "AF", "AS", "OC", "AF", "AF",
		"EU", "AS", "AF", "EU", "EU", "EU", "AF", "EU", "AF", "AF",
		"SA", "AF", "NA", "AS", "AF", "NA", "AF", "AN", "AF", "AS",
		"AS", "OC", "AS", "AF", "OC", "AS", "EU", "NA", "OC", "AS",
		"AF", "EU", "AF", "OC", "NA", "SA", "AS", "EU", "NA", "SA",
		"NA", "NA", "AS", "OC", "OC", "OC", "AS", "AF", "EU", "AF",
		"AF", "EU", "AF", "--", "--", "--", "EU", "EU", "EU", "EU",
		"NA", "NA", "NA", "AF",
	}
)

// Constants
const (
	maxRecordLength      = 4
	standardRecordLength = 3
	countryBegin         = 16776960
	structureInfoMaxSize = 20
	fullRecordLength     = 60
	segmentRecordLength  = 3

	// DB Types
	dbCountryEdition  = byte(1)
	dbCityEditionRev0 = byte(6)
	dbCityEditionRev1 = byte(2)
)

// These are some structs
type GeoIP struct {
	databaseSegment int    // No need to make an array of size 1
	recordLength    int    // Set to one of the constants above
	dbType          byte   // Store the database type
	data            []byte // All of the data from the DB file
}
type Location struct {
	CountryCode   string // If country ed. only country info is filled
	CountryCode3  string // If country ed. only country info is filled
	CountryName   string // If country ed. only country info is filled
	ContinentCode string // If country ed. only country info is filled
	Region        string
	City          string
	PostalCode    string
	Latitude      float32
	Longitude     float32
}

// Load the database file in memory, detect the db format and setup the GeoIP struct
func Load(filename string) (gi *GeoIP, err error) {
	// Try to open the requested file
	dbInfo, err := os.Lstat(filename)
	if err != nil {
		return
	}
	dbFile, err := os.Open(filename)
	defer dbFile.Close()
	if err != nil {
		return
	}

	// Copy the db into memory
	gi = new(GeoIP)
	gi.data = make([]byte, dbInfo.Size())
	dbFile.Read(gi.data)

	// Check the database type
	gi.dbType = dbCountryEdition           // Default the database to country edition
	gi.databaseSegment = countryBegin      // Default to country DB
	gi.recordLength = standardRecordLength // Default to country DB

	// Search for the DB type headers
	delim := make([]byte, 3)
	for i := 0; i < structureInfoMaxSize; i++ {
		delim = gi.data[len(gi.data)-i-3-1 : len(gi.data)-i-1]
		if int8(delim[0]) == -1 && int8(delim[1]) == -1 && int8(delim[2]) == -1 {
			gi.dbType = gi.data[len(gi.data)-i-1]
			// If we detect city edition set the correct segment offset
			if gi.dbType == dbCityEditionRev0 || gi.dbType == dbCityEditionRev1 {
				buf := make([]byte, segmentRecordLength)
				buf = gi.data[len(gi.data)-i-1+1 : len(gi.data)-i-1+4]
				gi.databaseSegment = 0
				for j := 0; j < segmentRecordLength; j++ {
					gi.databaseSegment += (int(buf[j]) << uint8(j*8))
				}
			}
			break
		}
	}

	// Support older formats
	if gi.dbType >= 106 {
		gi.dbType -= 105
	}

	// Reject unsupported formats
	if gi.dbType != dbCountryEdition && gi.dbType != dbCityEditionRev0 && gi.dbType != dbCityEditionRev1 {
		err = errors.New("Unsupported database format")
		return
	}

	return
}

// Lookup by IP address and return location
func (gi *GeoIP) GetLocationByIP(ip string) *Location {
	return gi.GetLocationByIPNum(AddrToNum(ip))
}

// Lookup by IP number and return location
func (gi *GeoIP) GetLocationByIPNum(ipNum uint32) *Location {
	// Perform the lookup on the database to see if the record is found
	offset := gi.lookupByIPNum(ipNum)

	// Check if the country was found
	if gi.dbType == dbCountryEdition && offset-countryBegin == 0 ||
		gi.dbType != dbCountryEdition && gi.databaseSegment == offset {
		return nil
	}

	// Create a generic location structure
	location := new(Location)

	// If the database is country
	if gi.dbType == dbCountryEdition {
		position := offset - countryBegin
		if int(position) < len(countryCode2) {
			location.CountryCode = countryCode2[position]
			location.CountryCode3 = countryCode3[position]
			location.CountryName = countryName[position]
			location.ContinentCode = countryContinentCode[position]
		}
		return location
	}

	// Find the max record length
	recPointer := offset + (2*gi.recordLength-1)*gi.databaseSegment
	recordEnd := recPointer + fullRecordLength
	if len(gi.data)-recPointer < fullRecordLength {
		recordEnd = len(gi.data)
	}

	// Read the country code/name first
	position := gi.data[recPointer]
	if int(position) < len(countryCode2) {
		location.CountryCode = countryCode2[position]
		location.CountryCode3 = countryCode3[position]
		location.CountryName = countryName[position]
		location.ContinentCode = countryContinentCode[position]
	}
	readLen := 1
	recPointer += 1

	// Get the region
	for readLen = 0; gi.data[recPointer+readLen] != '\000' &&
		recPointer+readLen < recordEnd; readLen++ {
	}
	if readLen != 0 {
		location.Region = string(gi.data[recPointer : recPointer+readLen])
	}
	recPointer += readLen + 1

	// Get the city
	for readLen = 0; gi.data[recPointer+readLen] != '\000' &&
		recPointer+readLen < recordEnd; readLen++ {
	}
	if readLen != 0 {
		location.City = string(gi.data[recPointer : recPointer+readLen])
	}
	recPointer += readLen + 1

	// Get the postal code
	for readLen = 0; gi.data[recPointer+readLen] != '\000' &&
		recPointer+readLen < recordEnd; readLen++ {
	}
	if readLen != 0 {
		location.PostalCode = string(gi.data[recPointer : recPointer+readLen])
	}
	recPointer += readLen + 1

	// Get the latitude
	coordinate := float32(0)
	for j := 0; j < 3; j++ {
		coordinate += float32(int32(gi.data[recPointer+j]) << uint8(j*8))
	}
	location.Latitude = float32(coordinate/10000 - 180)
	recPointer += 3

	// Get the longitude
	coordinate = 0
	for j := 0; j < 3; j++ {
		coordinate += float32(int32(gi.data[recPointer+j]) << uint8(j*8))
	}
	location.Longitude = float32(coordinate/10000 - 180)

	return location
}

// Read the database and return record position
func (gi *GeoIP) lookupByIPNum(ip uint32) int {
	buf := make([]byte, 2*maxRecordLength)
	x := make([]int, 2)
	offset := 0
	for depth := 31; depth >= 0; depth-- {
		for i := 0; i < 2*maxRecordLength; i++ {
			buf[i] = gi.data[(2*gi.recordLength*offset)+i]
		}
		for i := 0; i < 2; i++ {
			x[i] = 0
			for j := 0; j < gi.recordLength; j++ {
				var y int = int(buf[i*gi.recordLength+j])
				if y < 0 {
					y += 256
				}
				x[i] += (y << uint(j*8))
			}
		}
		if (ip & (1 << uint(depth))) > 0 {
			if x[1] >= gi.databaseSegment {
				return x[1]
			}
			offset = x[1]
		} else {
			if x[0] >= gi.databaseSegment {
				return x[0]
			}
			offset = x[0]
		}
	}
	return 0
}

//Development utility: Test lookup arrays for obvious size inconsistancy
func internalDataChecks() (bool, error) {
	if len(countryCode2) != len(countryCode3) ||
		len(countryCode3) != len(countryName) ||
		len(countryName) != len(countryContinentCode) {
		return false, errors.New("Internal consistancy error, lookup table size mismatch")
	}
	return true, nil
}

// Convert ip address to an int representation
func AddrToNum(ip string) uint32 {
	octet := uint32(0)
	ipnum := uint32(0)
	i := 3
	for j := 0; j < len(ip); j++ {
		c := byte(ip[j])
		if c == '.' {
			if octet > 255 {
				return 0
			}
			ipnum <<= 8
			ipnum += octet
			i--
			octet = 0
		} else {
			t := octet
			octet <<= 3
			octet += t
			octet += t
			c -= '0'
			if c > 9 {
				return 0
			}
			octet += uint32(c)
		}
	}
	if (octet > 255) || (i != 0) {
		return 0
	}
	ipnum <<= 8
	return uint32(ipnum + octet)
}
