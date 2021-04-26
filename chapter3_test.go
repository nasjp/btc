package main_test

import (
	"fmt"
	"testing"

	"github.com/nasjp/btc/ecc"
)

func TestPractice3_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		x, y int64
		want bool
	}{
		{192, 105, true},
		{17, 56, true},
		{200, 119, false},
		{1, 193, true},
		{42, 99, false},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(fmt.Sprintf("(%d,%d)", tt.x, tt.y), func(t *testing.T) {
			t.Parallel()

			if got := onCurveChapter4(t, tt.x, tt.y); tt.want != got {
				t.Errorf("(%d, %d) on curve? => %t, but want %t", tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func onCurveChapter4(t *testing.T, x, y int64) bool {
	const prime = 223
	a, err := ecc.NewFieldElement(0, prime)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ecc.NewFieldElement(7, prime)
	if err != nil {
		t.Fatal(err)
	}

	xx, err := ecc.NewFieldElement(x, prime)
	if err != nil {
		t.Fatal(err)
	}

	yy, err := ecc.NewFieldElement(y, prime)
	if err != nil {
		t.Fatal(err)
	}

	left, err := yy.Pow(2)
	if err != nil {
		t.Fatal(err)
	}

	right1, err := xx.Pow(3)
	if err != nil {
		t.Fatal(err)
	}

	right2, err := a.Mul(xx)
	if err != nil {
		t.Fatal(err)
	}

	right1and2, err := right1.Add(right2)
	if err != nil {
		t.Fatal(err)
	}

	right, err := right1and2.Add(b)
	if err != nil {
		t.Fatal(err)
	}

	return left.Eq(right)
}
