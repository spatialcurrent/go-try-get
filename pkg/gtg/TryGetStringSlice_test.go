// =================================================================
//
// Copyright (C) 2[]string{}19 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryGetStringSlice(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		assert.Equal(t, []string{"a", "b", "c"}, TryGetStringSlice(nil, "foo", []string{"a", "b", "c"}))
	})
	t.Run("MapStringInterface", func(t *testing.T) {
		in := map[string]interface{}{
			"foo":   []string{"a", "b", "c"},
			"hello": "world",
		}
		assert.Equal(t, []string{"a", "b", "c"}, TryGetStringSlice(in, "foo", []string{}))
		assert.Equal(t, []string{"world"}, TryGetStringSlice(in, "hello", []string{}))
	})
	t.Run("MapStringInterfacePointer", func(t *testing.T) {
		assert.Equal(t, []string{"a", "b", "c"}, TryGetStringSlice(&map[string]interface{}{"foo": []string{"a", "b", "c"}}, "foo", []string{}))
	})
	t.Run("MapStringStringSlice", func(t *testing.T) {
		assert.Equal(t, []string{"a", "b", "c"}, TryGetStringSlice(map[string][]string{"foo": []string{"a", "b", "c"}}, "foo", []string{}))
	})
	t.Run("MapStringFunc", func(t *testing.T) {
		in := map[string]interface{}{
			"foo":   func() []string { return []string{"a", "b", "c"} },
			"hello": func() string { return "world" },
		}
		assert.Equal(t, []string{"a", "b", "c"}, TryGetStringSlice(in, "foo", []string{}))
		assert.Equal(t, []string{"world"}, TryGetStringSlice(in, "hello", []string{}))
	})
	t.Run("StructField", func(t *testing.T) {
		assert.Equal(t, []string{"a", "b", "c"}, TryGetStringSlice(struct{ Foo []string }{Foo: []string{"a", "b", "c"}}, "Foo", []string{}))
	})
	t.Run("StructFieldPointer", func(t *testing.T) {
		assert.Equal(t, []string{"a", "b", "c"}, TryGetStringSlice(&struct{ Foo []string }{Foo: []string{"a", "b", "c"}}, "Foo", []string{}))
	})
	t.Run("StructMethodInt", func(t *testing.T) {
		assert.Equal(t, []string{"10"}, TryGetStringSlice(testCaseIntStruct{}, "Int", []string{}))
	})
	t.Run("StructMethodInt32", func(t *testing.T) {
		assert.Equal(t, []string{"10"}, TryGetStringSlice(testCaseIntStruct{}, "Int32", []string{}))
	})
	t.Run("StructMethodInt64", func(t *testing.T) {
		assert.Equal(t, []string{"10"}, TryGetStringSlice(testCaseIntStruct{}, "Int64", []string{}))
	})
}
