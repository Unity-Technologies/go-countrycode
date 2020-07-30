# go-countrycode

[![CI status](https://github.com/Unity-Technologies/go-countrycode/workflows/CI/badge.svg)](https://github.com/Unity-Technologies/go-countrycode/actions)
[![GoDoc](https://godoc.org/github.com/Unity-Technologies/go-countrycode?status.svg)](https://godoc.org/github.com/Unity-Technologies/go-countrycode)

Package `countrycode` provides utilities for representing countries in code, and handling their serializations and deserializations in a convenient way.

All deserializations will result in `CountryUndefined` if input data is not a recognized country code.

All the exported types are one word in memory and as such provide fast equality checks, hashing for usage as keys in maps. Conversions between the types are zero overhead (don't even escape the stack), serialization is worst case `O(1)`, and deserialization is worst case `O(n)`.

## License

MIT License. See [LICENSE](LICENSE) for more details.
