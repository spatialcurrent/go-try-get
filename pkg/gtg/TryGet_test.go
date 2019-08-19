// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTryGetMapStringInterface(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(map[string]interface{}{"foo": 10}, "foo", 0))
}

func TestTryGetMapStringInt(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(map[string]int{"foo": 10}, "foo", 0))
}

func TestTryGetMapStringFunc(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(map[string]interface{}{"foo": func() int { return 10 }}, "foo", 0))
}

func TestTryGetStructField(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(struct{ Foo int }{Foo: 10}, "Foo", 0))
}

func TestTryGetStructMethodInt(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(testCaseIntStruct{}, "Int", 0))
}

func TestTryGetStructMethodInt32(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(testCaseIntStruct{}, "Int32", 0))
}

func TestTryGetStructMethodInt64(t *testing.T) {
	assert.Equal(t, 10, TryGetInt(testCaseIntStruct{}, "Int64", 0))
}
