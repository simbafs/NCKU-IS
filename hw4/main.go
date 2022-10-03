package main

// reference: https://golangdocs.com/aes-encryption-decryption-in-golang

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
)

const (
	key    = "123456789"
	secret = "security"
	BS     = 16
)

func Zeropadding(s string) []byte {
	S := []byte(s)
	for len(S)%BS != 0 {
		S = append(S, '\000')
	}
	return S
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func EncryptAES(key []byte, plaintext []byte) string {
	// create cipher
	c, err := aes.NewCipher(key)
	CheckError(err)

	// allocate space for ciphered data
	out := make([]byte, len(plaintext))

	// encrypt
	c.Encrypt(out, []byte(plaintext))
	// return hex string
	return base64.StdEncoding.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) string {
	ciphertext, _ := base64.StdEncoding.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	return string(pt[:])
}

func main() {
	c := EncryptAES(Zeropadding(key), Zeropadding(secret))
	fmt.Println(c)
	p := DecryptAES(Zeropadding(key), c)
	fmt.Println(p)
}
