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

var int64Type = reflect.TypeOf(int64(0))

// TryGetInt64 returns the value from the object (a map or struct) by name.
//  - If the name refers to a struct method or func in a map, then call the function.
//    If the function only returns 1 value and it is convertible to the int64 type, then return the result,
//    otherwise return the value of the fallback parameter.
//  - If the name refers to a struct field or map key (and not a function) and it the value if convertible to the int64 type,
//    then return its value directly.
//  - If the value by name does not exist, then returns the value of the fallback parameter.
//
// Examples
//
//  - TryGetInt64(map[string]int{"yo":10}, "yo", int64(20)) == int64(10)
//  - TryGetInt64(map[string]int{"yo":10}, "hey", int64(20)) == int64(20)
//
func TryGetInt64(obj interface{}, name string, fallback int64) int64 {
	return tryGetInt64Value(reflect.ValueOf(obj), name, fallback)
}

func tryGetInt64Value(objectValue reflect.Value, name string, fallback int64) int64 {

	if !objectValue.IsValid() {
		return fallback
	}

	if !objectValue.CanInterface() {
		return fallback
	}

	objectKind := objectValue.Kind()

	if objectKind == reflect.Ptr {
		return tryGetInt64Value(objectValue.Elem(), name, fallback)
	}

	objectValue = reflect.ValueOf(objectValue.Interface()) // sets value to concerete type
	objectKind = objectValue.Kind()

	switch objectKind {
	case reflect.Struct:
		if field := objectValue.FieldByName(name); field.IsValid() {
			if fieldKind := field.Kind(); fieldKind == reflect.Int64 {
				return field.Interface().(int64)
			}
			if fieldType := field.Type(); fieldType.ConvertibleTo(int64Type) {
				return field.Convert(intType).Interface().(int64)
			}
		}
		if method := objectValue.MethodByName(name); method.IsValid() {
			if results := method.Call([]reflect.Value{}); len(results) == 1 {
				result := results[0]
				if resultKind := result.Kind(); resultKind == reflect.Int64 {
					return result.Interface().(int64)
				}
				if resultType := result.Type(); resultType.ConvertibleTo(int64Type) {
					return result.Convert(int64Type).Interface().(int64)
				}
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
		if valueKind == reflect.Int64 {
			return valueValue.Interface().(int64)
		}
		if valueType := valueValue.Type(); valueType.ConvertibleTo(intType) {
			return valueValue.Convert(int64Type).Interface().(int64)
		}
		if valueKind == reflect.Func {
			results := valueValue.Call([]reflect.Value{})
			if len(results) == 1 {
				result := results[0]
				if result.IsValid() && result.CanInterface() {
					resultKind := result.Kind()
					if resultKind == reflect.Int64 {
						return result.Interface().(int64)
					}
					if resultType := result.Type(); resultType.ConvertibleTo(int64Type) {
						return result.Convert(int64Type).Interface().(int64)
					}
					return fallback
				}
			}
		}
	}
	return fallback
}
