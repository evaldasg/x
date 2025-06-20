
package x

// Map is a generic immutable map.
type Map[K comparable, V any] struct {
	data map[K]V
}

// NewMap creates a new immutable map.
func NewMap[K comparable, V any](m map[K]V) Map[K, V] {
	copyMap := make(map[K]V, len(m))
	for k, v := range m {
		copyMap[k] = v
	}
	return Map[K, V]{data: copyMap}
}

// Get returns the value for key and a boolean flag.
func (m Map[K, V]) Get(key K) (V, bool) {
	v, ok := m.data[key]
	return v, ok
}

// MustGet returns the value for key or panics if key is missing.
func (m Map[K, V]) MustGet(key K) V {
	v, ok := m.Get(key)
	if !ok {
		panic("key not found")
	}
	return v
}

// Keys returns a slice with all keys in the map.
func (m Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(m.data))
	for k := range m.data {
		keys = append(keys, k)
	}
	return keys
}
