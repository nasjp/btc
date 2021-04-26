package ecc_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nasjp/btc/ecc"
)

func TestPointOnCurve(t *testing.T) {
	t.Parallel()

	const (
		a     = 0
		b     = 7
		prime = 223
	)

	tests := []struct {
		x      int64
		y      int64
		hasErr bool
	}{
		{192, 105, false},
		{17, 56, false},
		{1, 193, false},
		{200, 119, true},
		{42, 99, true},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(fmt.Sprintf("(x=%d,y=%d)", tt.x, tt.y), func(t *testing.T) {
			t.Parallel()

			_, err := ecc.NewPoint(
				mustFE(t, tt.x, prime),
				mustFE(t, tt.y, prime),
				mustFE(t, a, prime),
				mustFE(t, b, prime),
			)

			switch tt.hasErr {
			case true:
				if err == nil {
					t.Error("want err, but empty")
				}
			case false:
				if err != nil {
					t.Errorf("unexpected err: %v", err)
				}
			}
		})
	}
}

func TestPoint_Add(t *testing.T) {
	t.Parallel()

	const a, b int64 = 5, 7

	tests := []struct {
		name     string
		receiver *ecc.Point
		other    *ecc.Point
		want     *ecc.Point
	}{
		/*
			{"Add", &ecc.Point{3, 7, a, b}, &ecc.Point{-1, -1, a, b}, &ecc.Point{2, -5, a, b}},
			{"Inf_Receiver", &ecc.Point{ecc.Inf, ecc.Inf, a, b}, &ecc.Point{2, 5, a, b}, &ecc.Point{2, 5, a, b}},
			{"Inf_Other", &ecc.Point{2, 5, a, b}, &ecc.Point{ecc.Inf, ecc.Inf, a, b}, &ecc.Point{2, 5, a, b}},
			{"Inf_Return", &ecc.Point{2, 5, a, b}, &ecc.Point{2, -5, a, b}, &ecc.Point{ecc.Inf, ecc.Inf, a, b}},
			{"SamePoint", &ecc.Point{-1, -1, a, b}, &ecc.Point{-1, -1, a, b}, &ecc.Point{18, 77, a, b}},
		*/

		// {"Add", mustPoint(t, 3, 7, a, b, 223), mustPoint(t, -1, -1, a, b, 223), mustPoint(t, 2, -5, a, b, 223)},
		// {"Inf_Receiver", mustPoint(t, 0, 0, a, b, 223, true), mustPoint(t, 2, 5, a, b, 223), mustPoint(t, 2, 5, a, b, 223)},
		// {"Inf_Other", mustPoint(t, 2, 5, a, b, 223), mustPoint(t, 0, 0, a, b, 223, true), mustPoint(t, 2, 5, a, b, 223)},
		// {"Inf_Return", mustPoint(t, 2, 5, a, b, 223), mustPoint(t, 2, -5, a, b, 223), mustPoint(t, 0, 0, a, b, 223, true)},
		// {"SamePoint", mustPoint(t, -1, -1, a, b, 223), mustPoint(t, -1, -1, a, b, 223), mustPoint(t, 18, 77, a, b, 223)},
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

func mustPoint(t *testing.T, x, y, a, b, prime int64, inf ...bool) *ecc.Point {
	var ex, ey, ea, eb *ecc.FieldElement

	if !(len(inf) != 0 && inf[0]) {
		ex = mustFE(t, x, prime)
		ey = mustFE(t, y, prime)
	}

	ea = mustFE(t, a, prime)
	eb = mustFE(t, b, prime)

	fe, err := ecc.NewPoint(ex, ey, ea, eb)
	if err != nil {
		t.Fatal(err)
	}

	return fe
}
