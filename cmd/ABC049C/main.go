package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	for {
		if s = strings.Replace(s, "dreamereraser", "", 1); len(s) < 5 {
			break
		}
		if s = strings.Replace(s, "dreamererase", "", 1); len(s) < 5 {
			break
		}
		if s = strings.Replace(s, "dreameraser", "", 1); len(s) < 5 {
			break
		}
		if s = strings.Replace(s, "dreamerase", "", 1); len(s) < 5 {
			break
		}
		if s = strings.Replace(s, "eraseeraser", "", 1); len(s) < 5 {
			break
		}
		if s = strings.Replace(s, "eraseerase", "", 1); len(s) < 5 {
			break
		}
		if s = strings.Replace(s, "erasereraser", "", 1); len(s) < 5 {
			break
		}
		if s = strings.Replace(s, "erasererase", "", 1); len(s) < 5 {
			break
		}
		if s = strings.Replace(s, "dreamer", "", 1); len(s) < 5 {
			break
		}
		if s = strings.Replace(s, "eraser", "", 1); len(s) < 5 {
			break
		}
		if s = strings.Replace(s, "dream", "", 1); len(s) < 5 {
			break
		}
		if s = strings.Replace(s, "erase", "", 1); len(s) < 5 {
			break
		}
	}
	if len(s) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
