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
	X float64
	Y float64
	A float64
	B float64
}

func NewPoint(x, y, a, b float64) (*Point, error) {
	p := &Point{
		X: x,
		Y: y,
		A: a,
		B: b,
	}

	// 無限遠点
	if x == Inf && y == Inf {
		return p, nil
	}

	// 曲線上にあるか
	if math.Pow(y, 2) != math.Pow(x, 3)+a*x+b {
		return nil, errors.ValueErrorf("(%f, %f) is not on ther curve", x, y)
	}

	return p, nil
}

func (p *Point) String() string {
	if p.X == Inf {
		return fmt.Sprintf("Point(infinity)_%.2f_%.2f", p.A, p.B)
	}

	return fmt.Sprintf("Point(%.2f, %.2f)_%.2f_%.2f", p.X, p.Y, p.A, p.B)
}

func (p *Point) Eq(other *Point) bool {
	return true &&
		p.X == other.X &&
		p.Y == other.Y &&
		p.A == other.A &&
		p.B == other.B
}

func (p *Point) Ne(other *Point) bool {
	return !p.Eq(other)
}

func (p *Point) Add(other *Point) (*Point, error) {
	if p.A != other.A || p.B != other.B {
		return nil, errors.TypeErrorf("Points %s, %s are not on the same curve", p, other)
	}

	// x1 == infinity or x2 == infinity
	if p.X == Inf {
		return other, nil
	}

	if other.X == Inf {
		return p, nil
	}

	// x1 == x2 and y1 != y2
	if p.X == other.X && p.Y != other.Y {
		return NewPoint(Inf, Inf, p.A, p.Y)
	}

	// x1 != x2
	if p.X != other.X {
		s := (other.Y - p.Y) / (other.X - p.X)
		return p.add(other, s)
	}

	// p1 == p2 and y == 0
	if p.Eq(other) && p.Y == 0 {
		return NewPoint(Inf, Inf, p.A, p.Y)
	}

	// p1 == p2
	if p.Eq(other) {
		s := (3 * math.Pow(p.X, 2)) / 2 * p.Y
		return p.add(other, s)
	}

	return nil, errors.ValueError("Points isn't on the curve")
}

func (p *Point) add(other *Point, s float64) (*Point, error) {
	x := math.Pow(s, 2) - p.X - other.X
	y := s*(p.X-x) - p.Y
	return NewPoint(x, y, p.A, p.B)
}
