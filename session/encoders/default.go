package encoders

import (
	"encoding/base64"
)

var (
	DefaultEncoders = NewEncoders()
	ENCODE_CHARS = map[string]string{
		"+":"-",
		"/":"_",
	}
	DECODE_CHARS = map[string]string{
		"-":"+",
		"_":"/",
	}
)

type noop struct {
	opts Options
}

func NewEncoders(opts ...Option) Encoders {
	options := Options{}

	for _, o := range opts {
		o(&options)
	}

	return &noop{
		opts: options,
	}
}

func (b *noop) Init(opts ...Option) error {
	for _, o := range opts {
		o(&b.opts)
	}
	return nil
}

func (b *noop) Encode(str []byte) string {
	return base64.RawURLEncoding.EncodeToString(str)
}

func (b *noop) Decode(str string) ([]byte,error)  {
	return base64.RawURLEncoding.DecodeString(str)
}