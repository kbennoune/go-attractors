package main

import "flag"

type parameters struct {
	a        float64
	b        float64
	c        float64
	d        float64
	x0       float64
	y0       float64
	nMax     int
	filename string
}

func setParameters(p *parameters) {
	aPtr := flag.Float64("a", 1.0, "a parameter")
	bPtr := flag.Float64("b", 1.0, "a parameter")
	cPtr := flag.Float64("c", 1.0, "a parameter")
	dPtr := flag.Float64("d", 1.0, "a parameter")
	x0Ptr := flag.Float64("x", 0.0, "the starting x coordinate")
	y0Ptr := flag.Float64("y", 0.0, "the starting y coordinate")

	var nMax int
	flag.IntVar(&nMax, "n", 2, "the number of steps")

	filenamePtr := flag.String("filename", "data/default.dat", "The filename to store the data to.")
	flag.Parse()

	p.a = *aPtr
	p.b = *bPtr
	p.c = *cPtr
	p.d = *dPtr
	p.x0 = *x0Ptr
	p.y0 = *y0Ptr

	p.nMax = nMax
	p.filename = *filenamePtr
}
