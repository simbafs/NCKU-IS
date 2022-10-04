package main

import (
	"fmt"
	"os"

	"github.com/schollz/progressbar/v3"
)

// reference: https://golangdocs.com/aes-encryption-decryption-in-golang

const (
	ct     = "Wj3RQTGXWvIeIu5nEt2qYuYbHRhoNtJawk07R0oZWnI="
	bs     = 16
	prefix = "Na"
	max    = 128
	fo     = os.O_RDWR | os.O_CREATE | os.O_TRUNC
)

func genKey(a, b, c, d rune) []rune {
	return []rune(fmt.Sprintf("Hj%cN)%ctgZ%c9wrc%cm", a, b, c, d))
}

func resolve(from, to rune, add func(int) error, ptf, keyf *os.File) {
	index := 0
	for a := from; a < to; a++ {
		for b := rune(0); b < max; b++ {
			for c := rune(0); c < max; c++ {
				for d := rune(0); d < max; d++ {
					add(1)
					key := genKey(a, b, c, d)
					// fmt.Println([]byte(string(key)))
					pt := DecryptAES([]byte(string(key)), ct)
					if pt[:2] == prefix {
						fmt.Fprintln(ptf, index, Unpadding(pt))
						fmt.Fprintln(keyf, index, key)
						index++
					}
				}
			}
		}
	}
}

func main() {
	bar := progressbar.Default(max * max * max * max)
	ptf, err := os.OpenFile("plaintext", fo, 0644)
	CheckError(err)
	keyf, err := os.OpenFile("key", fo, 0644)
	CheckError(err)

	resolve(rune(0), rune(max), bar.Add, ptf, keyf)
}
