package main

import (
	"fmt"
	"math/rand"
	"my/meter"
	"time"
)

type sdfs struct {
	meter.Meter
}

func (a *sdfs) add() {
	for {
		//a.Count(1)
		a.Count(rand.Int63n(10))
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {

	a := sdfs{meter.New(1)}

	go a.add()

	go a.add()

	go a.add()

	for {
		fmt.Println(a.Get())
		time.Sleep(500 * time.Millisecond)
	}
}
