package money

// Currency provides information about a particular currency.
type Currency interface {
	// ISO 4217 currency code.
	Code() string

	// ISO 4217 three-digit numeric code.
	NumericCode() int

	// Number of digits in the currency minor unit.
	// This is the number of digits after the decimal separator.
	MinorUnit() int

	// Name of the currency.
	Name() string

	// Array of ISO 3166 country codes representing the countries
	// that use this currency.
	CountryCodes() []string
}

// currency implements the Currency interface.
type currency struct {
	code         string
	numericCode  int
	countryCodes []string
	minorUnit    int
	name         string
}

func (c currency) Code() string {
	return c.code
}

func (c currency) Name() string {
	return c.name
}

// String implements the fmt.Stringer interface.
func (c currency) String() string {
	return c.code
}

func (c currency) NumericCode() int {
	return c.numericCode
}

func (c currency) CountryCodes() []string {
	// make a copy
	codes := make([]string, len(c.countryCodes))
	copy(codes, c.countryCodes)
	return codes
}

func (c currency) MinorUnit() int {
	return c.minorUnit
}

// DefaultCurrency is the default currency used for serialization
// when a currency amount is not specified.
var DefaultCurrency Currency

var (
	// Australian Dollar
	AUD Currency = currency{
		code:        "AUD",
		numericCode: 36,
		minorUnit:   2,
		name:        "Australian Dollar",
		countryCodes: []string{
			"AU", // Australia
			"CX", // Christmas Island
			"CC", // Cocos (Keeling) Islands
			"HM", // Heard Island and McDonald Islands
			"KI", // Kiribati
			"NR", // Nauru
			"NF", // Norfolk Island
			"TV", // Tuvalu
		},
	}

	CAD Currency = currency{
		code:        "CAD",
		numericCode: 124,
		minorUnit:   2,
		name:        "Canadian Dollar",
		countryCodes: []string{
			"CA", // Canada
		},
	}

	CHF Currency = currency{
		code:        "CHF",
		numericCode: 756,
		minorUnit:   2,
		name:        "Swiss Franc",
		countryCodes: []string{
			"CH", // Switzerland
			"LI", // Liechtenstein
		},
	}

	EUR Currency = currency{
		code:        "EUR",
		numericCode: 978,
		minorUnit:   2,
		name:        "Euro",
		countryCodes: []string{
			"AX", // Åland Islands
			"AD", // Andorra
			"AT", // Austria
			"BE", // Belgium
			"CY", // Cyprus
			"EE", // Estonia
			"FI", // Finland
			"FR", // France
			"GF", // French Guiana
			"TF", // French Southern Territories
			"DE", // Germany
			"GR", // Greece
			"GP", // Guadeloupe
			"VA", // Holy See
			"IE", // Ireland
			"IT", // Italy
			"LV", // Latvia
			"LT", // Lithuania
			"LU", // Luxembourg
			"MT", // Malta
			"MQ", // Martinique
			"YT", // Mayotte
			"MC", // Monaco
			"ME", // Montenegro
			"NL", // The Netherlands
			"PT", // Portugal
			"RE", // Réunion
			"BL", // Saint Barthélemy
			"MF", // Saint Martin (French part)
			"PM", // Saint Pierre and Miquelon
			"SM", // San Marino
			"SK", // Slovakia
			"SI", // Slovenia
			"ES", // Spain
		},
	}

	GBP Currency = currency{
		code: "GBP",
		numericCode: 826,
		minorUnit:   2,
		name:        "Great Britain Pound",
		countryCodes: []string{
			"UK", // United Kingdom
		},
	}

	JPY Currency = currency{
		code:        "JPY",
		numericCode: 392,
		minorUnit:   0,
		name:        "Japanese Yen",
		countryCodes: []string{
			"JP", // Japan
		},
	}

	NZD Currency = currency{
		code:        "NZD",
		numericCode: 554,
		minorUnit:   2,
		name:        "New Zealand Dollar",
		countryCodes: []string{
			"NZ", // New Zealand
			"CK", // Cook Islands
			"NU", // Niue
			"PN", // Pitcairn
			"TK", // Tokelau
		},
	}

	USD Currency = currency{
		code:        "USD",
		numericCode: 840,
		minorUnit:   2,
		name:        "US Dollar",
		countryCodes: []string{
			"US", // United States of America
			"AS", // Americal Samoa
			"BQ", // Bonaire, Sint Eustatius and Saba
			"IO", // British Indian Ocean Territory
			"EC", // Ecuador
			"SV", // El Salvador
			"GU", // Guam
			"HT", // Haiti
			"MH", // Marshall Islands
			"FM", // Micronesia
			"MP", // Northern Mariana Islands
			"PW", // Palau
			"PA", // Panama
			"PR", // Puerto Rico
			"TL", // Timor-Leste
			"TC", // Turks and Caicos Islands
			"UM", // United States Minor Outlying
			"VG", // Virgin Islands (British)
			"VI", // Virgin Islands (U.S.)
		},
	}
)

func init() {
	// TODO: could determine using the locale.
	DefaultCurrency = AUD
}
