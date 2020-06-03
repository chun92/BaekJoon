package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	for i := 0; i < 3; i++ {
		fmt.Println(a * getNumOfDigit(b, i))
	}
	fmt.Println(a * b)
}

func getNumOfDigit(num int, digit int) int {
	tempNum := num
	for i := 0; i < digit; i++ {
		tempNum = tempNum / 10
	}
	return tempNum % 10
}
