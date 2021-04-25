package main_test

import (
	"testing"

	"github.com/nasjp/btc/ecc"
)

func TestPractice1(t *testing.T) {
	a, b := 5.0, 7.0
	t.Log(onCurve(2, 4, a, b))

	t.Log(onCurve(-1, -1, a, b))

	t.Log(onCurve(18, 77, a, b))

	t.Log(onCurve(5, 7, a, b))
}

func onCurve(x, y, a, b float64) bool {
	if _, err := ecc.NewPoint(x, y, a, b); err != nil {
		return false
	}
	return true
}

func TestPractice4(t *testing.T) {
	a, b := 5.0, 7.0
	p1, err := ecc.NewPoint(2, 5, a, b)
	if err != nil {
		t.Fatal(err)
	}

	p2, err := ecc.NewPoint(-1, -1, a, b)
	if err != nil {
		t.Fatal(err)
	}

	got, err := p1.Add(p2)
	if err != nil {
		t.Fatal(err)
	}

	want, err := ecc.NewPoint(3, -7, a, b)
	if err != nil {
		t.Fatal(err)
	}

	if !want.Eq(got) {
		t.Errorf("want: %s, got: %s", want, got)
	}
}
