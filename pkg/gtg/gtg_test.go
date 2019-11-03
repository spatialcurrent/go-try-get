// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

type testCaseIntStruct struct{}

func (t testCaseIntStruct) Int64() int64 {
	return int64(10)
}

func (t testCaseIntStruct) Int32() int32 {
	return int32(10)
}

func (t testCaseIntStruct) Int() int {
	return 10
}

type testCaseInterfaceSliceStruct struct{}

func (t testCaseInterfaceSliceStruct) Foo() []interface{} {
	return []interface{}{"a", "b", "c"}
}

type testCaseStringSliceStruct struct{}

func (t testCaseStringSliceStruct) Foo() []string {
	return []string{"a", "b", "c"}
}

type testCaseStringStruct struct{}

func (t testCaseStringStruct) Foo() string {
	return "bar"
}
