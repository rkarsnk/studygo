package main

import "fmt"

//int型変数iを1、jを2で初期化
var i, j int = 1, 2

func main() {
	//型を省略できる
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
