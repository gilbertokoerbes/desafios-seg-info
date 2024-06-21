package main

import (
	"fmt"
	"math/big"
)

type Point struct {
	X, Y *big.Int
}

// Curve parameters for y^2 = x^3 + 2x + 2 over a large prime field
var (
	A    = big.NewInt(2)
	B    = big.NewInt(2)
	P, _ = new(big.Int).SetString("2f5b6a0a4b7afbbe97d1f009d17424ed", 16) // example prime, replace with the actual prime for your curve
)

func addPoints(p1, p2 *Point) *Point {
	if p1 == nil {
		return p2
	}
	if p2 == nil {
		return p1
	}

	// λ = (y2 - y1) / (x2 - x1)
	lambda := new(big.Int).Sub(p2.Y, p1.Y)
	denom := new(big.Int).Sub(p2.X, p1.X)
	denom.ModInverse(denom, P)
	lambda.Mul(lambda, denom)
	lambda.Mod(lambda, P)

	// x3 = λ^2 - x1 - x2
	x3 := new(big.Int).Mul(lambda, lambda)
	x3.Sub(x3, p1.X)
	x3.Sub(x3, p2.X)
	x3.Mod(x3, P)

	// y3 = λ(x1 - x3) - y1
	y3 := new(big.Int).Sub(p1.X, x3)
	y3.Mul(lambda, y3)
	y3.Sub(y3, p1.Y)
	y3.Mod(y3, P)

	return &Point{X: x3, Y: y3}
}

func doublePoint(p *Point) *Point {
	// λ = (3x1^2 + a) / 2y1
	lambda := new(big.Int).Mul(p.X, p.X)
	lambda.Mul(lambda, big.NewInt(3))
	lambda.Add(lambda, A)
	denom := new(big.Int).Mul(p.Y, big.NewInt(2))
	denom.ModInverse(denom, P)
	lambda.Mul(lambda, denom)
	lambda.Mod(lambda, P)

	// x3 = λ^2 - 2x1
	x3 := new(big.Int).Mul(lambda, lambda)
	x3.Sub(x3, new(big.Int).Mul(p.X, big.NewInt(2)))
	x3.Mod(x3, P)

	// y3 = λ(x1 - x3) - y1
	y3 := new(big.Int).Sub(p.X, x3)
	y3.Mul(lambda, y3)
	y3.Sub(y3, p.Y)
	y3.Mod(y3, P)

	return &Point{X: x3, Y: y3}
}

func scalarMult(k *big.Int, p *Point) *Point {
	result := (*Point)(nil)
	addend := p

	for i := k.BitLen() - 1; i >= 0; i-- {
		if result != nil {
			result = doublePoint(result)
		}
		if k.Bit(i) == 1 {
			result = addPoints(result, addend)
		}
	}

	return result
}

func main() {
	// Example values for a and B (x, y)
	a, _ := new(big.Int).SetString("27e41b3246bec9b16e398115", 16)
	x, _ := new(big.Int).SetString("e81e9c7f9e80ef324834057e46d5776", 16)
	y, _ := new(big.Int).SetString("1ecfcbf7b267aca76fcc871c1b109954", 16)

	B := &Point{X: x, Y: y}
	V := scalarMult(a, B)

	fmt.Printf("V = (%s, %s)\n", V.X.Text(16), V.Y.Text(16))
}
