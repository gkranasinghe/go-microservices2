package data

import (
	"encoding/json"
	"io"
)

// ToJson serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not have to buffer the output into an memory slice of bytes
// this reduces allocations and overheads of the service
func ToJSON(i interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(i)
}

// FromJson deserializes the object from JSON string
// in an io.Reader to given interface
func FromJSON(i interface{}, r io.Reader) error {
	return json.NewDecoder(r).Decode(i)
}
