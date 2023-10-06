package main

type Ball struct {
	vX, vY, x, y float64
}

func (b *Ball) createBall(x, y, vX, vY float64) {
	b.vX = vX
	b.vY = vY
	b.x = x
	b.y = y
}

func (b *Ball) move() {
	b.x = b.x + b.vX
	b.y = b.y + b.vY
}
