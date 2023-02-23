/*
package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)
*/
/*
func main() {
	var sc string
	fmt.Scan(&sc)
	sle := strings.Split(sc, "")
	fmt.Println(sle)
	len := len(sle)
	var count int = 0
	for i := 0; i < len; i++ {
		val, _ := strconv.Atoi(sle[i])
		count += val
	}
	fmt.Println(count)

	ss := "12"
	fmt.Println(ss[0])
	fmt.Println(reflect.TypeOf(ss))
}
*/
package main

import (
	"fmt"
)

func removeAllPiens(word string) (noPienWord rune) {
	for _, char := range word {
		if char == '🥺' { // Error!!
			continue
		}
		noPienWord += char // Error!!
	}
	return
}

func main() {
	result := removeAllPiens("🥺fizz🥺buzz🥺")
	fmt.Println(result)
	fmt.Printf("%c \n", result)
}
