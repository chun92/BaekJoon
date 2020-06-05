package main

import (
	"bufio"
	"os"
	"strconv"
)

const NumOfHamburger = 3
const NumOfDrink = 2

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	minHamburger := 2000
	for i := 0; i < NumOfHamburger; i++ {
		in.Scan()
		price, _ := strconv.Atoi(in.Text())
		if price < minHamburger {
			minHamburger = price
		}
	}

	minDrink := 2000
	for i := 0; i < NumOfDrink; i++ {
		in.Scan()
		price, _ := strconv.Atoi(in.Text())
		if price < minDrink {
			minDrink = price
		}
	}

	out.WriteString(strconv.Itoa(minHamburger+minDrink-50) + "\n")
	out.Flush()
}
