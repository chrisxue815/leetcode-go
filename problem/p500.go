package problem

var keyToRow = map[byte]int{
	'Q': 1,
	'W': 1,
	'E': 1,
	'R': 1,
	'T': 1,
	'Y': 1,
	'U': 1,
	'I': 1,
	'O': 1,
	'P': 1,
	'A': 2,
	'S': 2,
	'D': 2,
	'F': 2,
	'G': 2,
	'H': 2,
	'J': 2,
	'K': 2,
	'L': 2,
	'Z': 3,
	'X': 3,
	'C': 3,
	'V': 3,
	'B': 3,
	'N': 3,
	'M': 3,
}

func toUpper(b byte) byte {
	if b >= 'a' {
		return b - 'a' + 'A'
	}
	return b
}

// O(n) time. O(1) space. Hash table.
func findWords(words []string) []string {
	result := make([]string, 0)

	for _, word := range words {
		if len(word) == 0 {
			continue
		}

		row := keyToRow[toUpper(word[0])]
		valid := true

		for i := 1; i < len(word); i++ {
			if keyToRow[toUpper(word[i])] != row {
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
