package main

import "itib/lab6/neuron"

func main() {
	n := neuron.NewNeuron([]float64{-3}, []float64{-0.1}, 1, 2, 1, 1)
	n.Train()
}
