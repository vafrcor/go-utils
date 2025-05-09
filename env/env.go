package env

import (
	"encoding/json"
	"os"

	"github.com/spf13/cast"
)

func GetEnvValue(key string, dtype string, defaultValue any) any {
	value, ok := os.LookupEnv(key)
	useDefault := false
	if (!ok || value == "") && defaultValue != nil {
		useDefault = true
	}
	if useDefault {
		return defaultValue
	} else {
		return ConvertEnvValueToAnyDataType(dtype, value)
	}
}

func ConvertEnvValueToAnyDataType(dtype string, value string) any {
	switch dtype {
	case "string":
		return cast.ToString(value)
	case "int":
		return cast.ToInt(value)
	case "int32":
		return cast.ToInt32(value)
	case "int64":
		return cast.ToInt64(value)
	case "float64":
		return cast.ToFloat64(value)
	case "float32":
		return cast.ToFloat32(value)
	case "bool":
		return cast.ToBool(value)
	case "string_map":
		return cast.ToStringMap(value)
	case "string_map_string":
		return cast.ToStringMapString(value)
	case "string_map_int":
		return cast.ToStringMapInt(value)
	case "string_map_bool":
		return cast.ToStringMapBool(value)
	case "string_map_string_slice":
		return cast.ToStringMapStringSlice(value)
	case "slice":
		return cast.ToSlice(value)
	case "json_slice":
		v := []interface{}{}
		err := json.Unmarshal([]byte(value), &v)
		if err != nil {
			panic(err)
		}
		return v
	case "json_map_any":
		v := map[string]interface{}{}
		err := json.Unmarshal([]byte(value), &v)
		if err != nil {
			panic(err)
		}
		return v
	default:
		return value
	}
}
