package ecc

import (
	"fmt"

	"github.com/nasjp/btc/errors"
)

// FieldElementは単一の有限体要素
type FieldElement struct {
	Num   int64
	Prime int64
}

func NewFieldElement(num int64, prime int64) (*FieldElement, error) {
	if num >= prime || num < 0 {
		return nil, errors.ValueErrorf("Num %d not in field range 0 to %d", num, prime-1)
	}

	return &FieldElement{
		Num:   num,
		Prime: prime,
	}, nil
}

func (fe *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", fe.Prime, fe.Num)
}

func (fe *FieldElement) Eq(other *FieldElement) bool {
	if other == nil {
		return false
	}

	return fe.Num == other.Num && fe.Prime == other.Prime
}

func (fe *FieldElement) Ne(other *FieldElement) bool {
	return !fe.Eq(other)
}

func (fe *FieldElement) Add(other *FieldElement) (*FieldElement, error) {
	if fe.Prime != other.Prime {
		return nil, errors.TypeError("Can't add two numbers in different Fields'")
	}

	num := (fe.Num + other.Num) % fe.Prime
	if num < 0 {
		num += fe.Prime
	}

	return NewFieldElement(num, fe.Prime)
}

func (fe *FieldElement) Sub(other *FieldElement) (*FieldElement, error) {
	if fe.Prime != other.Prime {
		return nil, errors.TypeError("Can't sub two numbers in different Fields'")
	}

	num := (fe.Num - other.Num) % fe.Prime
	if num < 0 {
		num += fe.Prime
	}

	return NewFieldElement(num, fe.Prime)
}

func (fe *FieldElement) Mul(other *FieldElement) (*FieldElement, error) {
	if fe.Prime != other.Prime {
		return nil, errors.TypeError("Can't mul two numbers in different Fields'")
	}

	num := (fe.Num * other.Num) % fe.Prime
	if num < 0 {
		num += fe.Prime
	}

	return NewFieldElement(num, fe.Prime)
}

func (fe *FieldElement) Pow(exponent int64) (*FieldElement, error) {
	e := exponent % (fe.Prime - 1)
	if e < 0 {
		e += fe.Prime - 1
	}

	num := bigpow(fe.Num, e, fe.Prime)
	return NewFieldElement(num, fe.Prime)
}

func (fe *FieldElement) Div(other *FieldElement) (*FieldElement, error) {
	if fe.Prime != other.Prime {
		return nil, errors.TypeError("Can't div two numbers in different Fields'")
	}

	num := fe.Num * bigpow(other.Num, fe.Prime-2, fe.Prime) % fe.Prime
	return NewFieldElement(num, fe.Prime)
}

func bigpow(n int64, e int64, m int64) int64 {
	n %= m
	var res int64 = 1

	for e > 0 {
		if e&1 == 1 {
			res = res * n % m
		}

		n = n * n % m
		e >>= 1
	}

	return res
}
