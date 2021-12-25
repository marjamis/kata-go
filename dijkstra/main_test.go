package main

import (
	"fmt"
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
		source      string
		destination string
		expected    int
	}{
		{
			"./test_data/example.csv",
			"Node0",
			"Node6",
			19,
		},
	}

	for _, v := range tests {
		nodes := generateNodeMap(ReadString(v.filename))
		workflow(nodes, v.source, v.destination)
		assert.Equal(t, v.expected, nodes[v.destination].TraversalDetails.Distance)
	}
}

func TestDijkstraMassiveTest(t *testing.T) {
	tests := []struct {
		source      string
		destination string
		expected    int
	}{
		{"Node1", "Node1", 0},
		{"Node1", "Node2", 9},
		{"Node1", "Node3", 4},
		{"Node1", "Node4", 13},
		{"Node1", "Node5", 14},
		{"Node1", "Node6", 11},
		{"Node1", "Node7", 14},
		{"Node1", "Node8", 28},
		{"Node1", "Node9", 26},
		{"Node1", "Node10", 16},
		{"Node1", "Node11", 19},
		{"Node1", "Node12", 25},

		{"Node2", "Node1", 10},
		{"Node2", "Node2", 0},
		{"Node2", "Node3", 5},
		{"Node2", "Node4", 4},
		{"Node2", "Node5", 5},
		{"Node2", "Node6", 12},
		{"Node2", "Node7", 15},
		{"Node2", "Node8", 29},
		{"Node2", "Node9", 27},
		{"Node2", "Node10", 17},
		{"Node2", "Node11", 20},
		{"Node2", "Node12", 26},

		{"Node3", "Node1", 15},
		{"Node3", "Node2", 5},
		{"Node3", "Node3", 0},
		{"Node3", "Node4", 9},
		{"Node3", "Node5", 10},
		{"Node3", "Node6", 7},
		{"Node3", "Node7", 10},
		{"Node3", "Node8", 24},
		{"Node3", "Node9", 22},
		{"Node3", "Node10", 12},
		{"Node3", "Node11", 15},
		{"Node3", "Node12", 21},

		{"Node4", "Node1", 56},
		{"Node4", "Node2", 46},
		{"Node4", "Node3", 41},
		{"Node4", "Node4", 0},
		{"Node4", "Node5", 1},
		{"Node4", "Node6", 34},
		{"Node4", "Node7", 31},
		{"Node4", "Node8", 29},
		{"Node4", "Node9", 27},
		{"Node4", "Node10", 34},
		{"Node4", "Node11", 37},
		{"Node4", "Node12", 36},

		{"Node5", "Node1", 73},
		{"Node5", "Node2", 63},
		{"Node5", "Node3", 58},
		{"Node5", "Node4", 67},
		{"Node5", "Node5", 0},
		{"Node5", "Node6", 51},
		{"Node5", "Node7", 54},
		{"Node5", "Node8", 58},
		{"Node5", "Node9", 56},
		{"Node5", "Node10", 46},
		{"Node5", "Node11", 41},
		{"Node5", "Node12", 35},

		{"Node6", "Node1", 22},
		{"Node6", "Node2", 12},
		{"Node6", "Node3", 7},
		{"Node6", "Node4", 16},
		{"Node6", "Node5", 7},
		{"Node6", "Node6", 0},
		{"Node6", "Node7", 3},
		{"Node6", "Node8", 17},
		{"Node6", "Node9", 15},
		{"Node6", "Node10", 5},
		{"Node6", "Node11", 8},
		{"Node6", "Node12", 14},

		{"Node7", "Node1", 25},
		{"Node7", "Node2", 15},
		{"Node7", "Node3", 10},
		{"Node7", "Node4", 19},
		{"Node7", "Node5", 4},
		{"Node7", "Node6", 3},
		{"Node7", "Node7", 0},
		{"Node7", "Node8", 20},
		{"Node7", "Node9", 18},
		{"Node7", "Node10", 8},
		{"Node7", "Node11", 11},
		{"Node7", "Node12", 17},

		{"Node8", "Node1", 27},
		{"Node8", "Node2", 17},
		{"Node8", "Node3", 12},
		{"Node8", "Node4", 21},
		{"Node8", "Node5", 6},
		{"Node8", "Node6", 5},
		{"Node8", "Node7", 2},
		{"Node8", "Node8", 0},
		{"Node8", "Node9", 2},
		{"Node8", "Node10", 5},
		{"Node8", "Node11", 8},
		{"Node8", "Node12", 14},

		{"Node9", "Node1", 29},
		{"Node9", "Node2", 19},
		{"Node9", "Node3", 14},
		{"Node9", "Node4", 23},
		{"Node9", "Node5", 8},
		{"Node9", "Node6", 7},
		{"Node9", "Node7", 4},
		{"Node9", "Node8", 2},
		{"Node9", "Node9", 0},
		{"Node9", "Node10", 7},
		{"Node9", "Node11", 10},
		{"Node9", "Node12", 16},

		{"Node10", "Node1", 27},
		{"Node10", "Node2", 17},
		{"Node10", "Node3", 12},
		{"Node10", "Node4", 21},
		{"Node10", "Node5", 12},
		{"Node10", "Node6", 5},
		{"Node10", "Node7", 8},
		{"Node10", "Node8", 12},
		{"Node10", "Node9", 10},
		{"Node10", "Node10", 0},
		{"Node10", "Node11", 3},
		{"Node10", "Node12", 9},

		{"Node11", "Node1", 32},
		{"Node11", "Node2", 22},
		{"Node11", "Node3", 17},
		{"Node11", "Node4", 26},
		{"Node11", "Node5", 14},
		{"Node11", "Node6", 10},
		{"Node11", "Node7", 13},
		{"Node11", "Node8", 17},
		{"Node11", "Node9", 15},
		{"Node11", "Node10", 5},
		{"Node11", "Node11", 0},
		{"Node11", "Node12", 6},

		{"Node12", "Node1", 38},
		{"Node12", "Node2", 28},
		{"Node12", "Node3", 23},
		{"Node12", "Node4", 32},
		{"Node12", "Node5", 8},
		{"Node12", "Node6", 16},
		{"Node12", "Node7", 19},
		{"Node12", "Node8", 23},
		{"Node12", "Node9", 21},
		{"Node12", "Node10", 11},
		{"Node12", "Node11", 6},
		{"Node12", "Node12", 0},
	}

	for _, v := range tests {
		t.Run(fmt.Sprintf("Source: %s Destination %s", v.source, v.destination), func(t *testing.T) {
			nodes := generateNodeMap(ReadString("./test_data/massive.csv"))
			workflow(nodes, v.source, v.destination)
			assert.True(t, nodes[v.destination].TraversalDetails.Visited)
			assert.Equal(t, v.expected, nodes[v.destination].TraversalDetails.Distance)
		})
	}
}

