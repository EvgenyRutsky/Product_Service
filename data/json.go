package data

import (
	"encoding/json"
	"io"
)

func (p *Product) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(p)
}

func (p *Products) ToJSON (w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}