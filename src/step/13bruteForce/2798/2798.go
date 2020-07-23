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
	in.Scan()
	m, _ := strconv.Atoi(in.Text())

	array := make([]int, n)

	for i := 0; i < n; i++ {
		in.Scan()
		num, _ := strconv.Atoi(in.Text())
		array[i] = num
	}
	quickSort(0, n, array)

	res := 0

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				sum := array[i] + array[j] + array[k]
				if sum <= m {
					res = max(res, sum)
				}
			}
		}
	}

	out.WriteString(strconv.Itoa(res) + "\n")
	out.Flush()
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func quickSort(left int, right int, array []int) {
	if left == right {
		return
	} else if left > right {
		return
	} else if left < 0 {
		return
	} else {
		pivot := array[left]
		leftCount := left
		for i := left; i < right; i++ {
			if array[i] < pivot {
				for j := i; j > leftCount; j-- {
					swap(&array[j], &array[j-1])
				}
				leftCount++
			}
		}
		quickSort(left, leftCount, array)
		quickSort(leftCount+1, right, array)
	}
}
