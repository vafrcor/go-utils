package variable

import "time"

// PStringValue returns the value of the string pointer passed in or
// "" if the pointer is nil.
func PStringValue(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// PInt returns a pointer to the integer value passed in.
func PInt(v int) *int {
	return &v
}

// PBool returns a pointer to the bool value passed in.
func PBool(v bool) *bool {
	return &v
}

// PboolValue returns the value of the bool pointer passed in or
// false if the pointer is nil.
func PBoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// PString returns a pointer to the string value passed in.
func PString(v string) *string {
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
