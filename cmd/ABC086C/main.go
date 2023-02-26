package main

import (
	"fmt"
)

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func main() {
	var n int
	t := make([]int, 110000)
	x := make([]int, 110000)
	y := make([]int, 110000)
	fmt.Scan(&n)
	for i := 1; i < n+1; i++ {
		fmt.Scan(&t[i], &x[i], &y[i])
	}
	flag := true
	for i := 0; i < n; i++ {
		dt := t[i+1] - t[i]
		dst := abs(x[i+1]-x[i]) + abs(y[i+1]-y[i])
		if dst > dt {
			flag = false
		}
		if dst%2 != dt%2 {
			flag = false
		}
	}
	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
