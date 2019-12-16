package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func main() {
	// Prepare some data to insert into the template.
	type service struct {
		Name      string
		Table     string
		Footer    string
		Stats     string
		Objects   []string
		OuterLoop []int
	}

	table := "| a | b |\n| --- | --- |\n| 1 | 2 |\n"
	stats := "| C | D |\n| --- | --- |\n| 1a | 2a |\n"

	tmpl := template.Must(template.ParseFiles("./generic.md.templ", "./data.md.templ"))
	var o bytes.Buffer
	tmpl.Execute(&o, &service{
		Name:      "Here is the name.",
		Table:     table,
		Footer:    "Some Text to take us out.",
		Stats:     stats,
		Objects:   []string{"f", "g", "h"},
		OuterLoop: []int{0, 1, 2, 3, 4},
	})
	fmt.Println(o.String())
}
