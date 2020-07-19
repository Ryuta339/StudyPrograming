package main

import "fmt"


/* ================================ */

type Observer interface {
	Update (generator *NumberGenerator)
}

/* ================================ */

type DigitObserver struct {
}

func (dgo *DigitObserver) Update (generator *NumberGenerator) {
	fmt.Print ("DegitObserver:")
	fmt.Println (generator.GetNumber())
}

/* ================================ */

type GraphObserver struct {
}

func (gro *GraphObserver) Update (generator *NumberGenerator) {
	fmt.Print ("GraphObserver:")
	count := generator.GetNumber ()
	for i:=0; i<count; i++ {
		fmt.Print ("*")
	}
	fmt.Println ("")
}


