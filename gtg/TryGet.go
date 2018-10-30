// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

import (
	"reflect"
)

// TryGet returns the value from the object (a map or struct) by name.
//  - If the name refers to a struct method or func in a map, then call the function.
//    If the function only returns 1 value, then return the result,
//    otherwise return the value of the fallback parameter.
//  - If the name refers to a struct field or map key (and not a function), then return its value directly.
//  - If the value by name does not exist, then returns the value of the fallback parameter.
//
// Examples
//
//  - TryGet(map[string]string{"yo":"yo"}, "yo", "what") == "yo"
//
func TryGet(obj interface{}, name string, fallback interface{}) interface{} {

	objectType := reflect.TypeOf(obj)
	objectValue := reflect.ValueOf(obj)
	if objectType.Kind() == reflect.Ptr {
		objectType = objectType.Elem()
		objectValue = objectValue.Elem()
	}

	switch objectType.Kind() {
	case reflect.Struct:
		if _, ok := objectType.FieldByName(name); ok {
			return objectValue.FieldByName(name).Interface()
		}
		if _, ok := objectType.MethodByName(name); ok {
			fn := objectValue.MethodByName(name)
			results := fn.Call([]reflect.Value{})
			if len(results) == 1 {
				return results[0].Interface()
			}
		}
	case reflect.Map:
		value := reflect.ValueOf(obj).MapIndex(reflect.ValueOf(name))
		if value.IsValid() && !value.IsNil() {
			actual := value.Interface()
			actualType := reflect.TypeOf(actual)
			if actualType.Kind() == reflect.Func {
				results := reflect.ValueOf(actual).Call([]reflect.Value{})
				if len(results) == 1 {
					return results[0].Interface()
				}
			} else {
				return actual
			}
		}
	}

	return fallback
}
