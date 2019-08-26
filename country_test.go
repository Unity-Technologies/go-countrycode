package countrycode_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/Applifier/go-countrycode"
)

func TestFormat_Deserialize(t *testing.T) {
	tests := []struct {
		test        string
		countryCode string
		format      countrycode.Format
		country     countrycode.Country
	}{
		{
			"FormatAlpha2LowerCase",
			"fi",
			countrycode.FormatAlpha2LowerCase,
			countrycode.FIN,
		},
		{
			"FormatAlpha2",
			"US",
			countrycode.FormatAlpha2,
			countrycode.USA,
		},
		{
			"FormatAlpha3",
			"USA",
			countrycode.FormatAlpha3,
			countrycode.USA,
		},
		{
			"invalid format",
			"USA",
			countrycode.FormatAlpha2,
			countrycode.CountryUndefined,
		},
		{
			"invalid country",
			"AAA",
			countrycode.FormatAlpha3,
			countrycode.CountryUndefined,
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			requireEqual(t, tt.country, tt.format.Deserialize(tt.countryCode))
		})
	}
}

func TestFormat_Serialize(t *testing.T) {
	tests := []struct {
		test        string
		country     countrycode.Country
		format      countrycode.Format
		countryCode string
	}{
		{
			"FormatAlpha2LowerCase",
			countrycode.FIN,
			countrycode.FormatAlpha2LowerCase,
			"fi",
		},
		{
			"FormatAlpha2",
			countrycode.USA,
			countrycode.FormatAlpha2,
			"US",
		},
		{
			"FormatAlpha3",
			countrycode.USA,
			countrycode.FormatAlpha3,
			"USA",
		},
		{
			"invalid country",
			countrycode.CountryUndefined,
			countrycode.FormatAlpha3,
			"---",
		},
	}

	for _, tt := range tests {
		t.Run(tt.test, func(t *testing.T) {
			requireEqual(t, tt.countryCode, tt.format.Serialize(tt.country))
		})
	}
}

func TestUnmarshaling(t *testing.T) {
	t.Run("Alpha3 map", func(t *testing.T) {
		// Arrange
		type x struct {
			Y map[countrycode.CountryAlpha3]int `json:"y"`
		}
		expected := x{
			Y: map[countrycode.CountryAlpha3]int{
				countrycode.USA.Alpha3(): 123,
				countrycode.FIN.Alpha3(): 234,
			},
		}
		input := `{"y": {"USA": 123, "FIN": 234}}`
		var received x

		// Act
		requireNoError(t, json.Unmarshal([]byte(input), &received))

		// Assert
		requireEqual(t, expected, received)
	})
	t.Run("Alpha2 map", func(t *testing.T) {
		// Arrange
		type x struct {
			Y map[countrycode.CountryAlpha2]int `json:"y"`
		}
		expected := x{
			Y: map[countrycode.CountryAlpha2]int{
				countrycode.USA.Alpha2(): 123,
				countrycode.FIN.Alpha2(): 234,
			},
		}
		input := `{"y": {"US": 123, "FI": 234}}`
		var received x

		// Act
		requireNoError(t, json.Unmarshal([]byte(input), &received))

		// Assert
		requireEqual(t, expected, received)
	})
	t.Run("Alpha2LowerCase map", func(t *testing.T) {
		// Arrange
		type x struct {
			Y map[countrycode.CountryAlpha2LowerCase]int `json:"y"`
		}
		expected := x{
			Y: map[countrycode.CountryAlpha2LowerCase]int{
				countrycode.USA.Alpha2LowerCase(): 123,
				countrycode.FIN.Alpha2LowerCase(): 234,
			},
		}
		input := `{"y": {"us": 123, "fi": 234}}`
		var received x

		// Act
		requireNoError(t, json.Unmarshal([]byte(input), &received))

		// Assert
		requireEqual(t, expected, received)
	})
}

func TestMarshaling(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     "Alpha3 list",
			input:    []countrycode.CountryAlpha3{countrycode.USA.Alpha3(), countrycode.FIN.Alpha3()},
			expected: `["USA","FIN"]`,
		},
		{
			name:     "Alpha2 list",
			input:    []countrycode.CountryAlpha2{countrycode.USA.Alpha2(), countrycode.FIN.Alpha2()},
			expected: `["US","FI"]`,
		},
		{
			name:     "Alpha2LowerCase list",
			input:    []countrycode.CountryAlpha2LowerCase{countrycode.USA.Alpha2LowerCase(), countrycode.FIN.Alpha2LowerCase()},
			expected: `["us","fi"]`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.input)
			requireNoError(t, err)
			requireEqual(t, tc.expected, string(data))
		})
	}
}

