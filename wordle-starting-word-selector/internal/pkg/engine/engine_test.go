package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	defaultTestingDictionary = Words{
		"hello",
		"no",
		"pause",
		"mo",
		"yurts",
	}
)

func TestGetScrabbleValue(t *testing.T) {
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

func TestFilterLength(t *testing.T) {
	tests := []struct {
		word     string
		length   int
		expected bool
	}{
		{"testing", 7, true},
		{"no", 2, true},
		{"here", 3, false},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, filterLength(test.word, test.length))
	}
}

func TestFilterScrabbleValue(t *testing.T) {
	tests := []struct {
		word     string
		value    int
		expected bool
	}{
		{"testing", 8, true},
		{"no", 2, true},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, filterScrabbleValue(test.word, test.value))
	}
}

func TestFilterDuplicateLetters(t *testing.T) {
	tests := []struct {
		word     string
		expected bool
	}{
		{"testing", false},
		{"no", true},
		{"pause", true},
		{"hello", false},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, filterDuplicateLetters(test.word, 0))
	}
}

func TestFilter(t *testing.T) {
	t.Run("All true", func(t *testing.T) {
		trueCheck := func(word string, value int) bool {
			return true
		}
		words := defaultTestingDictionary
		filteredWords := words.filter(trueCheck, 0)
		assert.ElementsMatch(t, defaultTestingDictionary, filteredWords)
	})

	t.Run("All false", func(t *testing.T) {
		falseCheck := func(word string, value int) bool {
			return false
		}
		words := defaultTestingDictionary
		filteredWords := words.filter(falseCheck, 0)
		assert.ElementsMatch(t, Words{}, filteredWords)
	})
}
