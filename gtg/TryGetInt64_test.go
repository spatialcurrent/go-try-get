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

type testCaseTryGetInt64 struct {
	Object   interface{}
	Name     string
	Fallback int64
	Output   int64
}

func TestTryGetInt64(t *testing.T) {

	testCases := []testCaseTryGetInt64{
		testCaseTryGetInt64{Object: map[string]interface{}{"a": 10}, Name: "a", Fallback: int64(20), Output: int64(10)},
		testCaseTryGetInt64{Object: map[string]interface{}{"a": func() int { return 10 }}, Name: "a", Fallback: int64(20), Output: int64(10)},
		testCaseTryGetInt64{Object: struct{ Foo int }{Foo: 10}, Name: "Foo", Fallback: int64(20), Output: int64(10)},
		testCaseTryGetInt64{Object: intStruct{}, Name: "Foo", Fallback: int64(20), Output: int64(10)},
		testCaseTryGetInt64{Object: intStruct{}, Name: "Bar", Fallback: int64(20), Output: int64(10)},
		testCaseTryGetInt64{Object: intStruct{}, Name: "Lol", Fallback: int64(20), Output: int64(10)},
	}

	for _, testCase := range testCases {

		got := TryGetInt64(testCase.Object, testCase.Name, testCase.Fallback)
		if !reflect.DeepEqual(got, testCase.Output) {
			t.Errorf("TryGetInt64(%v, %v, %v) == %v (%v), want %v (%s)", testCase.Object, testCase.Name, testCase.Fallback, got, reflect.TypeOf(got), testCase.Output, reflect.TypeOf(testCase.Output))
		}

	}

}