func TestConversions(t *testing.T) {
	requireTrue(t, countrycode.USA == countrycode.USA) // nolint: staticcheck
	requireFalse(t, countrycode.USA == countrycode.FIN)
	requireTrue(t, countrycode.USA.Alpha3() == countrycode.USA.Alpha3()) // nolint: staticcheck
	requireFalse(t, countrycode.USA.Alpha3() == countrycode.FIN.Alpha3())
	requireTrue(t, countrycode.USA.Alpha2() == countrycode.USA.Alpha2()) // nolint: staticcheck
	requireFalse(t, countrycode.USA.Alpha2() == countrycode.FIN.Alpha2())
	requireTrue(t, countrycode.USA.Alpha2LowerCase() == countrycode.USA.Alpha2LowerCase()) // nolint: staticcheck
	requireFalse(t, countrycode.USA.Alpha2LowerCase() == countrycode.FIN.Alpha2LowerCase())
	requireTrue(t, countrycode.USA.Alpha3().Country == countrycode.USA.Alpha3().Country) // nolint: staticcheck
	requireFalse(t, countrycode.USA.Alpha3().Country == countrycode.FIN.Alpha3().Country)
	requireTrue(t, countrycode.USA.Alpha2().Country == countrycode.USA.Alpha2().Country) // nolint: staticcheck
	requireFalse(t, countrycode.USA.Alpha2().Country == countrycode.FIN.Alpha2().Country)
	requireTrue(t, countrycode.USA.Alpha2LowerCase().Country == countrycode.USA.Alpha2LowerCase().Country) // nolint: staticcheck
	requireFalse(t, countrycode.USA.Alpha2LowerCase().Country == countrycode.FIN.Alpha2LowerCase().Country)
	requireTrue(t, countrycode.USA.Alpha3().Alpha2().Alpha2LowerCase().Country == countrycode.USA)
	requireTrue(t, countrycode.USA.Alpha2().Alpha3().Alpha2LowerCase().Country == countrycode.USA)
	requireTrue(t, countrycode.USA.Alpha2LowerCase().Alpha3().Country == countrycode.USA)
	requireTrue(t, countrycode.USA.Alpha2LowerCase().Alpha2().Country == countrycode.USA)
}

func TestDebug(t *testing.T) {
	requireEqual(t, fmt.Sprintf("%#v", countrycode.USA), "countrycode.Country{USA}")
	requireEqual(t, fmt.Sprintf("%#v", countrycode.USA.Alpha3()), "USA")
	requireEqual(t, fmt.Sprintf("%#v", countrycode.USA.Alpha2()), "US")
	requireEqual(t, fmt.Sprintf("%#v", countrycode.USA.Alpha2LowerCase()), "us")
}

func ExampleFormat_twoWayConversion() {
	const countryCodeFormat = countrycode.FormatAlpha3

	type externalData struct {
		Country string
	}

	type internalData struct {
		Country countrycode.Country
	}

	toInternalData := func(d externalData) internalData {
		return internalData{Country: countryCodeFormat.Deserialize(d.Country)}
	}

	toExternalData := func(d internalData) externalData {
		return externalData{Country: countryCodeFormat.Serialize(d.Country)}
	}

	edata := externalData{Country: "FIN"}
	idata := toInternalData(edata)
	edata = toExternalData(idata)
	// Output: {FIN}
	fmt.Println(edata)
}

func ExampleFormat_Deserialize() {
	country := countrycode.FormatAlpha2.Deserialize("US")
	// Output: true
	fmt.Println(country == countrycode.USA)
}

func ExampleFormat_Serialize() {
	countryCodeAlpha2 := countrycode.FormatAlpha2.Serialize(countrycode.USA)
	// Output: US
	fmt.Println(countryCodeAlpha2)
}

func Example_unmarshalJSON() {
	data := []byte(`
		{
			"USA": 1234,
			"FIN": 2345
		}
	`)
	var mapping map[countrycode.CountryAlpha3]int
	_ = json.Unmarshal(data, &mapping)
	// Output: map[countrycode.CountryAlpha3]int{FIN:2345, USA:1234}
	fmt.Printf("%#v\n", mapping)
}

func Example_marshalJSON() {
	mapping := map[countrycode.CountryAlpha3]int{
		countrycode.USA.Alpha3(): 1234,
		countrycode.FIN.Alpha3(): 2345,
	}
	data, _ := json.Marshal(mapping)
	// Output: {"FIN":2345,"USA":1234}
	fmt.Println(string(data))
}

func requireEqual(tb testing.TB, expected, received interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(expected, received) {
		tb.Fatalf("expected %#v, received %#v", expected, received)
	}
}

func requireNoError(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatalf("expected no error, got %#v", err)
	}
}

func requireTrue(tb testing.TB, v bool) {
	tb.Helper()
	if !v {
		tb.Fatal("expected the value to be true")
	}
}

func requireFalse(tb testing.TB, v bool) {
	tb.Helper()
	if v {
		tb.Fatal("expected the value to be false")
	}
}
