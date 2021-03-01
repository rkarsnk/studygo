package main

import "fmt"

//goの関数は複数の戻り値を返すことができる
//下の例では、2つのstringを返す
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "Go")
	fmt.Println(a, b)
}
