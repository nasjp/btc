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

	return newFieldElement(num, prime), nil
}

func newFieldElement(num int64, prime int64) *FieldElement {
	return &FieldElement{
		Num:   num,
		Prime: prime,
	}
}

func (fe *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", fe.Prime, fe.Num)
}

func (fe *FieldElement) Eq(other *FieldElement) bool {
	return fe.Num == other.Num && fe.Prime == other.Prime
}

func (fe *FieldElement) Ne(other *FieldElement) bool {
	return !fe.Eq(other)
}

func (fe *FieldElement) Add(other *FieldElement) (*FieldElement, error) {
	if err := fe.validate(other); err != nil {
		return nil, err
	}

	return fe.add(other), nil
}

func (fe *FieldElement) Sub(other *FieldElement) (*FieldElement, error) {
	if err := fe.validate(other); err != nil {
		return nil, err
	}

	return fe.sub(other), nil
}

func (fe *FieldElement) Mul(other *FieldElement) (*FieldElement, error) {
	if err := fe.validate(other); err != nil {
		return nil, err
	}

	return fe.mul(other), nil
}

func (fe *FieldElement) Pow(exponent int64) (*FieldElement, error) {
	return fe.pow(exponent), nil
}

func (fe *FieldElement) Div(other *FieldElement) (*FieldElement, error) {
	if fe.Prime != other.Prime {
		return nil, errors.TypeError("Can't div two numbers in different Fields'")
	}

	return fe.div(other), nil
}

func (fe *FieldElement) add(other *FieldElement) *FieldElement {
	num := (fe.Num + other.Num) % fe.Prime
	if num < 0 {
		num += fe.Prime
	}

	return newFieldElement(num, fe.Prime)
}

func (fe *FieldElement) sub(other *FieldElement) *FieldElement {
	num := (fe.Num - other.Num) % fe.Prime
	if num < 0 {
		num += fe.Prime
	}

	return newFieldElement(num, fe.Prime)
}

func (fe *FieldElement) mul(other *FieldElement) *FieldElement {
	num := (fe.Num * other.Num) % fe.Prime
	if num < 0 {
		num += fe.Prime
	}

	return newFieldElement(num, fe.Prime)
}

func (fe *FieldElement) pow(exponent int64) *FieldElement {
	e := exponent % (fe.Prime - 1)
	if e < 0 {
		e += fe.Prime - 1
	}

	num := bigpow(fe.Num, e, fe.Prime)
	return newFieldElement(num, fe.Prime)
}

func (fe *FieldElement) div(other *FieldElement) *FieldElement {
	num := fe.Num * bigpow(other.Num, fe.Prime-2, fe.Prime) % fe.Prime
	return newFieldElement(num, fe.Prime)
}

func (fe *FieldElement) validate(other *FieldElement) error {
	if fe.Prime != other.Prime {
		return errors.TypeErrorf("Can't use two numbers in different Fields'")
	}

	return nil
}

func (fe *FieldElement) times(num int64) *FieldElement {
	t := fe
	for i := int64(2); i <= num; i++ {
		t = t.add(fe)
	}

	return t
}

func samePrimes(fes ...*FieldElement) bool {
	if len(fes) == 0 {
		return true
	}

	prime := fes[0].Prime

	for _, fe := range fes {
		if prime != fe.Prime {
			return false
		}
	}

	return true
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
