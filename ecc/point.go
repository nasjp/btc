package ecc

import (
	"fmt"
	"math"

	"github.com/nasjp/btc/errors"
)

var (
	Inf = math.Inf(0)
)

type Point struct {
	X     *FieldElement
	Y     *FieldElement
	A     *FieldElement
	B     *FieldElement
	IsInf bool
}

func NewPoint(x, y, a, b *FieldElement) (*Point, error) {
	if !samePrimes(x, y, a, b) {
		return nil, errors.ValueError("prime are not same")
	}

	p := &Point{
		X: x,
		Y: y,
		A: a,
		B: b,
	}

	// x がnilのときに無限遠点とする
	if x == nil {
		p.IsInf = true
	}

	// 無限遠点は返却
	if p.IsInf {
		return p, nil
	}

	// 曲線上にあるか

	if !onCurve(x, y, a, b) {
		return nil, errors.ValueErrorf("(%s, %s) is not on ther curve", x, y)
	}

	return p, nil
}

func onCurve(x, y, a, b *FieldElement) bool {
	// y^2 == x^3 + a * x + b
	return y.pow(2).Eq(
		x.pow(3).add(a.mul(x)).add(b),
	)
}

func (p *Point) String() string {
	if p.IsInf {
		return fmt.Sprintf("Point(infinity)_%s_%s", p.A, p.B)
	}

	return fmt.Sprintf("Point(%s, %s)_%s_%s", p.X, p.Y, p.A, p.B)
}

func (p *Point) Eq(other *Point) bool {
	return true &&
		p.X.Eq(other.X) &&
		p.Y.Eq(other.Y) &&
		p.A.Eq(other.A) &&
		p.B.Eq(other.B)
}

func (p *Point) Ne(other *Point) bool {
	return !p.Eq(other)
}

func (p *Point) Add(other *Point) (*Point, error) {
	if p.A.Ne(other.A) || p.B.Ne(other.B) {
		return nil, errors.TypeErrorf("Points %s, %s are not on the same curve", p, other)
	}

	// x1 == infinity or x2 == infinity
	if p.IsInf {
		return other, nil
	}

	if other.IsInf {
		return p, nil
	}

	// x1 == x2 and y1 != y2
	if p.X.Eq(other.X) && p.Y.Ne(other.Y) {
		return NewPoint(nil, nil, p.A, p.B)
	}

	// x1 != x2
	if p.X.Ne(other.X) {
		// s = (other.Y - p.Y) / (other.X - p.X)
		s := other.Y.sub(p.Y).div(other.X.sub(p.X))
		return p.add(other, s)
	}

	// p1 == p2 and y == 0
	if p.Eq(other) && p.Y.Eq(newFieldElement(0, p.Y.Prime)) {
		return NewPoint(nil, nil, p.A, p.B)
	}

	// p1 == p2
	if p.Eq(other) {
		// s = (3*p.X^2 + p.A) / 2*p.Y
		s := p.X.pow(2).times(3).add(p.A).div(p.Y.times(2))
		return p.add(other, s)
	}

	return nil, errors.ValueError("Points isn't on the curve")
}

func (p *Point) add(other *Point, s *FieldElement) (*Point, error) {
	// x := s^2 - (p.X + other.X)
	x := s.pow(2).div(p.X).div(other.X)
	// y := s * (p.X-x) - p.Y
	y := s.mul(p.X.div(x)).div(p.Y)

	return NewPoint(x, y, p.A, p.B)
}
