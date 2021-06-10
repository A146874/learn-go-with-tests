package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum = sum + number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

// BOOKMARK: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#write-the-test-first-3
