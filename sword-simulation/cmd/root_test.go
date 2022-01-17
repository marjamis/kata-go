package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicSimulation(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{
			2,
			1,
		},
		{
			3,
			3,
		},
		// This example answer from: https://www.geeksforgeeks.org/lucky-alive-person-circle/
		{
			100,
			73,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, run(test.input).Number)
	}
}
