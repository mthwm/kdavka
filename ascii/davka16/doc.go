// Package davka16 provides functionality to generate ASCII-formatted dávka 16 files
// for Czech healthcare settlements of spa therapeutic rehabilitation care.
// It follows the specification from "Datové rozhraní - individuální doklady" version 6.2.XXXIa.
// Output is encoded in CP852 (PC Latin-2) as required by the specification.
//
// Usage:
//   - Define Header, Document, and associated structs.
//   - Use Generate() to produce the file content as []byte in CP852 encoding.
//   - Optionally write to a file or io.Writer.
//
// Dependencies:
//   - golang.org/x/text/encoding/charmap (for CP852 encoding)
//
// Example:
//
//	header := davka16.Header{...}
//	docs := []davka16.Document{...}
//	data, err := davka16.Generate(header, docs)
//	if err != nil { ... }
//	ioutil.WriteFile("KDAVKA.001", data, 0644)
package davka16
