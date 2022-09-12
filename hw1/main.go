package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var spaceMap = []rune{
	'\u0020',
	'\u00a0',
	'\u2000',
	'\u2001',
	'\u2002',
	'\u2003',
	'\u2004',
	'\u2005',
	'\u2006',
	'\u2007',
	'\u2008',
	'\u2009',
	'\u200a',
	'\u202f',
	'\u205f',
	'\u3000',
}

func hide() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stderr, "Enter what you want to hide, only ascii availabe: ")
	secret, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}
	secret = strings.TrimRight(secret, "\n")
	fmt.Fprintf(os.Stderr, "Enter the carrier text: ")
	carrier, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}
	carrier = strings.TrimRight(carrier, "\n")

	secretSpace := ""
	for _, c := range []rune(secret) {
		if c > unicode.MaxASCII {
			fmt.Fprintf(os.Stderr, "error: only ascii availabe\n")
			return
		}
		firstSpace := spaceMap[c/16]
		secondSpace := spaceMap[c%16]
		secretSpace = fmt.Sprintf("%s%c%c", secretSpace, firstSpace, secondSpace)
	}

	carrierList := strings.Split(carrier, " ")
	index := 0 // index to carrierList

	for _, c := range []rune(secretSpace) {
		fmt.Printf("%s%c", carrierList[index], c)
		index = (index + 1) % len(carrierList)
	}
	fmt.Printf(carrierList[index])
}

func extract() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	secretSpace := ""
	for _, c := range []rune(text) {
		if unicode.IsSpace(c) {
			secretSpace = fmt.Sprintf("%s%c", secretSpace, c)
		}
	}
	secret := ""
	for i := 0; i < len(secretSpace); i += 2 {
		firstSpace := rune(secretSpace[i])
		secondSpace := rune(secretSpace[i+1])
		firstIndex := 0
		secondIndex := 0
		for i, c := range spaceMap {
			if c == firstSpace {
				firstIndex = i
			}
			if c == secondSpace {
				secondIndex = i
			}
		}
		secret = fmt.Sprintf("%s%c", secret, firstIndex*16+secondIndex)
	}
	fmt.Printf(secret)
}

func main() {
	// hide()
	extract()
}
