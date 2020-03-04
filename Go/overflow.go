package main

import (
	"fmt"
)

/*
func main () {
	ui_1 := uint32(400000000)
	ui_2 := uint32(4000000000)

	fmt.Printf ("%d + %d = %d\n", ui_1, ui_2, ui_1+ui_2);
}
*/

/*
func integers () func () int {
	i := 0
	return func () int {
		i += 1;
		return i
	}
}

func main () {
	ints := integers ();

	fmt.Println (ints());
	fmt.Println (ints());
	fmt.Println (ints());
	
	otherInts := integers ();
	fmt.Println (otherInts());
}
*/

func sub () {
	for {
		fmt.Println ("sub loop");
	}
}

func main () {
	go sub ();
	for {
		fmt.Println ("main loop");
	}
}
