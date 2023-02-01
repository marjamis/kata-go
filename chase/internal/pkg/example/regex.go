package example

import (
	"fmt"
	"regexp"
)

const str = "data:marjamis@EXAMPLE.COM"

func regexpRun() {
	fmt.Println(regexp.MustCompile(`^data:(.*)@.*$`).FindStringSubmatch(str)[1])
}

func init() {
	category := GetCategories().AddCategory("regex")

	category.AddExample("submatch",
		CategoryExample{
			Description: "Find string submatch",
			Function:    stringsRun,
		})
}
