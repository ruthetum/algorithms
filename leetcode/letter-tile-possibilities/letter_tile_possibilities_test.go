package letter_tile_possibilities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumTilePossibilities(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		input := "AAB"
		actual := numTilePossibilities(input)
		expected := 8
		assert.Equal(t, expected, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		input := "AAABBC"
		actual := numTilePossibilities(input)
		expected := 188
		assert.Equal(t, expected, actual)
	})

	t.Run("case 3", func(t *testing.T) {
		input := "V"
		actual := numTilePossibilities(input)
		expected := 1
		assert.Equal(t, expected, actual)
	})
}