func TestDijkstraMissingRoutes(t *testing.T) {
	tests := []struct {
		source      string
		destination string
		expected    int
	}{
		{"Node13", "Node1", 0},
		{"Node13", "Node2", 0},
		{"Node13", "Node3", 0},
		{"Node13", "Node4", 0},
		{"Node13", "Node5", 0},
		{"Node13", "Node6", 0},
		{"Node13", "Node7", 0},
		{"Node13", "Node8", 0},
		{"Node13", "Node9", 0},
		{"Node13", "Node10", 0},
		{"Node13", "Node11", 0},
		{"Node13", "Node12", 0},
	}

	for _, v := range tests {
		t.Run(fmt.Sprintf("Source: %s Destination %s", v.source, v.destination), func(t *testing.T) {
			nodes := generateNodeMap(ReadString("./test_data/massive.csv"))
			workflow(nodes, v.source, v.destination)
			assert.False(t, nodes[v.destination].TraversalDetails.Visited)
			assert.Equal(t, v.expected, nodes[v.destination].TraversalDetails.Distance)
		})
	}

	t.Run("Check that even if a node doesn't have a path to any destination it still is visited itself", func(t *testing.T) {
		nodes := generateNodeMap(ReadString("./test_data/massive.csv"))
		workflow(nodes, "Node13", "Node1")
		assert.True(t, nodes["Node13"].TraversalDetails.Visited)
		assert.Equal(t, 0, nodes["Node13"].TraversalDetails.Distance)
	})
}
