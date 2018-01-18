package three_sum

import ()

func threeSum(numbers []int) [][]int {
	calculator := TreeSumCalculator{numbers, [][]int{}}
	calculator.CalculateSums()
	return calculator.result
}

type TreeSumCalculator struct {
	numbers []int
	result  [][]int
}

func (self *TreeSumCalculator) CalculateSums() {
	self.calculateSum(0, 1, len(self.numbers)-1)
}

func (self *TreeSumCalculator) calculateSum(index_a, index_b, index_c int) {
	a, b, c := self.numbers[index_a], self.numbers[index_b], self.numbers[index_c]

	sum := a + b + c

	switch {
	case sum == 0:
		self.result = append(self.result, []int{a, b, c})
		return
	}

}
