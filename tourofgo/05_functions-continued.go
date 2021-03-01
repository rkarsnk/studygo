package main

import "fmt"

// 2つ以上の引き数が同じ型である場合、型の記述は最後に省略することができる
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(26, 73))
}
