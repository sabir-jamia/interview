package main

import (
	"fmt"
	"interview/pkg/graph"
)

func main() {
	g := graph.New(13)
	g.AddEdge(0, 5)
	g.AddEdge(4, 3)
	g.AddEdge(0, 1)
	g.AddEdge(9, 12)
	g.AddEdge(6, 4)
	g.AddEdge(5, 4)
	g.AddEdge(0, 2)
	g.AddEdge(11, 12)
	g.AddEdge(9, 10)
	g.AddEdge(0, 6)
	g.AddEdge(7, 8)
	g.AddEdge(9, 11)
	g.AddEdge(5, 3)
	dfs := graph.NewDFS(g, 0)
	for v := 0; v < g.V; v++ {
		if dfs.Marked(v) {
			fmt.Print(v, "-->")
		}
	} 

	fmt.Println("")
	fmt.Println(dfs.Count())
}