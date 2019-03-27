package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	// seed properly
	rand.Seed(time.Now().UTC().UnixNano())

	// if string passed as param, return them
	if len(os.Args) >= 2 {
		for i := range os.Args {
			os.Args[i] = mocking(os.Args[i])
		}
		print(strings.Join(os.Args[1:], " "))
		os.Exit(0)
	}

	// else, get stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(mocking(scanner.Text()))
	}
}

func mocking(str string) string {
	var result string
	var flip bool
	for _, c := range str {
		if unicode.IsLetter(c) {
			if rand.Float32() < 0.6 {
				flip = !flip
			}
			if flip {
				c = unicode.ToUpper(c)
			} else {
				c = unicode.ToLower(c)
			}
		}
		result = result + string(c)
	}
	return result
}
