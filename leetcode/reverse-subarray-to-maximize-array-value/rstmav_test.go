package reverse_subarray_to_maximize_array_value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxValueAfterReverse(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		input := []int{2, 3, 1, 5, 4}
		actual := maxValueAfterReverse(input)
		expected := 10
		assert.Equal(t, expected, actual)
	})

	t.Run("case 2", func(t *testing.T) {
		input := []int{2, 4, 9, 24, 2, 1, 10}
		actual := maxValueAfterReverse(input)
		expected := 68
		assert.Equal(t, expected, actual)
	})
}
