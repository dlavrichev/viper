package main

import (
"fmt"
"math"
)
func add(x, y int) int {
    return x+y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool = true, false, true

func main() {
    fmt.Println("Hello, Image")
    fmt.Println(math.Pi)
    fmt.Println(add(24, 7))
    a, b := swap("hello", "world")
    fmt.Println(a, b)
    fmt.Println(split(17))
    var i int = 5
    fmt.Println(i, c, python, java)
    var dec, big, small = true, false, "no!"
    fmt.Println(dec, big, small)
    k := 5 // short assignment
    println(k)

    //basic types: bool, string,
    //int  int8  int16  int32  int64
    //uint uint8 uint16 uint32 uint64 uintptr
    //byte // alias for uint8
    //rune // alias for int32
     // represents a Unicode code point
    //float32 float64
    //complex64 complex128

    //var i int = 42
    //var f float64 = float64(i)
    //var u uint = uint(f)
       //equals to
    //i := 42
    //f := float64(i)
    //u := uint(f)
}