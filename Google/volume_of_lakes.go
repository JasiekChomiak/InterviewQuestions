// Link: https://techdevguide.withgoogle.com/resources/former-interview-question-volume-of-lakes/#!
// Also can be found at: https://leetcode.com/problems/trapping-rain-water/submissions/
package Google

func intMin(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func volume(heights []int) int {
	n := len(heights)
	prefixMax := make([]int, n + 1)
	suffixMax := make([]int, n + 1)

	currentMax := 0
	for i := 0; i < n; i++ {
		if heights[i] > currentMax {
			currentMax = heights[i]
		}

		prefixMax[i + 1] = currentMax
	}

	currentMax = 0
	for i := n - 1; i >= 0; i-- {
		if heights[i] > currentMax {
			currentMax = heights[i]
		}

		suffixMax[i] = currentMax
	}

	totalVolume := 0
	for i, v := range heights {
		if prefixMax[i] > v && v < suffixMax[i] {
			totalVolume += intMin(prefixMax[i], suffixMax[i]) - v
		}
	}

	return totalVolume
}

func main() {
	x := volume([]int{1,3,2,4,1,3,1,4,5,2,2,1,4,2,2})
	println(x)
}