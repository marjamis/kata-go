package main

/*
Adapted from: https://www.freecodecamp.org/news/dijkstras-shortest-path-algorithm-visual-introduction/
*/

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/generic"
)

// ReadString will take a filename and return the contents as a string
func ReadString(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}

	return string(data[:len(data)-1])
}

// Edge contains information related to the edge, connection between two nodes, of a graph
type Edge struct {
	Weight      int
	Source      Node
	Destination Node
	Directed    bool
}

// Node contains information related to the node of a graph
type Node struct {
	Name             string
	Edges            []Edge
	TraversalDetails TraversalDetails
}

// TraversalDetails contains details about the a node from the source node
type TraversalDetails struct {
	Distance      int
	ShortestPath  []string
	Visited       bool
	IsDistanceSet bool
}

type nodemap map[string]Node

// SortKV is the struct for the custom search of shortest distances for the shortest path
type SortKV struct {
	Key   string
	Value int
}

// SortKVs is an array of the struct for the custom search for the shortest path
type SortKVs []SortKV

func (s SortKVs) Len() int {
	return len(s)
}

func (s SortKVs) Less(i, j int) bool {
	return s[i].Value > s[j].Value
}

func (s SortKVs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func createDiagram(nodes nodemap, destinationNode Node) {
	image, err := diagram.New(diagram.Label("dijkstra Nodes"), diagram.Filename("dijkstra"))
	if err != nil {
		log.Fatal(err)
	}

	imageNodes := map[string]*diagram.Node{}
	for k := range nodes {
		imageNodes[k] = generic.Place.Datacenter().Label(k)
	}

	connectionAlreadyExists := map[string]bool{}
	for _, v := range nodes {
		for _, edge := range v.Edges {
			if _, ok := connectionAlreadyExists[edge.Source.Name+"To"+edge.Destination.Name]; !ok {
				image.Connect(imageNodes[edge.Source.Name], imageNodes[edge.Destination.Name], func(o *diagram.EdgeOptions) {
					o.Forward = true
					// If the edge is directed it means the reverse shouldn't be automatically applied. If it's not directed then a reverse should be applied
					o.Reverse = !edge.Directed
					o.Label = strconv.Itoa(edge.Weight)
				})
				connectionAlreadyExists[edge.Source.Name+"To"+edge.Destination.Name] = true
			}
		}
	}

	// Placing the line for shortest path
	for i := range destinationNode.TraversalDetails.ShortestPath[:len(destinationNode.TraversalDetails.ShortestPath)-1] {
		s := destinationNode.TraversalDetails.ShortestPath[i]
		d := destinationNode.TraversalDetails.ShortestPath[i+1]

		image.Connect(imageNodes[s], imageNodes[d], func(o *diagram.EdgeOptions) {
			o.Color = "#ff0000"
		})
	}

	if err := image.Render(); err != nil {
		log.Fatal(err)
	}
}

// generateNodeMap generates the map structure based off a preformatted string input
func generateNodeMap(data string) nodemap {
	nodes := map[string]Node{}

	edges := []struct {
		Weight      int
		Source      string
		Destination string
		Directed    bool
	}{}

	for _, dataRow := range strings.Split(data, "\n") {
		dataRowDetails := strings.Split(dataRow, ",")
		var directed bool
		if strings.Compare(string(dataRowDetails[3]), "false") == 0 {
			directed = false
		} else if strings.Compare(string(dataRowDetails[3]), "true") == 0 {
			directed = true
		}
		// Simple example code hence ignoring the err
		edgeWeight, _ := strconv.Atoi(dataRowDetails[2])

		edges = append(edges, struct {
			Weight      int
			Source      string
			Destination string
			Directed    bool
		}{
			edgeWeight,
			string(dataRowDetails[0]),
			string(dataRowDetails[1]),
			directed,
		})
	}

	for _, edge := range edges {
		// Creates the source node if not already in the map
		_, ok := nodes[edge.Source]
		if !ok {
			nodes[edge.Source] = Node{
				Name: edge.Source,
			}
		}

		// Creates the destination node if not already in the map
		_, ok = nodes[edge.Destination]
		if !ok {
			nodes[edge.Destination] = Node{
				Name: edge.Destination,
			}
		}

		// Creates the edge based on the inputs read from the string data
		newEdge := Edge{
			Weight:      edge.Weight,
			Source:      nodes[edge.Source],
			Destination: nodes[edge.Destination],
			Directed:    edge.Directed,
		}

		// Appends the edge to the source as the path is definitely available in this direction
		node := nodes[edge.Source]
		node.Edges = append(node.Edges, newEdge)
		nodes[edge.Source] = node

		// If the edge isn't directed, mean traffic can go in both directions, this adds the edge to the destination as well
		// for a bidirectional path
		if !edge.Directed {
			newEdge = Edge{
				Weight: edge.Weight,
				// This is swapped to ensure that if the edge isn't directed the reverse edge is applied to the correct source/destination, i.e. the opposite way
				Source:      nodes[edge.Destination],
				Destination: nodes[edge.Source],
				Directed:    edge.Directed,
			}

			node = nodes[edge.Destination]
			node.Edges = append(node.Edges, newEdge)
			nodes[edge.Destination] = node
		}
	}

	return nodes
}

// shortestDistanceNode returns the name of the node with the smallest weight which hasn't been marked as visited
func shortestDistanceNode(nodes nodemap) (string, bool) {
	var kvs SortKVs
	for k, v := range nodes {
		// The only ones we care about here are nodes that haven't been marked as visited and have had their distance explicitly set
		if !nodes[k].TraversalDetails.Visited && nodes[k].TraversalDetails.IsDistanceSet {
			kvs = append(kvs, SortKV{
				Key:   k,
				Value: v.TraversalDetails.Distance,
			})
		}
	}

	sort.Sort(sort.Reverse(kvs))

	// Check that the array isn't empy and if it is return that the shortest distance wasn't found
	if len(kvs) < 1 {
		return "", false
	}

	return kvs[0].Key, nodes[kvs[0].Key].TraversalDetails.IsDistanceSet
}

func (nodes nodemap) print() {
	var keys []string

	for k := range nodes {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Println("Distance:")
	for _, v := range keys {
		fmt.Printf("%s: %d V: %t\n", v, nodes[v].TraversalDetails.Distance, nodes[v].TraversalDetails.Visited)
	}
}

// dijkstra is the core of the algorithm
func dijkstra(nodes nodemap, path *[]string) {
	// Loop through each node in the path (as they're already precalculated) to find the next node with the smallest weight.
	for _, pathNode := range *path {
		// For each node in the path we need to loop through it's edges to find the next node candidate with the smallest weight
		for _, edge := range nodes[pathNode].Edges {
			// If the destination node of the edge is already marked as visited we skip as this path has already been calculated
			if !nodes[edge.Destination.Name].TraversalDetails.Visited {
				// If the weight of the current edge to this destination node isn't set or is smaller than whats currently tracked it updates to the shorter path
				// Otherwise it performs no additional operation as we already have set the shortest path details
				proposedDistance := edge.Weight + nodes[pathNode].TraversalDetails.Distance
				if !nodes[edge.Destination.Name].TraversalDetails.IsDistanceSet || (proposedDistance < nodes[edge.Destination.Name].TraversalDetails.Distance) {
					var tn = nodes[edge.Destination.Name]
					var td = tn.TraversalDetails

					td.Distance = proposedDistance
					td.ShortestPath = append(nodes[pathNode].TraversalDetails.ShortestPath, edge.Destination.Name)
					td.IsDistanceSet = true

					tn.TraversalDetails = td
					nodes[edge.Destination.Name] = tn
				}
			}
		}
	}

	// Select the node closest to the source node based on the currently known distances
	shortest, foundShortest := shortestDistanceNode(nodes)

	// Checks if the shortest distance had been found, if it hasn't it means there are still nodes available BUT they may not be connected to the main set of nodes, i.e. not processed for this path
	if foundShortest {
		// Mark this shortest node as visited
		var tn = nodes[shortest]
		var td = tn.TraversalDetails

		td.Visited = true

		tn.TraversalDetails = td
		nodes[shortest] = tn

		// Add it to the path as a node that has already been checked for the smallest weight
		*path = append(*path, shortest)
	}
}

func workflow(nodes nodemap, source string, destination string) error {
	// Simple check on if both the source and destination exist within the nodemap
	_, doesSourceNodeExist := nodes[source]
	_, doesDestinationNodeExist := nodes[destination]
	if !doesSourceNodeExist || !doesDestinationNodeExist {
		return fmt.Errorf("Source (%s: %t) or destination (%s: %t) not available in provided node map", source, doesSourceNodeExist, destination, doesDestinationNodeExist)
	}

	// Set the defaults for the source node
	var td = nodes[source].TraversalDetails
	var node = nodes[source]

	td.Distance = 0
	td.ShortestPath = []string{
		source,
	}
	td.Visited = true
	td.IsDistanceSet = true

	node.TraversalDetails = td
	nodes[source] = node

	path := []string{source}
	// This will loop through the algorthm to ensure all the nodes have a distance value from the source node, if they can
	for i := 0; i < len(nodes); i++ {
		dijkstra(nodes, &path)
	}

	return nil
}

func main() {
	filename := flag.String("f", "./test_data/example.csv", "specify the file of node data")
	sourceNode := flag.String("s", "Node0", "specify the source node")
	destinationNode := flag.String("d", "Node6", "specify the destination node")
	flag.Parse()

	nodes := generateNodeMap(ReadString(*filename))
	err := workflow(nodes, *sourceNode, *destinationNode)
	if err != nil {
		fmt.Printf("Application can't proceed due to error: %s\n", err)
		fmt.Println(nodes)
		return
	}

	nodes.print()
	fmt.Printf("Distance between %s and %s is: %d with the shortest path being: %+v\n", *sourceNode, *destinationNode, nodes[*destinationNode].TraversalDetails.Distance, nodes[*destinationNode].TraversalDetails.ShortestPath)

	createDiagram(nodes, nodes[*destinationNode])
}
