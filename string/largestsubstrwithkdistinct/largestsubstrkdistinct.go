package largestsubstrwithkdistinct

func FindLargest(s string, k int) int {
	if len(s) == 0 || len(s) < k {
		return 0
	}
	m := map[byte]int{}
	start, length := 0, 0

	for end := 0; end < len(s); end++ {
		ch := s[end]
		if _, ok := m[ch]; ok {
			m[ch]++
		} else {
			m[ch] = 1
		}
		for len(m) > k {
			c := s[start]
			if _, ok := m[c]; ok {
				m[c]--
			}
			if _, ok := m[c]; ok && m[c] == 0 {
				delete(m, c)
			}
			start++
		}

		length = max(length, end-start+1)

	}
	return length
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
