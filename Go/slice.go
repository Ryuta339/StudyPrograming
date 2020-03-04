package main

import (
	"fmt"
)

func main () {
	/* A */
	a := [4]int{0,1,2,3}
	s := a[1:3]
	fmt.Println (len(s))
	fmt.Println (cap(s))
	
	fmt.Println (a)
	fmt.Println (s)

	/* B */
	s = append (s, 4)
	fmt.Println (len(s))
	fmt.Println (cap(s))
	s[1] = 10
	a[0] = 9
	
	fmt.Println (a)
	fmt.Println (s)

	/* C */
	s2 := a[:]
	s2 = append (s2, 4)
	a[0] = 8
	fmt.Println (s2)
}
