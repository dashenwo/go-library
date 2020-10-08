package storage

type Options struct {
	// 设置存储驱动存储的前缀
	Prefix string
	// 设置session来源
	Source string
	// 设置cookie里面设置的id
	Id string
	// 设置节点信息
	Nodes []string
	// 设置存储的auth信息
	Auth string
	// 最大等待时间
	MaxLockWait int
	// 设置循环等待时间
	SpinLockWait int
}
// Option sets values in Options
type Option func(o *Options)
// 设置存储层的前缀
func Prefix(prefix string) Option {
	return func(o *Options) {
		o.Prefix = prefix
	}
}

// 设置存储层id
func Id(id string) Option {
	return func(o *Options) {
		o.Id = id
	}
}

// session来源，根据这个设置redis目录
func Source(source string) Option {
	return func(o *Options) {
		o.Source = source
	}
}

// 节点信息
func Nodes(n ...string) Option {
	return func(o *Options) {
		o.Nodes = n
	}
}

// 设置存储驱动的auth信息
func Auth(a string) Option {
	return func(o *Options) {
		o.Auth = a
	}
}

// 设置锁最大等待时间
func MaxLockWait(time int) Option {
	return func(o *Options) {
		o.MaxLockWait = time
	}
}

// 设置循环等待时间
func SpinLockWait(time int) Option {
	return func(o *Options) {
		o.SpinLockWait = time
	}
}