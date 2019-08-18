// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestTryGetInt64MapStringInterface(t *testing.T) {
	assert.Equal(t, int64(10), TryGetInt64(map[string]interface{}{"foo": int64(10)}, "foo", int64(0)))
}

func TestTryGetInt64MapStringInterfaceFallback(t *testing.T) {
	assert.Equal(t, int64(10), TryGetInt64(map[string]interface{}{}, "foo", int64(10)))
}

func TestTryGetInt64MapStringInt(t *testing.T) {
	assert.Equal(t, int64(10), TryGetInt64(map[string]int64{"foo": int64(10)}, "foo", int64(0)))
}

func TestTryGetInt64MapStringFunc(t *testing.T) {
	assert.Equal(t, int64(10), TryGetInt64(map[string]interface{}{"foo": func() int64 { return int64(10) }}, "foo", int64(0)))
}

func TestTryGetInt64StructField(t *testing.T) {
	assert.Equal(t, int64(10), TryGetInt64(struct{ Foo int64 }{Foo: int64(10)}, "Foo", int64(0)))
}

func TestTryGetInt64StructMethodInt(t *testing.T) {
	assert.Equal(t, int64(10), TryGetInt64(testCaseIntStruct{}, "Int", int64(0)))
}

func TestTryGetInt64StructMethodInt32(t *testing.T) {
	assert.Equal(t, int64(10), TryGetInt64(testCaseIntStruct{}, "Int32", int64(0)))
}

func TestTryGetInt64StructMethodInt64(t *testing.T) {
	assert.Equal(t, int64(10), TryGetInt64(testCaseIntStruct{}, "Int64", int64(0)))
}
