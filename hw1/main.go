package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var SpaceMap = []rune{
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
	// build reverse map from spaceMap
	for i, c := range SpaceMap {
		spaceMapReverse[c] = i
	}
}

func Hide() {
	// input
	// secret
	reader := bufio.NewReader(os.Stdin)
	fmt.Fprintf(os.Stderr, "Enter what you want to hide, only ascii availabe: ")
	secret, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}
	// carrier
	secret = strings.TrimRight(secret, "\n")
	fmt.Fprintf(os.Stderr, "Enter the carrier text: ")
	carrier, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}
	carrier = strings.TrimRight(carrier, "\n")

	// split and remove space
	carrierListN := strings.Split(carrier, " ")
	carrierList := []string{}
	index := 0 // index to carrierList
	getCarrier := func() string {
		s := carrierList[index]
		index = (index + 1) % len(carrierList)
		return s
	}
	for _, s := range carrierListN {
		if s != "" {
			carrierList = append(carrierList, strings.Trim(s, " "))
		}
	}

	// hide
	for _, c := range []rune(secret) {
		if c > unicode.MaxASCII {
			fmt.Fprintf(os.Stderr, "error: only ascii availabe\n")
			return
		}
		firstSpace := SpaceMap[c/16]
		secondSpace := SpaceMap[c%16]
		fmt.Printf("%s%c%s%c", getCarrier(), firstSpace, getCarrier(), secondSpace)
	}
	fmt.Printf(getCarrier())
}

func Extract() {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return
	}

	isFirst := true
	char := '\000'
	for _, c := range []rune(text) {
		if index, ok := spaceMapReverse[c]; ok {
			if isFirst {
				char = rune(index * 16)
				isFirst = false
			} else {
				char += rune(index)
				fmt.Printf("%c", char)
				isFirst = true
			}
		}
	}
}

func Help() {
	fmt.Println("hide: hide the secret in the carrier text")
	fmt.Println("extract: extract the secret from the carrier text")
}

func main() {
	subCMD := map[string]func(){
		"hide":    Hide,
		"extract": Extract,
		"help":    Help,
	}
	if len(os.Args) < 2 {
		Help()
	} else if f, ok := subCMD[os.Args[1]]; ok {
		f()
	} else {
		Help()
	}
}
