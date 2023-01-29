package divide_two_integers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivide(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		actual := divide(10, 3)
		expected := 3
		assert.Equal(t, expected, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		actual := divide(7, -3)
		expected := -2
		assert.Equal(t, expected, actual)
	})

	t.Run("case 3", func(t *testing.T) {
		actual := divide(-2147483648, -1)
		expected := 2147483647
		assert.Equal(t, expected, actual)
	})
}
