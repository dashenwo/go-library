package encoders

type Encoders interface {
	Init(opts ...Option) error
	Encode(str []byte) string
	Decode(str string) ([]byte,error)
}

type Option func(*Options)