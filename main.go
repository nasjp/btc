package main

import (
	"fmt"
	"os"

	"github.com/nasjp/btc/ecc"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if err := practice2_1(); err != nil {
		return err
	}

	if err := practice2_2(); err != nil {
		return err
	}

	if err := practice2_3(); err != nil {
		return err
	}

	if err := practice2_4(); err != nil {
		return err
	}

	return nil
}

func practice2_1() error {
	a, err := ecc.NewFieldElement(44, 57)
	if err != nil {
		return err
	}

	b, err := ecc.NewFieldElement(33, 57)
	if err != nil {
		return err
	}

	c, err := a.Add(b)
	if err != nil {
		return err
	}

	fmt.Println(c)

	return nil
}

func practice2_2() error {
	a, err := ecc.NewFieldElement(9, 57)
	if err != nil {
		return err
	}

	b, err := ecc.NewFieldElement(29, 57)
	if err != nil {
		return err
	}

	c, err := a.Sub(b)
	if err != nil {
		return err
	}

	fmt.Println(c)

	return nil
}

func practice2_3() error {
	a, err := ecc.NewFieldElement(17, 57)
	if err != nil {
		return err
	}

	b, err := ecc.NewFieldElement(42, 57)
	if err != nil {
		return err
	}

	c, err := ecc.NewFieldElement(49, 57)
	if err != nil {
		return err
	}

	d, err := a.Add(b)
	if err != nil {
		return err
	}

	e, err := d.Add(c)
	if err != nil {
		return err
	}

	fmt.Println(e)

	return nil
}

func practice2_4() error {
	a, err := ecc.NewFieldElement(52, 57)
	if err != nil {
		return err
	}

	b, err := ecc.NewFieldElement(30, 57)
	if err != nil {
		return err
	}

	c, err := ecc.NewFieldElement(38, 57)
	if err != nil {
		return err
	}

	d, err := a.Sub(b)
	if err != nil {
		return err
	}

	e, err := d.Sub(c)
	if err != nil {
		return err
	}

	fmt.Println(e)

	return nil
}
