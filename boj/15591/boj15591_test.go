package main

import (
	"fmt"
	"testing"
)

func TestBFS(t *testing.T) {
	// given
	N := 4

	graph = make([]Graph, N+1)
	for i := range graph {
		graph[i] = Graph{
			Nodes: []Node{},
		}
	}

	inputs1 := []struct {
		p int
		q int
		r int
	}{
		{p: 1, q: 2, r: 3},
		{p: 2, q: 3, r: 2},
		{p: 2, q: 4, r: 4},
	}

	for _, item := range inputs1 {
		graph[item.p].Nodes = append(graph[item.p].Nodes, Node{item.q, item.r})
		graph[item.q].Nodes = append(graph[item.q].Nodes, Node{item.p, item.r})
	}

	inputs2 := []struct {
		k int
		v int
	}{
		{k: 1, v: 2},
		{k: 4, v: 1},
		{k: 3, v: 1},
	}

	// when
	actual := make([]int, 0)
	for _, item := range inputs2 {
		visited = make([]bool, N+1)
		count := BFS(item.k, item.v)
		actual = append(actual, count)
	}

	// then
	expected := []int{3, 0, 2}
	if fmt.Sprintf("%v", actual) == fmt.Sprintf("%v", expected) {
		t.Log("success")
	} else {
		t.Errorf("failed, expected=%v, actual=%v", expected, actual)
	}
}
