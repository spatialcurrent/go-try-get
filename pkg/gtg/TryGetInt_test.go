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

func TestTryGetIntMapStringInterface(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(map[string]interface{}{"foo": 10}, "foo", 0))
}

func TestTryGetIntMapStringInterfaceFallback(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(map[string]interface{}{}, "foo", 10))
}

func TestTryGetIntMapStringInt(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(map[string]int{"foo": 10}, "foo", 0))
}

func TestTryGetIntMapStringFunc(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(map[string]interface{}{"foo": func() int { return 10 }}, "foo", 0))
}

func TestTryGetIntStructField(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(struct{ Foo int }{Foo: 10}, "Foo", 0))
}

func TestTryGetIntStructMethodInt(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(testCaseIntStruct{}, "Int", 0))
}

func TestTryGetIntStructMethodInt32(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(testCaseIntStruct{}, "Int32", 0))
}

func TestTryGetIntStructMethodInt64(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(testCaseIntStruct{}, "Int64", 0))
}
