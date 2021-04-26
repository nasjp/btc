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

	p := func(t *testing.T, x, y int64) *ecc.Point { return mustPoint(t, x, y, 0, 7, 223) }

	tests := []struct {
		receiver *ecc.Point
		other    *ecc.Point
		want     *ecc.Point
	}{
		{p(t, 170, 142), p(t, 60, 139), p(t, 220, 181)},
		{p(t, 47, 71), p(t, 17, 56), p(t, 215, 68)},
		{p(t, 143, 98), p(t, 76, 66), p(t, 47, 71)},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(fmt.Sprintf("(%d,%d)+(%d,%d)", tt.receiver.X.Num, tt.receiver.Y.Num, tt.other.X.Num, tt.other.Y.Num), func(t *testing.T) {
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
