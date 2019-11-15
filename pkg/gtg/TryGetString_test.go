// =================================================================
//
// Copyright (C) 2""19 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryGetString(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		assert.Equal(t, "bar", TryGetString(nil, "foo", "bar"))
	})
	t.Run("MapStringInterface", func(t *testing.T) {
		assert.Equal(t, "bar", TryGetString(map[string]interface{}{"foo": "bar"}, "foo", ""))
	})
	t.Run("MapStringInterfacePointer", func(t *testing.T) {
		assert.Equal(t, "bar", TryGetString(&map[string]interface{}{"foo": "bar"}, "foo", ""))
	})
	t.Run("MapStringString", func(t *testing.T) {
		assert.Equal(t, "bar", TryGetString(map[string]string{"foo": "bar"}, "foo", ""))
	})
	t.Run("MapStringFunc", func(t *testing.T) {
		assert.Equal(t, "bar", TryGetString(map[string]interface{}{"foo": func() string { return "bar" }}, "foo", ""))
	})
	t.Run("StructField", func(t *testing.T) {
		assert.Equal(t, "bar", TryGetString(struct{ Foo string }{Foo: "bar"}, "Foo", ""))
	})
	t.Run("StructFieldPointer", func(t *testing.T) {
		assert.Equal(t, "bar", TryGetString(&struct{ Foo string }{Foo: "bar"}, "Foo", ""))
	})
	t.Run("StructMethodInt", func(t *testing.T) {
		assert.Equal(t, "10", TryGetString(testCaseIntStruct{}, "Int", ""))
	})
	t.Run("StructMethodInt32", func(t *testing.T) {
		assert.Equal(t, "10", TryGetString(testCaseIntStruct{}, "Int32", ""))
	})
	t.Run("StructMethodInt64", func(t *testing.T) {
		assert.Equal(t, "10", TryGetString(testCaseIntStruct{}, "Int64", ""))
	})
}
