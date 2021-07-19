package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var buf bytes.Buffer
	for i := 0; i < 500; i++ {
		buf.WriteString("z")
	}
	fmt.Println(buf.String())

	s := "hello"
	fmt.Println(len(s))
	fmt.Println(s[0])        // The first char in bytes rather than a character
	fmt.Printf("%q\n", s[0]) // The first char as quoted verb will escape non-printable characters
	fmt.Printf("%b\n", s[0]) //The first char in binary

	fmt.Println("hello\nnow") // Control characters used properly
	fmt.Println(`hello\nnow`) // Verbatim output

	string := "Oh I do like to be beside the seaside,"
	fmt.Println(strings.ToUpper(string))
	fmt.Println(strings.Replace(string, "seaside", "bar", -1))
	fmt.Println(strings.Index(string, "the"))
}
