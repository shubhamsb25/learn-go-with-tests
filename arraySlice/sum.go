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

func SumAllTails(numberCollections ...[]int) []int {
	tailSums := []int{}

	for _, numbers := range numberCollections {
		var tail int
		if len(numbers) == 0 {
			tail = 0
		} else {
			tail = Sum(numbers[1:])
		}
		tailSums = append(tailSums, tail)
	}

	return tailSums
}
