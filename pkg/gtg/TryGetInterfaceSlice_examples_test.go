// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

import (
	"fmt"
)

// This example shows you can get a string from a map value.
func ExampleTryGetInterfaceSlice_mapValue() {
	in := map[string]interface{}{
		"foo": []interface{}{"a", "b", "c"},
	}
	out := TryGetInterfaceSlice(in, "foo", []interface{}{})
	fmt.Println(out)
	// Output: [a b c]
}

// This example shows you can get a string from a map function.
func ExampleTryGetInterfaceSlice_mapFunc() {
	in := map[string]interface{}{
		"foo": func() []interface{} {
			return []interface{}{"a", "b", "c"}
		},
	}
	out := TryGetInterfaceSlice(in, "foo", []interface{}{})
	fmt.Println(out)
	// Output: [a b c]
}

// This example shows you can get a string from a struct.
func ExampleTryGetInterfaceSlice_struct() {
	in := struct{ Foo []interface{} }{
		Foo: []interface{}{"a", "b", "c"},
	}
	out := TryGetInterfaceSlice(in, "Foo", []interface{}{})
	fmt.Println(out)
	// Output: [a b c]
}
