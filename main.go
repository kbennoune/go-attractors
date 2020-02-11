package main

import (
	"fmt"
)

func main() {
	p := parameters{}
	setParameters(&p)
	fmt.Printf("Running with params: %+v\n", p)

	writer, finalizer := fileWriter(p.filename)

	defer finalizer()

	initial := iteration{
		x: p.x0,
		y: p.y0,
		t: 0,
	}

	calc := calculation{last: initial, parameters: p}

	for n := 1; n <= p.nMax; n++ {
		calc.next(writer)
	}
}
