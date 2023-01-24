package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	X         int
	Y         int
	Direction int
}

const (
	nodeX = iota
	nodeY
	nodeDir
)

const (
	white = iota
	red
	blue
)

var (
	dirX    = []int{0, 0, 1, -1}
	dirY    = []int{1, -1, 0, 0}
	reverse = []int{1, 0, 3, 2}
)

const countLimit = 1000

var (
	nodes  []Node
	states [][][]int
)

func Solution(n, k int, pos, pieces [][]int) int {
	states = make([][][]int, len(pos))
	for i, rows := range pos {
		states[i] = make([][]int, len(pos[i]))
		for j, row := range rows {
			states[i][j] = make([]int, 0)
			states[i][j] = append(states[i][j], row)
		}
	}

	for i, piece := range pieces {
		x := piece[nodeX] - 1
		y := piece[nodeY] - 1
		dir := piece[nodeDir]
		nodes = append(nodes, Node{x, y, dir})
		states[x][y] = append(states[x][y], i)
	}

	flag := true
	count := 0
	for flag {
		count++
		if count > countLimit {
			break
		}

		for i := 0; i < k; i++ {
			node := nodes[i]
			x := node.X
			y := node.Y
			dir := node.Direction

			if states[x][y][0] != i {
				continue
			}

			nx := x + dirX[dir]
			ny := y + dirY[dir]

			if !inRange(n, nx, ny) || pos[nx][ny] == blue {
				nDir := reverse[dir]
				nx = x + dirX[nDir]
				ny = y + dirY[nDir]
			}

			if !inRange(n, nx, ny) || pos[nx][ny] == blue {
				continue
			} else if pos[nx][ny] == red {
				for j := len(states[x][y]) - 1; j >= 0; j-- {
					temp := states[x][y][j]
					states[nx][ny] = append(states[nx][ny], temp)
					nodes[temp].X = nx
					nodes[temp].Y = ny
				}
				states[x][y] = make([]int, 0)
			} else {
				for j := 0; j < len(states[x][y]); j++ {
					temp := states[x][y][j]
					states[nx][ny] = append(states[nx][ny], temp)
					nodes[temp].X = nx
					nodes[temp].Y = ny
				}
				states[x][y] = make([]int, 0)
			}

			if len(states[nx][ny]) >= 4 {
				flag = false
				break
			}
		}
	}

	if flag {
		return -1
	}
	return count
}

func inRange(n, x, y int) bool {
	if x >= 0 && y >= 0 && x < n && y < n {
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, k int
	_, _ = fmt.Fscanln(reader, &n, &k)

	pos := make([][]int, n)
	for i := 0; i < n; i++ {
		pos[i] = make([]int, n)
		for j := 0; j < n; j++ {
			_, _ = fmt.Fscan(reader, &pos[i][j])
		}
	}

	pieces := make([][]int, k)
	for i := 0; i < k; i++ {
		pieces[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			_, _ = fmt.Fscan(reader, &pieces[i][j])
		}
	}

	count := Solution(n, k, pos, pieces)
	_, _ = fmt.Fprint(writer, count)
	_ = writer.Flush()
}
