package main

import (
	"fmt"
//	"math"
//	"math/rand"
)

// func main () {
///	rand.Seed (12345);
///	fmt.Println ("My favorite number is", rand.Intn (10));
///}

func add (x, y int) int {
	return x+y
}

func swap (x, y string) (string, string) {
	return y,x
}

func split (sum int) (x, y int) {
	x = sum * 4 / 9;
	y = sum - x
	return
}

var c, python, java = true, false, "no"
// c2, python2, java2 := true, false, "no"		// Don't arrow outside a function

func main () {
//	fmt.Printf ("Now you have %g problems.\n", math.Sqrt (7));
//	fmt.Println ( math.Pi);
//	fmt.Println(add (42,31))
//	a, b := swap ("world", "hello");
//	fmt.Println (a,b);
//	fmt.Println (split (43));
	var i int
	j := 5
	c2, python2, java2 := true, false, "no"
	fmt.Println (i, c, python, java)
	fmt.Println (j, c2, python2, java2)
}
