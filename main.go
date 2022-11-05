package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read err: %v\n", err)
		os.Exit(1)
	}

	s := string(b)

	s = strings.ReplaceAll(s, "th", "þ")
	s = strings.ReplaceAll(s, "Th", "Þ")
	s = strings.ReplaceAll(s, "ond", "⁊")
	s = strings.ReplaceAll(s, "and", "⁊")
	s = strings.ReplaceAll(s, "And", "⁊")
	s = strings.ReplaceAll(s, "ss", "ß")
	s = strings.ReplaceAll(s, "s", "ſ")

	fmt.Printf(s)
}
