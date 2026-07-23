package mongodb

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPrimitiveAIntoSliceOfString(t *testing.T) {
	tests := []struct {
		name string
		data primitive.A
		want []string
	}{
		{
			name: "all strings",
			data: primitive.A{"alpha", "beta", "gamma"},
			want: []string{"alpha", "beta", "gamma"},
		},
		{
			name: "mixed values preserve positions",
			data: primitive.A{"alpha", 42, true, nil, "omega"},
			want: []string{"alpha", "", "", "", "omega"},
		},
		{
			name: "empty",
			data: primitive.A{},
			want: []string{},
		},
		{
			name: "nil",
			data: nil,
			want: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrimitiveAIntoSliceOfString(tt.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrimitiveAIntoSliceOfString() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestPrimitiveAIntoInterface(t *testing.T) {
	t.Run("returns primitive array", func(t *testing.T) {
		data := primitive.A{"alpha", 42, true}

		got := PrimitiveAIntoInterface(data)
		gotData, ok := got.(primitive.A)
		if !ok {
			t.Fatalf("PrimitiveAIntoInterface() type = %T, want primitive.A", got)
		}
		if !reflect.DeepEqual(gotData, data) {
			t.Errorf("PrimitiveAIntoInterface() = %#v, want %#v", gotData, data)
		}
	})

	t.Run("preserves nil", func(t *testing.T) {
		var data primitive.A

		got := PrimitiveAIntoInterface(data)
		gotData, ok := got.(primitive.A)
		if !ok {
			t.Fatalf("PrimitiveAIntoInterface() type = %T, want primitive.A", got)
		}
		if gotData != nil {
			t.Errorf("PrimitiveAIntoInterface() = %#v, want nil primitive.A", gotData)
		}
	})
}

func TestPprimitiveDate(t *testing.T) {
	want := primitive.DateTime(1_721_862_400_000)

	got := PprimitiveDate(want)
	if got == nil {
		t.Fatal("PprimitiveDate() returned nil")
	}
	if *got != want {
		t.Errorf("*PprimitiveDate() = %v, want %v", *got, want)
	}

	*got++
	if want != primitive.DateTime(1_721_862_400_000) {
		t.Errorf("PprimitiveDate() pointer aliases input: input changed to %v", want)
	}
}
