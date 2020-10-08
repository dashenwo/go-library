package storage

import "github.com/dashenwo/go-library/session/encoders"

type Options struct {
	Prefix       string            // Set the prefix of the storage store
	Nodes        []string          // Set node information of storage driver
	Auth         []string          // Set the auth information of the storage driver, such as password
	MaxLockWait  int               // Set the maximum wait time of storage drive
	SpinLockWait int               // Set storage drive cycle wait time
	Encoders     encoders.Encoders // Session data serialization
}

// Option sets values in Options
type Option func(o *Options)

func Prefix(prefix string) Option {
	return func(o *Options) {
		o.Prefix = prefix
	}
}

func Nodes(n ...string) Option {
	return func(o *Options) {
		o.Nodes = n
	}
}

func Auth(a ...string) Option {
	return func(o *Options) {
		o.Auth = a
	}
}

func MaxLockWait(time int) Option {
	return func(o *Options) {
		o.MaxLockWait = time
	}
}

func SpinLockWait(time int) Option {
	return func(o *Options) {
		o.SpinLockWait = time
	}
}

func Encoders(encoders encoders.Encoders) Option {
	return func(o *Options) {
		o.Encoders = encoders
	}
}
