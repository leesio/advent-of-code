package main

type Index struct {
	input        []int
	sums         map[int]int
	sumsByCursor [][]int
	cursor       int
	stepSize     int
}

func NewIndex(input []int, stepSize int) *Index {
	in := &Index{
		input:        input,
		sums:         make(map[int]int),
		sumsByCursor: make([][]int, len(input)),
		cursor:       0,
		stepSize:     stepSize,
	}
	for i := 0; i < stepSize; i++ {
		for j := i + 1; j < stepSize; j++ {
			sum := input[i] + input[j]
			in.sumsByCursor[i] = append(in.sumsByCursor[i], sum)
			in.sums[sum]++
		}
	}
	in.cursor = stepSize
	return in
}

func (in *Index) window() []int {
	windowStart := in.cursor - in.stepSize
	return in.input[windowStart:in.cursor]
}
func (in *Index) nextWindow() []int {
	windowStart := in.cursor + 1 - in.stepSize
	return in.input[windowStart : in.cursor+1]
}

func (in *Index) oldWindowStartAt() int {
	return in.cursor - in.stepSize
}
func (in *Index) newWindowStartAt() int {
	return in.cursor + 1 - in.stepSize
}

func (in *Index) Step() {
	// remove stale values from the index
	oldNums := in.sumsByCursor[in.oldWindowStartAt()]
	for i := 0; i < len(oldNums); i++ {
		in.sums[oldNums[i]]--
		if in.sums[oldNums[i]] == 0 {
			delete(in.sums, oldNums[i])
		}
	}

	// add new values to index
	nextWindow := in.nextWindow()
	newVal := nextWindow[len(nextWindow)-1]
	for v, val := range nextWindow[:len(nextWindow)-1] {
		sum := val + newVal
		in.sumsByCursor[in.newWindowStartAt()+v] = append(
			in.sumsByCursor[in.newWindowStartAt()+v],
			sum,
		)
		in.sums[sum]++
	}
	in.cursor++
}

func (in *Index) Contains(num int) bool {
	if c, ok := in.sums[num]; ok {
		return c > 0
	}
	return false
}

func PartOneAlternative(input []int, stepSize int) int {
	in := NewIndex(input, stepSize)
	for _, val := range input[stepSize:] {
		if !in.Contains(val) {
			return val
		}
		in.Step()
	}
	return -1
}
