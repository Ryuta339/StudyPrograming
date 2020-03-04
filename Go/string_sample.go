package main

import "fmt"

func main () {
	s := ""
	fmt.Printf ("%v,%v\t%v\n", len(s), &s, s);
	for _, v := range []string{"A", "B", "C"} {
		s += v
		fmt.Printf ("%v,%v\t%v\n", len(s), &s, s);
	}

	s2 := ""
	fmt.Printf ("%v,%v\t%v\n", len(s2), &s2, s2);
	for _, v := range []string{"A", "B", "C"} {
		s2 += v
		fmt.Printf ("%v,%v\t%v\n", len(s2), &s2, s2);
	}
}
