// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

import (
	"reflect"
	"testing"
)

type testCaseString struct {
	Object   interface{}
	Name     string
	Fallback string
	Output   string
}

type testCaseStringStruct struct{}

func (t testCaseStringStruct) Foo() string {
	return "bar"
}

func TestTryGetString(t *testing.T) {

	testCases := []testCaseString{
		testCaseString{Object: map[string]interface{}{"a": "x"}, Name: "a", Fallback: "", Output: "x"},
		testCaseString{Object: map[string]string{"a": "foo"}, Name: "a", Fallback: "bar", Output: "foo"},
		testCaseString{Object: map[string]interface{}{"a": func() string { return "x" }}, Name: "a", Fallback: "", Output: "x"},
		testCaseString{Object: struct{ Foo string }{Foo: "bar"}, Name: "Foo", Fallback: "", Output: "bar"},
		testCaseString{Object: testCaseStringStruct{}, Name: "Foo", Fallback: "", Output: "bar"},
	}

	for _, testCase := range testCases {

		got := TryGetString(testCase.Object, testCase.Name, testCase.Fallback)
		if !reflect.DeepEqual(got, testCase.Output) {
			t.Errorf("TryGetString(%v, %v, %v) == %v (%v), want %v (%s)", testCase.Object, testCase.Name, testCase.Fallback, got, reflect.TypeOf(got), testCase.Output, reflect.TypeOf(testCase.Output))
		}

	}

}
