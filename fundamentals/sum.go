package main

// Sum returns sum of all numbers in the slice
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll returns a slice of sums of its arguments
func SumAll(numbers ...[]int) (sums []int) {
	for _, slice := range numbers {
		sums = append(sums, Sum(slice))
	}
	return sums
}

// SumAllTails returns a sum of tails of each slice
func SumAllTails(numbers ...[]int) (sums []int) {
	for _, slice := range numbers {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
