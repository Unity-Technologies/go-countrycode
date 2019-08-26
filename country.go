// Package countrycode provides utilities for representing countries in code,
// and handling their serializations and deserializations in a convenient way.
//
// All deserializations will result in `CountryUndefined` if input data is not
// a recognized country code.
//
// All the exported types are one word in memory and as such provide fast
// equality checks, hashing for usage as keys in maps. Conversions between the
// types are zero overhead (don't even escape the stack), serialization is
// worst case O(1), and deserialization is worst case O(n).
package countrycode // import "github.com/Applifier/go-countrycode"

import "unsafe"

// Country represents the country codes of a country.
type Country struct {
	code code
}

// CountryUndefined represents an undefined country.
var CountryUndefined = Country{}

// Alpha3 returns the Country in a CountryAlpha3 serializer.
func (c Country) Alpha3() CountryAlpha3 {
	return CountryAlpha3{Country: c}
}

// Alpha2 returns the Country in a CountryAlpha2 serializer.
func (c Country) Alpha2() CountryAlpha2 {
	return CountryAlpha2{Country: c}
}

// Alpha2LowerCase returns the Country in a CountryAlpha2LowerCase serializer.
func (c Country) Alpha2LowerCase() CountryAlpha2LowerCase {
	return CountryAlpha2LowerCase{Country: c}
}

// GoString implements fmt.GoStringer for debugging purposes.
func (c Country) GoString() string {
	return "countrycode.Country{" + FormatAlpha3.Serialize(c) + "}"
}

func (c Country) marshalTextWithFormat(format Format) ([]byte, error) {
	str := format.Serialize(c)
	return *(*[]byte)(unsafe.Pointer(&str)), nil
}

func (c *Country) unmarshalTextWithFormat(format Format, text []byte) error {
	*c = format.Deserialize(*(*string)(unsafe.Pointer(&text)))
	return nil
}
