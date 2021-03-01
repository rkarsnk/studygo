package main

//fmt,math/randパッケージのimport
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite Number is", rand.Intn(10))
}
