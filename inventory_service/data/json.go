package data

import (
	"encoding/json"
	"io"
)

//ToJSON serializes the given interface into string based  json format

func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(i)
}

//FromJSON desirializes the objects from JSON string
// in an io.Reader to the given inteface

func FromJSON(i interface{}, r io.Reader) error {
	d := json.Decoder(r)
	return d.Decode(i)
}
