package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.acmicpc.net/problem/15591
//
//keyword
// - Spanning Tree, MST, Kruskal, BFS

var (
	graph   []Graph
	visited []bool
	writer  *bufio.Writer
)

type Graph struct {
	Nodes []Node
}

type Node struct {
	Index int
	Value int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)

	var N, Q, p, q, r, k, v int
	fmt.Fscanln(reader, &N, &Q)

	graph = make([]Graph, N+1)
	for i := range graph {
		graph[i] = Graph{
			Nodes: []Node{},
		}
	}

	for i := 0; i < N-1; i++ {
		fmt.Fscanln(reader, &p, &q, &r)
		graph[p].Nodes = append(graph[p].Nodes, Node{q, r})
		graph[q].Nodes = append(graph[q].Nodes, Node{p, r})
	}

	for i := 0; i < Q; i++ {
		fmt.Fscanln(reader, &k, &v)
		visited = make([]bool, N+1)
		count := BFS(k, v)
		fmt.Fprintf(writer, "%d\n", count)
	}
	writer.Flush()
}

func BFS(k, v int) int {
	count := 0
	visited[v] = true
	queue := []int{v}

	for len(queue) != 0 {
		front := queue[0]
		queue = queue[1:]

		nodes := graph[front].Nodes

		for i := 0; i < len(nodes); i++ {
			if !visited[nodes[i].Index] && nodes[i].Value >= k {
				visited[nodes[i].Index] = true
				queue = append(queue, nodes[i].Index)
				count++
			}
		}
	}
	return count
}
