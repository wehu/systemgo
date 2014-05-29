package systemgo

import (
	"fmt"
)

// DSL

func Initial(body func()) {
	New(body)
}

func Always(es []EventIntf, body func()) {
	New(func(){
		for {
			Wait(es...)
			body()
		}
	})
}

func On(es ...EventIntf) []EventIntf {
	return es
}

func Run(t int, body func()){
	body()
	fmt.Printf("Simulation start\n")
	Simulate(t)
	fmt.Printf("Simulation finished\n")
}
