package slices

import (
	"fmt"
	"math/rand"
	"slices"
	"strings"
	"time"
)

func SliceShuffle[T any](src []T) []T {
	// Copy the original slice
	dst := make([]T, len(src))
	copy(dst, src)

	// Shuffle the copy
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(dst), func(i, j int) {
		dst[i], dst[j] = dst[j], dst[i]
	})

	return dst
}

func SafeInt64ToUint64(input []int64) []uint64 {
	result := make([]uint64, 0, len(input))

	for _, v := range input {
		if v >= 0 {
			result = append(result, uint64(v))
		}
	}

	return result
}

func StringWithSeperatorToStringSlice(input string, separator string) []string {
	sl := strings.Split(input, separator)
	// Iterate and trim spaces
	for i, s := range sl {
		sl[i] = strings.TrimSpace(s)
	}
	return sl
}

func IsValidSliceOfStringValues(whitelist []string, input []string) error {
	// Check if all values in the input slice are present in the whitelist
	for _, value := range input {
		if !slices.Contains(whitelist, value) {
			return fmt.Errorf("invalid value (%s)", value)
		}
	}
	return nil
}

func ConvertSliceIntoInterfaceSlice[T any](items []T) []any {
	result := make([]any, len(items))

	for i, v := range items {
		result[i] = v
	}

	return result
}

func TrimSpaceForSliceOfStrings(values []string) []string {
	for i := range values {
		values[i] = strings.TrimSpace(values[i])
	}
	return values
}
