package algorithms

import "fmt"

// Assume that most graphs use in the course are parse graph,
// we'll use Adjacent List to present a graph

// Ex:
// "A": ["B", "C"]
// "B": ["A", "D"]
// "C": ["A", "D"]
// "D": ["B", "C"]

type UGraph map[string][]string

func (ug UGraph) InsertNode(node string, connections ...string) {
	ug[node] = connections
	l := len(connections)

	for i := 0; i < l; i++ {
		if _, exist := ug[connections[i]]; !exist {
			fmt.Println("Node " + connections[i] + " does not exist in current graph, deprecating...")

			connections[l-1], connections[i] = connections[i], connections[l-1]
			connections = connections[:l-1]
			i--
			l--

			fmt.Println("Deprecated")
		} else {
			ug[connections[i]] = append(ug[connections[i]], node)
		}
	}
}

type DGraph map[string][]string

func (dg DGraph) InsertNode(node string, connections ...string) {
	dg[node] = connections
	l := len(connections)

	for i := 0; i < l; i++ {
		if _, exist := dg[connections[i]]; !exist {
			fmt.Println("Node " + connections[i] + " does not exist in current graph, deprecating...")

			connections[l-1], connections[i] = connections[i], connections[l-1]
			connections = connections[:l-1]
			i--
			l--

			fmt.Println("Deprecated")
		}
	}
}
