package main

import "fmt"

// 戻り値にx,yのように名前をつけると、関数の最初で
//x, y int
//同様に扱え、かつreturnで変数を指定しなくてもOK
//Readablityに悪影響あるので、注意
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(23))
}
