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

var spaceMapReverse = map[rune]int{}

func init() {
	for i, c := range spaceMap {
		spaceMapReverse[c] = i
	}
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
	fmt.Println(carrierList[index])
}

func extract() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	// decode space to int
	secretSpace := []rune{}
	for _, c := range []rune(text) {
		if !unicode.IsSpace(c) {
			continue
		}
		if index, ok := spaceMapReverse[c]; ok {
			secretSpace = append(secretSpace, rune(index))
		}
	}

	// TODO: merge two for loop
	for i := 0; i < len(secretSpace); i += 2 {
		fmt.Printf("%c", secretSpace[i]*16+secretSpace[i+1])
	}
}

func help() {
	fmt.Println("hide: hide the secret in the carrier text")
	fmt.Println("extract: extract the secret from the carrier text")
}

func main() {
	subCMD := map[string]func(){
		"hide":    hide,
		"extract": extract,
		"help":    help,
	}
	if len(os.Args) < 2 {
		help()
	} else if f, ok := subCMD[os.Args[1]]; ok {
		f()
	} else {
		help()
	}
}
