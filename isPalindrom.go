package main

import "fmt"

func IsPalindrom(num int) bool {
	if num < 0 {
		return false
	}

	original := num
	reversed := 0

	for num > 0 {
		fmt.Println("number ", num)
		lastDigit := num % 10
		fmt.Println("last digit ", lastDigit)
		reversed = reversed*10 + lastDigit
		fmt.Println("reverse ", reversed)
		num /= 10
	}

	return original == reversed
}
