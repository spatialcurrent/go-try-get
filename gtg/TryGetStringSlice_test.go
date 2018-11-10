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

type testCaseStringSlice struct {
	Object   interface{}
	Name     string
	Fallback []string
	Output   []string
}

type testCaseStringSliceStruct struct{}

func (t testCaseStringSliceStruct) Foo() []string {
	return []string{"x", "y", "z"}
}

func TestTryGetStringSlice(t *testing.T) {

	testCases := []testCaseStringSlice{
		testCaseStringSlice{
			Object:   map[string]interface{}{"a": []string{"x", "y", "z"}},
			Name:     "a",
			Fallback: []string{},
			Output:   []string{"x", "y", "z"},
		},
		testCaseStringSlice{
			Object:   map[string]interface{}{"a": func() []string { return []string{"x", "y", "z"} }},
			Name:     "a",
			Fallback: []string{},
			Output:   []string{"x", "y", "z"},
		},
		testCaseStringSlice{
			Object:   struct{ Foo []string }{Foo: []string{"x", "y", "z"}},
			Name:     "Foo",
			Fallback: []string{},
			Output:   []string{"x", "y", "z"},
		},
		testCaseStringSlice{
			Object:   testCaseStringSliceStruct{},
			Name:     "Foo",
			Fallback: []string{},
			Output:   []string{"x", "y", "z"},
		},
	}

	for _, testCase := range testCases {

		got := TryGetStringSlice(testCase.Object, testCase.Name, testCase.Fallback)
		if !reflect.DeepEqual(got, testCase.Output) {
			t.Errorf("TryGetStringSlice(%v, %v, %v) == %v (%v), want %v (%s)", testCase.Object, testCase.Name, testCase.Fallback, got, reflect.TypeOf(got), testCase.Output, reflect.TypeOf(testCase.Output))
		}

	}

}
