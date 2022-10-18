package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
)

// help funcitons

// det returns the determinant of the matrix
func det(a, b, c, d, e, f, g, h, i int64) int64 {
	return a*(e*i-f*h) - b*(d*i-f*g) + c*(d*h-e*g)
}

// toInt convert string to int64
func toInt(s string) int64 {
	i, _ := strconv.Atoi(s)
	return int64(i)
}

// toDec convert 62 base to decimal
func toDec(s string) int64 {
	n := int64(0)
	for k, v := range s {
		c := int64(v)
		switch {
		case v >= '0' && v <= '9':
			c -= '0'
		case v >= 'a' && v <= 'z':
			c -= 'a' - 10
		case v >= 'A' && v <= 'Z':
			c -= 'A' - 36
		default:
			panic("Invalid character")
		}
		n += int64(c) * int64(math.Pow(62, float64(len(s)-k-1)))
	}
	return n
}

// to62 convert decimal to 62 base
func to62(n int64) string {
	return big.NewInt(n).Text(62)
}

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

func (e *Equation) CalcText(x, y byte) string {
	return to62(e.A*int64(x) + e.B*int64(y) + e.C)
}

func (e1 *Equation) Solve(e2 *Equation, X, Y int64) (rune, rune) {
	d := det(e1.A, e1.B, e1.C,
		e2.A, e2.B, e2.C,
		0, 0, 1)
	if d == 0 {
		panic("No solution")
	}
	x := det(X, e1.B, e1.C,
		Y, e2.B, e2.C,
		1, 0, 1) / d
	y := det(e1.A, X, e1.C,
		e2.A, Y, e2.C,
		0, 1, 1) / d
	return rune(x), rune(y)
}

func Encrypt(e1, e2 *Equation, secret string) string {
	if len(secret)%2 == 1 {
		secret += " "
	}

	var encrypted string
	for i := 0; i < len(secret); i += 2 {
		x, y := secret[i], secret[i+1]
		X, Y := e1.CalcText(x, y), e2.CalcText(x, y)
		// fmt.Println(x, y, X, Y)
		lX, lY := to62(int64(len(X))), to62(int64(len(Y)))
		encrypted += lX + X + lY + Y
	}

	return encrypted
}

func Decrypt(e1, e2 *Equation, secret string) string {
	s := ""

	for i := 0; i < len(secret); {
		lX := toDec(string(secret[i]))
		i++
		X := toDec(secret[i : i+int(lX)])
		// fmt.Println("\t")
		i += int(lX)
		lY := toDec(string(secret[i]))
		i++
		Y := toDec(secret[i : i+int(lY)])
		i += int(lY)
		x, y := e1.Solve(e2, X, Y)
		// fmt.Println(lX, X, lY, Y, x, y)
		s += string(x) + string(y)
	}

	return s
}

func main() {
	// for i := 0; i < 500; i++ {
	//     fmt.Println(i, to62(int64(i)))
	// }
	fn := map[string]func(*Equation, *Equation, string) string{
		"encrypt": Encrypt,
		"decrypt": Decrypt,
	}

	if len(os.Args) < 9 {
		fmt.Println("Usage: main.go <encrypt|decrypt> <A1> <B1> <C1> <A2> <B2> <C2> <secret>")
		return
	}

	f, ok := fn[os.Args[1]]
	if !ok {
		fmt.Fprintf(os.Stderr, "Unknown function: %s", os.Args[1])
		os.Exit(1)
	}
	e1 := NewEquation(toInt(os.Args[2]), toInt(os.Args[3]), toInt(os.Args[4]))
	e2 := NewEquation(toInt(os.Args[5]), toInt(os.Args[6]), toInt(os.Args[7]))
	secret := os.Args[8]

	fmt.Println(f(e1, e2, secret))
	// // e1 := NewEquation(1, 2, 3)
	// // e2 := NewEquation(4, 5, 6)
	// // fmt.Println(e1.Solve(e2, 1121, 3797))
}
