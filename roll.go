package main

import (
	"math/rand"
	"time"
)

type Roll [2]int

func (r *Roll) Roll() {
	r[0] = rand.Intn(6) + 1
	r[1] = rand.Intn(6) + 1
}

func (r *Roll) Animate() {
	a, b := r[0], r[1]
	for i := 7; i < 15; i++ {
		r[0], r[1] = i, i
		win.Invalidate()
		time.Sleep(time.Millisecond * 50)
	}
	r[0], r[1] = a, b
	win.Invalidate()
}

func (r *Roll) AnimateRoll() {
	r.Roll()
	r.Animate()
}
