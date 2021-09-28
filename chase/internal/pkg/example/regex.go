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
	examples := ExampleRuns{
		{"FindStringSubmatch", regexpRun},
	}
	GetMyExamples().Add("regex", examples.runExamples)
}
