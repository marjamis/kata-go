package main

import (
	"fmt"
	"regexp"
)

const string = "data:marjamis@EXAMPLE.COM"

func main() {
	fmt.Println(regexp.MustCompile(`^data:(.*)@.*$`).FindStringSubmatch(string)[1])
}
