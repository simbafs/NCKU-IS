package main

import (
	"crypto/aes"
	"encoding/base64"
	"strings"
)

func Zeropadding(s string) []byte {
	S := []byte(s)
	for len(S)%bs != 0 {
		S = append(S, '\000')
	}
	return S
}

func Unpadding(s string) string {
	return strings.Trim(s, "\000")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

// unused
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
