package main

type Racket struct {
	x, y, size int
	rotation   bool
}

func (r *Racket) createRacket(x, y, size int, rotation bool) {
	r.x = x
	r.y = y
	r.size = size
	r.rotation = true
}
