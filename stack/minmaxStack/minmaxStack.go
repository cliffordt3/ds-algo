package minmaxStack

type MinMaxStack struct {
	stack       []int
	minMaxStack []entry
}

type entry struct {
	min int
	max int
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{
		stack:       []int{},
		minMaxStack: []entry{},
	}
}

func (stack *MinMaxStack) Peek() int {
	return stack.stack[len(stack.stack)-1]
}

func (stack *MinMaxStack) Pop() int {
	stack.minMaxStack = stack.minMaxStack[:len(stack.minMaxStack)-1]
	val := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return val
}

func (stack *MinMaxStack) Push(number int) {
	newMinMax := entry{
		min: number,
		max: number,
	}
	if len(stack.minMaxStack) > 0 {
		lastMinMax := stack.minMaxStack[len(stack.minMaxStack)-1]
		newMinMax.min = min(lastMinMax.min, number)
		newMinMax.max = max(lastMinMax.max, number)

	}
	stack.minMaxStack = append(stack.minMaxStack, newMinMax)
	stack.stack = append(stack.stack, number)

}

func (stack *MinMaxStack) GetMin() int {
	return stack.minMaxStack[len(stack.minMaxStack)-1].min
}

func (stack *MinMaxStack) GetMax() int {
	return stack.minMaxStack[len(stack.minMaxStack)-1].max
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
