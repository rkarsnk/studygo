package main

import (
	"fmt"
	"math/cmplx"
)

//varのfactored 宣言
var (
	ToBe      bool       = false
	NormalInt uint32     = 1<<32 - 1
	MaxInt    uint64     = 1<<64 - 1
	z         complex128 = cmplx.Sqrt(-5 + 12i)
)

/*
出力結果
Type: bool Value: false
Type: uint32 Value: 4294967295
Type: uint64 Value: 18446744073709551615
Type: complex128 Value: (2+3i)
*/
func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", NormalInt, NormalInt)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

/*
Go の基本型
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

int, uint, uintptr 型は、32-bitのシステムでは32 bitで、64-bitのシステムでは64 bitです。

byte // uint8 の別名

rune // int32 の別名
     // Unicode のコードポイントを表す

float32 float64

complex64 complex128
*/
