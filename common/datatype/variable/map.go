package variable

import (
	"encoding/json"
	"reflect"
)

// Get map element value by key
func GetValueFromMap(data map[string]interface{}, key string, defaultValue any) any {
	field, ok := data[key]
	if !ok {
		return defaultValue
	} else {
		return field
	}
}

func ConvertAnyIntoMapStringOfInterfaceUsingJSON(input any) (map[string]interface{}, error) {
	var output map[string]interface{}
	inputJSON, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(inputJSON, &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func ConvertAnyMapIntoInterfaceMap[T any](m map[string]T) map[string]interface{} {
	result := make(map[string]interface{}, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

func MergeMapSoi(m1, m2 map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	// copy first map
	for k, v := range m1 {
		result[k] = v
	}

	// overwrite / add from second map
	for k, v := range m2 {
		result[k] = v
	}

	return result
}

func IsMapWithStringKey(v interface{}) bool {
	t := reflect.TypeOf(v)

	if t == nil {
		return false
	}

	return t.Kind() == reflect.Map &&
		t.Key().Kind() == reflect.String
}

func MapSoiff64ToMapF64(src map[string]interface{}) map[string]float64 {
	dst := make(map[string]float64, len(src))

	for k, v := range src {
		if f, ok := v.(float64); ok {
			dst[k] = f
		}
	}

	return dst
}
