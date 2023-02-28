package main

import "strconv"

func Atoi(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		panic(e)
	}
	return i
}

func Itoa(i int) string {
	return strconv.Itoa(i)
}

func LowerBound(v int, sl []int) int {
	if len(sl) == 0 {
		return -1
	}
	idx := bs(0, len(sl)-1, func(c int) bool {
		return sl[c] < v
	})
	return idx
}

func UpperBound(v int, sl []int) int {
	if len(sl) == 0 {
		return 0
	}
	idx := bs(0, len(sl)-1, func(c int) bool {
		return sl[c] <= v
	})
	return idx + 1
}

func bs(ok, ng int, f func(int) bool) int {
	if !f(ok) {
		return -1
	}
	if f(ng) {
		return ng
	}
	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2
		if f(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
