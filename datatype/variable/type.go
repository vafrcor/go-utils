package variable

import (
	"reflect"
)

func CheckType(e interface{}) string {
	v := reflect.TypeOf(e)
	kind := v.Kind()
	ret := "unknown"
	if kind == reflect.String {
		ret = "string"
	} else if kind == reflect.Bool {
		ret = "bool"
	} else if kind == reflect.Int {
		ret = "int"
	} else if kind == reflect.Int32 {
		ret = "int32"
	} else if kind == reflect.Int64 {
		ret = "int64"
	} else if kind == reflect.Float32 {
		ret = "float32"
	} else if kind == reflect.Float64 {
		ret = "float64"
	} else if kind == reflect.Slice && v.Elem().Kind() == reflect.String {
		ret = "[]string"
	} else if kind == reflect.Slice && v.Elem().Kind() == reflect.Int {
		ret = "[]int"
	} else if kind == reflect.Slice && v.Elem().Kind() == reflect.Int32 {
		ret = "[]int32"
	} else if kind == reflect.Slice && v.Elem().Kind() == reflect.Int64 {
		ret = "[]int64"
	} else if kind == reflect.Slice && v.Elem().Kind() == reflect.Float32 {
		ret = "[]float32"
	} else if kind == reflect.Slice && v.Elem().Kind() == reflect.Float64 {
		ret = "[]float64"
	} else if kind == reflect.Slice && v.Elem().Kind() == reflect.Bool {
		ret = "[]bool"
	} else if kind == reflect.Slice && v.Elem().Kind() == reflect.Interface {
		ret = "[]interface{}"
	}
	return ret
}
