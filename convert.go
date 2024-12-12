// Copyright (c) 2024 Daniel Pickens
// This file is part of lorna which is released under MIT license.
// See file LICENSE for full license details.

package lorna

import (
	"reflect"
)

// GetRealValue returns real value of reflect.Value
// Required for JSON Marshalling
func GetRealValue(val *reflect.Value) interface{} {
	if val == nil {
		return nil
	}
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return val.Int()
	case reflect.String:
		return val.String()
	case reflect.Bool:
		return val.Bool()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return val.Uint()
	case reflect.Float32, reflect.Float64:
		return val.Float()
	case reflect.Slice, reflect.Map:
		return val.Interface()
	default:
		return nil
	}
}
