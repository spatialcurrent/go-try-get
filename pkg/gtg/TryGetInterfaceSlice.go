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
				out := make([]interface{}, 0, length)
				for i := 0; i < length; i++ {
					out = append(out, values.Index(i).Interface())
				}
				return out
			}
		} else if _, ok := objectType.MethodByName(name); ok {
			fn := objectValue.MethodByName(name)
			results := fn.Call([]reflect.Value{})
			if len(results) == 1 {
				arr := results[0].Interface()
				if t := reflect.TypeOf(arr); t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
					values := reflect.ValueOf(arr)
					length := values.Len()
					out := make([]interface{}, 0, length)
					for i := 0; i < length; i++ {
						out = append(out, values.Index(i).Interface())
					}
					return out
				}
			}
		}
	case reflect.Map:
		value := reflect.ValueOf(obj).MapIndex(reflect.ValueOf(name))
		if value.IsValid() && !value.IsNil() {
			actual := value.Interface()
			if out, ok := actual.([]interface{}); ok {
				return out
			}
			actualType := reflect.TypeOf(actual)
			if actualType.Kind() == reflect.Func {
				results := reflect.ValueOf(actual).Call([]reflect.Value{})
				if len(results) == 1 {
					arr := results[0].Interface()
					if t := reflect.TypeOf(arr); t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
						values := reflect.ValueOf(arr)
						length := values.Len()
						out := make([]interface{}, 0, length)
						for i := 0; i < length; i++ {
							out = append(out, values.Index(i).Interface())
						}
						return out
					}
				}
			} else if actualType.Kind() == reflect.Array || actualType.Kind() == reflect.Slice {
				values := reflect.ValueOf(actual)
				length := values.Len()
				out := make([]interface{}, 0, length)
				for i := 0; i < length; i++ {
					out = append(out, values.Index(i).Interface())
				}
				return out
			}
		}
	}

	return fallback
}
