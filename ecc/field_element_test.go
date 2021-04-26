package ecc_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nasjp/btc/ecc"
)

func TestFieldElement_Eq(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		receiver *ecc.FieldElement
		other    *ecc.FieldElement
		want     bool
	}{
		{"Eq", &ecc.FieldElement{2, 31}, &ecc.FieldElement{2, 31}, true},
		{"Not Eq", &ecc.FieldElement{2, 31}, &ecc.FieldElement{15, 31}, false},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.receiver.Eq(&ecc.FieldElement{Num: tt.other.Num, Prime: tt.other.Prime}); got != tt.want {
				t.Errorf("FieldElement.Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldElement_Ne(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		receiver *ecc.FieldElement
		other    *ecc.FieldElement
		want     bool
	}{
		{"Ne", &ecc.FieldElement{2, 31}, &ecc.FieldElement{2, 31}, false},
		{"Not Ne(=Eq)", &ecc.FieldElement{2, 31}, &ecc.FieldElement{15, 31}, true},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.receiver.Ne(&ecc.FieldElement{Num: tt.other.Num, Prime: tt.other.Prime}); got != tt.want {
				t.Errorf("FieldElement.Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldElement_Add(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		receiver *ecc.FieldElement
		other    *ecc.FieldElement
		want     *ecc.FieldElement
	}{
		{"Add", &ecc.FieldElement{2, 31}, &ecc.FieldElement{15, 31}, &ecc.FieldElement{17, 31}},
		{"WrapAround", &ecc.FieldElement{17, 31}, &ecc.FieldElement{21, 31}, &ecc.FieldElement{7, 31}},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.receiver.Add(&ecc.FieldElement{Num: tt.other.Num, Prime: tt.other.Prime})
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Add() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFieldElement_Sub(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		receiver *ecc.FieldElement
		other    *ecc.FieldElement
		want     *ecc.FieldElement
	}{
		{"Sub", &ecc.FieldElement{29, 31}, &ecc.FieldElement{4, 31}, &ecc.FieldElement{25, 31}},
		{"WrapAround", &ecc.FieldElement{15, 31}, &ecc.FieldElement{30, 31}, &ecc.FieldElement{16, 31}},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.receiver.Sub(&ecc.FieldElement{Num: tt.other.Num, Prime: tt.other.Prime})
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Sub() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFieldElement_Mul(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		receiver *ecc.FieldElement
		other    *ecc.FieldElement
		want     *ecc.FieldElement
	}{
		{"Mul", &ecc.FieldElement{24, 31}, &ecc.FieldElement{19, 31}, &ecc.FieldElement{22, 31}},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.receiver.Mul(&ecc.FieldElement{Num: tt.other.Num, Prime: tt.other.Prime})
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Mul() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFieldElement_Pow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		receiver *ecc.FieldElement
		exponent int64
		want     *ecc.FieldElement
	}{
		{"Pow", &ecc.FieldElement{17, 31}, 3, &ecc.FieldElement{15, 31}},
		{"Big", &ecc.FieldElement{17, 31}, 24, &ecc.FieldElement{4, 31}},
		{"Minus", &ecc.FieldElement{17, 31}, -3, &ecc.FieldElement{29, 31}},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.receiver.Pow(tt.exponent)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Pow() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestFieldElement_Div(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		receiver *ecc.FieldElement
		other    *ecc.FieldElement
		want     *ecc.FieldElement
	}{
		{"Div", &ecc.FieldElement{3, 31}, &ecc.FieldElement{24, 31}, &ecc.FieldElement{4, 31}},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.receiver.Div(&ecc.FieldElement{Num: tt.other.Num, Prime: tt.other.Prime})
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Mul() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
