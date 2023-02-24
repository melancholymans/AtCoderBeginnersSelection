package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	l := len(s)
	c1 := strings.Count(s, "dream")
	c2 := strings.Count(s, "dreamer")
	c3 := strings.Count(s, "erase")
	c4 := strings.Count(s, "eraser")
	if l == c1*5+c2*7+c3*5+c4*6 {
		fmt.Println("YES")
		fmt.Println("C1", c1, "C2", c2, "C3", c3, "C4", c4)
	} else {
		fmt.Println("NO")
		fmt.Println("C1", c1, "C2", c2, "C3", c3, "C4", c4)
	}
}
