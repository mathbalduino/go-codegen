package tsImports

import "testing"

func TestNew(t *testing.T) {
	t.Run("Should never return a nil pointer", func(t *testing.T) {
		if New() == nil {
			t.Fatalf("Expected to be not nil")
		}
	})
	t.Run("Should return a slice with len equals to 0 and cap equals to 5", func(t *testing.T) {
		i := New()
		if len(*i) != 0 {
			t.Fatalf("Expected the len to be 0")
		}
		if cap(*i) != 5 {
			t.Fatalf("Expected the cap to be 5")
		}
	})
}
