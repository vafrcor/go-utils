package variable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValueFromMap(t *testing.T) {
	a := map[string]interface{}{
		"foo": "bar",
	}
	assert.Equal(t, "bar", GetValueFromMap(a, "foo", "unknown"))
}

func TestConvertAnyIntoMapStringOfInterfaceUsingJSON(t *testing.T) {
	convert1, err1 := ConvertAnyIntoMapStringOfInterfaceUsingJSON(map[string]string{"foo": "bar"})
	assert.NoError(t, err1)
	assert.Equal(t, map[string]interface{}{"foo": "bar"}, convert1)

	convert2, err2 := ConvertAnyIntoMapStringOfInterfaceUsingJSON(map[string]interface{}{"foo": 1, "bar": 2, "three": false})
	assert.NoError(t, err2)
	assert.Equal(t, map[string]interface{}{"foo": float64(1), "bar": float64(2), "three": false}, convert2)

	convert3, err3 := ConvertAnyIntoMapStringOfInterfaceUsingJSON(map[string]interface{}{
		"foo": 1,
		"bar": 2,
		"three": []map[string]interface{}{
			{"one": "two"},
		},
	})
	assert.NoError(t, err3)
	assert.Equal(t, map[string]interface{}{"bar": float64(2), "foo": float64(1), "three": []interface{}{map[string]interface{}{"one": "two"}}}, convert3)
}

func TestConvertAnyMapIntoInterfaceMap(t *testing.T) {
	a := map[string]interface{}{
		"foo": "bar",
		"boo": int(1),
	}
	b := ConvertAnyMapIntoInterfaceMap(a)
	assert.Equal(t, map[string]interface{}{"foo": "bar", "boo": int(1)}, b)

	c := map[string]string{
		"foo": "bar",
		"boo": "barr",
	}
	d := ConvertAnyMapIntoInterfaceMap(c)
	assert.Equal(t, map[string]interface{}{"foo": "bar", "boo": "barr"}, d)
}

func TestMergeMapSoi(t *testing.T) {
	testCases := []struct {
		Name     string
		M1       map[string]interface{}
		M2       map[string]interface{}
		Expected map[string]interface{}
	}{
		{
			Name: "simple_merge",
			M1: map[string]interface{}{
				"foo": "bar",
			},
			M2: map[string]interface{}{
				"bar": "foo",
			},
			Expected: map[string]interface{}{"foo": "bar", "bar": "foo"},
		},
		{
			Name: "overwrite_value",
			M1: map[string]interface{}{
				"foo": "bar",
			},
			M2: map[string]interface{}{
				"foo": "boo",
			},
			Expected: map[string]interface{}{"foo": "boo"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expected, MergeMapSoi(tc.M1, tc.M2))
		})
	}
}

func TestIsMapWithStringKey(t *testing.T) {
	a := map[string]interface{}{}
	b := map[string]string{}
	c := []string{}

	assert.True(t, IsMapWithStringKey(a))  // true
	assert.True(t, IsMapWithStringKey(b))  // true
	assert.False(t, IsMapWithStringKey(c)) // false
}

func TestMapSoiff64ToMapF64(t *testing.T) {
	tests := []struct {
		name string
		src  map[string]interface{}
		want map[string]float64
	}{
		{
			name: "all float64",
			src: map[string]interface{}{
				"transferred": 1.0,
				"success":     2.0,
				"failed":      7.0,
			},
			want: map[string]float64{
				"transferred": 1,
				"success":     2,
				"failed":      7,
			},
		},
		{
			name: "empty map",
			src:  map[string]interface{}{},
			want: map[string]float64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapSoiff64ToMapF64(tt.src)

			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("MapSoifToMapF64() = %v, want %v", got, tt.want)
			// }
			assert.Equal(t, tt.want, got)

		})
	}
}
