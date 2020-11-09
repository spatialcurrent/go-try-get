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

var intType = reflect.TypeOf(0)

// TryGetInt returns the value from the object (a map or struct) by name.
//  - If the name refers to a struct method or func in a map, then call the function.
//    If the function only returns 1 value and it is convertible to the int type, then return the result,
//    otherwise return the value of the fallback parameter.
//  - If the name refers to a struct field or map key (and not a function) and it the value if convertible to the int type,
//    then return its value directly.
//  - If the value by name does not exist, then returns the value of the fallback parameter.
//
// Examples
//
//  - TryGetInt(map[string]int{"yo":10}, "yo", 20) == 10
//  - TryGetInt(map[string]int{"yo":10}, "hey", 20) == 20
//
func TryGetInt(obj interface{}, name string, fallback int) int {
	return tryGetIntValue(reflect.ValueOf(obj), name, fallback)
}

func tryGetIntValue(objectValue reflect.Value, name string, fallback int) int {

	if !objectValue.IsValid() {
		return fallback
	}

	if !objectValue.CanInterface() {
		return fallback
	}

	objectKind := objectValue.Kind()

	if objectKind == reflect.Ptr {
		return tryGetIntValue(objectValue.Elem(), name, fallback)
	}

	objectValue = reflect.ValueOf(objectValue.Interface()) // sets value to concerete type
	objectKind = objectValue.Kind()

	switch objectKind {
	case reflect.Struct:
		if field := objectValue.FieldByName(name); field.IsValid() {
			if fieldKind := field.Kind(); fieldKind == reflect.Int {
				return field.Interface().(int)
			}
			if fieldType := field.Type(); fieldType.ConvertibleTo(intType) {
				return field.Convert(intType).Interface().(int)
			}
		}
		if method := objectValue.MethodByName(name); method.IsValid() {
			if results := method.Call([]reflect.Value{}); len(results) == 1 {
				result := results[0]
				if resultKind := result.Kind(); resultKind == reflect.Int {
					return result.Interface().(int)
				}
				if resultType := result.Type(); resultType.ConvertibleTo(intType) {
					return result.Convert(intType).Interface().(int)
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
		if valueKind == reflect.Int {
			return valueValue.Interface().(int)
		}
		if valueType := valueValue.Type(); valueType.ConvertibleTo(intType) {
			return valueValue.Convert(intType).Interface().(int)
		}
		if valueKind == reflect.Func {
			results := valueValue.Call([]reflect.Value{})
			if len(results) == 1 {
				result := results[0]
				if result.IsValid() && result.CanInterface() {
					resultKind := result.Kind()
					if resultKind == reflect.Int {
						return result.Interface().(int)
					}
					if resultType := result.Type(); resultType.ConvertibleTo(intType) {
						return result.Convert(intType).Interface().(int)
					}
				}
			}
		}
	}
	return fallback
}
