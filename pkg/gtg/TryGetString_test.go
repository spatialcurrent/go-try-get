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

func TestTryGetStringMapStringInterface(t *testing.T) {
	assert.Equal(t, "bar", TryGetString(map[string]interface{}{"foo": "bar"}, "foo", ""))
}

func TestTryGetStringMapStringInterfaceFallback(t *testing.T) {
	assert.Equal(t, "bar", TryGetString(map[string]interface{}{}, "foo", "bar"))
}

func TestTryGetStringMapStringString(t *testing.T) {
	assert.Equal(t, "bar", TryGetString(map[string]string{"foo": "bar"}, "foo", ""))
}

func TestTryGetStringMapStringFunc(t *testing.T) {
	assert.Equal(t, "bar", TryGetString(map[string]interface{}{"foo": func() string { return "bar" }}, "foo", ""))
}

func TestTryGetStringStructField(t *testing.T) {
	assert.Equal(t, "bar", TryGetString(struct{ Foo string }{Foo: "bar"}, "Foo", ""))
}

func TestTryGetStringStructMethod(t *testing.T) {
	assert.Equal(t, "bar", TryGetString(testCaseStringStruct{}, "Foo", ""))
}
