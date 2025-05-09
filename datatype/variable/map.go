package variable

// Get map element value by key
func GetValueFromMap(data map[string]interface{}, key string, defaultValue any) any {
	field, ok := data[key]
	if !ok {
		return defaultValue
	} else {
		return field
	}
}
