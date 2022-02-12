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

func TestFilterWordLength(t *testing.T) {
	words = defaultTestingDictionary
	tests := []struct {
		len int
	}{
		{5},
		{2},
	}

	for _, test := range tests {
		words := words
		words = words.filterLength(test.len)
		for _, word := range words {
			assert.Len(t, word, test.len)
		}
	}
}

func TestFilterScrabbleValue(t *testing.T) {
	words = defaultTestingDictionary
	tests := []struct {
		value int
	}{
		{8},
		{2},
	}

	for _, test := range tests {
		words := words
		words = words.filterScrabbleValue(test.value)
		for _, word := range words {
			assert.Equal(t, test.value, getScrabbleValue(word))
		}
	}
}

func TestContainsDuplicates(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"hello", true},
		{"heaph", true},
		{"pause", false},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, hasDuplicateLetters(test.input))
	}
}

func TestFilterDuplicateLetters(t *testing.T) {
	words := Words{
		"hello",
		"no",
		"pause",
		"mo",
	}
	expected := Words{
		"no",
		"pause",
		"mo",
	}
	assert.ElementsMatch(t, expected, words.filterDuplicateLetters())
}
