package identifiers

type Identifiers interface {
	Init(opts ...Option) error
	Generate() string
}

type Option func(*Options)