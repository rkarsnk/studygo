package main

import (
	"fmt"
	"math"
)

func main() {
	//math.piだとエラーになる、外部パッケージから参照する場合、大文字から始まる
	fmt.Println(math.Pi)
}
