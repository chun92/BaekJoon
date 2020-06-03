package main

import "fmt"

func main() {
	var x, y int
	fmt.Scanf("%d %d", &x, &y)
	fmt.Printf("%g", float64(x)/float64(y))
}
