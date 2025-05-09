package variable

import "go.mongodb.org/mongo-driver/bson/primitive"

// PstringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func PstringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// Pint returns a pointer to the integer value passed in.
func Pint(v int) *int {
	return &v
}

// Pbool returns a pointer to the bool value passed in.
func Pbool(v bool) *bool {
	return &v
}

// PboolValue returns the value of the bool pointer passed in or
// false if the pointer is nil.
func PboolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// Pstring returns a pointer to the string value passed in.
func Pstring(v string) *string {
	return &v
}

// String returns a pointer to the mongodb primitive.DateTime value passed in.
func PprimitiveDate(v primitive.DateTime) *primitive.DateTime {
	return &v
}
