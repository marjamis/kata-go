package example

import (
	"fmt"
	"regexp"
)

const str = "data:marjamis@EXAMPLE.COM"

func RegexpRun() {
	fmt.Println(regexp.MustCompile(`^data:(.*)@.*$`).FindStringSubmatch(str)[1])
}

func init() {
	GetMyExamples().Add("regex", RegexpRun)
}
