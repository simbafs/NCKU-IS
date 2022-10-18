package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/schollz/progressbar/v3"
)

// reference:
// https://golangdocs.com/aes-encryption-decryption-in-golang
// https://blog.clarence.tw/2020/12/28/golang_implements_aes_ecb_and_pkcs7_pkgs5/

const (
	ct        = "Wj3RQTGXWvIeIu5nEt2qYuYbHRhoNtJawk07R0oZWnI=" // ciphertext
	bs        = 16                                             // blocksize
	prefix    = "Na"                                           // prefix of plaintext
	max       = 128                                            // max rune to try in key
	fo        = os.O_RDWR | os.O_CREATE | os.O_TRUNC           // file option be used what open 'plaintext', 'key' file
	printable = true                                           // if printable is true, plaintext will be print only if all of it is printable
)

// modified this
func genKey(a, b, c, d rune) []rune {
	return []rune(fmt.Sprintf("Hj%cN)%ctgZ%c9wrc%cm", a, b, c, d))
}

func resolve(wg *sync.WaitGroup, from, to rune, add func(int) error, ptf, keyf *os.File) {
	index := 0
	fmt.Printf("%d ~ %d\n", from, to)
	for a := from; a < to; a++ {
		for b := rune(0); b < max; b++ {
			for c := rune(0); c < max; c++ {
				for d := rune(0); d < max; d++ {
					add(1)
					key := genKey(a, b, c, d)
					pt := Unpadding(DecryptAES([]byte(string(key)), []byte(ct)))
					if pt[:2] == prefix && (!printable || AllIsPrint(pt)) {
						fmt.Fprintln(ptf, index, pt)
						fmt.Fprintln(keyf, index, key)
						index++
					}
				}
			}
		}
	}

	fmt.Printf("%d ~ %d done\n", from, to)
	wg.Done()
}

func main() {
	bar := progressbar.Default(max * max * max * max)
	ptf, err := os.OpenFile("plaintext", fo, 0644)
	CheckError(err)
	keyf, err := os.OpenFile("key", fo, 0644)
	CheckError(err)

	n := 4
	d := max / n
	wg := &sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n-1; i++ {
		go resolve(wg, rune(i*d), rune(i*d+d), bar.Add, ptf, keyf)
	}
	go resolve(wg, rune(max-d), rune(max), bar.Add, ptf, keyf)

	// goroutune := make([]byte, 100000)
	// k := runtime.Stack(goroutune, true)
	// fmt.Println(k, string(goroutune))

	wg.Wait()
}
