package main

//import( ~ )のことを、factored import statementと呼ぶ
//import "fmt"
//import "math"
//書くこともできるが、factoredの方がより良い
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
