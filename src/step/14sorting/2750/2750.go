package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	in.Split(bufio.ScanWords)
	in.Scan()
	n, _ := strconv.Atoi(in.Text())

	array := make([]int, n)

	for i := 0; i < n; i++ {
		in.Scan()
		num, _ := strconv.Atoi(in.Text())
		array[i] = num
	}

	quickSort(0, n, n, array)

	for i := 0; i < n; i++ {
		out.WriteString(strconv.Itoa(array[i]) + "\n")
	}
	out.Flush()
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func quickSort(left int, right int, n int, array []int) {
	if left == right {
		return
	}

	if left < 0 {
		return
	}

	if right > n {
		return
	}

	pivot := array[left]
	index := left
	for i := left + 1; i < right; i++ {
		if array[i] < pivot {
			for j := i; j > index; j-- {
				swap(&array[j], &array[j-1])
			}
			index++
		}
	}
	quickSort(left, index, n, array)
	quickSort(index+1, right, n, array)
}
