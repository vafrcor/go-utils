package slices

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SampleStruct struct {
	Number int
	Name   string
}

func TestSliceShuffle(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("original: %v\n", nums)
	nums2 := SliceShuffle(nums)
	fmt.Printf("reshuffled: %v\n\n", nums2)
	assert.Len(t, nums, len(nums2))
	assert.ElementsMatch(t, nums, nums) // same items, any order

	words := []string{"alpha", "beta", "gamma"}
	fmt.Printf("original: %v\n", words)
	words2 := SliceShuffle(words)
	fmt.Printf("reshuffled: %v\n\n", words2)
	assert.Len(t, words, len(words2))
	assert.ElementsMatch(t, words, words2) // same items, any order

	users := []SampleStruct{
		{1, "Alice"},
		{2, "Bob"},
		{3, "Charlie"},
	}
	fmt.Printf("original: %v\n", users)
	users2 := SliceShuffle(users)
	fmt.Printf("reshuffled: %v\n\n", users2)
	assert.Len(t, users, len(users2))
	assert.ElementsMatch(t, users, users2) // same items, any order

	single := []int{1}
	fmt.Printf("original: %v\n", single)
	single2 := SliceShuffle(single)
	fmt.Printf("reshuffled: %v\n\n", single2)
	assert.Len(t, single, len(single2))
	assert.ElementsMatch(t, single, single2) // same items, any order

}

func TestSafeInt64ToUint64(t *testing.T) {
	a := []int64{1, 2, 3, -1, -2, -3}
	b := SafeInt64ToUint64(a)
	assert.Equal(t, []uint64{1, 2, 3}, b)
}

func TestStringWithSeperatorToStringSlice(t *testing.T) {
	assert.Equal(t, []string{"foo", "bar"}, StringWithSeperatorToStringSlice("foo, bar", ","))
}

func TestIsValidSliceOfStringValues(t *testing.T) {
	assert.NoError(t, IsValidSliceOfStringValues([]string{"foo", "bar"}, []string{"foo", "bar"}))
	assert.Equal(t, errors.New("invalid value (baz)"), IsValidSliceOfStringValues([]string{"foo", "bar"}, []string{"foo", "bar", "baz"}))
}

func TestConvertSliceIntoInterfaceSlice(t *testing.T) {
	slice := []string{"foo", "bar"}
	interfaceSlice := ConvertSliceIntoInterfaceSlice(slice)
	assert.Equal(t, []interface{}{"foo", "bar"}, interfaceSlice)
}

func TestTrimSpaceForSliceOfStrings(t *testing.T) {
	source := []string{" foo ", " bar "}
	trimmed := TrimSpaceForSliceOfStrings(source)
	assert.Equal(t, []string{"foo", "bar"}, trimmed)
}

func TestStringToUint64Slice(t *testing.T) {
	ids, _ := StringToUint64Slice("1,2,3")
	assert.Equal(t, []uint64{1, 2, 3}, ids)
}
