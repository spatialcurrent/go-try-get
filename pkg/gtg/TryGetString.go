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
	return tryGetStringValue(reflect.ValueOf(obj), name, fallback)
}

func tryGetStringValue(objectValue reflect.Value, name string, fallback string) string {

	if !objectValue.IsValid() {
		return fallback
	}

	if !objectValue.CanInterface() {
		return fallback
	}

	objectKind := objectValue.Kind()

	if objectKind == reflect.Ptr {
		return tryGetStringValue(objectValue.Elem(), name, fallback)
	}

	objectValue = reflect.ValueOf(objectValue.Interface()) // sets value to concerete type
	objectKind = objectValue.Kind()

	switch objectKind {
	case reflect.Struct:
		if field := objectValue.FieldByName(name); field.IsValid() {
			if fieldKind := field.Kind(); fieldKind == reflect.String {
				return field.Interface().(string)
			}
			return fmt.Sprint(field.Interface())
		}
		if method := objectValue.MethodByName(name); method.IsValid() {
			if results := method.Call([]reflect.Value{}); len(results) == 1 {
				result := results[0]
				if resultKind := result.Kind(); resultKind == reflect.String {
					return result.Interface().(string)
				}
				return fmt.Sprint(result.Interface())
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
		if valueKind == reflect.String {
			return valueValue.Interface().(string)
		}
		if valueKind == reflect.Func {
			results := valueValue.Call([]reflect.Value{})
			if len(results) == 1 {
				result := results[0]
				if result.IsValid() && result.CanInterface() {
					resultKind := result.Kind()
					if resultKind == reflect.String {
						return result.Interface().(string)
					}
					return fmt.Sprint(result.Interface())
				}
			}
			return fallback
		}
		return fmt.Sprint(valueValue.Interface())
	}

	return fallback
}
