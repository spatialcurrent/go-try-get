// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package gtg

import (
	"fmt"
	"reflect"
)

// TryGetString returns the value from the object (a map or struct) by name.
//  - If the name refers to a struct method or func in a map, then call the function.
//    If the function only returns 1 value, then return the result as a string using `fmt.Sprint`,
//    otherwise return the value of the fallback parameter.
//  - If the name refers to a struct field or map key (and not a function),
//    then return the value by name as a string using `fmt.Sprint`.
//  - If the value by name does not exist, then returns the value of the fallback parameter.
//
// Examples
//
//  - TryGetString(map[string]int{"yo":"yo"}, "yo", "what") == "yo"
//  - TryGetString(map[string]string{"yo":"yo"}, "hey", "what") == "what"
//
func TryGetString(obj interface{}, name string, fallback string) string {

	objectType := reflect.TypeOf(obj)
	objectValue := reflect.ValueOf(obj)
	if objectType.Kind() == reflect.Ptr {
		objectType = objectType.Elem()
		objectValue = objectValue.Elem()
	}

	switch objectType.Kind() {
	case reflect.Struct:
		if _, ok := objectType.FieldByName(name); ok {
			return fmt.Sprint(objectValue.FieldByName(name).Interface())
		} else if _, ok := objectType.MethodByName(name); ok {
			fn := objectValue.MethodByName(name)
			results := fn.Call([]reflect.Value{})
			if len(results) == 1 {
				return fmt.Sprint(results[0].Interface())
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
					return fmt.Sprint(results[0].Interface())
				}
			} else {
				return fmt.Sprint(actual)
			}
		}
	}

	return fallback
}
