package main

/*
Adapted from: https://www.freecodecamp.org/news/dijkstras-shortest-path-algorithm-visual-introduction/
*/

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/generic"
)

func ReadString(file string) string {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}

	return string(data)
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
	Name  string
	Edges []Edge
}

// Distance contains details about the a node from the source node
type Distance struct {
	Distance int
	Visited  bool
}

type distancemap map[string]Distance

func createDiagram(nodes map[string]Node) {
	image, err := diagram.New(diagram.Label("dijkstra Nodes"), diagram.Filename("dijkstra"), diagram.Direction("LR"))
	if err != nil {
		log.Fatal(err)
	}

	image_nodes := map[string]*diagram.Node{}
	for k, _ := range nodes {
		image_nodes[k] = generic.Place.Datacenter().Label(k)
	}

	connectionAlreadyExists := map[string]bool{}
	for _, v := range nodes {
		for _, edge := range v.Edges {
			// TODO this check will need improvement when multi-directions are implemented
			if _, ok := connectionAlreadyExists[edge.Source.Name+"To"+edge.Destination.Name]; !ok {
				image.Connect(image_nodes[edge.Source.Name], image_nodes[edge.Destination.Name], diagram.Bidirectional())
				connectionAlreadyExists[edge.Source.Name+"To"+edge.Destination.Name] = true
			}
		}
	}

	// TODO Add the distance map when supported

	if err := image.Render(); err != nil {
		log.Fatal(err)
	}
}

// generateNodeMap generates the map structure based off a preformatted string input
func generateNodeMap(data string) map[string]Node {
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

		// TODO improve the logic on this so in can be unidirectional and different weights
		// Creates the edge based on the inputs read from the string data
		edge := Edge{
			Weight:      edge.Weight,
			Source:      nodes[edge.Source],
			Destination: nodes[edge.Destination],
			Directed:    edge.Directed,
		}

		// Appends the edge to the source and destination node as the path is usable in both directions
		node := nodes[edge.Source.Name]
		node.Edges = append(node.Edges, edge)
		nodes[edge.Source.Name] = node

		node = nodes[edge.Destination.Name]
		node.Edges = append(node.Edges, edge)
		nodes[edge.Destination.Name] = node
	}

	return nodes
}

// shortestDistanceNode returns the name of the node with the smallest weight which hasn't been marked as visited
func shortestDistanceNode(distances map[string]Distance) string {
	// TODO remove the arbitary large value
	var smallestWeight int = 999999
	var nodeNameOfSmallestWeight string
	for k, v := range distances {
		if v.Distance != 0 && v.Distance != -1 && v.Distance < smallestWeight && !distances[k].Visited {
			smallestWeight = v.Distance
			nodeNameOfSmallestWeight = k
		}
	}

	return nodeNameOfSmallestWeight
}

func (d distancemap) print() {
	fmt.Println("Distance:")
	for k, v := range d {
		fmt.Printf("%s: %d V: %t\n", k, v.Distance, v.Visited)
	}
}

// dijkstra is the core of the algorithm
func dijkstra(nodes map[string]Node, distances map[string]Distance, path *[]string) {

	// Loop through each node in the path (as they're already precalculated) to find the next node with the smallest weight.
	for _, v := range *path {
		// For each node in the path we need to loop through it's edges to find the next node candidate with the smallest weight
		for _, edge := range nodes[v].Edges {
			// TODO likely an efficiency issue has it checks all paths edges again just doesn't update them?
			// If the destination node of the edge is already marked as visited we skip as this path has already been calculated
			if !distances[edge.Destination.Name].Visited {
				// If the weight of the current edge to this destination node is smaller than what we're currently tracking update to this shorter path
				// else ignore as we have a shorter distance already
				if edge.Weight+distances[v].Distance < distances[edge.Destination.Name].Distance {
					distances[edge.Destination.Name] = Distance{
						Distance: edge.Weight + distances[v].Distance,
						Visited:  distances[edge.Destination.Name].Visited,
					}
				}

			}
		}
	}

	// Select the node closest to the source node based on the currently known distances
	shortest := shortestDistanceNode(distances)

	// Mark this shortest node as visited
	distances[shortest] = Distance{
		Distance: distances[shortest].Distance,
		Visited:  true,
	}

	// Add it to the path as a node that has already been checked for the smallest weight
	*path = append(*path, shortest)
}

func workflow(nodes map[string]Node, source string, destination string) int {
	// Simple check on if both the source and destination exist within the nodemap
	_, ok := nodes[source]
	_, ok2 := nodes[destination]
	if !ok || !ok2 {
		fmt.Println("Source or detination not available in provided node map")
		return -1
	}

	// Create the map of Distances. This tracks the distance from the source to the all other nodes through the execution
	distances := distancemap{}
	path := []string{source}
	for k, _ := range nodes {
		// As this is the source it will have a distance of 0 and will be marked as visited
		if k == source {
			distances[k] = Distance{
				Distance: 0,
				Visited:  true,
			}
		} else {
			// TODO remove the arbitary large number method
			// For all other nodes defaults are set
			distances[k] = Distance{
				Distance: 99999,
			}
		}
	}

	// This will loop through the algorthm to ensure all the nodes have a distance value from the source node
	// As each node needs to be in the path for the algorithm to end this simply checks the length of the path versuses the number of nodes
	// TODO improve this especially for when a node may not be able to reach all other nodes
	for i := 0; i < len(nodes); i++ {
		dijkstra(nodes, distances, &path)
	}

	// TODO find where the mysterious node is added and remove
	distances.print()

	return distances[destination].Distance
}

func main() {
	output := ReadString("./test_data/example.csv")
	// TODO Removing the last char due to a blank line. fix this so I dont have to do this
	nodes := generateNodeMap(output[0 : len(output)-1])

	source := "Node0"
	destination := "Node6"
	fmt.Printf("Distance between %s and %s is: %d\n", source, destination, workflow(nodes, source, destination))

	createDiagram(nodes)
}
