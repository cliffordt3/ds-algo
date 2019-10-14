package minNumberOfCoinsForChange

import (
	"math"
)

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	// Write your code here.
	numCoins := make([]int, n+1)
	for i := range numCoins {
		numCoins[i] = math.MaxInt32
	}
	numCoins[0] = 0

	for _, denom := range denoms {
		for amount := range numCoins {
			if denom <= amount {
				numCoins[amount] = min(numCoins[amount], 1+numCoins[amount-denom])
			}
		}
	}
	if numCoins[n] == math.MaxInt32 {
		return -1
	}
	return numCoins[n]
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
