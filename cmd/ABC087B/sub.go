package main

import (
	"fmt"
)

var a, b, c, x int
var count int

func main() {
	fmt.Scan(&a, &b, &c, &x)
	count = 0
	m := make(map[string]int)

	for i := 0; i <= a; i++ {
		if (x - 500*i) == 0 {
			m["500"] += 1
			count += 1
			break
		}
		loop100(m, x-500*i)
	}
	fmt.Println(count)
	//fmt.Println(m)
}

func loop100(m map[string]int, x int) {
	if x <= 0 {
		return
	}
	for j := 0; j <= b; j++ {
		if (x - 100*j) == 0 {
			m["100"] += 1
			count += 1
			break
		}
		loop50(m, x-100*j)
	}
}

func loop50(m map[string]int, x int) {
	if x <= 0 {
		return
	}
	for k := 0; k <= c; k++ {
		if (x - 50*k) == 0 {
			m["50"] += 1
			count += 1
			break
		}
	}
}
