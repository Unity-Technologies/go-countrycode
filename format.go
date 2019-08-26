package countrycode

// Format represents a specific country code format with a specific
// serialization.
type Format int

const (
	// FormatAlpha3 is an ISO-3166-1 Alpha-3 country code.
	FormatAlpha3 Format = iota
	// FormatAlpha2 is an ISO-3166-1 Alpha-2 country code.
	FormatAlpha2
	// FormatAlpha2LowerCase is a lowercased ISO-3166-1 Alpha-2 country code.
	FormatAlpha2LowerCase
	formatsCount
)

// Serialize the specified Country into a country code string of the Format.
func (f Format) Serialize(country Country) string {
	return codes[country.code][f]
}

// Deserialize the specified country code string of the Format into a Country.
func (f Format) Deserialize(countryCode string) Country {
	return countries[f][countryCode]
}

var countries = func() (l [formatsCount]map[string]Country) {
	for f := Format(0); f < formatsCount; f++ {
		l[f] = make(map[string]Country, len(codes))
		for j, countryCodes := range codes {
			l[f][countryCodes[f]] = Country{code: code(j)}
		}
	}
	return
}()
