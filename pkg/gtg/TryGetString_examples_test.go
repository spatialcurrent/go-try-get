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
func ExampleTryGetString_mapValue() {
	in := map[string]interface{}{
		"foo": "bar",
	}
	out := TryGetString(in, "foo", "")
	fmt.Println(out)
	// Output: bar
}

// This example shows you can get a string from a map function.
func ExampleTryGetString_mapFunc() {
	in := map[string]interface{}{
		"foo": func() string {
			return "bar"
		},
	}
	out := TryGetString(in, "foo", "")
	fmt.Println(out)
	// Output: bar
}

// This example shows you can get a string from a struct.
func ExampleTryGetString_struct() {
	in := struct{ Foo string }{
		Foo: "bar",
	}
	out := TryGetString(in, "Foo", "")
	fmt.Println(out)
	// Output: bar
}
