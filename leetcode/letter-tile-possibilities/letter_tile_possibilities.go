package letter_tile_possibilities

// https://leetcode.com/problems/letter-tile-possibilities/

const (
	letterSize = 26
	asciiOfA   = int('A')
)

func numTilePossibilities(tiles string) int {
	arr := make([]int, letterSize)

	charArr := []rune(tiles)
	for _, char := range charArr {
		ascii := int(char)
		arr[ascii-asciiOfA]++
	}
	return dfs(arr)
}

func dfs(arr []int) int {
	sum := 0
	for i := 0; i < letterSize; i++ {
		if arr[i] == 0 {
			continue
		}
		sum++
		arr[i]--
		sum += dfs(arr)
		arr[i]++
	}
	return sum
}
