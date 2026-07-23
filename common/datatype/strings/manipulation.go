package strings

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Customer String Replacer with dynamic placeholder and replacer
func StringDynamicReplacer(text string, replacers ...string) string {
	replacer := strings.NewReplacer(replacers...)
	return replacer.Replace(text)
}

func ConvertAnyToJSONString(data any) string {
	// Convert the struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("[ConvertAnyToJSONString] error: ", err)
		return ""
	}
	return string(jsonData)
}

func ToString(v any) string {
	switch val := v.(type) {

	case string:
		return val

	case bool:
		return strconv.FormatBool(val)

	case int:
		return strconv.Itoa(val)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)

	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)

	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)

	default:
		return fmt.Sprint(val) // fallback
	}
}

func StringJoinSlice[T any](items []T, separator string) string {
	if len(items) == 0 {
		return ""
	}

	parts := make([]string, len(items))
	for i, v := range items {
		parts[i] = fmt.Sprint(v)
	}

	return strings.Join(parts, separator)
}

func StringLeftPad(s string, length int, char string) string {
	for len(s) < length {
		s = char + s
	}
	return s
}
