package mongodb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Primitive A into []string
func PrimitiveAIntoSliceOfString(data primitive.A) []string {
	strSlice := make([]string, len(data))
	for i, v := range data {
		if str, ok := v.(string); ok {
			strSlice[i] = str
		} else {
			fmt.Println("Error: Non-string value found")
			continue
		}
	}
	return strSlice
}

// Primitive A into Interface{}
func PrimitiveAIntoInterface(data primitive.A) interface{} {
	return data
}
