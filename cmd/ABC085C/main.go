package main

import (
	"fmt"
	"time"
)

func main() {
	var n, t int
	fmt.Scan(&n, &t)
	now := time.Now()
	for x := 0; x <= n; x++ {
		for y := 0; y <= n-x; y++ {
			for z := 0; z <= n-x-y; z++ {
				if (10000*x+5000*y+1000*z == t) && (x+y+z == n) {
					fmt.Println(x, y, z)
					goto Z
				}
			}
		}
	}
	fmt.Println("-1 -1 -1")
Z:
	fmt.Printf("%vms\n", time.Since(now).Milliseconds())
}
