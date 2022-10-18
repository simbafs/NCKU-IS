package main

import (
	"crypto/aes"
	"encoding/base64"
	"strings"
	"unicode"
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

func AllIsPrint(s string) bool {
	for _, c := range s {
		if !unicode.IsPrint(c) {
			return false
		}
	}
	return true
}

// unused
func EncryptAES(key, plaintext []byte) string {
	// create cipher
	cipher, err := aes.NewCipher(key)
	CheckError(err)

	// allocate space for ciphered data
	ciphertext := make([]byte, len(plaintext))

	// encrypt
	for start, end := 0, bs; start < len(plaintext); start, end = start+bs, end+bs {
		cipher.Encrypt(ciphertext[start:end], plaintext[start:end])
	}
	// return hex string
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func DecryptAES(key, ct []byte) string {
	ciphertext, err := base64.StdEncoding.DecodeString(string(ct))
	CheckError(err)

	c, err := aes.NewCipher(key)
	CheckError(err)

	plaintext := make([]byte, len(ciphertext))
	for start, end := 0, bs; start < len(ciphertext); start, end = start+bs, end+bs {
		c.Decrypt(plaintext[start:end], ciphertext[start:end])
	}

	return string(plaintext)
}
