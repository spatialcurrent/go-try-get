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

// TryGetString returns the value from the object (a map or struct) by name.
//  - If the name refers to a struct method or func in a map, then call the function.
//    If the function only returns 1 value, then return the result as a string using `fmt.Sprint`,
//    otherwise return the value of the fallback parameter.
//  - If the name refers to a struct field or map key (and not a function),
//    then return the value by name as a string using `fmt.Sprint`.
//  - If the value by name does not exist, then returns the value of the fallback parameter.
//
func TryGetString(obj interface{}, name string, fallback string) string {

	objectValue := reflect.ValueOf(obj)
	for reflect.TypeOf(objectValue.Interface()).Kind() == reflect.Ptr {
		objectValue = objectValue.Elem()
	}
	objectValue = reflect.ValueOf(objectValue.Interface()) // sets value to concerete type
	objectType := objectValue.Type()

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
					return fmt.Sprint(results[0].Interface())
				}
			} else {
				return fmt.Sprint(actual)
			}
		}
	}

	return fallback
}
