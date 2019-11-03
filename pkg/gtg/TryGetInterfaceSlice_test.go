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

func TestTryGetInterfaceSliceMapStringInterface(t *testing.T) {
	obj := map[string]interface{}{"foo": []interface{}{"a", "b", "c"}}
	assert.Equal(t, []interface{}{"a", "b", "c"}, TryGetInterfaceSlice(obj, "foo", []interface{}{}))
}

func TestTryGetInterfaceSliceMapStringString(t *testing.T) {
	obj := map[string][]interface{}{"foo": []interface{}{"a", "b", "c"}}
	assert.Equal(t, []interface{}{"a", "b", "c"}, TryGetInterfaceSlice(obj, "foo", []interface{}{}))
}

func TestTryGetInterfaceSliceMapStringFunc(t *testing.T) {
	obj := map[string]interface{}{"foo": func() []interface{} { return []interface{}{"a", "b", "c"} }}
	assert.Equal(t, []interface{}{"a", "b", "c"}, TryGetInterfaceSlice(obj, "foo", []interface{}{}))
}

func TestTryGetInterfaceSliceStructField(t *testing.T) {
	obj := struct{ Foo []interface{} }{Foo: []interface{}{"a", "b", "c"}}
	assert.Equal(t, []interface{}{"a", "b", "c"}, TryGetInterfaceSlice(obj, "Foo", []interface{}{}))
}

func TestTryGetInterfaceSliceStructMethod(t *testing.T) {
	assert.Equal(t, []interface{}{"a", "b", "c"}, TryGetInterfaceSlice(testCaseInterfaceSliceStruct{}, "Foo", []interface{}{}))
}
