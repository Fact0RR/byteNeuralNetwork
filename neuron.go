package main

import (
	"encoding/json"

	"math/rand"
)

type Neuron struct {
	w   []int8
	out int8
}

func (n *Neuron) New(out int8, sizeOfNextFiel int) {
	n.w = make([]int8, sizeOfNextFiel)
	n.out = out
}

func (n *Neuron) setRandomW() {
	for i := 0; i < len(n.w); i++ {
		n.w[i] = int8(rand.Intn(127+128) - 128)
	}
}

func (n *Neuron) getWString() string {
	//return string(n.w)
	s, _ := json.Marshal(n.w)

	return string(s)
}
