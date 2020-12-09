package problem

var keyToRow = []int{3, 4, 4, 3, 2, 3, 3, 3, 2, 3, 3, 3, 4, 4, 2, 2, 2, 2, 3, 2, 2, 4, 2, 4, 2, 4}

func getRow(b byte) int {
	if b >= 'a' {
		return keyToRow[b-'a']
	}
	return keyToRow[b-'A']
}

// O(n) time. O(1) space. Array.
func findWords(words []string) []string {
	result := make([]string, 0, len(words))

	for _, word := range words {
		if len(word) == 0 {
			continue
		}

		row := getRow(word[0])
		valid := true

		for i := 1; i < len(word); i++ {
			if getRow(word[i]) != row {
				valid = false
				break
			}
		}

		if valid {
			result = append(result, word)
		}
	}

	return result
}
