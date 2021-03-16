// Link: https://techdevguide.withgoogle.com/resources/former-interview-question-find-longest-word/#!

package Google

func isSubword(positions [][]int, word string) bool {
	index := 0

	for _, c := range word {
		charIndex := c - 'a'
		if positions[charIndex][index] == -1 {
			return false
		}

		index = positions[charIndex][index] + 1
	}

	return true
}

// Denote:
// n - the length of the main word
// m - the number of words in the array
// This solution uses O(n) space (26(n+1) + c) and works in time O(n + m) (26n + m)
// One could also think of a solution where the space is only n, but I think the
// time complexity could become O(nm) then.
func longestWord(word string, words []string) string {
	wordLen := len(word)
	positions := make([][]int, 26)
	currentPositions := make([]int, 26)

	for i := 0; i < 26; i++ {
		positions[i] = make([]int, wordLen + 1)
		for j := 0; j <= wordLen; j++ {
			positions[i][j] = -1
		}
		currentPositions[i] = -1
	}

	for i, c := range word {
		charIndex := c - 'a'
		for j := currentPositions[charIndex] + 1; j <= i; j++ {
			positions[charIndex][j] = i
		}

		currentPositions[charIndex] = i
	}

	maxWord := ""

	for _, w := range words {
		if isSubword(positions, w) {
			if len(w) > len(maxWord) {
				maxWord = w
			}
		}
	}

	return maxWord
}

