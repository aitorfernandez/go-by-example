package main

import "fmt"

type vec struct {
	X, Y float64
}

type Ball struct {
	Pos, Speed vec
}

func (b *Ball) update() {
	if b.Pos.X < 0 {
		b.Pos.X = 0
	} else if b.Pos.X > 200 {
		b.Speed.X = 200
	}

	b.Speed.X *= -1

	b.Pos.X *= b.Speed.X
}

func update(b *Ball) {
	// ...
}
func (b *Ball) draw() {
	fmt.Printf("Pos: %v, Vel: %v\n", b.Pos, b.Speed)
}

func main() {
	// ball := Ball{Pos: vec{X: 1, Y: 2}, Speed: vec{X: 3, Y: 4}}
	// or
	ball := Ball{Pos: vec{1, 2}, Speed: vec{3, 4}}
	tick := 0
	for tick < 10 {
		ball.update()
		ball.draw()

		tick++
	}

	// without methods, pointer and function
	b := Ball{}
	update(&b)
}
