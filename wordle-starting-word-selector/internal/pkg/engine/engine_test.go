package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScrabbleValue(t *testing.T) {
	tests := []struct {
		input                 string
		expectedScrabbleValue int
	}{
		{"hello", 8},
		{"no", 2},
	}

	for _, test := range tests {
		assert.Equal(t, test.expectedScrabbleValue, getScrabbleValue(test.input))
	}
}

func TestGetSelectedWord(t *testing.T) {
	tests := []struct {
		inputLen           int
		inputScrabbleValue int
	}{
		{5, 8},
		{2, 2},
	}

	for _, test := range tests {
		word, _ := getWord(test.inputLen, test.inputScrabbleValue)
		assert.Len(t, word, test.inputLen)
		assert.Equal(t, test.inputScrabbleValue, getScrabbleValue(word))
	}

	t.Run("No match for selection criteria", func(t *testing.T) {
		tests := []struct {
			inputLen           int
			inputScrabbleValue int
		}{
			{5, 5},
		}

		for _, test := range tests {
			_, err := getWord(test.inputLen, test.inputScrabbleValue)
			if assert.Error(t, err) {
				assert.Equal(t, "No available word", err.Error())
			}
		}
	})
}

func TestFilterWordLength(t *testing.T) {
	tests := []struct {
		len int
	}{
		{5},
		{2},
	}

	for _, test := range tests {
		wordList := dict
		wordList = wordList.filterLen(test.len)
		for _, word := range wordList {
			assert.Len(t, word, test.len)
		}
	}
}

func TestFilterScrabbleValue(t *testing.T) {
	tests := []struct {
		value int
	}{
		{8},
		{2},
	}

	for _, test := range tests {
		wordList := dict
		wordList = wordList.filterScrabbleValue(test.value)
		for _, word := range wordList {
			assert.Equal(t, test.value, getScrabbleValue(word))
		}
	}
}
