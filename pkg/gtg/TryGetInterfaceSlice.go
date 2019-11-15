// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

import (
	"reflect"
)

func arrayToInterfaceSlice(array reflect.Value) []interface{} {
	if arrayElemKind := array.Type().Elem().Kind(); arrayElemKind == reflect.Interface {
		return array.Slice(0, array.Len()).Interface().([]interface{})
	}
	out := make([]interface{}, 0, array.Len())
	for i := 0; i < array.Len(); i++ {
		out = append(out, array.Index(i).Interface())
	}
	return out
}

func sliceToInterfaceSlice(slc reflect.Value) []interface{} {
	if sliceElemKind := slc.Type().Elem().Kind(); sliceElemKind == reflect.Interface {
		return slc.Interface().([]interface{})
	}
	out := make([]interface{}, 0, slc.Len())
	for i := 0; i < slc.Len(); i++ {
		out = append(out, slc.Index(i).Interface())
	}
	return out
}

func valueToInterfaceSlice(value reflect.Value) []interface{} {
	kind := value.Kind()
	if kind == reflect.String {
		return []interface{}{value.Interface()}
	}
	if kind == reflect.Array {
		return arrayToInterfaceSlice(value)
	}
	if kind == reflect.Slice {
		return sliceToInterfaceSlice(value)
	}
	return []interface{}{value.Interface()}
}

// TryGetInterfaceSlice returns the value from the object (a map or struct) by name.
//  - If the name refers to a struct method or func in a map, then call the function.
//    If the function only returns 1 value, then return the result as a 1 element string slice,
//    otherwise return the value of the fallback parameter.
//  - If the name refers to a struct field or map key (and not a function),
//    then return the value by name as a interface slice.
//  - If the value by name does not exist, then returns the value of the fallback parameter.
//
// Examples
//
//  - TryGetInterfaceSlice(map[string]string{"yo":[]interface{}{"a", "b", "c"}}, "yo", "what") == []interface{}{"a", "b", "c"}
//
func TryGetInterfaceSlice(obj interface{}, name string, fallback []interface{}) []interface{} {
	return tryGetInterfaceSliceValue(reflect.ValueOf(obj), name, fallback)
}

func tryGetInterfaceSliceValue(objectValue reflect.Value, name string, fallback []interface{}) []interface{} {

	if !objectValue.IsValid() {
		return fallback
	}

	if !objectValue.CanInterface() {
		return fallback
	}

	objectKind := objectValue.Kind()

	if objectKind == reflect.Ptr {
		return tryGetInterfaceSliceValue(objectValue.Elem(), name, fallback)
	}

	objectValue = reflect.ValueOf(objectValue.Interface()) // sets value to concerete type
	objectKind = objectValue.Kind()

	switch objectKind {
	case reflect.Struct:
		if field := objectValue.FieldByName(name); field.IsValid() {
			return valueToInterfaceSlice(field)
		}
		if method := objectValue.MethodByName(name); method.IsValid() {
			if results := method.Call([]reflect.Value{}); len(results) == 1 {
				return valueToInterfaceSlice(results[0])
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
		if valueKind == reflect.Func {
			if results := valueValue.Call([]reflect.Value{}); len(results) == 1 {
				if result := results[0]; result.IsValid() && result.CanInterface() {
					return valueToInterfaceSlice(result)
				}
			}
			return fallback
		}
		return valueToInterfaceSlice(valueValue)
	}
	return fallback
}
