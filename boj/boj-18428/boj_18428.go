package boj_18428

import (
	"bufio"
	"fmt"
	"os"
)

// https://www.acmicpc.net/problem/18428

type Node struct {
	X int
	Y int
}

var (
	dx      = []int{0, 0, 1, -1}
	dy      = []int{1, -1, 0, 0}
	N       int
	hallway [][]string
	nodes   []Node
	result  bool
)

const (
	wallLimit = 3
	TEACHER   = "T"
	STUDENT   = "S"
	OBSTACLE  = "O"
	EMPTY     = "X"
)

func solution(n int, h [][]string) bool {
	N = n
	hallway = h
	nodes = make([]Node, 0)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if hallway[i][j] == STUDENT {
				nodes = append(nodes, Node{i, j})
			}
		}
	}

	dfs(0)
	return result
}

func dfs(wall int) {
	if wall == wallLimit {
		bfs()
		return
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if hallway[i][j] == EMPTY {
				hallway[i][j] = OBSTACLE
				dfs(wall + 1)
				hallway[i][j] = EMPTY
			}
		}
	}
}

func bfs() {
	queue := make([]Node, 0)

	copyHallway := make([][]string, N)
	check := make([][]bool, N)
	for i := 0; i < N; i++ {
		check[i] = make([]bool, N)
		copyHallway[i] = make([]string, N)

		for j := 0; j < N; j++ {
			copyHallway[i][j] = hallway[i][j]

			if copyHallway[i][j] == TEACHER {
				queue = append(queue, Node{i, j})
				check[i][j] = true
			}
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		x := node.X
		y := node.Y

		for i := 0; i < 4; i++ {
			nx := x + dx[i]
			ny := y + dy[i]

			for inRange(nx, ny) {
				if copyHallway[nx][ny] != OBSTACLE {
					check[nx][ny] = true
					nx += dx[i]
					ny += dy[i]
				} else {
					break
				}
			}
		}
	}

	isDone := true
	for _, node := range nodes {
		if check[node.X][node.Y] == true {
			isDone = false
			break
		}
	}
	if isDone {
		result = isDone
	}
}

func inRange(x, y int) bool {
	return x >= 0 && y >= 0 && x < N && y < N
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n int
	_, _ = fmt.Fscanln(reader, &n)

	h := make([][]string, n)
	for i := 0; i < n; i++ {
		h[i] = make([]string, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Fscan(reader, &h[i][j])
		}
	}

	isAble := solution(n, h)
	result := "NO"
	if isAble {
		result = "YES"
	}
	_, _ = fmt.Fprint(writer, result)
	_ = writer.Flush()
}
