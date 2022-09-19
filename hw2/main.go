package main

import (
	"fmt"
	"math/big"
)

type Equation struct {
	A int64
	B int64
	C int64
}

func NewEquation(A, B, C int64) *Equation {
	return &Equation{
		A: A,
		B: B,
		C: C,
	}
}

func (e *Equation) Calc(x, y int64) int64 {
	return e.A*x + e.B*y + e.C
}

func (e *Equation) CalcText(x, y int64) string {
	return big.NewInt(e.Calc(x, y)).Text(62)
}

func main() {
	var a, b, c, d, e, f int64
	var secret, encrypted string
	fmt.Print("key(six int number): ")
	fmt.Scanf("%d %d %d %d %d %d", &a, &b, &c, &d, &e, &f)
	e1 := NewEquation(a, b, c)
	e2 := NewEquation(d, e, f)
	fmt.Print("secret: ")
	fmt.Scanln(&secret)

	if len(secret)%2 == 1 {
		secret += " "
	}

	for i := 0; i < len(secret); i += 2 {
		x, y := int64(secret[i]), int64(secret[i+1])
		encrypted += e1.CalcText(x, y) + " " + e2.CalcText(x, y) + " "
	}

	fmt.Println(encrypted)
}
