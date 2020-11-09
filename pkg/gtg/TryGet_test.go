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

func TestTryGet(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		assert.Equal(t, 10, TryGet(nil, "foo", 10))
	})
	t.Run("MapStringInterface", func(t *testing.T) {
		assert.Equal(t, 10, TryGet(map[string]interface{}{"foo": 10}, "foo", 0))
	})
	t.Run("MapStringInterfacePointer", func(t *testing.T) {
		assert.Equal(t, 10, TryGet(&map[string]interface{}{"foo": 10}, "foo", 0))
	})
	t.Run("MapStringInt", func(t *testing.T) {
		assert.Equal(t, 10, TryGet(map[string]int{"foo": 10}, "foo", 0))
	})
	t.Run("MapStringFunc", func(t *testing.T) {
		assert.Equal(t, 10, TryGet(map[string]interface{}{"foo": func() int { return 10 }}, "foo", 0))
	})
	t.Run("StructField", func(t *testing.T) {
		assert.Equal(t, 10, TryGet(struct{ Foo int }{Foo: 10}, "Foo", 0))
	})
	t.Run("StructFieldPointer", func(t *testing.T) {
		assert.Equal(t, 10, TryGet(&struct{ Foo int }{Foo: 10}, "Foo", 0))
	})
	t.Run("StructMethodInt", func(t *testing.T) {
		assert.Equal(t, 10, TryGet(testCaseIntStruct{}, "Int", 0))
	})
	t.Run("StructMethodInt32", func(t *testing.T) {
		assert.Equal(t, int32(10), TryGet(testCaseIntStruct{}, "Int32", 0))
	})
	t.Run("StructMethodInt64", func(t *testing.T) {
		assert.Equal(t, int64(10), TryGet(testCaseIntStruct{}, "Int64", 0))
	})
}
