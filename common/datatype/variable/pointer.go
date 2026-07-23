package variable

import "time"

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

// PUint64 returns a pointer to the Uint64 value passed in.
func PUint64(s uint64) *uint64 {
	return &s
}

// PTime returns a pointer to the Time value passed in.
func PTime(s time.Time) *time.Time {
	return &s
}
