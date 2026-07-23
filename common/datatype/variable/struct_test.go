package variable

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestCopyStruct(t *testing.T) {
	type TestStruct1 struct {
		Foo string
		Bar int
	}

	type TestStruct2 struct {
		Foo string
		Bar int
	}

	type TestStruct3 struct {
		Foo      string
		Bar      int
		Metadata *structpb.Struct
	}

	type TestStruct4 struct {
		Foo      string
		Bar      int
		Metadata map[string]interface{} `json:"metadata"`
	}

	// case-1: simple struct
	a := TestStruct1{
		Foo: "foo",
		Bar: 1,
	}
	b := TestStruct2{}
	err := CopyStruct(&a, &b)
	assert.NoError(t, err)
	assert.Equal(t, TestStruct2{Foo: "foo", Bar: 1}, b)

	// case-2: struct with custom type
	c := TestStruct3{
		Foo: "foo",
		Bar: 1,
		Metadata: &structpb.Struct{
			Fields: map[string]*structpb.Value{
				"scoo": {Kind: &structpb.Value_StringValue{StringValue: "bido"}},
			},
		},
	}
	d := TestStruct4{}
	err2 := CopyStruct(&c, &d)
	assert.NoError(t, err2)
	assert.Equal(t, TestStruct4{Foo: "foo", Bar: 1, Metadata: nil}, d)

}

func TestStructToMap(t *testing.T) {
	type User struct {
		ID        int    `json:"id"`
		FirstName string `json:"first_name"`
		Email     string `json:"email"`
		IsActive  bool   `json:"is_active"`
	}
	user := User{
		ID:        1,
		FirstName: "John",
		Email:     "john@example.com",
		IsActive:  true,
	}

	m, err := StructToMap(user)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"id":         1.,
		"first_name": "John",
		"email":      "john@example.com",
		"is_active":  true,
	}, m)
}
