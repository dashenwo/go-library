package maps

// Map interface that all maps implement
type Map interface {
	Put(key interface{}, value interface{})
	Get(key interface{}) (value interface{}, found bool)
	Remove(key interface{})
	Keys() []interface{}
	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
}

// BidiMap interface that all bidirectional maps implement (extends the Map interface)
type BidiMap interface {
	GetKey(value interface{}) (key interface{}, found bool)
	Map
}
