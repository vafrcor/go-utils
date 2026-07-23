package variable

// convert pointer Uint into pointer Uint64
func VariablePuintIntoPuint64(i *uint) *uint64 {
	var x *uint64 = nil
	if i != nil {
		t := uint64(*i)
		x = &t
	}
	return x
}

// convert pointer Uint64 into pointer Uint
func VariablePuint64IntoPuint(i *uint64) *uint {
	var x *uint = nil
	if i != nil {
		t := uint(*i)
		x = &t
	}
	return x
}

// Convert `[]interface{}` into `[]string`
func ConvertSoifIntoSos(input []interface{}) []string {
	strSlice := make([]string, len(input))
	for i, v := range input {
		strSlice[i] = v.(string)
	}
	return strSlice
}

// Convert `[]interface{}` into `[]int`
func ConvertSoifIntoSoi(input []interface{}) []int {
	intSlice := make([]int, len(input))
	for i, v := range input {
		intSlice[i] = v.(int)
	}
	return intSlice
}

// Convert `[]interface{}` into `[]int64`
func ConvertSoifIntoSoi64(input []interface{}) []int64 {
	intSlice := make([]int64, len(input))
	for i, v := range input {
		intSlice[i] = v.(int64)
	}
	return intSlice
}

// Convert `[]interface` into `[]float64`
func ConvertSoifIntoSof64(input []interface{}) []float64 {
	float64Slice := make([]float64, len(input))
	for i, v := range input {
		float64Slice[i] = v.(float64)
	}
	return float64Slice
}

// Convert `[]string` into `[]interface{}`
func ConvertSosIntoSoif(input []string) []interface{} {
	ifSlice := make([]interface{}, len(input))
	for i, v := range input {
		ifSlice[i] = v
	}
	return ifSlice
}

// Convert `[]int64` into `[]interface{}`
func ConvertSoi64IntoSoif(input []int64) []interface{} {
	ifSlice := make([]interface{}, len(input))
	for i, v := range input {
		ifSlice[i] = v
	}
	return ifSlice
}

// Convert `[]*struct` into `[]interface{}`
func ConvertSogtIntoSoif[V interface{}](input []*V) []interface{} {
	ifSlice := make([]interface{}, len(input))
	for i, v := range input {
		ifSlice[i] = v
	}
	return ifSlice
}

// Convert `map[string]string` into `map[string]interface{}`
func ConvertMapSosIntoMapSoI(original map[string]string) map[string]interface{} {
	converted := make(map[string]interface{})
	for key, value := range original {
		converted[key] = value
	}
	return converted
}
