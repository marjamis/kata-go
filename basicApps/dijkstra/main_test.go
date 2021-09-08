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

func BenchmarkStringFromBuffer(b *testing.B) {
	terms := SortKVs{}

	for i := 0; i < 30; i++ {
		f := fuzz.New()
		object := SortKV{}
		f.Fuzz(&object)
		terms = append(terms, object)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		temp := terms
		sort.Sort(sort.Reverse(temp))
	}
}

func TestDijkstra(t *testing.T) {
	tests := []struct {
		filename    string
		expected    int
		source      string
		destination string
	}{
		{
			"./test_data/example.csv",
			19,
			"Node0",
			"Node6",
		},
		{
			"./test_data/massive.csv",
			1,
			"Node0",
			"Node3",
		},
	}

	for _, v := range tests {
		nodes := generateNodeMap(ReadString(v.filename))
		distance, _ := workflow(nodes, v.source, v.destination)
		assert.Equal(t, v.expected, distance)
	}
}
