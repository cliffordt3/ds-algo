package productsumspecialarray

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.

func ProductSum(array []interface{}) int {
	return helper(array, 1)
}

func helper(array SpecialArray, multiplier int) int {
	sum := 0
	for _, elem := range array {
		if cast, ok := elem.(SpecialArray); ok {
			sum += helper(cast, multiplier+1)
		} else if cast, ok := elem.(int); ok {
			sum += cast
		}

	}
	return sum
}
