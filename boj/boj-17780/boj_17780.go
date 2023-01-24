package boj_17780

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	X         int
	Y         int
	Direction int
}

const (
	white = iota
	red
	blue
)

var reverse = []int{1, 0, 3, 2}

var direction = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

const countLimit = 1000

func Solution(n, k int, pos, pieces [][]int) int {
	nodes := make([]Node, 0)

	var depth [][][]int
	for i := 0; i < n; i++ {
		depth = append(depth, make([][]int, 0))
		for j := 0; j < n; j++ {
			depth[i] = append(depth[i], make([]int, 0))
		}
	}

	for i, piece := range pieces {
		x := piece[0] - 1
		y := piece[1] - 1
		dir := piece[2] - 1

		node := Node{x, y, dir}
		nodes = append(nodes, node)
		depth[x][y] = append(depth[x][y], i)
	}

	return move(n, k, pos, nodes, depth, 0)
}

func move(n, length int, pos [][]int, nodes []Node, depth [][][]int, count int) int {
	if count > countLimit {
		return -1
	}

	if isDone() {
		return count
	}

	for i := 0; i < length; i++ {
		node := nodes[i]
		x := node.X
		y := node.Y
		dir :=

		if depth[x][y][0] != i {
			continue
		}

		nx := x + direction[dir][0]
		ny := y + direction[dir][1]

		if inPos(nx, ny, n) || pos[nx][ny] == blue {
			node.Direction = reverse[node.Direction]
			nx = x + direction[dir][0]
			ny = y + direction[dir][1]
		}

		if inPos(nx, ny, n) || pos[nx][ny] == blue {
			continue
		}

		if pos[nx][ny] == red {
			for


			for (int j = state[r][c].size() - 1; j >= 0; j--) {
				int tmp = state[r][c].get(j);
				state[nr][nc].add(tmp);
				horses[tmp].r = nr;
				horses[tmp].c = nc;
			}
			state[r][c].clear();
		} else {
			// 모든 말이 이동
			for (int j = 0; j < state[r][c].size(); j++) {
				int tmp = state[r][c].get(j);
				state[nr][nc].add(tmp);
				horses[tmp].r = nr;
				horses[tmp].c = nc;
			}
			state[r][c].clear();
		}

		// 이동한 곳에 말이 4개 이상 있는가?
		if (state[nr][nc].size() >= 4) {
			flag = false;
			break;
		}
	}

}

func inPos(x, y, n int) bool {
	if x >= 0 && y >= 0 && x < n && y < n {
		return true
	}
	return false
}

func isDone(pieces [][]int) bool {
	x := pieces[0][0]
	y := pieces[0][1]

	for _, piece := range pieces {
		if x != piece[0] || y != piece[1] {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, k int
	_, _ = fmt.Fscanln(reader, &n, &k)

	var p string
	pos := make([][]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscanln(reader, &p)
		row := strings.Split(p, " ")
		for _, r := range row {
			v, _ := strconv.Atoi(r)
			values := make([]int, 0)
			values = append(values, v)
			pos[i] = values
		}
	}

	pieces := make([][]int, k)
	for i := 0; i < k; i++ {
		_, _ = fmt.Fscanln(reader, &p)
		row := strings.Split(p, " ")
		for _, r := range row {
			v, _ := strconv.Atoi(r)
			values := make([]int, 0)
			values = append(values, v)
			pieces[i] = values
		}
	}

	count := Solution(n, k, pos, pieces)
	_, _ = fmt.Fprint(writer, count)
	_ = writer.Flush()
}
