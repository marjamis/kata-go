package main

import (
	"sort"
	"testing"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func TestCustomSort(t *testing.T) {
	terms := SortKVs{}

	for i := 0; i < 30; i++ {
		f := fuzz.New()
		object := SortKV{}
		f.Fuzz(&object)
		terms = append(terms, object)
	}

	t.Run("Length Check", func(t *testing.T) {
		assert.Len(t, terms, 30)
	})

	t.Run("Sort Output Check", func(t *testing.T) {
		sort.Sort(sort.Reverse(terms))
		allTrue := true
		for i := 0; i < 30-1; i++ {
			if !(terms[i].Value < terms[i+1].Value) {
				allTrue = false
				break
			}
		}
		assert.True(t, allTrue)
	})
}
