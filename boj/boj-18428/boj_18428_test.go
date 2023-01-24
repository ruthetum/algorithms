package boj_18428

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		n := 5
		hallway := [][]string{
			{"X", "S", "X", "X", "T"},
			{"T", "X", "S", "X", "X"},
			{"X", "X", "X", "X", "X"},
			{"X", "T", "X", "X", "X"},
			{"X", "X", "T", "X", "X"},
		}

		actual := solution(n, hallway)
		assert.Equal(t, true, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		n := 4
		hallway := [][]string{
			{"S", "S", "S", "T"},
			{"X", "X", "X", "X"},
			{"X", "X", "X", "X"},
			{"T", "T", "T", "X"},
		}

		actual := solution(n, hallway)
		assert.Equal(t, false, actual)
	})
}
