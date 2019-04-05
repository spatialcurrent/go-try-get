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

type testCaseTryGet struct {
	Object   interface{}
	Name     string
	Fallback interface{}
	Output   interface{}
}

type testCaseIntStruct struct{}

func (t testCaseIntStruct) Foo() int {
	return 10
}

func TestTryGet(t *testing.T) {

	testCases := []testCaseTryGet{
		testCaseTryGet{Object: map[string]interface{}{"a": 10}, Name: "a", Fallback: 20, Output: 10},
		testCaseTryGet{Object: map[string]interface{}{"a": "foo"}, Name: "a", Fallback: "bar", Output: "foo"},
		testCaseTryGet{Object: map[string]interface{}{"a": func() int { return 10 }}, Name: "a", Fallback: 20, Output: 10},
		testCaseTryGet{Object: struct{ Foo int }{Foo: 10}, Name: "Foo", Fallback: 20, Output: 10},
		testCaseTryGet{Object: testCaseIntStruct{}, Name: "Foo", Fallback: 20, Output: 10},
	}

	for _, testCase := range testCases {

		got := TryGet(testCase.Object, testCase.Name, testCase.Fallback)
		if !reflect.DeepEqual(got, testCase.Output) {
			t.Errorf("TryGetString(%v, %v, %v) == %v (%v), want %v (%s)", testCase.Object, testCase.Name, testCase.Fallback, got, reflect.TypeOf(got), testCase.Output, reflect.TypeOf(testCase.Output))
		}

	}

}
