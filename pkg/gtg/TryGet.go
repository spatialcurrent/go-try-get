// =================================================================
//
// Copyright (C) 2020 Spatial Current, Inc. - All Rights Reserved
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
func TryGet(object interface{}, name string, fallback interface{}) interface{} {
	return tryGetValue(reflect.ValueOf(object), name, fallback)
}

func tryGetValue(objectValue reflect.Value, name string, fallback interface{}) interface{} {

	if !objectValue.IsValid() {
		return fallback
	}

	if !objectValue.CanInterface() {
		return fallback
	}

	objectKind := objectValue.Kind()

	if objectKind == reflect.Ptr {
		return tryGetValue(objectValue.Elem(), name, fallback)
	}

	objectValue = reflect.ValueOf(objectValue.Interface()) // sets value to concerete type
	objectKind = objectValue.Kind()

	switch objectKind {
	case reflect.Struct:
		if field := objectValue.FieldByName(name); field.IsValid() {
			return field.Interface()
		}
		if method := objectValue.MethodByName(name); method.IsValid() {
			if results := method.Call([]reflect.Value{}); len(results) == 1 {
				return results[0].Interface()
			}
		}
	case reflect.Map:
		valueValue := objectValue.MapIndex(reflect.ValueOf(name))
		if !valueValue.IsValid() {
			return fallback
		}
		if !valueValue.CanInterface() {
			return fallback
		}
		valueValue = reflect.ValueOf(valueValue.Interface()) // sets value to concerete type
		valueKind := valueValue.Kind()
		if !valueValue.IsValid() {
			return fallback
		}
		if valueKind == reflect.Func {
			results := valueValue.Call([]reflect.Value{})
			if len(results) == 1 {
				if results[0].IsValid() && results[0].CanInterface() {
					return results[0].Interface()
				}
				return fallback
			}
			return fallback
		}
		return valueValue.Interface()
	}

	return fallback
}
