package main_test

import (
	"testing"

	"github.com/nasjp/btc/ecc"
)

func TestPractice1_3(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		a, err := ecc.NewFieldElement(44, 57)
		if err != nil {
			t.Fatal(err)
		}

		b, err := ecc.NewFieldElement(33, 57)
		if err != nil {
			t.Fatal(err)
		}

		c, err := a.Add(b)
		if err != nil {
			t.Fatal(err)
		}

		if want, got := int64(20), c.Num; want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})

	t.Run("2", func(t *testing.T) {
		a, err := ecc.NewFieldElement(9, 57)
		if err != nil {
			t.Fatal(err)
		}

		b, err := ecc.NewFieldElement(29, 57)
		if err != nil {
			t.Fatal(err)
		}

		c, err := a.Sub(b)
		if err != nil {
			t.Fatal(err)
		}

		if want, got := int64(37), c.Num; want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})

	t.Run("3", func(t *testing.T) {
		a, err := ecc.NewFieldElement(17, 57)
		if err != nil {
			t.Fatal(err)
		}

		b, err := ecc.NewFieldElement(42, 57)
		if err != nil {
			t.Fatal(err)
		}

		c, err := ecc.NewFieldElement(49, 57)
		if err != nil {
			t.Fatal(err)
		}

		d, err := a.Add(b)
		if err != nil {
			t.Fatal(err)
		}

		e, err := d.Add(c)
		if err != nil {
			t.Fatal(err)
		}

		if want, got := int64(51), e.Num; want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})

	t.Run("4", func(t *testing.T) {
		a, err := ecc.NewFieldElement(52, 57)
		if err != nil {
			t.Fatal(err)
		}

		b, err := ecc.NewFieldElement(30, 57)
		if err != nil {
			t.Fatal(err)
		}

		c, err := ecc.NewFieldElement(38, 57)
		if err != nil {
			t.Fatal(err)
		}

		d, err := a.Sub(b)
		if err != nil {
			t.Fatal(err)
		}

		e, err := d.Sub(c)
		if err != nil {
			t.Fatal(err)
		}

		if want, got := int64(41), e.Num; want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})
}

func TestPractice7(t *testing.T) {
	for _, p := range []int64{7, 11, 17, 31} {
		nums := make([]int64, 0, p-1)
		for i := int64(1); i < p; i++ {
			f, err := ecc.NewFieldElement(i, p)
			if err != nil {
				t.Fatal(err)
			}

			f, err = f.Pow(p - 1)
			if err != nil {
				t.Fatal(err)
			}
			nums = append(nums, f.Num)
		}
		for _, n := range nums {
			if want, got := int64(1), n; want != got {
				t.Errorf("want: %d, got: %d", want, got)
			}
		}
	}
}

func TestPractice8(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		a, err := ecc.NewFieldElement(3, 31)
		if err != nil {
			t.Fatal(err)
		}

		b, err := ecc.NewFieldElement(24, 31)
		if err != nil {
			t.Fatal(err)
		}

		c, err := a.Div(b)
		if err != nil {
			t.Fatal(err)
		}

		if want, got := int64(4), c.Num; want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})

	t.Run("2", func(t *testing.T) {
		a, err := ecc.NewFieldElement(17, 31)
		if err != nil {
			t.Fatal(err)
		}

		b, err := a.Pow(31 - 3 - 1)
		if err != nil {
			t.Fatal(err)
		}

		if want, got := int64(29), b.Num; want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})

	t.Run("3", func(t *testing.T) {
		a, err := ecc.NewFieldElement(4, 31)
		if err != nil {
			t.Fatal(err)
		}

		b, err := a.Pow(31 - 4 - 1)
		if err != nil {
			t.Fatal(err)
		}

		c, err := ecc.NewFieldElement(11, 31)
		if err != nil {
			t.Fatal(err)
		}

		d, err := b.Mul(c)
		if err != nil {
			t.Fatal(err)
		}

		if want, got := int64(13), d.Num; want != got {
			t.Errorf("want: %d, got: %d", want, got)
		}
	})
}
