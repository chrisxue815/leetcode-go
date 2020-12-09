package problem

import (
	"sort"
	"strconv"
)

type IndexedNum struct {
	index int
	num   int
}

// O(nlog(n)) time. O(log(n)) space. Built-in quicksort.
func findRelativeRanks(nums []int) []string {
	indexedNums := make([]IndexedNum, len(nums))

	for index, num := range nums {
		indexedNums[index] = IndexedNum{index, num}
	}

	sort.Slice(indexedNums, func(i, j int) bool {
		return indexedNums[i].num > indexedNums[j].num
	})

	result := make([]string, len(nums))
	for rank, indexedNum := range indexedNums {
		var rankStr string
		switch rank {
		case 0:
			rankStr = "Gold Medal"
		case 1:
			rankStr = "Silver Medal"
		case 2:
			rankStr = "Bronze Medal"
		default:
			rankStr = strconv.FormatInt(int64(rank+1), 10)
		}
		result[indexedNum.index] = rankStr
	}

	return result
}
