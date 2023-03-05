package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
gpprog氏のコードから持ってきた
問題はABC086C - Traveling

	func main(){
		 for i := 0; i < n; i++ {
	    	t, x, y := rl(), rl(), rl()
			...
		 }
		 のように使う
	}
*/
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	buf := make([]byte, 0)
	sc.Buffer(buf, 131072)
}

func rl() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

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

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func Sqrt(i int) int {
	return int(math.Sqrt(float64(i)))
}

func Permutation(n int, k int) int {
	if k > n || k <= 0 {
		panic(fmt.Sprintf("invalid param n:%v k:%v", n, k))
	}
	v := 1
	for i := 0; i < k; i++ {
		v *= (n - i)
	}
	return v
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func Divisor(n int) ([]int, map[int]int) {
	sqrtn := int(math.Sqrt(float64(n)))
	c := 2
	divisor := []int{}
	divisorm := make(map[int]int)
	for {
		if n%2 != 0 {
			break
		}
		divisor = append(divisor, 2)
		divisorm[2]++
		n /= 2
	}
	c = 3
	for {
		if n%c == 0 {
			divisor = append(divisor, c)
			divisorm[c]++
			n /= c
		} else {
			c += 2
			if c > sqrtn {
				break
			}
		}
	}
	if n != 1 {
		divisor = append(divisor, n)
		divisorm[n]++
	}
	return divisor, divisorm
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
	for Abs(ok-ng) > 1 {
		mid := (ok + ng) / 2
		if f(mid) {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func Hasbit(a int, n int) bool {
	return (a>>uint(n))&1 == 1
}

func Nthbit(a int, n int) int {
	return int((a >> uint(n)) & 1)
}

func Popcount(a int) int {
	return bits.OnesCount(uint(a))
}

func Bitlen(a int) int {
	return bits.Len(uint(a))
}

func Debugbit(n int) string {
	r := ""
	for i := Bitlen(n) - 1; i >= 0; i-- {
		if n&(1<<i) != 0 {
			r += "1"
		} else {
			r += "0"
		}
	}
	return r
}

func ToLowerCase(s string) string {
	return strings.ToLower(s)
}

func ToUpperCase(s string) string {
	return strings.ToUpper(s)
}

func IsLower(b byte) bool {
	return 'a' <= b && b <= 'z'
}

func IsUpper(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

func Sorti(sl []int) {
	sort.Sort(sort.IntSlice(sl))
}

func Sortir(sl []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(sl)))
}

func Sorts(sl []string) {
	sort.Slice(sl, func(i, j int) bool {
		return sl[i] < sl[j]
	})
}

func Reversei(sl []int) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}

func Reverses(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}

func Uniquei(sl []int) []int {
	hist := make(map[int]struct{})
	j := 0
	rsl := make([]int, len(sl))
	for i := 0; i < len(sl); i++ {
		if _, ok := hist[sl[i]]; ok {
			continue
		}
		rsl[j] = sl[i]
		hist[sl[i]] = struct{}{}
		j++
	}
	return rsl[:j]
}

type cusum struct {
	l int
	s []int
}

func newcusum(sl []int) *cusum {
	c := &cusum{}
	c.l = len(sl)
	c.s = make([]int, len(sl)+1)
	for i, v := range sl {
		c.s[i+1] = c.s[i] + v
	}
	return c
}

// get sum f <= i && i <= t
func (c *cusum) getRange(f, t int) int {
	if f > t || f >= c.l {
		return 0
	}
	return c.s[t+1] - c.s[f]
}

// get sum 0 to i
func (c *cusum) get(i int) int {
	return c.s[i+1]
}

func Sli(l, m int, def int) [][]int {
	sl := make([][]int, l)
	for i := 0; i < l; i++ {
		sl[i] = make([]int, m)
		for j := 0; j < m; j++ {
			sl[i][j] = def
		}
	}
	return sl
}
