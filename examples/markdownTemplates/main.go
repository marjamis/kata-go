package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"toLower": toLower,
	"toUpper": toUpper,
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func toLower(s string) string {
	return strings.ToLower(s)
}

func main() {
	// Prepare some data to insert into the template.
	table := "| a | b |\n| --- | --- |\n| 1 | 2 |\n"
	stats := "| C | D |\n| --- | --- |\n| 1a | 2a |\n"
	service := struct {
		Name      string
		Table     string
		Footer    string
		Stats     string
		Objects   []string
		OuterLoop []int
	}{
		Name:      "TestName",
		Table:     table,
		Footer:    "This is the end.",
		Stats:     stats,
		Objects:   []string{"f", "g", "h"},
		OuterLoop: []int{0, 1, 2, 3, 4},
	}

	//New("generic.md.templ") - The template name and file name are the same to prevent the creation of a new template but instead use one from New with that name. Names can be defined as well as be the filename. Also template.ExecuteTemplate() lets you specify an explicit filename to use as the base rather than the first being parsed.

	//ParseFiles is used instead of ParseGlob to show that while different "content" templates being defined in different files are being used the content keyword is being use in generic.md.templ and is simply filled in with the content template depending on which file is parsed. If ParseGlob is used the last one of the name would be used and would break this test.

	//Note: These examples show a lot of different ways to have the templating setup which is powerful but all the different tests do also make it a bit hard to follow. This exemplifies for production code this has to be properly thought out for easy reading and debugging.

	tmpl := template.Must(template.New("base.md.templ").Funcs(funcMap).ParseFiles("./templates/base.md.templ", "./templates/stats.md.templ", "./templates/context.md.templ", "./templates/content1.md.templ"))
	var o bytes.Buffer
	tmpl.Execute(&o, service)
	fmt.Printf("Template Example 1\n%s\n\n", o.String())

	tmpl2 := template.Must(template.New("base.md.templ").Funcs(funcMap).ParseFiles("./templates/base.md.templ", "./templates/stats.md.templ", "./templates/context.md.templ", "./templates/content2.md.templ"))
	var o2 bytes.Buffer
	//Setting to nil to show the if statement portion
	service.Objects = nil
	tmpl2.Execute(&o2, service)
	fmt.Printf("Template Example 2\n%s\n\n", o2.String())
}
