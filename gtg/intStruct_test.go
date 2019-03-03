// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

type intStruct struct{}

func (t intStruct) Foo() int64 {
	return int64(10)
}

func (t intStruct) Bar() int32 {
	return int32(10)
}

func (t intStruct) Lol() int {
	return 10
}
