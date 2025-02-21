package content

import (
	"encoding/json"
	"encoding/xml"
	"io"
)

type xmlCodec struct{}

func (c *xmlCodec) Decode(rd io.Reader, target any) error {
	return xml.NewDecoder(rd).Decode(target)
}

func (c *xmlCodec) Encode(w io.Writer, obj any) error {
	return xml.NewEncoder(w).Encode(obj)
}

type jsonCodec struct{}

func (c *jsonCodec) Decode(rd io.Reader, target any) error {
	return json.NewDecoder(rd).Decode(target)
}

func (c *jsonCodec) Encode(w io.Writer, s any) error {
	return json.NewEncoder(w).Encode(s)
}
