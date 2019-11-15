// =================================================================
//
// Copyright (C) 2[]interface{}{}19 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryGetInterfaceSlice(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		assert.Equal(t, []interface{}{"a", "b", "c"}, TryGetInterfaceSlice(nil, "foo", []interface{}{"a", "b", "c"}))
	})
	t.Run("MapStringInterface", func(t *testing.T) {
		in := map[string]interface{}{
			"foo":   []interface{}{"a", "b", "c"},
			"hello": "world",
		}
		assert.Equal(t, []interface{}{"a", "b", "c"}, TryGetInterfaceSlice(in, "foo", []interface{}{}))
		assert.Equal(t, []interface{}{"world"}, TryGetInterfaceSlice(in, "hello", []interface{}{}))
	})
	t.Run("MapStringInterfacePointer", func(t *testing.T) {
		assert.Equal(t, []interface{}{"a", "b", "c"}, TryGetInterfaceSlice(&map[string]interface{}{"foo": []interface{}{"a", "b", "c"}}, "foo", []interface{}{}))
	})
	t.Run("MapStringStringSlice", func(t *testing.T) {
		assert.Equal(t, []interface{}{"a", "b", "c"}, TryGetInterfaceSlice(map[string][]interface{}{"foo": []interface{}{"a", "b", "c"}}, "foo", []interface{}{}))
	})
	t.Run("MapStringFunc", func(t *testing.T) {
		in := map[string]interface{}{
			"foo":   func() []interface{} { return []interface{}{"a", "b", "c"} },
			"hello": func() string { return "world" },
		}
		assert.Equal(t, []interface{}{"a", "b", "c"}, TryGetInterfaceSlice(in, "foo", []interface{}{}))
		assert.Equal(t, []interface{}{"world"}, TryGetInterfaceSlice(in, "hello", []interface{}{}))
	})
	t.Run("StructField", func(t *testing.T) {
		assert.Equal(t, []interface{}{"a", "b", "c"}, TryGetInterfaceSlice(struct{ Foo []interface{} }{Foo: []interface{}{"a", "b", "c"}}, "Foo", []interface{}{}))
	})
	t.Run("StructFieldPointer", func(t *testing.T) {
		assert.Equal(t, []interface{}{"a", "b", "c"}, TryGetInterfaceSlice(&struct{ Foo []interface{} }{Foo: []interface{}{"a", "b", "c"}}, "Foo", []interface{}{}))
	})
	t.Run("StructMethodInt", func(t *testing.T) {
		assert.Equal(t, []interface{}{10}, TryGetInterfaceSlice(testCaseIntStruct{}, "Int", []interface{}{}))
	})
	t.Run("StructMethodInt32", func(t *testing.T) {
		assert.Equal(t, []interface{}{int32(10)}, TryGetInterfaceSlice(testCaseIntStruct{}, "Int32", []interface{}{}))
	})
	t.Run("StructMethodInt64", func(t *testing.T) {
		assert.Equal(t, []interface{}{int64(10)}, TryGetInterfaceSlice(testCaseIntStruct{}, "Int64", []interface{}{}))
	})
}
