package main

import (
	"fmt"
	"reflect"
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

func TestMax(t *testing.T) {
	a := 123
	b := 456
	got := Max(a, b)
	if got != 456 {
		t.Errorf("expect 456, but %d", got)
	}
}

func TestMin(t *testing.T) {
	a := 123
	b := 456
	got := Min(a, b)
	if got != 123 {
		t.Errorf("expect 123, but %d", got)
	}
}

func TestAbs(t *testing.T) {
	a := -123
	got := Abs(a)
	if got != 123 {
		t.Errorf("expect 123, but %d", got)
	}
	a = 456
	got = Abs(a)
	if got != 456 {
		t.Errorf("expect 123, but %d", got)
	}
}

func TestPow(t *testing.T) {
	a := -123
	got := Pow(a, 2)
	if got != 15129 {
		t.Errorf("expect 15129, but %d", got)
	}
	a = 456
	got = Pow(a, 3)
	if got != 94818816 {
		t.Errorf("expect 123, but %d", got)
	}
}

func TestSqrt(t *testing.T) {
	a := 16
	got := Sqrt(a)
	if got != 4 {
		t.Errorf("expect 4, but %d", got)
	}
	a = 5
	got = Sqrt(a)
	if got != 2 {
		t.Errorf("expect 2, but %d", got)
	}
}

/*
順列
p(n,k)=n!/(n-k)!
*/
func TestPermutation(t *testing.T) {
	got := Permutation(7, 3)
	if got != 210 {
		t.Errorf("expect 210, but %d", got)
	}
	got = Permutation(3, 1)
	if got != 3 {
		t.Errorf("expect 210, but %d", got)
	}
}

/*
最大公約数
最大公約数とは、すべての公約数を約数にもつ公約数である。
24=2*2*2*3
30=2*3**5
Gcd=2*3=6
24と30を共に割り切れる約数のうち最大の値
*/
func TestGcd(t *testing.T) {
	got := Gcd(24, 30)
	if got != 6 {
		t.Errorf("expect 6, but %d", got)
	}
}

/*
除数
与えられた数の約数を並べて返す。同時にmap形式で約数:その数でも返す
*/
func TestDivisor(t *testing.T) {
	n, m := Divisor(24)
	if !reflect.DeepEqual(n, []int{2, 2, 2, 3}) {
		t.Errorf("expect [2 2 2 3], but %v", n)
	}
	if !reflect.DeepEqual(m, map[int]int{2: 3, 3: 1}) {
		t.Errorf("expect [2 2 2 3], but %v", n)
	}

}

/*
指定した数値より下側のインデックスを返す。指定した数値自身は含まない
V=10を指定した場合index=4(数値9がいる場所)を返す
v=8を指定した場合index=2(数値5がいる場所)を返す
つまり指定した数値未満のindexを返す
*/
func TestLowerBound(t *testing.T) {
	s := []int{1, 4, 5, 8, 9, 13, 17}
	i := LowerBound(8, s)
	fmt.Println(i)
}

/*
指定した数値より上側のインデックスを返す。指定した数値自身は含まない
V=10を指定した場合index=5(数値13がいる場所)を返す
v=8を指定した場合index=4(数値9がいる場所)を返す
つまり指定した数値超のindexを返す
*/
func TestUpperBound(t *testing.T) {
	s := []int{1, 4, 5, 8, 9, 13, 17}
	i := UpperBound(10, s)
	fmt.Println(i)
}

/*
第一引数のビット表現で第二引数の位置にあるビットが１であればtrue0ならfalseを返す
11 = 1011
右から０始まりとする
*/
func TestHasbit(t *testing.T) {
	got := Hasbit(11, 0)
	if got != true {
		t.Errorf("expect 1, but %v", got)
	}
	got = Hasbit(11, 1)
	if got != true {
		t.Errorf("expect 1, but %v", got)
	}
	got = Hasbit(11, 2)
	if got != false {
		t.Errorf("expect 1, but %v", got)
	}
}

/*
第一引数のビット表現で第二引数の位置にあるビットを返す
11 = 1011
右から０始まりとする
*/
func TestNthbit(t *testing.T) {
	got := Nthbit(11, 0)
	if got != 1 {
		t.Errorf("expect 1, but %v", got)
	}
	got = Nthbit(11, 1)
	if got != 1 {
		t.Errorf("expect 1, but %v", got)
	}
	got = Nthbit(11, 2)
	if got != 0 {
		t.Errorf("expect 0, but %v", got)
	}
}

/*
引数を２進数で表現した時立っているビット数を返す
*/
func TestPopcount(t *testing.T) {
	got := Popcount(11)
	if got != 3 {
		t.Errorf("expect 3, but %v", got)
	}
}

/*
引数を２進数表現にしたときの長さを返す
11 = 1011なので長さは４となる
*/
func TestPBitlen(t *testing.T) {
	got := Bitlen(11)
	if got != 4 {
		t.Errorf("expect 4, but %v", got)
	}
}

/*
引数を２進数表現にして文字列で表示する
*/
func TestDebugbit(t *testing.T) {
	got := Debugbit(11)
	fmt.Println(got)
	if got != "1011" {
		t.Errorf("expect 1011, but %s", got)
	}
}

/*
引数文字列を全て小文字にする
*/
func TestToLowerCase(t *testing.T) {
	got := ToLowerCase("AtCoder")
	fmt.Println(got)
	if got != "atcoder" {
		t.Errorf("expect atcoder, but %s", got)
	}
}

/*
引数文字列を全て大文字にする
*/
func TestToUpperCase(t *testing.T) {
	got := ToUpperCase("AtCoder")
	fmt.Println(got)
	if got != "ATCODER" {
		t.Errorf("expect ATCODER, but %s", got)
	}
}

/*
byte型とはuin8の別名で定義されている。つまり１バイトのことである
ファイル等の読み書きなどに出てくる。byte[]型はstringと相互変換できる
golangでは文字列と文字は違う扱いなので注意が必要
*/
/*
引数のbyte型変数が小文字であればtrueを返す
*/
func TestIsLower(t *testing.T) {
	got := IsLower(byte('A'))
	fmt.Println(got)
	if got != false {
		t.Errorf("expect false, but %v", got)
	}
	got = IsLower(byte('a'))
	fmt.Println(got)
	if got != true {
		t.Errorf("expect true, but %v", got)
	}
}

/*
引数のbyte型変数が小文字であればtrueを返す
*/
func TestIsUpper(t *testing.T) {
	got := IsUpper(byte('A'))
	fmt.Println(got)
	if got != true {
		t.Errorf("expect true, but %v", got)
	}
	got = IsUpper(byte('a'))
	fmt.Println(got)
	if got != false {
		t.Errorf("expect false, but %v", got)
	}
}

/*
golangでソートするときはsortパッケージを使う
int型のスライスをソートするときはsort.Ints(昇順)を使う
sort.Ints関数はsort.Sort(sort.IntSlice(s))を読んでいるだけ
逆順にしたいときはsort.Reverseを使う。
これからテストする関数はsort.Intsと同じことをしているだけ
引数として渡す[]intスライスは関数内で変更されるので返り値はない
*/
func TestSorti(t *testing.T) {
	sl := []int{1, 8, 17, 13, 4, 5, -5}
	Sorti(sl)
	if !reflect.DeepEqual(sl, []int{-5, 1, 4, 5, 8, 13, 17}) {
		t.Errorf("expect [-5, 1, 4, 5, 8, 13, 17], but %v", sl)
	}
}

/*
 */
func TestSortir(t *testing.T) {
	sl := []int{1, 8, 17, 13, 4, 5, -5}
	Sortir(sl)
	if !reflect.DeepEqual(sl, []int{17, 13, 8, 5, 4, 1, -5}) {
		t.Errorf("expect [1, 4, 5, 8, 13, 17], but %v", sl)
	}
}

func TestSorts(t *testing.T) {
	sl := []string{"1", "8", "atcoder", "beginner", "selection", "ATCODER", "BEGINNER"}
	Sorts(sl)
	if !reflect.DeepEqual(sl, []string{"1", "8", "ATCODER", "BEGINNER", "atcoder", "beginner", "selection"}) {
		t.Errorf("expect [1, 8, ATCODER, BEGINNER, atcoder, beginner, selection], but %v", sl)
	}
}

/*
int型のスライスを並びの逆順にする（数値の大小ではない）
*/
func TestReversei(t *testing.T) {
	sl := []int{1, 8, 17, 13, 4, 5, -5}
	Reversei(sl)
	if !reflect.DeepEqual(sl, []int{-5, 5, 4, 13, 17, 8, 1}) {
		t.Errorf("expect [-5 5 4 13 17 8 1], but %v", sl)
	}
}

/*
文字列を逆順にする（並びの逆順）
*/
func TestReverses(t *testing.T) {
	sg := "atcoder"
	s := Reverses(sg)
	if s != "redocta" {
		t.Errorf("expect redocta, but %v", s)
	}
}

/*
[]int型のスライスのなかから重複を除いたスライスを返す
*/
func TestUniquei(t *testing.T) {
	sg := []int{5, 8, 17, 13, 4, 5, -5}
	sl := Uniquei(sg)
	if !reflect.DeepEqual(sl, []int{5, 8, 17, 13, 4, -5}) {
		t.Errorf("expect [-5 5 4 13 17 8 1], but %v", sl)
	}
}
