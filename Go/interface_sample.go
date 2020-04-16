package main

import "fmt"

type InterfaceA interface {
	doSomething ()
}

type InterfaceB interface {
	doSomething ()
	doNothing ()
}

type StructureA struct {
	val int
}

func (a *StructureA) doSomething () {
	fmt.Println (a.val)
}

type StructureB struct {
	val rune
}

func (b *StructureB) doSomething () {
	fmt.Printf ("%c\n", b.val)
	fmt.Printf ("%c\n", b.val)
}

func (b *StructureB) doNothing () {
}

func main () {
	var ia1, ia2 InterfaceA;
	var ib1 InterfaceB;

	ia1 = &StructureA { val: 0x41 }
	ia2 = &StructureB { val: 0x41 }
	ib1 = &StructureB { val: 0x41 }

	ia1.doSomething ()
	ia2.doSomething ()
	ib1.doSomething ()
}
