package jsonhelper

import (
	"encoding/json"
	"reflect"
)

// MarshalAll marshals the given value to JSON, ignoring omitempty for all fields.
func MarshalAll(v any) ([]byte, error) {
	result := structToMap(v)
	return json.Marshal(result)
}

// structToMap converts a struct to a map recursively
func structToMap(v interface{}) interface{} {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil
		}
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		m := make(map[string]interface{})
		typ := val.Type()
		for i := 0; i < val.NumField(); i++ {
			field := typ.Field(i)
			if field.IsExported() {
				m[field.Name] = structToMap(val.Field(i).Interface())
			}
		}
		return m
	case reflect.Map:
		m := make(map[string]interface{})
		for _, k := range val.MapKeys() {
			m[k.String()] = structToMap(val.MapIndex(k).Interface())
		}
		return m
	case reflect.Slice, reflect.Array:
		s := make([]interface{}, val.Len())
		for i := 0; i < val.Len(); i++ {
			s[i] = structToMap(val.Index(i).Interface())
		}
		return s
	default:
		return v
	}
}
