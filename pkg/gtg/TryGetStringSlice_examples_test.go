// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

import (
	"fmt"
)

// This example shows you can get a string from a map value.
func ExampleTryGetStringSlice_mapValue() {
	in := map[string]interface{}{
		"foo": []string{"a", "b", "c"},
	}
	out := TryGetStringSlice(in, "foo", []string{})
	fmt.Println(out)
	// Output: [a b c]
}

// This example shows you can get a string from a map function.
func ExampleTryGetStringSlice_mapFunc() {
	in := map[string]interface{}{
		"foo": func() []string {
			return []string{"a", "b", "c"}
		},
	}
	out := TryGetStringSlice(in, "foo", []string{})
	fmt.Println(out)
	// Output: [a b c]
}

// This example shows you can get a string from a struct.
func ExampleTryGetStringSlice_struct() {
	in := struct{ Foo []string }{
		Foo: []string{"a", "b", "c"},
	}
	out := TryGetStringSlice(in, "Foo", []string{})
	fmt.Println(out)
	// Output: [a b c]
}
