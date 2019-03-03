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

// TryGetStringArray returns the value from the object (a map or struct) by name.
//  - If the name refers to a struct method or func in a map, then call the function.
//    If the function only returns 1 value, then return the result as a string array using `fmt.Sprint`,
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

	objectType := reflect.TypeOf(obj)
	objectValue := reflect.ValueOf(obj)
	if objectType.Kind() == reflect.Ptr {
		objectType = objectType.Elem()
		objectValue = objectValue.Elem()
	}

	switch objectType.Kind() {
	case reflect.Struct:
		if _, ok := objectType.FieldByName(name); ok {
			arr := objectValue.FieldByName(name).Interface()
			if t := reflect.TypeOf(arr); t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
				values := reflect.ValueOf(arr)
				length := values.Len()
				strs := make([]string, 0, length)
				for i := 0; i < length; i++ {
					strs = append(strs, fmt.Sprint(values.Index(i).Interface()))
				}
				return strs
			}
		} else if _, ok := objectType.MethodByName(name); ok {
			fn := objectValue.MethodByName(name)
			results := fn.Call([]reflect.Value{})
			if len(results) == 1 {
				arr := results[0].Interface()
				if t := reflect.TypeOf(arr); t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
					values := reflect.ValueOf(arr)
					length := values.Len()
					strs := make([]string, 0, length)
					for i := 0; i < length; i++ {
						strs = append(strs, fmt.Sprint(values.Index(i).Interface()))
					}
					return strs
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
					arr := results[0].Interface()
					if t := reflect.TypeOf(arr); t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
						values := reflect.ValueOf(arr)
						length := values.Len()
						strs := make([]string, 0, length)
						for i := 0; i < length; i++ {
							strs = append(strs, fmt.Sprint(values.Index(i).Interface()))
						}
						return strs
					}
				}
			} else if actualType.Kind() == reflect.Array || actualType.Kind() == reflect.Slice {
				values := reflect.ValueOf(actual)
				length := values.Len()
				strs := make([]string, 0, length)
				for i := 0; i < length; i++ {
					strs = append(strs, fmt.Sprint(values.Index(i).Interface()))
				}
				return strs
			}
		}
	}

	return fallback
}
