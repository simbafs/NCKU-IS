package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"

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

func resolve(wg *sync.WaitGroup, from, to rune, add func(int) error, index *int64, ptf, keyf *os.File) {
	for a := from; a < to; a++ {
		for b := rune(0); b < max; b++ {
			for c := rune(0); c < max; c++ {
				for d := rune(0); d < max; d++ {
					add(1)
					key := genKey(a, b, c, d)
					// fmt.Println([]byte(string(key)))
					pt := DecryptAES([]byte(string(key)), ct)
					if pt[:2] == prefix {
						fmt.Fprintln(ptf, *index, Unpadding(pt))
						fmt.Fprintln(keyf, *index, key)
						atomic.AddInt64(index, 1)
					}
				}
			}
		}
	}
	wg.Done()
}

func main() {
	bar := progressbar.Default(max * max * max * max)
	ptf, err := os.OpenFile("plaintext", fo, 0644)
	CheckError(err)
	keyf, err := os.OpenFile("key", fo, 0644)
	CheckError(err)

	index := int64(0)

	resolve(nil, rune(0), rune(max), bar.Add, &index, ptf, keyf)
	// wg := &sync.WaitGroup{}
	// n := 1
	// d := max / n
	// wg.Add(n)
	// for i := 0; i < n; i++ {
	//     if i == n-1 {
	//         // fmt.Println(d*i, max)
	//         go resolve(nil, rune(d*i), rune(max), bar.Add, &index, ptf, keyf)
	//     } else {
	//         // fmt.Println(d*i, d*i+d)
	//         go resolve(nil, rune(d*i), rune(d*i+d), bar.Add, &index, ptf, keyf)
	//     }
	// }
	// fmt.Println(N)
	// wg.Wait()
}
