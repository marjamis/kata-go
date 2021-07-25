package main

import (
	"testing"
)

func testData() map[string]Node {
	output := ReadString("./test_data/example.csv")
	// TODO Removing the last char due to a blank line. fix this so I dont have to do this
	return generateNodeMap(output[0 : len(output)-1])
}

func TestGeneral(t *testing.T) {
	//TOD implement proper testing prior to future implementations, now that the algorithm is understood, for TDD
}
