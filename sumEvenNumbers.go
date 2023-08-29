package main

func SumEvenNumbers(numbers []int) int {
	count := 0
	for _, number := range numbers {
		if number%2 == 0 {
			count += number
		}
	}
	return count
}
