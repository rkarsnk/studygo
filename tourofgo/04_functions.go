package main

import "fmt"

//変数の型名は、変数の後ろに書く
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(27, 73))
}
