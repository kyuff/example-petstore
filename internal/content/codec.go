package content

import (
	"io"
	"net/http"
)

const (
	defaultEncoding   = "application/json"
	contentTypeHeader = "Content-Type"
)

var contentTypes = map[string]encoding{
	defaultEncoding:   &jsonCodec{},
	"application/xml": &xmlCodec{},
}

type encoding interface {
	Decode(rd io.Reader, target any) error
	Encode(w io.Writer, s any) error
}

func New[T any](r *http.Request, w http.ResponseWriter) *Codec[T] {
	var (
		contentType = r.Header.Get(contentTypeHeader)
		e           encoding
	)

	e, ok := contentTypes[contentType]
	if !ok {
		e = contentTypes[defaultEncoding]
		contentType = defaultEncoding
	}

	return &Codec[T]{
		r:           r,
		w:           w,
		contentType: contentType,
		encoding:    e,
	}
}

type Codec[T any] struct {
	r           *http.Request
	w           http.ResponseWriter
	contentType string
	encoding    encoding
}

func (rw *Codec[T]) Read() (T, error) {
	var t T
	err := rw.encoding.Decode(rw.r.Body, &t)
	if err != nil {
		return t, err
	}

	return t, nil
}

func (rw *Codec[T]) Write(httpStatus int, content any) {
	rw.w.WriteHeader(httpStatus)
	if content == nil {
		return
	}
	rw.w.Header().Set(contentTypeHeader, rw.contentType)
	_ = rw.encoding.Encode(rw.w, content)
	// TODO handle(log) error
}
