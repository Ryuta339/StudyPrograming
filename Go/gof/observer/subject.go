package main

import (
	"math/rand"
	"time"
//	"fmt"
)



/* ================================ */

type NumberGenerator struct {
	number int
	observers []Observer
}

func (ng *NumberGenerator) AddObserver (observer Observer) {
	ng.observers = append (ng.observers, observer)
}

func (ng *NumberGenerator) DeleteObserver (observer Observer) {
	var i int = 0
	for i<len(ng.observers) && observer!=ng.observers[i] {
		i ++;
	}
	// does not exist
	if i>=len(ng.observers) {
		return;
	}
	s := append (ng.observers[:i], ng.observers[i+1:]...)
	ng.observers = make ([]Observer, len(s))
	copy (ng.observers, s);
}

func (ng *NumberGenerator) NotifyObservers () {
	for _, observer := range ng.observers {
		observer.Update (ng)
	}
}

func (ng *NumberGenerator) GetNumber () int {
	return ng.number;
}


/* ================================ */

type RandomNumberGenerator struct {
	*NumberGenerator
}

func (rng *RandomNumberGenerator) Execute () {
	for i:=0; i<50; i++ {
		rand.Seed (time.Now().UnixNano())
		rng.number = rand.Intn(50);
		rng.NotifyObservers ()
	}
}

func NewRandomNumberGenerator () *RandomNumberGenerator {
	return &RandomNumberGenerator {
		NumberGenerator: &NumberGenerator {number: 0},
	}
}


