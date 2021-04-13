package main

import (
	"fmt"

	"rsc.io/quote"

	"math/rand"
)

// func add(x int, y int) int {
// you can omit the types all but the last if they are the same
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var jl, rl bool
var i, j = 1, 2

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("My lucky number is: ", rand.Intn(10))

	// call code from an external package
	fmt.Println(quote.Go())

	// call a function with args
	fmt.Println(add(2, 4))

	// multiple results
	a, b := swap("go", "land")
	fmt.Println(a, b)

	// naked return
	fmt.Println(split(17))

	// variables with initializers
	var cl, pl = true, "Oh!"
	k := 3 // short assignment in func only
	fmt.Println(i, j, k, cl, pl, jl, rl)
}
