package main

import "math"

type calculation struct {
	last iteration
	parameters
}

type iteration struct {
	x float64
	y float64
	t int
}

func (calc *calculation) next(writer func(calculation)) {
	params := calc.parameters
	last := calc.last

	next := iteration{t: last.t + 1}
	next.x = math.Sin(params.b*last.y) + params.c*math.Sin(params.b*last.x)
	next.y = math.Sin(params.a*last.x) + params.d*math.Sin(params.a*last.y)

	calc.last = next

	writer(*calc)
}
