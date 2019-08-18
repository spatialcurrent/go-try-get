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

func TestTryGetStringSliceMapStringInterface(t *testing.T) {
	obj := map[string]interface{}{"foo": []string{"a", "b", "c"}}
	assert.Equal(t, []string{"a", "b", "c"}, TryGetStringSlice(obj, "foo", []string{}))
}

func TestTryGetStringSliceMapStringString(t *testing.T) {
	obj := map[string][]string{"foo": []string{"a", "b", "c"}}
	assert.Equal(t, []string{"a", "b", "c"}, TryGetStringSlice(obj, "foo", []string{}))
}

func TestTryGetStringSliceMapStringFunc(t *testing.T) {
	obj := map[string]interface{}{"foo": func() []string { return []string{"a", "b", "c"} }}
	assert.Equal(t, []string{"a", "b", "c"}, TryGetStringSlice(obj, "foo", []string{}))
}

func TestTryGetStringSliceStructField(t *testing.T) {
	obj := struct{ Foo []string }{Foo: []string{"a", "b", "c"}}
	assert.Equal(t, []string{"a", "b", "c"}, TryGetStringSlice(obj, "Foo", []string{}))
}

func TestTryGetStringSliceStructMethod(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, TryGetStringSlice(testCaseStringSliceStruct{}, "Foo", []string{}))
}
