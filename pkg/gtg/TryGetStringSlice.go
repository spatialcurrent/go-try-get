// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

import (
	"fmt"
	"reflect"
)

func arrayToStringSlice(array reflect.Value) []string {
	if arrayElemKind := array.Type().Elem().Kind(); arrayElemKind == reflect.String {
		return array.Slice(0, array.Len()).Interface().([]string)
	}
	out := make([]string, 0, array.Len())
	for i := 0; i < array.Len(); i++ {
		out = append(out, fmt.Sprint(array.Index(i).Interface()))
	}
	return out
}

func sliceToStringSlice(slc reflect.Value) []string {
	if sliceElemKind := slc.Type().Elem().Kind(); sliceElemKind == reflect.String {
		return slc.Interface().([]string)
	}
	out := make([]string, 0, slc.Len())
	for i := 0; i < slc.Len(); i++ {
		out = append(out, fmt.Sprint(slc.Index(i).Interface()))
	}
	return out
}

func valueToStringSlice(value reflect.Value) []string {
	kind := value.Kind()
	if kind == reflect.String {
		return []string{value.Interface().(string)}
	}
	if kind == reflect.Array {
		return arrayToStringSlice(value)
	}
	if kind == reflect.Slice {
		return sliceToStringSlice(value)
	}
	return []string{fmt.Sprint(value.Interface())}
}

// TryGetStringSlice returns the value from the object (a map or struct) by name.
//  - If the name refers to a struct method or func in a map, then call the function.
//    If the function only returns 1 value, then return the result as a 1 element string slice using `fmt.Sprint`,
//    otherwise return the value of the fallback parameter.
//  - If the name refers to a struct field or map key (and not a function),
//    then return the value by name as a string array using `fmt.Sprint`.
//  - If the value by name does not exist, then returns the value of the fallback parameter.
//
// Examples
//
//  - TryGetStringSlice(map[string]string{"yo":[]string{"a", "b", "c"}}, "yo", "what") == []string{"a", "b", "c"}
//
func TryGetStringSlice(obj interface{}, name string, fallback []string) []string {
	return tryGetStringSliceValue(reflect.ValueOf(obj), name, fallback)
}

func tryGetStringSliceValue(objectValue reflect.Value, name string, fallback []string) []string {

	if !objectValue.IsValid() {
		return fallback
	}

	if !objectValue.CanInterface() {
		return fallback
	}

	objectKind := objectValue.Kind()

	if objectKind == reflect.Ptr {
		return tryGetStringSliceValue(objectValue.Elem(), name, fallback)
	}

	objectValue = reflect.ValueOf(objectValue.Interface()) // sets value to concerete type
	objectKind = objectValue.Kind()

	switch objectKind {
	case reflect.Struct:
		if field := objectValue.FieldByName(name); field.IsValid() {
			return valueToStringSlice(field)
		}
		if method := objectValue.MethodByName(name); method.IsValid() {
			if results := method.Call([]reflect.Value{}); len(results) == 1 {
				return valueToStringSlice(results[0])
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
			if results := valueValue.Call([]reflect.Value{}); len(results) == 1 {
				if result := results[0]; result.IsValid() && result.CanInterface() {
					return valueToStringSlice(result)
				}
			}
			return fallback
		}
		return valueToStringSlice(valueValue)
	}
	return fallback
}
