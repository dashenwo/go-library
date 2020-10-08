package hashmap

import (
	"encoding/json"
	"github.com/dashenwo/goutils/strings"
)

// ToJSON outputs the JSON representation of the map.
func (m *Map) ToJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	for key, value := range m.data {
		elements[strings.ToString(key)] = value
	}
	return json.Marshal(&elements)
}

// FromJSON populates the map from the input JSON representation.
func (m *Map) FromJSON(data []byte) error {
	elements := make(map[string]interface{})
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for key, value := range elements {
			m.data[key] = value
		}
	}
	return err
}
