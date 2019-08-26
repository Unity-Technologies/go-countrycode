package countrycode

// CountryAlpha3 is a Country in ISO-3166-1 Alpha-3 format.
type CountryAlpha3 struct {
	Country
}

// String implements fmt.Stringer.
func (c CountryAlpha3) String() string {
	return FormatAlpha3.Serialize(c.Country)
}

// GoString implements fmt.GoStringer.
func (c CountryAlpha3) GoString() string {
	return c.String()
}

// MarshalText implements encoding.TextMarshaler.
func (c CountryAlpha3) MarshalText() ([]byte, error) {
	return c.marshalTextWithFormat(FormatAlpha3)
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (c *CountryAlpha3) UnmarshalText(text []byte) error {
	*c = CountryAlpha3{}
	return c.unmarshalTextWithFormat(FormatAlpha3, text)
}
