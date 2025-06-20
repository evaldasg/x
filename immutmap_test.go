
package x

import (
	"testing"
)

func TestMap(t *testing.T) {
	m := NewMap(map[int]string{
		200: "OK",
		404: "Not Found",
	})

	if v := m.MustGet(200); v != "OK" {
		t.Errorf("expected OK, got %v", v)
	}

	if _, ok := m.Get(500); ok {
		t.Errorf("expected key 500 to be missing")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected MustGet to panic for missing key")
		}
	}()
	m.MustGet(500)

	keys := m.Keys()
	if len(keys) != 2 {
		t.Errorf("expected 2 keys, got %d", len(keys))
	}
}
