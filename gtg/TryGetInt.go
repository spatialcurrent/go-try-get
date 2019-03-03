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

	objectType := reflect.TypeOf(obj)
	objectValue := reflect.ValueOf(obj)
	if objectType.Kind() == reflect.Ptr {
		objectType = objectType.Elem()
		objectValue = objectValue.Elem()
	}

	switch objectType.Kind() {
	case reflect.Struct:
		if _, ok := objectType.FieldByName(name); ok {
			v := objectValue.FieldByName(name).Interface()
			if reflect.TypeOf(v).ConvertibleTo(intType) {
				return reflect.ValueOf(v).Convert(intType).Interface().(int)
			}
		} else if _, ok := objectType.MethodByName(name); ok {
			fn := objectValue.MethodByName(name)
			results := fn.Call([]reflect.Value{})
			if len(results) == 1 {
				result := results[0].Interface()
				resultType := reflect.TypeOf(result)
				if resultType.Kind() == reflect.Int {
					return result.(int)
				} else if resultType.ConvertibleTo(intType) {
					return reflect.ValueOf(result).Convert(intType).Interface().(int)
				}
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
					result := results[0].Interface()
					resultType := reflect.TypeOf(result)
					if resultType.Kind() == reflect.Int {
						return result.(int)
					} else if resultType.ConvertibleTo(intType) {
						return reflect.ValueOf(result).Convert(intType).Interface().(int)
					}
				}
			} else if actualType.Kind() == reflect.Int {
				return actual.(int)
			} else if actualType.ConvertibleTo(intType) {
				return reflect.ValueOf(actual).Convert(intType).Interface().(int)
			}
		}
	}

	return fallback
}
