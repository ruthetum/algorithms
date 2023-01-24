package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCases struct {
	N      int
	K      int
	Pos    [][]int
	Pieces [][]int
	Result int
}

func TestSolution(t *testing.T) {

	testCases := []TestCases{
		{
			N: 4,
			K: 4,
			Pos: [][]int{
				{0, 0, 2, 0},
				{0, 0, 1, 0},
				{0, 0, 1, 2},
				{0, 2, 0, 0},
			},
			Pieces: [][]int{
				{2, 1, 1},
				{3, 2, 3},
				{2, 2, 1},
				{4, 1, 2},
			},
			Result: -1,
		},
		{
			N: 4,
			K: 4,
			Pos: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			Pieces: [][]int{
				{1, 1, 1},
				{1, 2, 1},
				{1, 3, 1},
				{1, 4, 2},
			},
			Result: 1,
		},
		{
			N: 4,
			K: 4,
			Pos: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			Pieces: [][]int{
				{1, 1, 1},
				{1, 2, 1},
				{1, 3, 1},
				{2, 4, 3},
			},
			Result: 1,
		},
		{
			N: 4,
			K: 4,
			Pos: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			Pieces: [][]int{
				{1, 1, 1},
				{1, 2, 1},
				{1, 3, 1},
				{3, 3, 3},
			},
			Result: 2,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t,
			testCase.Result,
			Solution(testCase.N, testCase.K, testCase.Pos, testCase.Pieces))
	}
}
