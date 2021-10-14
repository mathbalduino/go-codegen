package goImports

import (
	"reflect"
	"testing"
)

func TestMergeImports(t *testing.T) {
	t.Run("If the given GoImports is empty, return an empty map and dont change the receiver imports", func(t *testing.T) {
		i := &GoImports{"", map[string]string{"alias1": "pkg/1"}}
		m := i.MergeImports(&GoImports{})
		if len(m) != 0 {
			t.Fatalf("Expected an empty map")
		}
		if !reflect.DeepEqual(i.imports, map[string]string{"alias1": "pkg/1"}) {
			t.Fatalf("The receiver imports shouldn't be modified")
		}
	})
	t.Run("Merge the receiver with the given GoImports, returning an empty map when there's no clashes", func(t *testing.T) {
		i1 := &GoImports{"", map[string]string{
			"alias1": "pkg/1",
			"alias2": "pkg/2",
		}}
		i2 := &GoImports{"", map[string]string{
			"alias3": "pkg/3",
			"alias4": "pkg/4",
		}}
		m := i1.MergeImports(i2)
		if len(m) != 0 {
			t.Fatalf("Expected an empty map")
		}
		expected := map[string]string{
			"alias1": "pkg/1",
			"alias2": "pkg/2",
			"alias3": "pkg/3",
			"alias4": "pkg/4",
		}
		if !reflect.DeepEqual(i1.imports, expected) {
			t.Fatalf("Receiver imports expected to be merged")
		}
	})
	t.Run("Merge the receiver with the given GoImports, returning a map with all the fixed clashes", func(t *testing.T) {
		i1 := &GoImports{"", map[string]string{
			"alias1": "pkg/1",
			"alias2": "pkg/2",
			"alias4": "pkg/3",
		}}
		i2 := &GoImports{"", map[string]string{
			"alias3": "pkg/4",
			"alias4": "pkg/5",
			"alias5": "pkg/6",
		}}
		m := i1.MergeImports(i2)
		if !reflect.DeepEqual(m, map[string]string{"alias4": "alias4_2"}) {
			t.Fatalf("Expected a map containing the fixed clashes")
		}
		expected := map[string]string{
			"alias1":   "pkg/1",
			"alias2":   "pkg/2",
			"alias4":   "pkg/3",
			"alias3":   "pkg/4",
			"alias4_2": "pkg/5",
			"alias5":   "pkg/6",
		}
		if !reflect.DeepEqual(i1.imports, expected) {
			t.Fatalf("Receiver imports expected to be merged")
		}
	})
}
