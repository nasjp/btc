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
		{"Eq", mustFE(t, 2, 31), mustFE(t, 2, 31), true},
		{"Not Eq", mustFE(t, 2, 31), mustFE(t, 15, 31), false},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.receiver.Eq(mustFE(t, tt.other.Num, tt.other.Prime)); got != tt.want {
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
		{"Ne", mustFE(t, 2, 31), mustFE(t, 2, 31), false},
		{"Not Ne(=Eq)", mustFE(t, 2, 31), mustFE(t, 15, 31), true},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.receiver.Ne(mustFE(t, tt.other.Num, tt.other.Prime)); got != tt.want {
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
		{"Add", mustFE(t, 2, 31), mustFE(t, 15, 31), mustFE(t, 17, 31)},
		{"WrapAround", mustFE(t, 17, 31), mustFE(t, 21, 31), mustFE(t, 7, 31)},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.receiver.Add(mustFE(t, tt.other.Num, tt.other.Prime))
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
		{"Sub", mustFE(t, 29, 31), mustFE(t, 4, 31), mustFE(t, 25, 31)},
		{"WrapAround", mustFE(t, 15, 31), mustFE(t, 30, 31), mustFE(t, 16, 31)},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.receiver.Sub(mustFE(t, tt.other.Num, tt.other.Prime))
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
		{"Mul", mustFE(t, 24, 31), mustFE(t, 19, 31), mustFE(t, 22, 31)},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.receiver.Mul(mustFE(t, tt.other.Num, tt.other.Prime))
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
		{"Pow", mustFE(t, 17, 31), 3, mustFE(t, 15, 31)},
		{"Big", mustFE(t, 17, 31), 24, mustFE(t, 4, 31)},
		{"Minus", mustFE(t, 17, 31), -3, mustFE(t, 29, 31)},
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
		{"Div", mustFE(t, 3, 31), mustFE(t, 24, 31), mustFE(t, 4, 31)},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.receiver.Div(mustFE(t, tt.other.Num, tt.other.Prime))
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Mul() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func mustFE(t *testing.T, num, prime int64) *ecc.FieldElement {
	fe, err := ecc.NewFieldElement(num, prime)
	if err != nil {
		t.Fatal(err)
	}

	return fe
}
