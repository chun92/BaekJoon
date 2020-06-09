package main

import (
	"fmt"
)

func main() {
	var c rune
	fmt.Scanf("%c", &c)
	ascii := int(c)
	fmt.Println(ascii)
}
