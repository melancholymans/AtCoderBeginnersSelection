package main

import (
	"fmt"
	"testing"
)

func TestAtoi(t *testing.T) {
	s := "1234"
	got := Atoi(s)
	if got != 1234 {
		t.Errorf("expect 1234, but %d", got)
	}
}

func TestItoa(t *testing.T) {
	n := 1234
	got := Itoa(n)
	if got != "1234" {
		t.Errorf("expect 1234, but %s", got)
	}
}

func TestLowerBound(t *testing.T) {
	s := []int{1, 4, 5, 8, 9, 13, 17}
	i := LowerBound(15, s)
	fmt.Println(i)
}

func TestUpperBound(t *testing.T) {
	s := []int{1, 4, 5, 8, 9, 13, 17}
	i := UpperBound(15, s)
	fmt.Println(i)
}
