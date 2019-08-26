package countrycode

// CountryAlpha2LowerCase is a Country in lowercase ISO-3166-1 Alpha-2 format.
type CountryAlpha2LowerCase struct {
	Country
}

// String implements fmt.Stringer.
func (c CountryAlpha2LowerCase) String() string {
	return FormatAlpha2LowerCase.Serialize(c.Country)
}

// GoString implements fmt.GoStringer.
func (c CountryAlpha2LowerCase) GoString() string {
	return c.String()
}

// MarshalText implements encoding.TextMarshaler.
func (c CountryAlpha2LowerCase) MarshalText() ([]byte, error) {
	return c.marshalTextWithFormat(FormatAlpha2LowerCase)
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (c *CountryAlpha2LowerCase) UnmarshalText(text []byte) error {
	*c = CountryAlpha2LowerCase{}
	return c.unmarshalTextWithFormat(FormatAlpha2LowerCase, text)
}
