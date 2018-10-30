// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

import (
	"reflect"
	"testing"
)

type testCaseTryGetInt struct {
	Object   interface{}
	Name     string
	Fallback int
	Output   int
}

func TestTryGetInt(t *testing.T) {

	testCases := []testCaseTryGetInt{
		testCaseTryGetInt{Object: map[string]interface{}{"a": 10}, Name: "a", Fallback: 20, Output: 10},
		testCaseTryGetInt{Object: map[string]interface{}{"a": func() int { return 10 }}, Name: "a", Fallback: 20, Output: 10},
		testCaseTryGetInt{Object: struct{ Foo int }{Foo: 10}, Name: "Foo", Fallback: 20, Output: 10},
		testCaseTryGetInt{Object: intStruct{}, Name: "Foo", Fallback: 20, Output: 10},
		testCaseTryGetInt{Object: intStruct{}, Name: "Bar", Fallback: 20, Output: 10},
		testCaseTryGetInt{Object: intStruct{}, Name: "Lol", Fallback: 20, Output: 10},
	}

	for _, testCase := range testCases {

		got := TryGetInt(testCase.Object, testCase.Name, testCase.Fallback)
		if !reflect.DeepEqual(got, testCase.Output) {
			t.Errorf("TryGetInt(%v, %v, %v) == %v (%v), want %v (%s)", testCase.Object, testCase.Name, testCase.Fallback, got, reflect.TypeOf(got), testCase.Output, reflect.TypeOf(testCase.Output))
		}

	}

}
