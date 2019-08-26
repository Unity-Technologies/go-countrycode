package countrycode

// CountryAlpha2 is a Country in ISO-3166-1 Alpha-2 format.
type CountryAlpha2 struct {
	Country
}

// String implements fmt.Stringer.
func (c CountryAlpha2) String() string {
	return FormatAlpha2.Serialize(c.Country)
}

// GoString implements fmt.GoStringer.
func (c CountryAlpha2) GoString() string {
	return c.String()
}

// MarshalText implements encoding.TextMarshaler.
func (c CountryAlpha2) MarshalText() ([]byte, error) {
	return c.marshalTextWithFormat(FormatAlpha2)
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (c *CountryAlpha2) UnmarshalText(text []byte) error {
	*c = CountryAlpha2{}
	return c.unmarshalTextWithFormat(FormatAlpha2, text)
}
