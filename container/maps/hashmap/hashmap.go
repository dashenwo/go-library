package hashmap

import (
	"fmt"
	"github.com/dashenwo/go-library/utils/empty"
	"github.com/dashenwo/go-library/utils/rwmutex"
)

// Map holds the elements in go's native map
type Map struct {
	mu   rwmutex.RWMutex
	data map[interface{}]interface{}
}

// New instantiates a hash map.
func New(safe ...bool) *Map {
	return &Map{
		mu:   rwmutex.Create(safe...),
		data: make(map[interface{}]interface{}),
	}
}

// NewStrAnyMapFrom creates and returns a hash map from given map <data>.
// Note that, the param <data> map will be set as the underlying data map(no deep copy),
// there might be some concurrent-safe issues when changing the map outside.
func NewFrom(data map[interface{}]interface{}, safe ...bool) *Map {
	return &Map{
		mu:   rwmutex.Create(safe...),
		data: data,
	}
}

// Iterator iterates the hash map readonly with custom callback function <f>.
// If <f> returns true, then it continues iterating; or false to stop.
func (m *Map) Iterator(f func(k interface{}, v interface{}) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.data {
		if !f(k, v) {
			break
		}
	}
}

// Clone returns a new hash map with copy of current map data.
func (m *Map) Clone() *Map {
	return NewFrom(m.MapCopy(), !m.mu.IsSafe())
}

// MapCopy returns a copy of the underlying data of the hash map.
func (m *Map) MapCopy() map[interface{}]interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	data := make(map[interface{}]interface{}, len(m.data))
	for k, v := range m.data {
		data[k] = v
	}
	return data
}

// FilterEmpty deletes all key-value pair of which the value is empty.
// Values like: 0, nil, false, "", len(slice/map/chan) == 0 are considered empty.
func (m *Map) FilterEmpty() {
	m.mu.Lock()
	for k, v := range m.data {
		if empty.IsEmpty(v) {
			delete(m.data, k)
		}
	}
	m.mu.Unlock()
}

// FilterNil deletes all key-value pair of which the value is nil.
func (m *Map) FilterNil() {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range m.data {
		if empty.IsNil(v) {
			delete(m.data, k)
		}
	}
}

// Put inserts element into the map.
func (m *Map) Put(key interface{}, value interface{}) {
	m.mu.Lock()
	if m.data == nil {
		m.data = make(map[interface{}]interface{})
	}
	m.data[key] = value
	m.mu.Unlock()
}

// Sets batch sets key-values to the hash map.
func (m *Map) Puts(data map[interface{}]interface{}) {
	m.mu.Lock()
	if m.data == nil {
		m.data = data
	} else {
		for k, v := range data {
			m.data[k] = v
		}
	}
	m.mu.Unlock()
}

// Search searches the map with given <key>.
// Second return parameter <found> is true if key was found, otherwise false.
func (m *Map) Search(key interface{}) (value interface{}, found bool) {
	m.mu.RLock()
	if m.data != nil {
		value, found = m.data[key]
	}
	m.mu.RUnlock()
	return
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	m.mu.RLock()
	if m.data != nil {
		value, found = m.data[key]
	}
	m.mu.RUnlock()
	return
}

// Removes batch deletes values of the map by keys.
func (m *Map) Removes(keys []interface{}) {
	m.mu.Lock()
	if m.data != nil {
		for _, key := range keys {
			delete(m.data, key)
		}
	}
	m.mu.Unlock()
}

// Remove removes the element from the map by key.
func (m *Map) Remove(key string) (value interface{}) {
	m.mu.Lock()
	if m.data != nil {
		var ok bool
		if value, ok = m.data[key]; ok {
			delete(m.data, key)
		}
	}
	m.mu.Unlock()
	return
}

// Empty returns true if map does not contain any elements
func (m *Map) Empty() bool {
	m.mu.RLock()
	r := m.Size() == 0
	m.mu.RUnlock()
	return r
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	m.mu.RLock()
	length := len(m.data)
	m.mu.RUnlock()
	return length
}

// Keys returns all keys (random order).
func (m *Map) Keys() []interface{} {
	m.mu.RLock()
	var (
		keys  = make([]interface{}, len(m.data))
		index = 0
	)
	for key := range m.data {
		keys[index] = key
		index++
	}
	m.mu.RUnlock()
	return keys
}

// Values returns all values (random order).
func (m *Map) Values() []interface{} {
	m.mu.RLock()
	values := make([]interface{}, m.Size())
	count := 0
	for _, value := range m.data {
		values[count] = value
		count++
	}
	m.mu.RUnlock()
	return values
}

// Contains checks whether a key exists.
// It returns true if the <key> exists, or else false.
func (m *Map) Contains(key interface{}) bool {
	var ok bool
	m.mu.RLock()
	if m.data != nil {
		_, ok = m.data[key]
	}
	m.mu.RUnlock()
	return ok
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.mu.Lock()
	m.data = make(map[interface{}]interface{})
	m.mu.Unlock()
}

// Replace the data of the map with given <data>.
func (m *Map) Replace(data map[interface{}]interface{}) {
	m.mu.Lock()
	m.data = data
	m.mu.Unlock()
}

// Merge merges two hash maps.
// The <other> map will be merged into the map <m>.
func (m *Map) Merge(other *Map) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.data == nil {
		m.data = other.MapCopy()
		return
	}
	if other != m {
		other.mu.RLock()
		defer other.mu.RUnlock()
	}
	for k, v := range other.data {
		m.data[k] = v
	}
}

// String returns a string representation of container
func (m *Map) String() string {
	str := "HashMap\n"
	str += fmt.Sprintf("%v", m.data)
	return str
}
