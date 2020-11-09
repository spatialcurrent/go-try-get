// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryGetInt64(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		assert.Equal(t, int64(10), TryGetInt64(nil, "foo", int64(10)))
	})
	t.Run("MapStringInterface", func(t *testing.T) {
		assert.Equal(t, int64(10), TryGetInt64(map[string]interface{}{"foo": int64(10)}, "foo", 0))
	})
	t.Run("MapStringInterfacePointer", func(t *testing.T) {
		assert.Equal(t, int64(10), TryGetInt64(&map[string]interface{}{"foo": int64(10)}, "foo", 0))
	})
	t.Run("MapStringInt", func(t *testing.T) {
		assert.Equal(t, int64(10), TryGetInt64(map[string]int64{"foo": int64(10)}, "foo", 0))
	})
	t.Run("MapStringFunc", func(t *testing.T) {
		assert.Equal(t, int64(10), TryGetInt64(map[string]interface{}{"foo": func() int64 { return int64(10) }}, "foo", 0))
	})
	t.Run("StructField", func(t *testing.T) {
		assert.Equal(t, int64(10), TryGetInt64(struct{ Foo int64 }{Foo: int64(10)}, "Foo", 0))
	})
	t.Run("StructFieldPointer", func(t *testing.T) {
		assert.Equal(t, int64(10), TryGetInt64(&struct{ Foo int64 }{Foo: int64(10)}, "Foo", 0))
	})
	t.Run("StructMethodInt", func(t *testing.T) {
		assert.Equal(t, int64(10), TryGetInt64(testCaseIntStruct{}, "Int", 0))
	})
	t.Run("StructMethodInt32", func(t *testing.T) {
		assert.Equal(t, int64(10), TryGetInt64(testCaseIntStruct{}, "Int32", 0))
	})
	t.Run("StructMethodInt64", func(t *testing.T) {
		assert.Equal(t, int64(10), TryGetInt64(testCaseIntStruct{}, "Int64", 0))
	})
}
