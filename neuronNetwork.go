package main

import (
	"fmt"
	"os"
)

type NeuronNetwork struct {
	firstLayer  []Neuron
	middleLayer []Neuron
	finalNeuron Neuron
}

func (nn *NeuronNetwork) newNNwithRandom(sizeFirst, sizeMiddle int) {

	nn.firstLayer = make([]Neuron, sizeFirst)
	nn.middleLayer = make([]Neuron, sizeMiddle)

	for i := 0; i < len(nn.firstLayer); i++ {
		nn.firstLayer[i].New(0, sizeMiddle)
		nn.firstLayer[i].setRandomW()
	}
	for i := 0; i < len(nn.middleLayer); i++ {
		nn.middleLayer[i].New(0, 1)
		nn.middleLayer[i].setRandomW()
	}
}

func (nn *NeuronNetwork) newNN() {

	nn.firstLayer = make([]Neuron, 2)
	nn.middleLayer = make([]Neuron, 4)

	nn.firstLayer[0].w = []int8{3, -2, 3, -5}
	nn.firstLayer[1].w = []int8{-3, 2, -3, 5}

	nn.middleLayer[0].w = []int8{-1}
	nn.middleLayer[1].w = []int8{1}
	nn.middleLayer[2].w = []int8{-1}
	nn.middleLayer[3].w = []int8{1}
}

func (nn *NeuronNetwork) pushTheData(ry, by int8) {

	nn.firstLayer[0].out = ry
	nn.firstLayer[1].out = by
	for i := 0; i < len(nn.middleLayer); i++ { // на каждой итерации заполняем нейроны данными

		var sum int8

		for j := 0; j < len(nn.firstLayer); j++ {

			sum += nn.firstLayer[j].out * nn.firstLayer[j].w[i] // процесс суммирования всех множеств весов и выходов первого слоя

		}

		nn.middleLayer[i].out = sum // сразу вычисляем и записываем выходное значение нейрона

	}
	var sum int8
	for i := 0; i < len(nn.middleLayer); i++ {
		sum += nn.middleLayer[i].out * nn.middleLayer[i].w[0]
	}
	nn.finalNeuron.out = sum
}

func (nn *NeuronNetwork) startNN(vY, y float64, ry int) int8 {
	return 0
}

func (nn *NeuronNetwork) executeNN() {

	text := getNNString(nn)
	file, err := os.Create("weigts.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)
}

func getNNString(nn *NeuronNetwork) string {

	text := ""

	for i := 0; i < len(nn.firstLayer); i++ {
		text += nn.firstLayer[i].getWString()
	}
	text += "\n"
	for i := 0; i < len(nn.middleLayer); i++ {
		text += nn.middleLayer[i].getWString()
	}

	return text
}
