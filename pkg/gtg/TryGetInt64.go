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

	objectValue := reflect.ValueOf(obj)
	for reflect.TypeOf(objectValue.Interface()).Kind() == reflect.Ptr {
		objectValue = objectValue.Elem()
	}
	objectValue = reflect.ValueOf(objectValue.Interface()) // sets value to concerete type
	objectType := objectValue.Type()

	switch objectType.Kind() {
	case reflect.Struct:
		if _, ok := objectType.FieldByName(name); ok {
			v := objectValue.FieldByName(name).Interface()
			if reflect.TypeOf(v).ConvertibleTo(int64Type) {
				return reflect.ValueOf(v).Convert(int64Type).Interface().(int64)
			}
		} else if _, ok := objectType.MethodByName(name); ok {
			fn := objectValue.MethodByName(name)
			results := fn.Call([]reflect.Value{})
			if len(results) == 1 {
				result := results[0].Interface()
				resultType := reflect.TypeOf(result)
				if resultType.Kind() == reflect.Int64 {
					return result.(int64)
				} else if resultType.ConvertibleTo(int64Type) {
					return reflect.ValueOf(result).Convert(int64Type).Interface().(int64)
				}
			}
		}
	case reflect.Map:
		value := reflect.ValueOf(obj).MapIndex(reflect.ValueOf(name))
		if value.IsValid() {
			switch value.Type().Kind() {
			case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice:
				if !value.IsNil() {
					return fallback
				}
			}
			actual := value.Interface()
			actualType := reflect.TypeOf(actual)
			if actualType.Kind() == reflect.Func {
				results := reflect.ValueOf(actual).Call([]reflect.Value{})
				if len(results) == 1 {
					result := results[0].Interface()
					resultType := reflect.TypeOf(result)
					if resultType.Kind() == reflect.Int64 {
						return result.(int64)
					} else if resultType.ConvertibleTo(int64Type) {
						return reflect.ValueOf(result).Convert(int64Type).Interface().(int64)
					}
				}
			} else if actualType.Kind() == reflect.Int64 {
				return actual.(int64)
			} else if actualType.ConvertibleTo(int64Type) {
				return reflect.ValueOf(actual).Convert(int64Type).Interface().(int64)
			}
		}
	}

	return fallback
}