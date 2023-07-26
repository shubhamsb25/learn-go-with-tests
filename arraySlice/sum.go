package main

func Sum(numbers []int) int {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numberCollections ...[]int) []int {
	sums := []int{}

	for _, numbers := range numberCollections {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
