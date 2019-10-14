package findsmallest

import "math"

func FindSmallestSubString(searchString, patternString string) string {
	patternMap := map[byte]int{}

	for i := range patternString {
		ch := patternString[i]
		patternMap[ch] = 1
	}
	left, right, count := 0, 0, len(patternString)
	length := math.MaxInt32
	smallest := ""
	for right < len(searchString) {
		ch := searchString[right]
		if _, ok := patternMap[ch]; ok {
			count--
			patternMap[ch]--
		}
		right++
		for count == 0 {
			if right-left < length {
				length = right - left
				smallest = searchString[left:right]
			}
			ch := searchString[left]
			if _, ok := patternMap[ch]; ok {
				patternMap[ch]++
				count++
			}
			left++
		}
	}

	return smallest
}
