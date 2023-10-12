package main

import (
	"fmt"
	"math"
	"strconv"
)

func angkaNol(num int) bool {
	numStr := strconv.Itoa(num)
	for _, char := range numStr {
		if char == '0' {
			return true
		}
	}
	return false
}

func primeNumber(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(num))); i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func right(num int) bool {
	const code = 333
	lastThreeDigits := num % 1000
	return lastThreeDigits == code
}

func left(num int) bool {
	lastDigit := num % 10
	secondlastDigit := (num / 10) % 10

	return lastDigit == secondlastDigit+1 || lastDigit == secondlastDigit-1
}

func main() {
	var input int
	fmt.Print("Masukkan angka: ")
	_, err := fmt.Scan(&input)

	if err != nil {
		fmt.Println("Input tidak valid.")
		return
	}

	if !angkaNol(input) {
		if primeNumber(input) {
			if left(input) {
				fmt.Println("LEFT")
			} else {
				if right(input) {
					fmt.Println("RIGHT")
				} else {
					fmt.Println("CENTER")
				}
			}
		} else {
			fmt.Println("DEAD")
		}
	} else {
		fmt.Println("DEAD")
	}
}
