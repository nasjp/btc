package ecc_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nasjp/btc/ecc"
)

func TestPoint_Add(t *testing.T) {
	t.Parallel()

	const a, b = 5.0, 7.0

	tests := []struct {
		name     string
		receiver *ecc.Point
		other    *ecc.Point
		want     *ecc.Point
	}{
		{"Add", &ecc.Point{3, 7, a, b}, &ecc.Point{-1, -1, a, b}, &ecc.Point{2, -5, a, b}},
		{"Inf_Receiver", &ecc.Point{ecc.Inf, ecc.Inf, a, b}, &ecc.Point{2, 5, a, b}, &ecc.Point{2, 5, a, b}},
		{"Inf_Other", &ecc.Point{2, 5, a, b}, &ecc.Point{ecc.Inf, ecc.Inf, a, b}, &ecc.Point{2, 5, a, b}},
		{"Inf_Return", &ecc.Point{2, 5, a, b}, &ecc.Point{2, -5, a, b}, &ecc.Point{ecc.Inf, ecc.Inf, a, b}},
		{"SamePoint", &ecc.Point{-1, -1, a, b}, &ecc.Point{-1, -1, a, b}, &ecc.Point{18, 77, a, b}},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.receiver.Add(tt.other)

			if err != nil {
				t.Fatal(err)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Add() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
