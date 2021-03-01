package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string

	//編集に初期値を与えない場合、ゼロ値が与えられる
	//0 0 false ""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
